package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Listening on :8080")
	panic(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", os.Getenv("REDIR_SERVER"))
		r.URL.Scheme = "https"
		r.URL.Host = r.Host
		if r.URL.Host == "" {
			r.URL.Host = os.Getenv("REDIR_FALLBACK")
		}
		http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
	})))
}
