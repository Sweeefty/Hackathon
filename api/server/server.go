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

	styles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", styles)) //We link the css with http.Handle

	http.HandleFunc("/", handlers.Root)                        //We create the root page GET
	http.HandleFunc("/connection", handlers.ConnectionHandler) //We create the connection page GET
	http.HandleFunc("/getInventory", handlers.GetInventory)    //We create the get inventory page GET
	http.HandleFunc("/addInventory", handlers.AddInventory)    //We create the add inventory page POST

	Port := "8080"                                    //We choose our port
	fmt.Println("api started on port " + Port + " 🚀") //We print this when the server is online
	fmt.Println("http://localhost:" + Port + "/")
	http.ListenAndServe(":"+Port, nil) //We start the server}

}
