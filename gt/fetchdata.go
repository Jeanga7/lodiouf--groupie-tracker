package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	fmt.Println("II AAAM READDDDDIIIIIIIIIIINNNNNNNNNG  AAAAPPPPPIIII")
	r, err := http.Get(url)
	ErrorPage(err, 500)
	defer r.Body.Close()

	data, err1 := io.ReadAll(r.Body)
	ErrorPage(err1, 500)

	err1 = json.Unmarshal(data, &allDatas)
	ErrorPage(err1, 500)

	err2 := WriteCachedData(allDatas)
	ErrorPage(err2, 500)
}

func WriteCachedData(data interface{}) error {
	file, err := os.OpenFile(cacheFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	if err := decoder.Decode(data); err != nil {
		fmt.Println("Error decoding data:", err)
		return err
	}
	fmt.Println("Decoding successfully")
	return nil
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
				fmt.Println("reeeeeeeeeeddddddddddddAAAAAAAAAAAAPPPPPPPPIIIIIIIIIIII")
				ReadAPI(BaseURL+endpoint, data)
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

func RedableRelations(relation Relation) Relation {
	var newRelation Relation

	for _, val := range relation.Index {
		rel := struct {
			Id             int                 `json:"id"`
			DatesLocations map[string][]string `json:"datesLocations"`
		}{
			val.Id,
			TransformMap(val.DatesLocations),
		}
		newRelation.Index = append(newRelation.Index, rel)
	}
	return newRelation
}

func TransformMap(m map[string][]string) map[string][]string {
	newMap := make(map[string][]string)

	for i, val := range m {
		newMap[VillePays(i)] = FormatDates(val)

	}

	return newMap
}

func VillePays(s string) string {
	lieu := strings.Split(s, "-")
	ville, pays := strings.Title(RemoveUnderscores(lieu[0])), strings.ToUpper(RemoveUnderscores(lieu[1]))
	return ville + " (" + pays + ")"
}

func RemoveUnderscores(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

func FormatDates(date []string) []string {
	var newDate []string
	for _, val := range date {
		newDate = append(newDate, strings.ReplaceAll(strings.Trim(val, "*"), "-", "/"))
	}

	return newDate
}
