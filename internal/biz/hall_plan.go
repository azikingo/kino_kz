package biz

import (
	"context"
	"kinokz/internal/data"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type HallPlanUsecase struct {
	log          *log.Helper
	hallPlanRepo data.HallPlanRepo
}

func NewHallPlansUsecase(
	logger log.Logger,
	hallPlanRepo data.HallPlanRepo,
) *HallPlanUsecase {
	return &HallPlanUsecase{
		log:          log.NewHelper(log.With(logger, "module", "biz/hall_plan")),
		hallPlanRepo: hallPlanRepo,
	}
}

func (uc *HallPlanUsecase) GetSeats(ctx context.Context, sessionId, filmId, cityId int64, date time.Time) (*string, error) {
	reply := ""

	hallPlan, err := uc.hallPlanRepo.GetHallPlan(ctx, sessionId, filmId, cityId)
	if err != nil {
		return nil, err
	}

	row := int64(1)
	for _, place := range hallPlan.HallPlan.Places {
		if strconv.FormatInt(row, 10) != place.Row {
			reply += "\n"
		}
		if place.Status == 1 {
			reply += place.Place + " "
		} else {
			reply += "x "
		}
	}

	return &reply, nil
}
