package userDbConnect

import (
	"database/sql"
	"encoding/json"
	"github.com/erdinat/shoes_price_tracker/model"
	"log"
	"net/http"
)

var db *sql.DB

var user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Only POST and GET allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username or password is empty", http.StatusBadRequest)
		return
	}

	userExists := model.CheckUser(db, username, password)

	response := map[string]interface{}{
		"success": userExists,
		"message": "Successfully logged in",
	}

	if !userExists {
		response["message"] = "Username is not found"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SingupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request not allowed", http.StatusMethodNotAllowed)
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Username or email or password is empty", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("Error executing query:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
