package user_controller

import (
	"encoding/json"
	"net/http"
)

var users []User

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user = User{Id: "183gh9", Name: "Giovani"}
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
