package gt

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

var BaseURL = "https://groupietrackers.herokuapp.com/api/"

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(errors.New("PAGE NOT FOUND"), "", 404)
		return
	}

	switch r.Method {
	case "POST":
		ErrorPage(errors.New("MÉTHOD NOT ALLOWED"), "", 405)
		return
	default:
		tStart := time.Now()
		var data AllData
		parseDone := make(chan struct{})

		go func() {
			Parse()
			close(parseDone)
		}()
		
		<-parseDone

		data.Artists = *endpoints["artists"].(*[]Artist)
		data.Locations = *endpoints["locations"].(*Location)
		data.Dates = *endpoints["dates"].(*Date)
		data.Relations = RedableRelations(*endpoints["relation"].(*Relation))

		for _, val := range data.Artists {
			data.FinalArtists = append(data.FinalArtists, FinalArtist{val.Id, val.Image, val.Name, val.Members, val.CreationDate, val.FirstAlbum, data.Locations.Index[val.Id-1], data.Dates.Index[val.Id-1], data.Relations.Index[val.Id-1]})
		}

		t, err := template.ParseFiles("vues/index.html")
		ErrorPage(err, "", 500)

		tmpl := template.Must(t, err)

		err = tmpl.Execute(w, data)
		ErrorPage(err, "", 500)

		elapsed := time.Since(tStart)
		log.Printf("Parsing took %.8fs\n", elapsed.Seconds())
	}
}

func ErrorPage(err error, msg string, code int) {
	if err != nil {
		fmt.Println(strconv.Itoa(code) + " ERROR")
		panic(err)

	}
}
