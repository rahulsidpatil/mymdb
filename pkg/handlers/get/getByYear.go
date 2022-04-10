package get

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rahulsidpatil/mymdb/pkg/dal"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
)

// @Summary Fetch movies by year
// @Description Fetch movies by year
// @Produce  json
// @Param Year path int true "Movie Year"
// @Success 200 {object} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/Year/{Year} [get]
func GetByYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, ok := vars["Year"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid year value!!")
		return
	}
	movieYear, err := strconv.Atoi(year)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid year value!!")
		return
	}
	db := dal.GetMySQLDriver()
	logger.Sugar().Infof("Got this DB driver:%#v", db)
	m, err := db.GetMoviesByYear(movieYear)
	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, err.Error())
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}

// @Summary Fetch movies by year
// @Description Fetch movies by year
// @Produce  json
// @Param start path int true "start year"
// @Param end path int true "end year"
// @Success 200 {object} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/Year/start/{start}/end/{end} [get]
func GetByYearRange(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	start, ok := vars["start"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Absent param start!!")
		return
	}
	startYear, err := strconv.Atoi(start)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid startYear value!!")
		return
	}

	end, ok := vars["end"]
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Absent param end!!")
		return
	}
	endYear, err := strconv.Atoi(end)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid startYear value!!")
		return
	}

	db := dal.GetMySQLDriver()
	logger.Sugar().Infof("Got this DB driver:%#v", db)
	m, err := db.GetMoviesByYearRange(startYear, endYear)
	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, err.Error())
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}
