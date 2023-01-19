package router

import (
	"final-project/handler"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", handler.Login)
	app.Post("/api/v1/register", handler.StoreUser)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/api/v1/user", handler.GetUser)
	app.Get("/api/v1/user/:id", handler.GetUsers)
	app.Put("/api/v1/user/update/:id", handler.UpdateUser)
	app.Delete("/api/v1/user/:id", handler.DeleteUser)

	app.Get("/api/v1/product", handler.GetProduct)
	app.Put("/api/v1/product/update/:id", handler.UpdateProduct)
	app.Get("/api/v1/product/:id", handler.GetProducts)
	app.Post("/api/v1/product", handler.StoreProduct)
	app.Delete("/api/v1/product/:id", handler.DeleteProduct)

	app.Get("/api/v1/order", handler.GetOrder)
	app.Get("/api/v1/order/:id", handler.GetOrders)
	app.Post("/api/v1/order", handler.StoreOrder)
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
