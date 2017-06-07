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
	
	http.ListenAndServe(port, mux)
}

func main() {
	server(":8080")
}
