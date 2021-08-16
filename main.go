package main

import (
	"LastFM/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func trackByAlbumArtist(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)
	w.Header().Set("Content-Type", "application/json")
	artist := r.URL.Query().Get("artist")
	album := r.URL.Query().Get("album")
	res, err := usecase.GetBody(artist, album)
	res1, err := usecase.ArtistAlbum(res)
	u, err := json.Marshal(res1)

	if err != nil {
		log.Println(errors.Wrap(err, "Error request"))
	}

	w.Write(u)
}

func lastfmApi() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/info", trackByAlbumArtist).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	lastfmApi()
}
