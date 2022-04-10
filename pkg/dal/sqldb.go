package dal

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rahulsidpatil/mymdb/pkg/entities"
	"github.com/rahulsidpatil/mymdb/pkg/logging"
	"go.uber.org/zap"
)

var (
	dbdriver = "mysql"
	dbname   = "mDB"
	logger   = logging.NewLogger()
)

// MySQLDriver ... mysql db driver
type MySQLDriver struct {
	driver *sql.DB
}

// TODO: Do away with mysql raw queries

//GetMySQLDriver ... get a mysql DB driver
func GetMySQLDriver() *MySQLDriver {
	if os.Getenv("DB_DRIVER") != "" {
		dbdriver = os.Getenv("DB_DRIVER")
	}
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpasswd := os.Getenv("DB_PASSWD")
	if os.Getenv("DB_NAME") != "" {
		dbname = os.Getenv("DB_NAME")
	}
	dbaccessurl := dbuser + ":" + dbpasswd + "@" + "tcp(" + dbhost + ":" + dbport + ")" + "/" + dbname

	logger.Sugar().Infof("dbaccessurl:", dbaccessurl)
	var db MySQLDriver
	var err error
	db.driver, err = sql.Open(dbdriver, dbaccessurl)
	if err != nil {
		logger.Sugar().Errorf("Error connecting to database: ", err)
	}
	logger.Sugar().Infof("db connection successful!!", dbaccessurl)
	return &db
}

// GetMoviesWithMinRating ...
func (db *MySQLDriver) GetMoviesWithMinRating(minRating float32) ([]entities.Movie, error) {
	statement := fmt.Sprintf("SELECT * FROM mDB.movies WHERE Rated>=%f", minRating)

	rows, err := db.driver.Query(statement)
	if err != nil {
		logger.Sugar().Errorf("Error executiong query!!", err)
		return nil, err
	}
	defer rows.Close()
	return rowMapper(rows)
}

// GetMoviesWithMaxRating ...
func (db *MySQLDriver) GetMoviesWithMaxRating(maxRated float32) ([]entities.Movie, error) {
	statement := fmt.Sprintf("SELECT * FROM mDB.movies WHERE Rated<=%f", maxRated)
	rows, err := db.driver.Query(statement)
	if err != nil {
		logger.Sugar().Errorf("Error executiong query!!", err)
		return nil, err
	}
	defer rows.Close()
	return rowMapper(rows)
}

// GetMovieById ... get movie by ID
func (db *MySQLDriver) GetMovieById(Id string) (entities.Movie, error) {
	m := entities.Movie{}
	statement := fmt.Sprintf("SELECT * FROM mDB.movies WHERE Id=\"%s\"", Id)
	var genre string
	err := db.driver.QueryRow(statement).Scan(&m.Id, &m.Title, &m.Year, &m.Rated, &genre)
	if err != nil {
		return m, err
	}
	m.Genre = strings.Split(genre, ",")
	return m, nil
}

// GetMoviesByTitle ...
func (db *MySQLDriver) GetMoviesByTitle(Title string) (entities.Movie, error) {
	m := entities.Movie{}
	movieTitle := `%` + Title + `%`
	statement := fmt.Sprintf("SELECT * FROM mDB.movies WHERE Title like '%s'", movieTitle)
	var genre string
	err := db.driver.QueryRow(statement).Scan(&m.Id, &m.Title, &m.Year, &m.Rated, &genre)
	if err != nil {
		return m, err
	}
	m.Genre = strings.Split(genre, ",")
	return m, nil
}

// GetMoviesByYear ...
func (db *MySQLDriver) GetMoviesByYear(year int) ([]entities.Movie, error) {
	statement := fmt.Sprintf("SELECT * FROM mDB.movies WHERE Year=%d", year)
	logger.Sugar().Infof("Query statement:", statement)
	rows, err := db.driver.Query(statement)
	if err != nil {
		logger.Sugar().Errorf("Error executiong query!!", err)
		return nil, err
	}
	defer rows.Close()
	return rowMapper(rows)
}

// GetMoviesByYearRange ...
func (db *MySQLDriver) GetMoviesByYearRange(start, end int) ([]entities.Movie, error) {
	statement := fmt.Sprintf("SELECT * FROM mDB.movies WHERE Year>=%d AND Year<=%d", start, end)
	rows, err := db.driver.Query(statement)
	if err != nil {
		logger.Sugar().Errorf("Error executiong query!!", err)
		return nil, err
	}
	logger.Sugar().Infof("Rows fetched:", rows)
	defer rows.Close()
	return rowMapper(rows)
}

// AddMovie ...
func (db *MySQLDriver) AddMovie(m *entities.Movie) {
	statement := fmt.Sprintf("INSERT INTO mDB.movies(Id, Title, Year, Rated, Genre) VALUES ('%s','%s','%d','%f','%s');",
		m.Id, m.Title, m.Year, m.Rated, strings.Join(m.Genre, ","))
	_, err := db.driver.Exec(statement)
	if err != nil {
		logger.Sugar().Error("Error executing insert statement", zap.Error(err))
	}
	err = db.driver.QueryRow("SELECT LAST_INSERT_ID()").Scan(&m.Id)
	if err != nil {
		logger.Sugar().Error("Error finding last insert id", zap.Error(err))
	}
}

//GetAll ... get all movies
func (db *MySQLDriver) GetAll() ([]entities.Movie, error) {
	statement := "SELECT * FROM mDB.movies"
	rows, err := db.driver.Query(statement)
	if err != nil {
		logger.Sugar().Errorf("Error executiong query!!", err)
		return nil, err
	}

	defer rows.Close()
	return rowMapper(rows)
}

// GetMoviesByGenre ...
func (db *MySQLDriver) GetMoviesByGenre(genre []string) []entities.Movie {
	movies := make([]entities.Movie, 0)
	for _, g := range genre {
		movieGenre := `%` + g + `%`
		statement := fmt.Sprintf("SELECT * FROM mDB.movies WHERE Genre like '%s'", movieGenre)
		rows, err := db.driver.Query(statement)
		if err != nil {
			logger.Sugar().Errorf("Error executiong query!!", err)
			continue
		}
		defer rows.Close()
		m, err := rowMapper(rows)
		if err != nil {
			continue
		}
		if len(m) > 0 {
			movies = append(movies, m...)
		}
	}
	return movies
}

func rowMapper(rows *sql.Rows) ([]entities.Movie, error) {
	movies := []entities.Movie{}
	for rows.Next() {
		var m entities.Movie
		var genre string
		if err := rows.Scan(&m.Id, &m.Title, &m.Year, &m.Rated, &genre); err != nil {
			logger.Sugar().Errorf("Error scanning row:", err)
			return nil, err
		}
		m.Genre = strings.Split(genre, ",")
		movies = append(movies, m)
	}
	logger.Sugar().Infof("Sending following movies to user:", movies)
	return movies, nil
}
