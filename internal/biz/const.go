package biz

// ------------------------ Films ------------------------

type ResponseFilm struct {
	Message string `json:"message"`
	Result  Film   `json:"result"`
	Status  bool   `json:"status"`
}

type ResponseFilms struct {
	IsCache bool   `json:"is_cache"`
	Message string `json:"message"`
	Result  []Film `json:"result"`
	Status  bool   `json:"status"`
}

type Film struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	NameRus         string  `json:"name_rus"`
	NameOrigin      string  `json:"name_origin"`
	Duration        int64   `json:"duration"`
	Presentation    string  `json:"presentation"`
	RatingState     bool    `json:"rating_state"`
	AgeRestriction  int64   `json:"age_restriction"`
	Genre           string  `json:"genre"`
	Genres          []Genre `json:"genres"`
	PremiereKaz     string  `json:"premiere_kaz"`
	SmallPoster     string  `json:"small_poster"`
	Priority        int64   `json:"priority"`
	SessionId       int64   `json:"session_id"`
	KinopoiskId     string  `json:"kinopoisk_id"`
	KinopoiskRating string  `json:"kinopoisk_rating"`
	KinopoiskVotes  int64   `json:"kinopoisk_votes"`
	ImdbRating      string  `json:"imdb_rating"`
	ImdbVotes       int64   `json:"imdb_votes"`
	Rating          float64 `json:"rating"`
	Votes           int64   `json:"votes"`
}

type Genre struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	TitleKz  string `json:"title_kz"`
	TitleEng string `json:"title_eng"`
}

// ------------------------ Sessions ------------------------

type ResponseSessions struct {
	IsCache bool          `json:"is_cache"`
	Message string        `json:"message"`
	Result  SessionResult `json:"result"`
	Status  bool          `json:"status"`
}

type SessionResult struct {
	Sessions       []Sessions `json:"sessions"`
	AvailableDates []string   `json:"available_dates"`
}

type Sessions struct {
	Language  string `json:"language"`
	FreeSeats int64  `json:"free_seats"`
	Capacity  int64  `json:"capacity"`

	Session Session `json:"session"`
	Cinema  Cinema  `json:"cinema"`
	Hall    Hall    `json:"hall"`
}

type Session struct {
	Id             int64  `json:"id"`
	SessionId      int64  `json:"session_id"`
	CinemaId       int64  `json:"cinema_id"`
	MovieId        int64  `json:"movie_id"`
	HallId         int64  `json:"hall_id"`
	SessionType    int64  `json:"session_type"`
	LangId         int64  `json:"lang_id"`
	LangLabel      string `json:"lang_label"`
	Is3D           bool   `json:"is_3_d"`
	IsAp           bool   `json:"is_ap"`
	Adult          int64  `json:"adult"`
	Child          int64  `json:"child"`
	Student        int64  `json:"student"`
	Vip            int64  `json:"vip"`
	SessionDate    string `json:"session_date"`
	SessionDateTz  string `json:"session_date_tz"`
	IsAtmos        bool   `json:"is_atmos"`
	IsReald        bool   `json:"is_reald"`
	Hour           string `json:"hour"`
	Minutes        string `json:"minutes"`
	IsAtmos3D      bool   `json:"is_atmos3d"`
	IsImax         bool   `json:"is_imax"`
	IsFdx          bool   `json:"is_fdx"`
	ProviderLink   string `json:"provider_link"`
	SubSessionType int64  `json:"sub_session_type"`
}

type Cinema struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Phone       string  `json:"phone"`
	Order       int64   `json:"order"`
	SmallPoster string  `json:"small_poster"`
	BigPoster   string  `json:"big_poster"`
}

type Hall struct {
	Id         int64  `json:"id"`
	CinemaId   int64  `json:"cinema_id"`
	Name       string `json:"name"`
	NameParser string `json:"name_parser"`
	Color      string `json:"color"`
	Imax       bool   `json:"imax"`
	Fdx        bool   `json:"fdx"`
	Laser      bool   `json:"laser"`
	State      bool   `json:"state"`
	HallOrder  int64  `json:"hall_order"`
	ProviderId int64  `json:"provider_id"`
}

//// ------------------------ Hall Plan ------------------------
//
//type ResponseHallPlan struct {
//	Code    int64          `json:"code"`
//	Message string         `json:"message"`
//	Result  HallPlanResult `json:"result"`
//	Status  bool           `json:"status"`
//}
//
//type HallPlanResult struct {
//	SeanseId       string   `json:"seance_id"`
//	SeanseDate     string   `json:"seance_date"`
//	SeanseTime     string   `json:"seance_time"`
//	Duration       int64    `json:"duration"`
//	TheatreId      int64    `json:"theatre_id"`
//	TheatreName    string   `json:"theatre_name"`
//	MovieId        string   `json:"movie_id"`
//	MovieName      string   `json:"movie_name"`
//	Genres         string   `json:"genres"`
//	AgeRestriction string   `json:"age_restriction"`
//	Format         string   `json:"format"`
//	Language       string   `json:"language"`
//	HallName       string   `json:"hall_name"`
//	HallPlan       HallPlan `json:"hall_plan"`
//}
//
//type HallPlan struct {
//	Capacity int64 `json:"capacity"`
//	Width    int64 `json:"width"`
//	Height   int64 `json:"height"`
//	Places   []Place
//}
//
//type Place struct {
//	Id          string `json:"id"`
//	Row         string `json:"row"`
//	Place       string `json:"place"`
//	X           int64  `json:"x"`
//	Y           int64  `json:"y"`
//	Width       int64  `json:"width"`
//	Height      int64  `json:"height"`
//	PlaceTypeId string `json:"place_type_id"`
//	PlaceColor  string `json:"place_color"`
//	Status      int64  `json:"status"`
//}
