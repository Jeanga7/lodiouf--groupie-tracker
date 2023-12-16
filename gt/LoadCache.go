package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ReadAPI(url string, allDatas interface{}) error {
	r, err := http.Get(url)
	ErrorPage(err, "Erreur HTTP:", 500)
	defer r.Body.Close()

	data, err1 := io.ReadAll(r.Body)
	ErrorPage(err1, "", 500)

	err1 = json.Unmarshal(data, &allDatas)
	ErrorPage(err1, "", 500)
	_, errNetwork := http.Get(url)
	if errNetwork != nil {
		return fmt.Errorf("erreur lors de l'appel à l'API : %v", err)
	}

	return nil
}

func SaveDataToLocalFile(filename string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadDataFromCache(filename string, data interface{}) error {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, data)
	if err != nil {
		return err
	}

	return nil
}

func FetchAndCacheData() {
	var allDatas AllData

	ReadAPI(BaseURL, &allDatas)

	SaveDataToLocalFile("cached_data.json", allDatas)
}
