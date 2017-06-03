package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	EncrPass string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func isUniqueUser(u User){
	//Todo: create a method for this
}
func DB() *sql.DB{
	db, err := sql.Open("mysql", "root:Password!@/users")
	if err != nil {
		fmt.Print("could not access the database")
	}
	return db
}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		u := User{}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("error in parsing user information. try again"))
			return
		}

		db := DB()
		defer db.Close()
		//checking if email is already in database
		var chkemail string
		q := "SELECT * FROM users WHERE email=?"
		err = db.QueryRow(q, u.Email).Scan(&chkemail)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("error at query for duplicate email"))
			return
		}
		//dont take the input and recommend logging in with a forgot password button when
		// the user enters a signup email that is the same as one in the database
		if chkemail == u.Email {
			w.WriteHeader(400)
			w.Write([]byte("This email has already been used. Queue login?"))
			return
		}

		q = "INSERT INTO users VALUES(?, ?, ?, ?, ?)"
		_, err = db.Exec(q, u.ID, u.Email, u.EncrPass, u.Firstname, u.Lastname)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("error writing values in the "))

			return
		}
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect request type. Please do a post request"))
	}
}
func getUser(w http.ResponseWriter, r *http.Request) {
	//struct for body
	//{"id": 189}
	u := User{}
	//creating stuct & json parse
	type IdBody struct {
		ID int `json:"id"`
	}
	id := IdBody{}
	json.NewDecoder(r.Body).Decode(&id)
	db := DB()
	defer db.Close()
	//check post method
	if r.Method == http.MethodPost {
		//query, parse and encode
		q:= "SELECT * FROM users WHERE id=?"
		//var uid, ue, up, uf, ul string
		err  := db.QueryRow(q, id.ID).Scan(&u.ID, &u.Email, &u.EncrPass, &u.Firstname, &u.Lastname)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("err at query"))
			return
		}
		//u.ID = string(uid)
		//u.Email = ue
		//u.EncrPass = up
		//u.Firstname =uf
		//u.Lastname = ul
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(u)

	} else {
	w.WriteHeader(400)
	w.Write([]byte("Incorrect request type. Please do a post request"))
	}
}

func server (port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/adduser", addUser)
	mux.HandleFunc("/api/getuser", getUser)
	
	http.ListenAndServe(port, mux)
}

func main() {
	server(":8080")
}
