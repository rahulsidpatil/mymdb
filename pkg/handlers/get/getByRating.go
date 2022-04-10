package get

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rahulsidpatil/mymdb/pkg/dal"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
)

// @Summary Get all movie with min rating r
// @Description Get all Movies with min rating r
// @Accept  json
// @Produce  json
// @Param minRating path string true "Movie rating"
// @Success 200 {array} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/minRating/{minRating} [get]
func GetMovieWithMinRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rating, ok := vars["minRating"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Param minRating absent")
		return
	}
	minRating, err := strconv.ParseFloat(rating, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid movie rating")
		return
	}
	db := dal.GetMySQLDriver()
	logger.Sugar().Infof("Got this DB driver:%#v", db)
	m, err := db.GetMoviesWithMinRating(float32(minRating))
	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, err.Error())
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}

// @Summary Get all movie with max rating r
// @Description Get all Movies with max rating r
// @Produce  json
// @Param maxRating path string true "Movie rating"
// @Success 200 {array} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/maxRating/{maxRating} [get]
func GetMovieWithMaxRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rating, ok := vars["maxRating"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Param maxRating absent")
		return
	}
	maxRating, err := strconv.ParseFloat(rating, 32)
	if err != nil {

		utils.RespondWithError(w, http.StatusBadRequest, "Invalid movie rating")
		return
	}
	db := dal.GetMySQLDriver()
	logger.Sugar().Infof("Got this DB driver:%#v", db)
	m, err := db.GetMoviesWithMaxRating(float32(maxRating))
	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, err.Error())
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}
