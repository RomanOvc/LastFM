package usecase

import (
	"LastFM/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get_content(nameArtist, albumName string) string {

	nameArtistT := &url.URL{Path: nameArtist}
	nameArtistTEncoder := nameArtistT.String()
	albumNameT := &url.URL{Path: albumName}
	albumNameTEncoder := albumNameT.String()

	url := "http://ws.audioscrobbler.com/2.0/?method=album.getinfo&api_key=a8e35717f897047c80a1b447e9438370&artist=" + nameArtistTEncoder + "&album=" + albumNameTEncoder + "&format=json"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data *models.AlbumU

	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	var parsedData []models.AlbumTrack
	for i := 0; i < len(data.Album.Tracks.Track); i++ {
		parsedData = append(parsedData, models.AlbumTrack{Id: i + 1, URL: data.Album.Tracks.Track[i].Url, TrackName: data.Album.Tracks.Track[i].Name})

	}
	result := models.Result{AlbumName: string(data.Album.Name), AlbumArtist: data.Album.Artist, AlbumTrack: parsedData}
	u, err := json.Marshal(result)

	if err != nil {
		panic(err)
	}

	return string(u)
}
