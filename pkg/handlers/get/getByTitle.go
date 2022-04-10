package get

import (
	"fmt"
	"net/http"

	"github.com/eefret/gomdb"
	"github.com/gorilla/mux"
	"github.com/rahulsidpatil/mymdb/pkg/dal"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
	"go.uber.org/zap"
)

// @Summary Fetch movie by ID
// @Description Fetch movie by ID
// @Produce  json
// @Param Title path string true "Movie Title"
// @Success 200 {object} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/Title/{Title} [get]
func GetByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title, ok := vars["Title"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid movie Title")
		return
	}
	db := dal.GetMySQLDriver()
	logger.Sugar().Infof("Got this DB driver:%#v", db)
	m, err := db.GetMoviesByTitle(title)
	if err != nil {
		logger.Sugar().Error("Error fetching movie by title:", zap.Error(err))
		query := &gomdb.QueryData{Title: title}
		result, err := utils.MdbApi.MovieByTitle(query)
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
	utils.RespondWithJSON(w, http.StatusOK, m)
}
