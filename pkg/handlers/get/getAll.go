package get

import (
	"net/http"

	"github.com/rahulsidpatil/mymdb/pkg/dal"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
)

// @Summary Get all movies
// @Description Get all Movies
// @Success 200 {array} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies [get]
func GetAll(w http.ResponseWriter, r *http.Request) {
	db := dal.GetMySQLDriver()
	logger.Sugar().Infof("Got this DB driver:%#v", db)
	movies, err := db.GetAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, err.Error())
	}
	utils.RespondWithJSON(w, http.StatusOK, movies)
}
