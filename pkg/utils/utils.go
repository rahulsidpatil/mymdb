package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/eefret/gomdb"
	"github.com/rahulsidpatil/mymdb/pkg/entities"
)

var MdbApi *gomdb.OmdbApi

func init() {
	// This apiKey is needed to access the services of github.com/eefret/gomdb
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "823ef2af"
	}
	MdbApi = gomdb.Init(apiKey)
}

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON ...
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, code, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// ParseResults ... parse results received from github.com/eefret/gomdb
func ParseResults(result *gomdb.MovieResult) (*entities.Movie, error) {
	year, _ := strconv.Atoi(result.Year)
	rating, _ := strconv.ParseFloat(result.ImdbRating, 32)

	return &entities.Movie{
		Id:    result.ImdbID,
		Title: result.Title,
		Year:  year,
		Genre: strings.Split(result.Genre, ","),
		Rated: float32(rating),
	}, nil
}

// PopulateDB ... this function will be used to populate local movie database at runtime
// This is still a WIP
func PopulateDB() {
	query := &gomdb.QueryData{Title: "Avengers", SearchType: gomdb.MovieSearch}
	searchRes, err := MdbApi.Search(query)
	if err != nil {
		fmt.Println(err)
	}
	titles := make([]string, 0)
	for _, res := range searchRes.Search {
		titles = append(titles, res.Title)
	}
	for _, t := range titles {
		query := &gomdb.QueryData{Title: t}
		result, err := MdbApi.MovieByTitle(query)
		if err == nil {
			m, err := ParseResults(result)
			if err == nil {
				fmt.Println(m)
				//TODO: Add code for db inserts
			}
		}
	}
}
