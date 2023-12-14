package main

import (
	"fmt"
	"gt/gt"
	"log"
	"net/http"
	"time"
)

func main() {
	tStart := time.Now()
	port := ":8080"
	http.HandleFunc("/", gt.HandlePage)

	// Create MUX Routing
	fs := http.FileServer(http.Dir("./vues/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Server Online...")
	fmt.Println("Access to the app using this link : http://localhost" + port)

	elapsed := time.Since(tStart)
	log.Printf("Parsing took %.8fs\n", elapsed.Seconds())
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
