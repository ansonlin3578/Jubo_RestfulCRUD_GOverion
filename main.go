package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	
	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterTodoRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterTodoRoutes(router *mux.Router) {
	router.HandleFunc("/api/Todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/api/Todos/{id}", controllers.GetTodoById).Methods("GET")
	router.HandleFunc("/api/Todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/Todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/Todos/{id}", controllers.DeleteTodo).Methods("DELETE")
}
