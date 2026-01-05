package routes

import (
	"fmt"
	"log"
	"net/http"
)

func Task2Handler(w http.ResponseWriter, r *http.Request) {
	sessionID, ok := r.Context().Value(sessionKey).(string)
	if !ok {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	session := Sessions[sessionID]
	defer func() { Sessions[sessionID] = session }()

	if r.Method == http.MethodGet {
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<body>
				<h2>status = %s</h2>
				<form method="post" action="/task2">
					<label>status: <input type="text" name="status"></label><br>
					<input type="submit" value="Update">
				</form>
			</body>
			</html>
		`, session.task2Status)

		return
	}

	if r.Method == http.MethodPost {
		status := r.FormValue("status")
		session.task2Status = status
		log.Printf("Received status update for task2, session %s: %s", sessionID, status)

		if status == "wykonano" && r.Header.Get("Origin") == "http://localhost:8080" {
			session.Task2Completed = true
			log.Printf("Session %s completed Task 2", sessionID)
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

		http.Redirect(w, r, "/task2", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
