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

var BaseURL = "https://groupietrackers.herokuapp.com/api/"

func ReadAPI(url string, allDatas interface{}) {
	r, err := http.Get(url)
	HandleError(err)

	data, err1 := io.ReadAll(r.Body)
	HandleError(err1)

	json.Unmarshal(data, &allDatas)
}
