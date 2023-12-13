package gt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
)

var BaseURL = "https://groupietrackers.herokuapp.com/api/"

func ReadAPI(url string, allDatas interface{}) {
	r, err := http.Get(url)
	ErrorPage(err, 500)

	data, err1 := io.ReadAll(r.Body)
	ErrorPage(err1, 500)

	json.Unmarshal(data, &allDatas)
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(errors.New("PAGE NOT FOUND"), 404)
	} else {
		switch r.Method {
		case "POST":
			ErrorPage(errors.New("BAD METHOD"), 405)
		default:
			type Data struct {
				Artists []Artist
			}

			data := Data{FetchArtists()}
			t, err := template.ParseFiles("vues/index.html")
			ErrorPage(err, 500)

			tmpl := template.Must(t, err)

			err = tmpl.Execute(w, data)
			ErrorPage(err, 500)
		}
	}
}

func ErrorPage(err error, code int) {
	if err != nil {
		panic(err)
		fmt.Println(strconv.Itoa(code) + " ERROR")
	}
}
