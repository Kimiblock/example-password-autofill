package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	addr = ":8443" // Rev proxied by NGINX
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	log.Printf("Starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("Received username: %s, password: %s\n", username, password)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Password Login</title>
	</head>
	<body>
	<h1>Password Login</h1>
	<form method="POST">
	<label for="username">Username:</label>
	<input type="text" id="username" name="username" required>
	<br>
	<label for="password">Password:</label>
	<input type="password" id="password" name="password" autocomplete="current-password" required>
	<br>
	<button type="submit">Submit</button>
	</form>
	</body>
	</html>
	`
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
