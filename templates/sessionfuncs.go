package main
import (
	"net/http"
	"fmt"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type Session struct{
	ID       string `json:"id"`
	Tutoree  string `json:"tutoreeID"`
	Tutor    string `josn:"tutorID"`
	Time     string `json:"time"`
	Location string `json:"location"`
	Clockin  string `json:"clockin"`
	Clockout string `json:"clockout"`
}
func sDB() *sql.DB{
	db, err := sql.Open("mysql", "root:Password!@/session")
	if err != nil {
		fmt.Print("could not access the database")
	}
	return db
}
func addSession(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//parse info into dummy struct s
		s := Session{}
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("error in parsing session information. try again"))
			return
		}
		//open database
		db := sDB()
		defer db.Close()
		//enter the values in the database
		q := "INSERT INTO sessions(tutoreeID, tutorID, time, location) VALUES(?, ?, ?, ?)"
		_, err = db.Exec(q, s.Tutoree, s.Tutor, s.Time, s.Location)
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
func getAllTutorSessions(w http.ResponseWriter, r *http.Request) {
	//struct for body
	//{"id": 189, "access":"secret code"}
	s := Session{}
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
		q:= "SELECT * FROM sessions WHERE tutorID=?"
		err  := db.QueryRow(q, bod.ID).Scan(&s.ID, &s.Tutoree, &s.Tutor, &s.Time, &s.Location, &s.Clockin, &s.Clockout)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("err at query : check ID field"))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(s)

	} else {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect request type. Please do a post request"))
	}
}
func getSession(w http.ResponseWriter, r *http.Request) {
	//struct for body
	//{"id": 189, "access":"secret code"}
	s := Session{}
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
		q:= "SELECT * FROM sessions WHERE id=?"
		err  := db.QueryRow(q, bod.ID).Scan(&s.ID, &s.Time, &s.Tutoree, &s.Tutor)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("err at query : check ID field"))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(s)

	} else {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect request type. Please do a post request"))
	}
}
