package data

type Product struct {
	Id          string
	Name        string
	Description string
	Price       string
}

type ProductList struct {
	Products []Product
}
