package gt

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var BaseURL = "https://groupietrackers.herokuapp.com/api/"
var FileNames = []string{"vues/index.html", "vues/_header.html", "vues/_footer.html"}
var ErrorFileNames = []string{"vues/error.html", "vues/_header.html", "vues/_footer.html"}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(errors.New("PAGE NOT FOUND"), w, 404)
		return
	}

	switch r.Method {
	case "POST":
		ErrorPage(errors.New("METHOD NOT ALLOWED"), w, 405)
		return
	default:
		tStart := time.Now()
		var data AllData
		parseDone := make(chan struct{})

		go func() {
			Parse(w)
			close(parseDone)
		}()

		<-parseDone

		data.Artists = *endpoints["artists"].(*[]Artist)
		data.Locations = *endpoints["locations"].(*Location)
		data.Dates = *endpoints["dates"].(*Date)
		data.Relations = RedableRelations(*endpoints["relation"].(*Relation))

		for _, val := range data.Artists {
			data.FinalArtists = append(data.FinalArtists, FinalArtist{val.Id, val.Image, val.Name, []string{strings.Join(val.Members, ", ")}, val.CreationDate, FormatDates([]string{val.FirstAlbum})[0], data.Locations.Index[val.Id-1], data.Dates.Index[val.Id-1], data.Relations.Index[val.Id-1]})
		}

		t, err := (template.ParseFiles(FileNames...))
		//necessaire de verifier l'erreur
		if err != nil {
			ErrorPage(err, w, http.StatusInternalServerError)
			return
		}
		tmpl := template.Must(t, err)

		err = tmpl.Execute(w, data)
		if err != nil {
			ErrorPage(err, w, http.StatusInternalServerError)
			return
		}

		elapsed := time.Since(tStart)
		log.Printf("Parsing took %.4fs\n", elapsed.Seconds())
	}
}

func ErrorPage(err error, w http.ResponseWriter, code int) {
	if err != nil {
		errorName := map[int]string{
			http.StatusNotFound:            http.StatusText(http.StatusNotFound),
			http.StatusMethodNotAllowed:    http.StatusText(http.StatusMethodNotAllowed),
			http.StatusInternalServerError: http.StatusText(http.StatusInternalServerError),
			http.StatusBadRequest:          http.StatusText(http.StatusBadRequest),
		}

		w.WriteHeader(code)

		tmpl, err := template.ParseFiles(ErrorFileNames...)
		if err != nil {
			panic(err)
		}
		tmpl = template.Must(tmpl, err)
		data := map[string]string{
			"error_code": strconv.Itoa(code),
			"error_name": errorName[code],
		}
		tmpl.Execute(w, data)
	}
}
