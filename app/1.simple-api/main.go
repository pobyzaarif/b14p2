package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

var Users []User

func main() {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		mapPanic := map[string]string{"message": "panic"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(mapPanic)
	}

	router.GET("/users", GetUser)
	router.POST("/users", CreateUser)
	router.GET("/panic", PanicTest)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Create Endpoint
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	var user User
	_ = json.Unmarshal(body, &user)

	Users = append(Users, user)
	http.Error(w, "Created", http.StatusCreated)
}

// Get Endpoint
func GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// users, _ := json.Marshal(Users)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Users)
	// http.Error(w, string(users), http.StatusOK)
}

func PanicTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sliceOfString := []string{"edric", "febian", "firman"}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(sliceOfString[3])
}
