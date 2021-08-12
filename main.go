package main

import (
	"LastFM/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func albumArtistPost(w http.ResponseWriter, r *http.Request) { // trackByAlbumAndArtist
	w.Header().Set("Content-Type", "application/json")
	artist := r.URL.Query().Get("artist")
	album := r.URL.Query().Get("album")
	res := usecase.GetContent(artist, album)
	w.Write([]byte(res))
}

func handleRequests() { // нейм сменить api и т.
	myRouter := mux.NewRouter().StrictSlash(true) // router
	myRouter.HandleFunc("/info", albumArtistPost).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
