package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"
	
	"github.com/gorilla/mux"
	"time"
	"fmt"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo entities.Todo
	now := time.Now()
	t_string := now.String()
	todo.CreatedAt = t_string
	fmt.Println(t_string)
	json.NewDecoder(r.Body).Decode(&todo)
	database.Instance.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	if checkIfTodoExists(todoID) == false {
		json.NewEncoder(w).Encode("Todo Not Found!")
		return
	}
	var todo entities.Todo
	database.Instance.First(&todo, todoID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var products []entities.Todo
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	if checkIfTodoExists(todoID) == false {
		json.NewEncoder(w).Encode("Todo Not Found!")
		return
	}
	var todo entities.Todo
	database.Instance.First(&todo, todoID)
	json.NewDecoder(r.Body).Decode(&todo)
	database.Instance.Save(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todoID := mux.Vars(r)["id"]
	if checkIfTodoExists(todoID) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Todo Not Found!")
		return
	}
	var todo entities.Todo
	database.Instance.Delete(&todo, todoID)
	json.NewEncoder(w).Encode("Todo Deleted Successfully!")
}

func checkIfTodoExists(todoID string) bool {
	var todo entities.Todo
	database.Instance.First(&todo, todoID)
	if todo.ID == 0 {
		return false
	}
	return true
}
