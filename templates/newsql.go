package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
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

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {

		u := User{}

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			panic(err)
		}

		db, err := sql.Open("mysql", "root:Password!@/users")
		if err != nil {
			panic(err)
		}
		q := "SELECT * FROM users WHERE email='%s'"
		rows, err := db.Query(q, u.Email)
		if err != nil {
			panic(err)
		}

		if rows.Next() {
			//respond with error
			//
		}

		q = "INSERT INTO users VALUES(%d, %s)"
		_, err = db.Exec(q, u.ID, u.Email)
		if err != nil {
			//respond with error (on the server side)
			panic(err)
		}
		//new user
	} else {
		//return error
	}
}

func server (port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/adduser", addUser)
	
	http.ListenAndServe(port, mux)
}

func main() {
	
}
