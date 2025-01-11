package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Models
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Database struct {
	sync.Mutex
	users []*User
}

// In Memory Database for dev

var db = &Database{
	users: []*User{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 22},
		{ID: 3, Name: "Charlie", Age: 22},
	},
}

// Handlers

func HandleGetMessage(w http.ResponseWriter, r *http.Request) {
	message := struct {
		Content string `json:"content"`
	}{
		Content: `This is my example CRUD application for user managment`,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Lock()
	defer db.Unlock()
	if err := json.NewEncoder(w).Encode(db.users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
	}
}

func HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(idStr, err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db.Lock()
	defer db.Unlock()
	user, err := findUserByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func HandlePostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Age <= 0 {
		http.Error(w, "Name and Age are required and must be valid", http.StatusBadRequest)
		return
	}

	db.Lock()
	defer db.Unlock()

	user.ID = db.generateID()
	db.users = append(db.users, &user)
	// Why do this? is this for handling contracts?
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)
}

func HandlePutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateUser User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
	}

	db.Lock()
	defer db.Unlock()

	user, err := findUserByID(updateUser.ID)
	if err != nil {
		log.Println(updateUser.ID, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}

	if updateUser.Age > 0 {
		user.Age = updateUser.Age
	}

	json.NewEncoder(w).Encode(user)

}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(idStr, err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db.Lock()
	defer db.Unlock()

	index, err := findUserIndexByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db.users = append(db.users[:index], db.users[index+1:]...)
	w.WriteHeader(http.StatusNoContent)

}

// Helper Functions

func (d *Database) generateID() int {
	if len(d.users) == 0 {
		return 1
	}
	return d.users[len(d.users)-1].ID + 1
}

func findUserByID(id int) (*User, error) {
	for _, user := range db.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func findUserIndexByID(id int) (int, error) {
	for i, user := range db.users {
		if user.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("User not found")
}

// Application code
type App struct {
	Router *http.ServeMux
}

func NewApp() *App {
	return &App{
		Router: http.NewServeMux(),
	}
}

func (a *App) Run(port string) {
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func main() {
	app := NewApp()
	app.Router.Handle("/", http.FileServer(http.Dir("./static")))
	app.Router.HandleFunc("GET /message", HandleGetMessage)
	app.Router.HandleFunc("GET /user", HandleGetUser)
	// I understand why I am using path vars here
	app.Router.HandleFunc("GET /user/{id}", HandleGetUserByID)
	app.Router.HandleFunc("POST /user", HandlePostUser)

	// Why do I use path vars for both of these?
	app.Router.HandleFunc("PUT /user", HandlePutUser)
	app.Router.HandleFunc("DELETE /user/{id}", HandleDeleteUser)

	app.Run(":8080")
}
