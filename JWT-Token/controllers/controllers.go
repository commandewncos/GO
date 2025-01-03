package controllers

import (
	"encoding/json"
	"net/http"
	"token/authorization"
)

type Message struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Message: "Public route", Status: "200"}
	j, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginParams LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginParams); err != nil {
		message := Message{
			Message: "Invalid credentials",
			Status:  "401",
		}

		j, err := json.Marshal(message)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(j)
			return
		}

	}

	if loginParams.Username == "Wilson" && loginParams.Password == "123456" {

		token, err := authorization.GenerateJwtToken(loginParams.Username)

		if err != nil {
			http.Error(w, "not authorized", http.StatusInternalServerError)
			return
		} else {
			message := Message{
				Message: token,
				Status:  "200",
			}
			j, _ := json.Marshal(message)
			w.WriteHeader(http.StatusOK)
			w.Write(j)

		}

		return
	} else {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Message: "Secure route",
		Status:  "200",
	}

	j, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func AuthMiddlewarew(next http.HandlerFunc) http.HandlerFunc {
	message := Message{
		Message: "not authorized",
		Status:  "401",
	}

	j, _ := json.Marshal(message)

	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Api-Token")
		if len(token) == 0 {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(j)
			return
		}

		next.ServeHTTP(w, r)
	}
}
