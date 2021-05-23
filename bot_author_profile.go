package main

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"os"
)

const GetBotAuthorProfileCmd = "@author"

func (app *HiYaYaBot) ShowBotAuthorProfile(event *linebot.Event) error {
	// get github link.
	profileLink := os.Getenv("BotAuthorProfileLink")

	// create TextMessage.
	replyMessage := linebot.NewTextMessage(profileLink)

	// create Task.
	task := app.bot.ReplyMessage(event.ReplyToken, replyMessage)

	// execute the Task.
	_, err := task.Do()
	if err != nil {
		return err
	}

	return nil
}
