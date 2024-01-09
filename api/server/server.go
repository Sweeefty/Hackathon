package server

import (
	"Hackathon/api/server/handlers"
	"fmt"
	"net/http"
)

func StartApi() {
	//Swagger
	//http.Handle("/swagger/", httpSwagger.Handler(
	//	httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // This URL should point to your generated Swagger JSON file
	//))
	//Root
	http.HandleFunc("/", handlers.Root) //GET
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
