package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// BasicAuth middleware
func BasicAuth(next http.HandlerFunc, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println(auth)
		// Decode base64 username:password
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || pair[0] != username || pair[1] != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func Square(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Square 2*2=4")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Del("Authorization")
	r.Header.Del("Authorization")
	r.Header.Set("Authorization", "")
	w.Header().Set("Authorization", "")
	w.Write([]byte("Log out"))
}

func main() {
	username := "admin"
	password := "passwor"

	http.HandleFunc("/", BasicAuth(HelloHandler, username, password))
	http.HandleFunc("/square", BasicAuth(Square, username, password))
	http.HandleFunc("/logout", Logout)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
