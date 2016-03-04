package main

import (
	"fmt"
	"net/http"
	"os"
)

const portEnv = "PORT"
const defaultPort = "3000"

func port(defaultPort string) string {
	port := os.Getenv(portEnv)
	if port == "" {
		port = defaultPort
	}

	return fmt.Sprintf(":%s", port)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Success")
}

func main() {
	port := port(defaultPort)

	http.HandleFunc("/", handler)

	fmt.Printf("Server started on %s...\n", port)
	http.ListenAndServe(port, nil)
}
