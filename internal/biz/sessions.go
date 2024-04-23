package biz

import (
	"encoding/json"
	"fmt"
	"kinokz/internal/data"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type SessionsUsecase struct {
	log          *log.Helper
	usersRepo    data.UsersRepo
	hallPlanRepo data.HallPlanRepo
}

func NewSessionsUsecase(
	logger log.Logger,
	usersRepo data.UsersRepo,
	hallPlanRepo data.HallPlanRepo,
) *SessionsUsecase {
	return &SessionsUsecase{
		log:          log.NewHelper(log.With(logger, "module", "biz/sessions")),
		usersRepo:    usersRepo,
		hallPlanRepo: hallPlanRepo,
	}
}

func (uc *SessionsUsecase) GetSessions(isDetailed bool, filmId, cityId int64, date time.Time) ([]Sessions, error) {
	sessionDate := fmt.Sprintf("%v-%02d-%02d", date.Year(), int(date.Month()), date.Day())
	sessionsUrl := fmt.Sprintf("https://api.kino.kz/sessions/v1/movie/sessions?movie_id=%v&date=%v&city_id=%v", filmId, sessionDate, cityId)

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

	response := &ResponseSessions{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		panic(derr)
		return nil, derr
	}

	if isDetailed {
		for i, s := range response.Result.Sessions {
			hallPlan, err := uc.hallPlanRepo.GetHallPlan(nil, s.Session.SessionId, s.Cinema.Id, cityId)
			if err != nil {
				continue
			}

			response.Result.Sessions[i].Language = hallPlan.Language
			response.Result.Sessions[i].Capacity = hallPlan.HallPlan.Capacity

			freeSeatCount := 0
			for _, place := range hallPlan.HallPlan.Places {
				if place.Status == 1 {
					freeSeatCount++
				}
			}
			response.Result.Sessions[i].FreeSeats = int64(freeSeatCount)
		}
	}

	return response.Result.Sessions, nil
}
