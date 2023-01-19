package handler

import (
	"final-project/database"
	"final-project/repository"
	"final-project/structs"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	err, users := repository.GetAllUser(database.DbConnection)
	if err != nil {
		fmt.Println(users)
	}
	return c.JSON(fiber.Map{"status": "success", "data": users})
}

func GetUsers(c *fiber.Ctx) error {
	var user structs.User

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	user.ID = int(id)

	users, err := repository.GetAllUserById(database.DbConnection, user)
	return c.JSON(fiber.Map{"status": "success", "message": users})

}

func StoreUser(c *fiber.Ctx) error {
	var user structs.User

	err := c.BodyParser(&user)

	if err != nil {
		panic(err)
	}

	err = repository.InsertUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Successfully Insert User"})
}

func UpdateUser(c *fiber.Ctx) error {
	var user structs.User
	id, _ := strconv.Atoi(c.Params("id"))

	err := c.BodyParser(&user)
	if err != nil {
		panic(err)
	}

	user.ID = int(id)

	err = repository.UpdateUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Successfully Update User"})

}

func DeleteUser(c *fiber.Ctx) error {
	var user structs.User
	id, err := strconv.Atoi(c.Params("id"))

	user.ID = int(id)

	err = repository.DeleteUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Successfully Deleted User"})
}
