package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Location *Location `json:"location"`
}
type Location struct {
	State string `json:"state"`
	Zip string `json:"zip"`
}
var users []User
func getUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params:= mux.Vars(req)
	for _, item := range users{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}
func getUsersEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(users)
}
func createUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}
func deleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range users{
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}


func main() {
	fmt.Print("running")
	router := mux.NewRouter()
	users = append(users, User{ID: "1", Email: "jc3fake@gmail.com", Firstname: "Juan", Lastname: "C3", Location: &Location{State: "CA", Zip: "92129"}})
	users = append(users, User{ID: "2", Email: "zakydfake@gmail.com", Firstname: "Zaky", Lastname: "Demnianiek"})
	router.HandleFunc("/", handler)
	router.HandleFunc("/people", getUsersEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", getUserEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", createUserEndpoint).Methods("POST")
	router.HandleFunc("/people/delete/{id}", deleteUserEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to BagelBros CST service\n" +
		"ZakyD, requests work like this:\n" +
		"go to thisAddress/people to get a list of all users\n" +
		"go to thisAddress/people/idnumber to get user information for someone with a certain id. up until now these have been get requests\n" +
		"go to thisAddress/people/idnumber with a delete request to delete a certain user from the cache\n" +
		"go to thisAddress/people/idnumber with a post request to post a user. the body should look like this (how it looks like when you do a get user request): {\"id\":\"69\" ....} \n" +
		"KEEP IN MIND: this is not connected to the sql server yet so the information will get reset every time i rerun the program on my linux VM so the way of doing it might change for u in the future")
}
