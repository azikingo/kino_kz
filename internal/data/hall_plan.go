package data

import (
	"context"
	"encoding/json"
	"fmt"
	"kinokz/ent"
	"net/http"
)

type HallPlanRepo interface {
	GetHallPlan(ctx context.Context, sessionId, cinemaId, cityId int64) (*HallPlanResult, error)
}

type hallPlanRepo struct {
	db *ent.Client
}

func NewHallPlanRepo(d *Data) HallPlanRepo {
	return &hallPlanRepo{
		db: d.db,
	}
}

func (r *hallPlanRepo) GetHallPlan(ctx context.Context, sessionId, cinemaId, cityId int64) (*HallPlanResult, error) {
	sessionsUrl := fmt.Sprintf("https://api.kino.kz/mediator/seances/hall-plan-prices?seance_id=%v&cinema_id=%v&city_id=%v", sessionId, cinemaId, cityId)

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

	response := &ResponseHallPlan{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		panic(derr)
		return nil, derr
	}

	return &response.Result, nil
}
