package data

type Request string

const (
	CLUB_NAME Request = "club_name"
)

// ------------------------ Hall Plan ------------------------

type ResponseHallPlan struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Result  HallPlanResult `json:"result"`
	Status  bool           `json:"status"`
}

type HallPlanResult struct {
	SeanseId       string   `json:"seance_id"`
	SeanseDate     string   `json:"seance_date"`
	SeanseTime     string   `json:"seance_time"`
	Duration       int64    `json:"duration"`
	TheatreId      int64    `json:"theatre_id"`
	TheatreName    string   `json:"theatre_name"`
	MovieId        string   `json:"movie_id"`
	MovieName      string   `json:"movie_name"`
	Genres         string   `json:"genres"`
	AgeRestriction string   `json:"age_restriction"`
	Format         string   `json:"format"`
	Language       string   `json:"language"`
	HallName       string   `json:"hall_name"`
	HallPlan       HallPlan `json:"hall_plan"`
}

type HallPlan struct {
	Capacity int64 `json:"capacity"`
	Width    int64 `json:"width"`
	Height   int64 `json:"height"`
	Places   []Place
}

type Place struct {
	Id          string `json:"id"`
	Row         string `json:"row"`
	Place       string `json:"place"`
	X           int64  `json:"x"`
	Y           int64  `json:"y"`
	Width       int64  `json:"width"`
	Height      int64  `json:"height"`
	PlaceTypeId string `json:"place_type_id"`
	PlaceColor  string `json:"place_color"`
	Status      int64  `json:"status"`
}
