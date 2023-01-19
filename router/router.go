package router

import (
	"final-project/handler"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func SetupRoutes(api *fiber.App) {

	app := api.Group("/api/v1")
	app.Post("/login", handler.Login)
	app.Post("/register", handler.StoreUser)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/user", handler.GetUser)
	app.Get("/user/:id", handler.GetUsers)
	app.Put("/user/update/:id", handler.UpdateUser)
	app.Delete("/user/:id", handler.DeleteUser)

	app.Get("/product", handler.GetProduct)
	app.Put("/product/update/:id", handler.UpdateProduct)
	app.Get("/product/:id", handler.GetProducts)
	app.Post("/product", handler.StoreProduct)
	app.Delete("/product/:id", handler.DeleteProduct)

	app.Get("/order", handler.GetOrder)
	app.Get("/order/:id", handler.GetOrders)
	app.Post("/order", handler.StoreOrder)
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
