package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Args[1]
	for _, path := range os.Args[2:] {
		fs := http.FileServer(http.Dir(path))
		parts := strings.Split(path, string(os.PathSeparator))
		route := "/" + strings.Replace(parts[len(parts)-1], " ", "_", -1)
		http.Handle(route+"/", http.StripPrefix(route, fs))
	}

	log.Println(http.ListenAndServe(":"+port, nil))
}
