package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	switch user {
	case "username": // username
		return "$1$RzNAxMqf$2Cjw9950B0wEE3CIthyJj." // password  https://unix4lyfe.org/crypt/
	default:
		return ""
	}
}

func main() {
	authenticator := auth.NewBasicAuthenticator("agent", Secret)

	port := os.Args[1]
	for _, path := range os.Args[2:] {
		fs := http.FileServer(http.Dir(path))
		parts := strings.Split(path, string(os.PathSeparator))
		route := "/" + strings.Replace(parts[len(parts)-1], " ", "_", -1)
		http.HandleFunc(route+"/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
			http.StripPrefix(route, fs).ServeHTTP(w, &r.Request)
		}))
	}

	log.Println(http.ListenAndServe(":"+port, nil))
}
