package main

import (
	"HttpServerPureGolang/main/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {

	envError := godotenv.Load()
	if envError != nil {
		fmt.Print(envError)
	}

	router := mux.NewRouter()
	router.HandleFunc("/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/contacts", controllers.GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", controllers.GetContact).Methods("GET")

	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	fmt.Println("server started at port: " + port)

	serverLaunchError := http.ListenAndServe(":"+port, router)
	if serverLaunchError != nil {
		fmt.Print(serverLaunchError)
	}

}
