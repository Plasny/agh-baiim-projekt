package routes

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const sessionKey contextKey = "session"

type Session struct {
	Task1Completed bool `json:"task1_completed"`
	task1Password  string

	Task2Completed bool `json:"task2_completed"`
	Task3Completed bool `json:"task3_completed"`
	Task4Completed bool `json:"task4_completed"`
}

var Sessions = map[string]Session{}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sessionId string
		cookie, err := r.Cookie("session_id")
		if err != nil || cookie.Value == "" {
			sessionId = uuid.New().String()

			Sessions[sessionId] = Session{
				task1Password: "default",
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "session_id",
				Value: sessionId,
				Domain:   "", 
				Secure:   false, 
    			SameSite: http.SameSiteLaxMode,
				HttpOnly: true,
			})
		} else {
			sessionId = cookie.Value
		}

		ctx := context.WithValue(r.Context(), sessionKey, sessionId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
