package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type Profile struct {
	Name string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", Log(http.DefaultServeMux))
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
