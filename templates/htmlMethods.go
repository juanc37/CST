package main

import (
	"net/http"
	"fmt"
)

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
