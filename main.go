package main

import (
	"fmt"
	"gt/gt"
)

func main() {
	url := "https://groupietrackers.herokuapp.com/api/"
	allArtists := &[]gt.Artist{}
	gt.ReadAPI(url+"artists", allArtists)
	fmt.Println(allArtists)
}
