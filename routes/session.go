package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const sessionKey contextKey = "session"

type Session struct {
	Id string `json:"id"`

	Task1Completed bool `json:"task1_completed"`
	task1Password  string

	Task2Completed bool `json:"task2_completed"`
	task2Status    string

	Task3Completed bool `json:"task3_completed"`
	task3Role      string

	Task4Completed bool `json:"task4_completed"`
}

var Sessions = map[string]Session{}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// --- Handle CORS preflight requests -------------------------
		if r.Method == "OPTIONS" {
			http.Error(w, "CORS Preflight not allowed", http.StatusForbidden)
			return
		}

		// --- Manage session -----------------------------------------
		var sessionId string
		cookie, err := r.Cookie("session_id")

		if err != nil || cookie.Value == "" {
			sessionId = uuid.New().String()

			Sessions[sessionId] = Session{
				Id:            sessionId,
				task1Password: "default",
				task2Status:   "oczekujący na wykonanie",
				task3Role:     "guest",
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    sessionId,
				Domain:   "",
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
				HttpOnly: true,
			})
		} else {
			sessionId = cookie.Value
		}

		ctx := context.WithValue(r.Context(), sessionKey, sessionId)

		// --- Logging ------------------------------------------------
		log.Printf("Session ID: %s, Path: %s, Method: %s, Origin: %s", sessionId, r.URL.Path, r.Method, r.Header.Get("Origin"))

		// --- Call the next handler ----------------------------------
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
