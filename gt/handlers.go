package gt

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var BaseURL = "https://groupietrackers.herokuapp.com/api/"

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(errors.New("PAGE NON TROUVÉE"), 404)
		return
	}

	switch r.Method {
	case "POST":
		ErrorPage(errors.New("MÉTHODE NON AUTORISÉE"), 405)
		return
	default:
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
		data.Relations = *endpoints["relation"].(*Relation)

		t, err := template.ParseFiles("vues/index.html")
		ErrorPage(err, 500)

		tmpl := template.Must(t, err)

		err = tmpl.Execute(w, data)
		ErrorPage(err, 500)
	}
}

func ErrorPage(err error, code int) {
	if err != nil {
		fmt.Println(strconv.Itoa(code) + " ERROR")
		panic(err)
	}
}
