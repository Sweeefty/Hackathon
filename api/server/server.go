package server

import (
	"Hackathon/api/server/handlers"
	"fmt"
	"net/http"
)

func StartApi() {
	//404
	http.HandleFunc("/", handlers.Error404) //GET
	//Token
	http.HandleFunc("/createToken", handlers.SetAuthorization) //POST
	//doc
	http.Handle("/api", http.HandlerFunc(handlers.Doc)) //GET
	//Accounts
	http.Handle("/connection", handlers.Authorization(http.HandlerFunc(handlers.ConnectionHandler))) //GET
	http.Handle("/accountininfo", handlers.Authorization(http.HandlerFunc(handlers.GetAccountInfo))) //GET
	//Inventory
	http.Handle("/getInventory", handlers.Authorization(http.HandlerFunc(handlers.GetInventory))) //GET
	http.Handle("/addInventory", handlers.Authorization(http.HandlerFunc(handlers.AddInventory))) //POST
	//Products
	http.Handle("/getProducts", handlers.Authorization(http.HandlerFunc(handlers.GetProducts)))     //GET
	http.Handle("/getProduct", handlers.Authorization(http.HandlerFunc(handlers.GetProduct)))       //GET
	http.Handle("/addProduct", handlers.Authorization(http.HandlerFunc(handlers.AddProduct)))       //POST
	http.Handle("/deleteProduct", handlers.Authorization(http.HandlerFunc(handlers.DeleteProduct))) //DELETE
	//Bde
	http.Handle("/getBde", handlers.Authorization(http.HandlerFunc(handlers.GetBde))) //GET
	//Events 
	http.Handle("/getEvents", handlers.Authorization(http.HandlerFunc(handlers.GetEvents))) //GET
	http.Handle("/addEvent", handlers.Authorization(http.HandlerFunc(handlers.AddEvent))) //POST
	//Request
	http.Handle("/getRequest", handlers.Authorization(http.HandlerFunc(handlers.GetRequest)))       //GET
	http.Handle("/createRequest", handlers.Authorization(http.HandlerFunc(handlers.CreateRequest))) //POST
	http.Handle("/updateRequest", handlers.Authorization(http.HandlerFunc(handlers.UpdateRequest))) //UPDATE

	Port := "8080"                                    //We choose our port
	fmt.Println("api started on port " + Port + " 🚀") //We print this when the server is online
	fmt.Println("http://localhost:" + Port + "/")
	http.ListenAndServe(":"+Port, nil) //We start the server

}
