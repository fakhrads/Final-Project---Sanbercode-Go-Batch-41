package handler

import (
	"fmt"
	"time"

	"final-project/database"
	"final-project/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	type login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var body login
	err := c.BodyParser(&body)
	if err != nil {
		panic(err)
	}
	var user = structs.User{}

	err = database.DbConnection.QueryRow("SELECT username, password FROM tbl_users WHERE username=$1", body.Username).Scan(&user.Username, &user.Password)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	fmt.Println(user.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
