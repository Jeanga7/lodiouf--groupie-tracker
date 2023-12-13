package gt

import "sync"

var endpoints = map[string]interface{}{
	"artists":   &[]Artist{},
	"locations": &Location{},
	"dates":     &Date{},
	"relation":  &Relation{},
}

func Parse() {
	var wg sync.WaitGroup
	ch := make(chan string)

	for endpoint, data := range endpoints {
		wg.Add(1)
		go func(endpoint string, data interface{}) {
			defer wg.Done()
			ReadAPI(BaseURL+endpoint, data)
			ch <- endpoint
		}(endpoint, data)
	}

	// Attend que toutes les goroutines soient terminées
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Utilisez le canal pour synchroniser la fin de l'exécution de toutes les goroutines
	for range endpoints {
		<-ch
	}
}
