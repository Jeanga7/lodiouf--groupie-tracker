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

func ReadAPI(url string, allArtists interface{}) {
	r, err := http.Get(url)
	HandleError(err)

	data, err1 := io.ReadAll(r.Body)
	HandleError(err1)

	json.Unmarshal(data, &allArtists)
}
