package tools

import (
	"Hackathon/api/server/data"
	"database/sql"
	"fmt"
)

func GetProducts() interface{} {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 11")
		return false
	}
	defer db.Close()

	// Extract inventory from id
	rows, err := db.Query("SELECT id , name , description , price FROM products;")
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Product.go line 31")
		return false
	}
	Products := data.ProductList{}
	for rows.Next() {
		var Product data.Product
		rows.Scan(&Product.Id, &Product.Name, &Product.Description, &Product.Price)
		Products.Products = append(Products.Products, Product)
	}
	return Products
}

func GetProductsByBDE(bde_id string) interface{} {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 11")
		return false
	}
	defer db.Close()

	// Extract inventory from id
	rows, err := db.Query("SELECT id , name , description , price FROM products WHERE bde_id = ?;", bde_id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Product.go line 31")
		return false
	}
	Products := data.ProductList{}
	for rows.Next() {
		var Product data.Product
		rows.Scan(&Product.Id, &Product.Name, &Product.Description, &Product.Price)
		Products.Products = append(Products.Products, Product)
	}
	return Products
}

func GetProductById(id string) interface{} {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 11")
		return false
	}
	defer db.Close()

	// Extract inventory from id
	rows, err := db.Query("SELECT id , name , description , price FROM products WHERE id = ?;", id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Product.go line 31")
		return false
	}
	Products := data.ProductList{}
	for rows.Next() {
		var Product data.Product
		rows.Scan(&Product.Id, &Product.Name, &Product.Description, &Product.Price)
		Products.Products = append(Products.Products, Product)
	}
	return Products
}

func AddProduct(name string, description string, price string, bde_id string) bool {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Product.go line 47")
		return false
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO products (name , description , price , bde_id) VALUES (?,?,?,?)", name, description, price, bde_id)
	if err != nil {
		fmt.Println(err.Error())
		WriteErr("Error inserting into database tools/Product.go line 54")
		return false
	}
	return true
}

func DeleteProduct(id string) bool {
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Product.go line 63")
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT id FROM products WHERE id = ?", id)
	if err != nil {
		WriteErr("Error querying database tools/Product.go line 88")
		return false
	}

	if rows.Next() == false {
		return false
	}
	rows.Close()

	_, err = db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		WriteErr("Error deleting from database tools/Product.go line 88")
		return false
	}
	return true
}
