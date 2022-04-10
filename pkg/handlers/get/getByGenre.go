package get

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rahulsidpatil/mymdb/pkg/dal"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
)

// @Summary Fetch movies by Genre
// @Description Fetch movies by Genre
// @Produce  json
// @Param Genre path string true "Movie Genre"
// @Success 200 {object} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/Genre/{Genre} [get]
func GetByGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genre, ok := vars["Genre"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid movie id")
		return
	}
	movieGenres := strings.Split(genre, ",")
	db := dal.GetMySQLDriver()
	logger.Sugar().Infof("Got this DB driver:%#v", db)
	m := db.GetMoviesByGenre(movieGenres)
	if len(m) == 0 {
		utils.RespondWithError(w, http.StatusNoContent, "No movies!!")
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}
