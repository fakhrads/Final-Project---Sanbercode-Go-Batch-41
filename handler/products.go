package handler

import (
	"final-project/database"
	"final-project/repository"
	"final-project/structs"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	products, err := repository.GetAllProduct(database.DbConnection)
	if err != nil {
		fmt.Println(products)
	}
	return c.JSON(fiber.Map{"status": "success", "data": err})
}

func GetProducts(c *fiber.Ctx) error {
	var product structs.Product

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	product.ID = int(id)

	products, err := repository.GetAllProductById(database.DbConnection, product)
	fmt.Println(products)
	return c.JSON(fiber.Map{"status": "success", "message": products})

}

func StoreProduct(c *fiber.Ctx) error {
	var product structs.Product

	err := c.BodyParser(&product)

	if err != nil {
		panic(err)
	}

	err = repository.InsertProduct(database.DbConnection, product)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Successfully Insert Product"})
}

func UpdateProduct(c *fiber.Ctx) error {
	var product structs.Product
	id, _ := strconv.Atoi(c.Params("id"))

	err := c.BodyParser(&product)
	if err != nil {
		panic(err)
	}

	product.ID = int(id)

	err = repository.UpdateProducts(database.DbConnection, product)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Successfully Update Product"})

}

func DeleteProduct(c *fiber.Ctx) error {
	var product structs.Product
	id, err := strconv.Atoi(c.Params("id"))

	product.ID = int(id)

	err = repository.DeleteProduct(database.DbConnection, product)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Successfully Deleted Product"})
}
