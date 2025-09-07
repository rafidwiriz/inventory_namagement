package main

import (
	"fmt"
	"inventory_management/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string = ":5000"

func main() {
	err := handlers.DBConnection()
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets/"))

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	//Items
	router.HandleFunc("/items", handlers.GetAllItems).Methods("GET")
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	//router.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	//router.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	//Containers
	router.HandleFunc("/containers", handlers.GetAllContainers).Methods("GET")
	router.HandleFunc("/containers", handlers.CreateContainer).Methods("POST")
	router.HandleFunc("/containers/{id}", handlers.GetContainer).Methods("GET")
	//router.HandleFunc("/containers/{id}", handlers.UpdateContainer).Methods("PUT")
	//router.HandleFunc("/containers/{id}", handlers.DeleteContainer).Methods("DELETE")

	fmt.Println("Program dijalankan!")
	fmt.Printf("Akses program dengan menuju http://localhost%s menggunakan browser.\n", port)
	http.ListenAndServe(port, router)
}
