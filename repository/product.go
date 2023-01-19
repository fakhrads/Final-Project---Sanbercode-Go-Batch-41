package repository

import (
	"database/sql"
	"final-project/structs"
	"fmt"
	"time"
)

func GetAllProduct(db *sql.DB) (err error, results []structs.Product) {
	sql := "SELECT * FROM tbl_products"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product = structs.Product{}

		err = rows.Scan(&product.ID, &product.Product_name, &product.Product_description, &product.Product_stock, &product.Product_price, &product.Created_at, &product.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, product)

	}
	return

}

func GetAllProductById(db *sql.DB, product structs.Product) (results []structs.Product, err error) {
	sql := "SELECT * FROM tbl_products WHERE id= $1"

	rows, err := db.Query(sql, product.ID)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product = structs.Product{}

		err = rows.Scan(&product.ID, &product.Product_name, &product.Product_description, &product.Product_stock, &product.Product_price, &product.Created_at, &product.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, product)
		fmt.Println(results)

	}
	return

}

func InsertProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "INSERT INTO tbl_products (product_name, product_description, product_stock, product_price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	fmt.Println(product.Product_name)
	errs := db.QueryRow(sql, product.Product_name, product.Product_description, product.Product_stock, product.Product_price, time.Now(), time.Now())

	return errs.Err()
}

func UpdateProducts(db *sql.DB, product structs.Product) (err error) {
	sql := "UPDATE tbl_products SET product_name= $2, product_description= $3, product_stock= $4, product_price= $5, updated_at= $6 WHERE id= $1"

	errs := db.QueryRow(sql, product.ID, product.Product_name, product.Product_description, product.Product_stock, product.Product_price, time.Now())

	return errs.Err()
}

func DeleteProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "DELETE FROM tbl_products WHERE id= $1"

	errs := db.QueryRow(sql, product.ID)

	return errs.Err()
}
