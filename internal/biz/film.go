package biz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kinokz/internal/data"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type FilmsUsecase struct {
	log       *log.Helper
	usersRepo data.UsersRepo
}

func NewFilmsUsecase(
	logger log.Logger,
	usersRepo data.UsersRepo,
) *FilmsUsecase {
	return &FilmsUsecase{
		log:       log.NewHelper(log.With(logger, "module", "biz/film")),
		usersRepo: usersRepo,
	}
}

func (uc *FilmsUsecase) GetFilm(filmId int64) (*Film, error) {
	sessionsUrl := fmt.Sprintf("https://api.kino.kz/catalog/v1/movie?movie_id=%v", filmId)

	res, err := http.Get(sessionsUrl)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
		return nil, fmt.Errorf("status code: %v", res.Status)
	}

	response := &ResponseFilm{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		panic(derr)
		return nil, derr
	}

	return &response.Result, nil
}

func (uc *FilmsUsecase) GetFilms(cityId int64) ([]Film, error) {
	// url of the API
	postUrl := "https://api.kino.kz/sessions/v1/movies/find"

	// JSON body
	now := time.Now()
	startDate := fmt.Sprintf("%v-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	now = now.AddDate(0, 0, 7)
	endDate := fmt.Sprintf("%v-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	body := []byte(`{
				  "city_id": ` + strconv.FormatInt(cityId, 10) + `,
				  "start_date": "` + startDate + `",
				  "end_date": "` + endDate + `"
				}`)

	res, err := http.Post(postUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
		return nil, fmt.Errorf("status code: %v", res.Status)
	}

	response := &ResponseFilms{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		panic(derr)
		return nil, derr
	}

	return response.Result, nil
}
