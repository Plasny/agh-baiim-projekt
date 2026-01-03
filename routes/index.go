package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, ok := r.Context().Value(sessionKey).(string)
	// should never happen
	if !ok {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	session := Sessions[sessionID]
	templ := template.Must(template.ParseFiles("templates/index.html"))
	err := templ.ExecuteTemplate(w, "index", session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, ok := r.Context().Value(sessionKey).(string)
	if !ok {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	session := Sessions[sessionID]
	jsonText, err := json.Marshal(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonText)
}
