package repository

import (
	"database/sql"
	"final-project/structs"
	"fmt"
	"time"
)

func GetAllOrder(db *sql.DB) (results []structs.Orders, err error) {
	sql := "SELECT * FROM tbl_orders"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Orders{}

		err = rows.Scan(&order.ID, &order.Users_id, &order.Products_id, &order.Total_price, &order.Total_bought, &order.Created_at, &order.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, order)

	}
	return

}

func GetAllOrderById(db *sql.DB, order structs.Orders) (results []structs.Orders, err error) {
	sql := "SELECT * FROM tbl_orders WHERE id= $1"

	rows, err := db.Query(sql, order.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Orders{}

		err = rows.Scan(&order.ID, &order.Users_id, &order.Products_id, &order.Total_price, &order.Total_bought, &order.Created_at, &order.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, order)

	}
	return

}

func InsertOrder(db *sql.DB, order structs.Orders) (err error, err2 error) {
	var product = structs.Product{}
	ers := db.QueryRow("SELECT product_price FROM tbl_products WHERE id=$1", order.Products_id).Scan(&product.Product_price)
	if ers != nil {
		panic(ers)
	}

	fmt.Println(product.Product_price)
	sql := "INSERT INTO tbl_orders (users_id, products_id, total_price, total_bought, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

	errs := db.QueryRow(sql, order.Users_id, order.Products_id, order.Total_bought*product.Product_price, order.Total_bought, time.Now(), time.Now())

	sql2 := "UPDATE tbl_products SET product_stock = (product_stock - 1) WHERE id= $1"

	errs2 := db.QueryRow(sql2, order.Products_id)

	return errs.Err(), errs2.Err()
}
