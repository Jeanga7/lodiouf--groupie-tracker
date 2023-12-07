package gt

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func ReadAPI() []Artist {
	r, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	HandleError(err)

	data, err1 := io.ReadAll(r.Body)
	HandleError(err1)

	var allArtists []Artist

	json.Unmarshal(data, &allArtists)

	return allArtists
}
