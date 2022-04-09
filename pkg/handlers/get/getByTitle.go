package get

import (
	"fmt"
	"net/http"
	"os"

	"github.com/eefret/gomdb"
	"github.com/gorilla/mux"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
)

var api *gomdb.OmdbApi

func init() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "823ef2af"
	}
	api = gomdb.Init(apiKey)
}

// @Summary Fetch movie by ID
// @Description Fetch movie by ID
// @Accept  json
// @Produce  json
// @Param title path string true "Movie title"
// @Success 200 {object} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/{title} [get]
func GetByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title, ok := vars["title"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Title!!")
		return
	}
	query := &gomdb.QueryData{Title: title}
	result, err := api.MovieByTitle(query)
	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusNoContent, "")
	}

	m, err := utils.ParseResults(result)
	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, "")
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}
