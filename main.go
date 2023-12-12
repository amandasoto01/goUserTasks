package main

import (
	"github.com/amandasoto01/go-gorm-rest-api/db"
	"github.com/amandasoto01/go-gorm-rest-api/models"
	"github.com/amandasoto01/go-gorm-rest-api/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db.DbConnection()
	log.Println("Database connection created")

	db.DB.AutoMigrate(models.Task{})
	log.Println("Task table created")
	db.DB.AutoMigrate(models.User{})
	log.Println("User table created")

	//if errTask != nil {
	//	log.Print("Error in task creation: ", errTask)
	//	return
	//}
	//
	//if errUser != nil {
	//	log.Print("Error in user creation: ", errUser)
	//	return
	//}

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	// task routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
