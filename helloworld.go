package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//call handler
	http.HandleFunc("/", handle)
	//Do healthcheck
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 3001")
	//Listen an serve on port 8080
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {

	//if request path invalid return error
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//Check if env var has been set in yaml
	target := os.Getenv("TARGET")
	if target == "" {
		target = "Dev"
	}
	//else return message in responsewriter
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
