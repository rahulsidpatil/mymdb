package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/eefret/gomdb"
	"github.com/rahulsidpatil/mymdb/pkg/entities"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
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
