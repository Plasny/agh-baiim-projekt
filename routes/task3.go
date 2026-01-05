package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task3Data struct {
	Role string `json:"role"`
}

func Task3Handler(w http.ResponseWriter, r *http.Request) {
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
				<h2>role = %s</h2>
				<form id="roleForm">
					<select name="role" id="roleSelect">
						<option value="guest">guest</option>
						<option value="editor">editor</option>
					</select>
					<button type="submit">Update</button>
				</form>
				<script>
					const form = document.getElementById('roleForm');
					form.addEventListener('submit', (e) => {
						e.preventDefault();
						const select = document.getElementById('roleSelect');
						const role = select.value;
						fetch('/task3', {
							method: 'POST',
							body: JSON.stringify({ role }),
						})
							.then(res => { 
								alert('Role updated to ' + role)
								window.location.reload();
							})
							.catch(err => { console.error(err); });
					});
				</script>
			</body>
			</html>
		`, session.task3Role)
		return
	}

	if r.Method == http.MethodPost {
		jsonData := Task3Data{}
		err := json.NewDecoder(r.Body).Decode(&jsonData)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		session.task3Role = jsonData.Role
		log.Printf("Received role update for task3, session %s: %s", sessionID, jsonData.Role)

		if jsonData.Role == "admin" && r.Header.Get("Origin") == "http://localhost:8080" {
			session.Task3Completed = true
			log.Printf("Session %s completed Task 3", sessionID)
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

		http.Redirect(w, r, "/task3", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
