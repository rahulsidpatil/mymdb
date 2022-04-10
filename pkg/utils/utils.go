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
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "823ef2af"
	}
	MdbApi = gomdb.Init(apiKey)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, code, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(code)
	w.Write(response)
}

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
