package gt

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sync"
)

const cacheFilePath = "cache.json"


var endpoints = map[string]interface{}{
	"artists":   &[]Artist{},
	"locations": &Location{},
	"dates":     &Date{},
	"relation":  &Relation{},
}

func ReadAPI(url string, allDatas interface{}) {
	r, err := http.Get(url)
	ErrorPage(err, 500)
	defer r.Body.Close()

	data, err1 := io.ReadAll(r.Body)
	ErrorPage(err1, 500)

	err1 = json.Unmarshal(data, &allDatas)
	ErrorPage(err1, 500)
}

func WriteCachedData(data interface{}) error {
	file, err := os.Create(cacheFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func ReadCachedData(data interface{}) error {
	file, err := os.Open(cacheFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}

func Parse() {

	var wg sync.WaitGroup
	ch := make(chan string)

	for endpoint, data := range endpoints {
		wg.Add(1)
		go func(endpoint string, data interface{}) {
			defer wg.Done()

			err := ReadCachedData(data)
			if err != nil {
				ReadAPI(BaseURL+endpoint, data)

				err := WriteCachedData(data)
				ErrorPage(err, 500)
			}
			ch <- endpoint
		}(endpoint, data)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for range endpoints {
		<-ch
	}
}
