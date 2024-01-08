package server

import (
	"Hackathon/api/server/handlers"
	"fmt"
	"net/http"
)

func StartApi() {

	styles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", styles)) //We link the css with http.Handle
	http.HandleFunc("/", handlers.Root)                          //We create the main page , the only function who use a template

	Port := "8080"                                          //We choose port 8080
	fmt.Println("The serveur start on port " + Port + " 🔥") //We print this when the server is online
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":"+Port, nil) //We start the server}

}
