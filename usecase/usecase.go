package usecase

import (
	"LastFM/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var apiKey = "a8e35717f897047c80a1b447e9438370"
var urlAlbumGetInfo = "http://ws.audioscrobbler.com/2.0/"

func GetBody(nameArtist, albumName string) (*[]byte, error) {
	var (
		err error
	)
	req, err := http.NewRequest(http.MethodGet, urlAlbumGetInfo, nil)
	q := req.URL.Query()
	q.Add("method", "album.getinfo")
	q.Add("api_key", apiKey)
	q.Add("artist", nameArtist)
	q.Add("album", albumName)
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()
	res, err := http.Get(req.URL.String())
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error request")
	}

	return &body, nil
}

func ArtistAlbum(body *[]byte) (*models.Result, error) {
	var data models.AlbumU
	if err := json.Unmarshal(*body, &data); err != nil {
		return nil, errors.Wrap(err, "Error request")
	}
	var parsedData []models.AlbumTrack
	for i := range data.Album.Tracks.Track {
		parsedData = append(parsedData, models.AlbumTrack{Id: i + 1, URL: data.Album.Tracks.Track[i].Url, TrackName: data.Album.Tracks.Track[i].Name})
	}
	result := models.Result{AlbumName: string(data.Album.Name), AlbumArtist: data.Album.Artist, AlbumTrack: parsedData}
	return &result, nil
}
