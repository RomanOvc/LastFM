package main

import (
	"LastFM/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func trackByAlbumArtist(w http.ResponseWriter, r *http.Request) {

	var (
		u   []byte
		err error
	)
	defer func() {
		if err != nil {
			log.Println(err, "Error request")
			w.Write([]byte(err.Error()))
		} else {
			w.Write(u)
		}
	}()
	w.Header().Set("Content-Type", "application/json")
	artist := r.URL.Query().Get("artist")
	album := r.URL.Query().Get("album")

	res, err := usecase.GetBody(artist, album)
	if err != nil {
		return
	}

	res1, err := usecase.ArtistAlbum(res)
	if err != nil {
		return
	}

	u, err = json.Marshal(res1)
	if err != nil {
		return
	}

	return
}

func lastfmApi() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/info", trackByAlbumArtist).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	lastfmApi()
}
