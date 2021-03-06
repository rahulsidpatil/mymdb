package get

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rahulsidpatil/mymdb/pkg/dal"
	"github.com/rahulsidpatil/mymdb/pkg/logging"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
)

var (
	logger = logging.NewLogger()
)

// @Summary Fetch movie by ID
// @Description Fetch movie by ID
// @Produce  json
// @Param Id path string true "Movie Id"
// @Success 200 {object} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/Id/{Id} [get]
func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId, ok := vars["Id"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid movie id")
		return
	}
	db := dal.GetMySQLDriver()
	logger.Sugar().Infof("Got this DB driver:%#v", db)
	m, err := db.GetMovieById(movieId)
	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, err.Error())
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}
