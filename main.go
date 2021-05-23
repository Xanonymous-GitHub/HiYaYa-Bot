package main

import (
	"fmt"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
	"os"
)

type HiYaYaBot struct {
	bot            *linebot.Client
	defaultMessage *linebot.TextMessage
}

func NewHiYaYaBot(channelSecret, channelToken, defaultMessage string) (*HiYaYaBot, error) {
	bot, err := linebot.New(
		channelSecret,
		channelToken,
	)
	if err != nil {
		return nil, err
	}

	return &HiYaYaBot{
		bot:            bot,
		defaultMessage: linebot.NewTextMessage(defaultMessage),
	}, nil
}

func (app *HiYaYaBot) ShowDefaultMsg(event *linebot.Event) error {
	task := app.bot.ReplyMessage(event.ReplyToken, app.defaultMessage)
	_, err := task.Do()
	if err != nil {
		return err
	}
	return nil
}

func (app *HiYaYaBot) callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case GetHelpCmd:
					err := app.ShowHelpMsg(event)
					if err != nil {
						log.Println(err)
					}
				case GetBotAuthorProfileCmd:
					err := app.ShowBotAuthorProfile(event)
					if err != nil {
						log.Println(err)
					}
				default:
					err := app.ShowDefaultMsg(event)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}

func main() {
	app, err := NewHiYaYaBot(
		os.Getenv("ChannelSecret"),
		os.Getenv("ChannelAccessToken"),
		os.Getenv("DefaultTextMessage"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	server := fiber.New()
	server.Post("/callback", adaptor.HTTPHandlerFunc(app.callbackHandler))

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	_ = server.Listen(addr)
}
