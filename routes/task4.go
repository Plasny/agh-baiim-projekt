package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Task4Handler(w http.ResponseWriter, r *http.Request) {
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
				<h2>Delete account</h2>
				<button id="deleteBtn">Delete Account</button>
				<script>
					const deleteBtn = document.getElementById('deleteBtn');
					deleteBtn.addEventListener('click', () => {
						fetch('/task4', {
							method: 'DELETE',
						})
							.then(res => { alert('Account deleted'); })
							.catch(err => { console.error(err); });
					})
				</script>
			</body>
			</html>
		`)
		return
	}

	if r.Method == http.MethodDelete {
		log.Printf("Received account deletion for task4, session %s", sessionID)

		w.WriteHeader(http.StatusNoContent)
		return
	}

	
	if r.Method == http.MethodPost && strings.ToUpper(r.URL.Query().Get("_method")) == "DELETE" {
		session.Task4Completed = true
		log.Printf("Session %s completed Task 4", sessionID)

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

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
