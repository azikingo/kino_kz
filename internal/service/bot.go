package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	v1 "kinokz/api/helloworld/v1"
	"kinokz/internal/biz"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BotService struct {
	log *log.Helper

	v1.UnimplementedGreeterServer
	userBiz     *biz.UsersUsecase
	filmsBiz    *biz.FilmsUsecase
	sessionsBiz *biz.SessionsUsecase
}

// NewBotService new a greeter service.
func NewBotService(
	userBiz *biz.UsersUsecase,
	filmsBiz *biz.FilmsUsecase,
	sessionsBiz *biz.SessionsUsecase,
	logger log.Logger,
) *BotService {
	return &BotService{
		userBiz:     userBiz,
		filmsBiz:    filmsBiz,
		sessionsBiz: sessionsBiz,
		log:         log.NewHelper(log.With(logger, "module", "service/bot")),
	}
}

func (bs *BotService) BotRequest(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {
	message := update.Message
	reply := tgbotapi.Chattable(nil)

	// Check if we've gotten a message update.
	if message != nil {
		chatId := update.Message.Chat.ID
		tgUser := update.Message.From

		if message.Command() != "" {
			switch message.Command() {
			case "start":
				bs.RegisterUser(ctx, tgUser)
				reply = tgbotapi.NewMessage(chatId, "Welcome to Kino.kz bot! You can find there best films in Kino.kz also their schedule and availabilities.")

			case "films":
				films, err := bs.filmsBiz.GetFilms(2)
				if err != nil {
					bs.log.Errorf("GetFilms error: %v", err)
					return nil, err
				}

				reply = tgbotapi.NewMessage(update.Message.Chat.ID, bs.GetFilmsList(films))

			case "today":
				text := strings.ReplaceAll(message.Text, "/today ", "")

				msg, err := bs.GetSessions(time.Now(), text)
				if err != nil {
					bs.log.Errorf("GetSessions error: %v", err)
					return nil, err
				}

				reply = tgbotapi.NewMessage(update.Message.Chat.ID, *msg)

			case "tomorrow":
				text := strings.ReplaceAll(message.Text, "/today ", "")

				msg, err := bs.GetSessions(time.Now().Add(24*time.Hour), text)
				if err != nil {
					bs.log.Errorf("GetSessions error: %v", err)
					return nil, err
				}

				reply = tgbotapi.NewMessage(update.Message.Chat.ID, *msg)

			case "seats":
				text := strings.ReplaceAll(message.Text, "/seats ", "")

				msg, err := bs.GetSeats(text)
				if err != nil {
					bs.log.Errorf("GetSeats error: %v", err)
					return nil, err
				}

				reply = tgbotapi.NewMessage(update.Message.Chat.ID, *msg)

			case "help":
				reply = tgbotapi.NewMessage(chatId, "I did a list to you here for commands that I can do:\n"+
					"/start - Start the bot\n"+
					"/films - List films by rating\n"+
					"/help - List bot commands\n")

			default:
				reply = tgbotapi.NewMessage(chatId, "Unknown command, use one of the commands in the menu.")
			}

			return reply, nil
		} else if message.Text != "" {
			reply = tgbotapi.NewMessage(chatId, "Welcome to the Kino.kz bot! Here you can list /films and see more information about session.")
			return reply, nil
		}
	}

	return reply, nil
}

func (bs *BotService) RegisterUser(ctx context.Context, tgUser *tgbotapi.User) {
	_, err := bs.userBiz.CreateUser(ctx, tgUser)
	if err != nil {
		bs.log.Errorf("CreateUser error: %v", err)
		return
	}
}

// TODO: finish function about seats
func (bs *BotService) GetSeats(text string) (*string, error) {
	return nil, nil
}

func (bs *BotService) GetSessions(date time.Time, text string) (*string, error) {
	onlyImax := strings.Contains(strings.ToLower(text), "imax")
	isDetailed := strings.Contains(strings.ToLower(text), "detailed")
	filmId := int64(0)

	for _, t := range strings.Split(text, " ") {
		id, err := strconv.ParseInt(t, 10, 64)
		if err == nil {
			filmId = id
			break
		}
	}
	if filmId == 0 {
		bs.log.Errorf("ParseInt error: filmId=%v", filmId)
		return nil, fmt.Errorf("filmId not found")
	}

	film, err := bs.filmsBiz.GetFilm(filmId)
	if err != nil {
		bs.log.Errorf("GetFilm error: %v", err)
		return nil, err
	}

	sessions, err := bs.sessionsBiz.GetSessions(isDetailed, filmId, 2, date)
	if err != nil {
		bs.log.Errorf("GetSessions error: %v", err)
		return nil, err
	}

	reply := bs.GetSessionsList(onlyImax, isDetailed, *film, sessions)
	return &reply, nil
}

func (bs *BotService) GetFilmsList(films []biz.Film) string {
	reply := "Today's films by rating:\n"

	sort.Slice(films, func(i, j int) bool {
		return films[i].Rating > films[j].Rating
	})

	for i, film := range films {
		name := film.Name
		if name == "" {
			name = film.NameRus
		}
		if name == "" {
			name = film.NameOrigin
		}

		reply += fmt.Sprintf("%v) [%v] %s - %v\n", i+1, film.Id, name, film.Rating)
	}

	return reply
}

func (bs *BotService) GetSessionsList(onlyImax, isDetailed bool, film biz.Film, sessions []biz.Sessions) string {
	hour := film.Duration / 60
	minute := film.Duration % 60
	reply := fmt.Sprintf("Today's sessions of film \"%v\" (%vh %vm):\n\n", film.Name, hour, minute)

	mapSessions := make(map[string]string)
	for _, session := range sessions {
		imax := ""
		if session.Hall.Imax {
			imax = "IMAX"
		}
		if !onlyImax || session.Hall.Imax {
			if isDetailed {
				mapSessions[session.Cinema.Name] += fmt.Sprintf("(%v:%v) [%v] %v %v - (%v) places: [%v/%v] prices: [%v; %v; %v]\n", session.Session.Hour, session.Session.Minutes, session.Session.SessionId, session.Language, imax, session.Hall.Name, session.FreeSeats, session.Capacity, session.Session.Adult, session.Session.Student, session.Session.Child)
			}
			mapSessions[session.Cinema.Name] += fmt.Sprintf("(%v:%v) [%v] %v %v prices: [%v; %v; %v]\n", session.Session.Hour, session.Session.Minutes, session.Session.SessionId, session.Language, imax, session.Session.Adult, session.Session.Student, session.Session.Child)
		}
	}

	for cinema, session := range mapSessions {
		reply += fmt.Sprintf("[%v]\n%v\n", cinema, session)
	}

	return reply
}
