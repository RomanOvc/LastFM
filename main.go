package main

import (
	"LastFM/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func albumArtistPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	artist := r.URL.Query().Get("artist")
	album := r.URL.Query().Get("album")
	res := usecase.Get_content(string(artist), string(album))
	w.Write([]byte(res))
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/info", albumArtistPost).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
