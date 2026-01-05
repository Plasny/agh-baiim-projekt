package routes

import (
	"fmt"
	"log"
	"net/http"
)

func Task1Handler(w http.ResponseWriter, r *http.Request) {
	sessionID, ok := r.Context().Value(sessionKey).(string)
	if !ok {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	session := Sessions[sessionID]
	defer func() { Sessions[sessionID] = session }()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	password := r.URL.Query().Get("password")

	if password != "" {
		session.task1Password = password
	}
	if session.task1Password == "abc12345" && r.Referer() == "http://localhost:8080/" {
		session.Task1Completed = true
		log.Printf("Session %s completed Task 1", sessionID)
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<body>
				<h2>Good job</h2>
			</body>
			</html>
		`)
		return
	}

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<body>
			<h2>Settings Menu</h2>
			<form method="get" action="/task1">
				<label>Password: <input type="password" name="password"></label><br>
				<input type="submit" value="Update">
			</form>
		</body>
		</html>
	`)
}
