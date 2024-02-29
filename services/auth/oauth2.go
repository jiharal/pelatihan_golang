package auth

import (
	"log"
	"net/http"
	"os"
	"pelatihan_golang/src"

	"github.com/gorilla/mux"
)

func StartServerOauth2() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	src.InitRouterMux(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", r)
	log.Printf("Server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
