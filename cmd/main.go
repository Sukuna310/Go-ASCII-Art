package main

import (
	server "ascii-art-web-stylize/cmd/internal/server"
	"fmt"
	"log"
	"net/http"
)

type customMux struct {
	mux *http.ServeMux
}

func (m *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if the route is registered
	for p, handler := range map[string]http.Handler{
		"/ascii-art": http.HandlerFunc(server.ResultHandler),
		"/":          http.HandlerFunc(server.MainHandler),
	} {
		if r.URL.Path == p {
			handler.ServeHTTP(w, r)
			return
		}
	}

	// Check for static file routes
	if len(r.URL.Path) > len("/assets/") && r.URL.Path[:8] == "/assets/" {
		http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))).ServeHTTP(w, r)
		return
	}
	if len(r.URL.Path) > len("/templates/") && r.URL.Path[:11] == "/templates/" {
		http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))).ServeHTTP(w, r)
		return
	}

	// Handle 404 for undefined routes
	server.NotFoundHandler(w, r)
}

func main() {
	fmt.Print("Server running on http://localhost:8080 \nTo stop the server press Ctrl+C\n")

	mux := &customMux{mux: http.NewServeMux()}

	log.Fatal(http.ListenAndServe(":8080", mux))
}
