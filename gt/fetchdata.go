package gt

func FetchArtists() []Artist {
	allArtists := []Artist{}
	ReadAPI(BaseURL+"artists", &allArtists)
	return allArtists
}

func FetchLocations() []Location {
	allLocations := []Location{}
	ReadAPI(BaseURL+"locations", &allLocations)
	return allLocations
}

func FetchDates() []Date {
	allDates := []Date{}
	ReadAPI(BaseURL+"dates", &allDates)
	return allDates
}

func FetchRelations() []Relation {
	allRelations := []Relation{}
	ReadAPI(BaseURL+"relations", &allRelations)
	return allRelations
}
