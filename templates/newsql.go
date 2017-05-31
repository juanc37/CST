package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	EncrPass string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func isUniqueUser(u User){

}
func DB() *sql.DB{
	db, err := sql.Open("mysql", "root:Password!@/users")
	if err != nil {
		panic(err)
	}
	return db
}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		u := User{}

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			panic(err)
		}

		db := DB()
		q := "SELECT * FROM users WHERE email=?"
		rows, err := db.Query(q, u.Email)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if rows.Next() {
			//respond with error
			//
		}

		q = "INSERT INTO users VALUES(?, ?, ?, ?, ?)"
		_, err = db.Exec(q, u.ID, u.Email, u.EncrPass, u.Firstname, u.Lastname)
		if err != nil {
			//respond with error (on the server side)
			panic(err)
		}
		w.WriteHeader(200)
		w.Write([]byte("gg"))
	} else {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect request type. Please do a post request"))
	}
}
func getUser(w http.ResponseWriter, r *http.Request) {
	params:= mux.Vars(r)
	db := DB()
	defer db.Close()
	if r.Method == http.MethodPost {
		q:= "SELECT * FROM users WHERE id=?"
		w.Write([]byte(params["id"]))
		rows, err  := db.Query(q, params["id"])
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		json.NewEncoder(w).Encode(rows.Next())

	} else {
	w.WriteHeader(400)
	w.Write([]byte("Incorrect request type. Please do a post request"))
	}
}

func server (port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/adduser", addUser)
	mux.HandleFunc("/api/getuser/{id}", getUser)
	
	http.ListenAndServe(port, mux)
}

func main() {
	server(":8080")
}
