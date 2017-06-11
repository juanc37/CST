package main
import (
	"net/http"
	"fmt"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type qUser struct {
	ID string `json:"id"`
	Email string `json:"email"`
	EncrPass string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Accesscode string `json:"access"`
}
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	EncrPass  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>CST</title>
	</head>
	<style>
		html, body, h1{
		padding: 0;
		border: 0;
		margin: 0;
		box-sizing: border-box;
	}
	body{
		justify-content: center;
		align-items: center;
		background-image: url("http://qualityfence.com/wp-content/uploads/2015/04/blueback1.jpg");
		height: 100vh;
	}
	</style>
	<h1>Welcome to the Coffee Shop Tutors website!</h1>
	<body>
	<p>ZakyD, requests work like this:</p>
	<p>go to thisAddress/people to get a list of all users</p>
	<p>go to thisAddress/people/idnumber to get user information for someone with a certain id. up until now these have been get requests</p>
	<p>go to thisAddress/people/idnumber with a delete request to delete a certain user from the cache</p>
	<p>go to thisAddress/people/idnumber with a post request to post a user. the body should look like this (how it looks like when you do a get user request): {\"id\":\"69\" ....}</p>
	<p>KEEP IN MIND: this is not connected to the sql server yet so the information will get reset every time i rerun the program on my linux VM so the way of doing it might change for u in the future")</p>
	</body>
	</html>`)
}
func isUniqueUser(u qUser, db *sql.DB, w http.ResponseWriter) bool{
	u1 := User{}
	//checking if email is already in database
	q := "SELECT * FROM users WHERE email=?"
	err := db.QueryRow(q, u.Email).Scan(&u1.ID, &u1.Email, &u1.EncrPass, &u1.Firstname, &u1.Lastname)
	if err == nil {
		//dont take the input and recommend logging in with a forgot password button when
		// the user enters a signup email that is the same as one in the database
		if u1.Email == u.Email {
			w.WriteHeader(400)
			w.Write([]byte("This email has already been used. Queue login?"))
			return false
		}
	} else if err == sql.ErrNoRows{
		return true
	} else {
		w.WriteHeader(400)
		w.Write([]byte("error at query for email"))
		return false
	}
	return false
}

func uDB() *sql.DB{
	db, err := sql.Open("mysql", "root:Password!@/users")
	if err != nil {
		fmt.Print("could not access the database")
	}
	return db
}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		u := qUser{}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("error in parsing user information. try again"))
			return
		}
		if u.Accesscode != "dong"{
			w.WriteHeader(400)
			w.Write([]byte("icorrect passcode to create a user"))
			return
		}
		db := uDB()
		defer db.Close()
		////checking if email is already in database
		if isUniqueUser(u, db, w) == false{
			return
		}
		//enter the values in the database
		q := "INSERT INTO users(email, password, firstname, lastname) VALUES(?, ?, ?, ?)"
		_, err = db.Exec(q, u.Email, u.EncrPass, u.Firstname, u.Lastname)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("error when writing info to database. (incorrect format?) "))
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
	//{"id": 189, "access":"secret code"}
	u := User{}
	//creating stuct & json parse
	type IdBody struct {
		ID int `json:"id"`
		Accesscode string `json:"access"`
	}
	bod := IdBody{}
	err := json.NewDecoder(r.Body).Decode(&bod)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error in parsing user information. try again"))
		return
	}
	if bod.Accesscode != "wiener"{
		w.WriteHeader(400)
		w.Write([]byte("icorrect passcode to access users"))
		return
	}
	db := uDB()
	defer db.Close()
	//check post method
	if r.Method == http.MethodPost {
		//query, parse and encode
		q:= "SELECT * FROM users WHERE id=?"
		err  := db.QueryRow(q, bod.ID).Scan(&u.ID, &u.Email, &u.EncrPass, &u.Firstname, &u.Lastname)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("err at query : check ID field"))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(u)

	} else {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect request type. Please do a post request"))
	}
}

