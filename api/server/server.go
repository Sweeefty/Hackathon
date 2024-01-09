package server

import (
	"Hackathon/api/server/handlers"
	"fmt"
	"net/http"
)

func StartApi() {
	//404
	http.HandleFunc("/", handlers.Error404) //GET
	//doc
	http.HandleFunc("/api", handlers.Doc) //GET
	//Accounts
	http.HandleFunc("/connection", handlers.ConnectionHandler) //GET
	http.HandleFunc("/accountininfo", handlers.GetAccountInfo) //GET
	//Inventory
	http.HandleFunc("/getInventory", handlers.GetInventory) //GET
	http.HandleFunc("/addInventory", handlers.AddInventory) //POST
	//Products
	http.HandleFunc("/getProducts", handlers.GetProducts)     //GET
	http.HandleFunc("/getProduct", handlers.GetProduct)       //GET
	http.HandleFunc("/addProduct", handlers.AddProduct)       //POST
	http.HandleFunc("/deleteProduct", handlers.DeleteProduct) //DELETE
	//Bde
	http.HandleFunc("/getBde", handlers.GetBde) //GET

	Port := "8080"                                    //We choose our port
	fmt.Println("api started on port " + Port + " 🚀") //We print this when the server is online
	fmt.Println("http://localhost:" + Port + "/")
	http.ListenAndServe(":"+Port, nil) //We start the server

}
