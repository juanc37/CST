package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func server (port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/api/adduser", addUser)
	mux.HandleFunc("/api/getuser", getUser)
	mux.HandleFunc("/api/addsession", addSession)
	mux.HandleFunc("/api/getsession", getSession)
	mux.HandleFunc("/api/gettutorsessions", getAllTutorSessions)
	mux.HandleFunc("/api/gettutoreesessions", getAllTutoreeSessions)
	
	http.ListenAndServe(port, mux)
}

func main() {
	server(":8080")
}
