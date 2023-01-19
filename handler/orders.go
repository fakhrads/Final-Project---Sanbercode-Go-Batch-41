package handler

import (
	"final-project/database"
	"final-project/repository"
	"final-project/structs"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetOrder(c *fiber.Ctx) error {
	orders, err := repository.GetAllOrder(database.DbConnection)
	if err != nil {
		fmt.Println(orders)
	}
	return c.JSON(fiber.Map{"status": "success", "data": err})
}

func GetOrders(c *fiber.Ctx) error {
	var order structs.Orders

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	order.ID = int(id)

	products, err := repository.GetAllOrderById(database.DbConnection, order)
	fmt.Println(products)
	return c.JSON(fiber.Map{"status": "success", "message": products})

}

func StoreOrder(c *fiber.Ctx) error {
	var order structs.Orders

	err := c.BodyParser(&order)

	if err != nil {
		panic(err)
	}

	err = repository.InsertOrder(database.DbConnection, order)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Successfully Insert Order"})
}
