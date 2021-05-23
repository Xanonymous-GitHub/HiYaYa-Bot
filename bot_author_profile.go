package main

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"os"
)

const GetBotAuthorProfileCmd = "@author"

func (app *HiYaYaBot) GetBotAuthorProfile() *linebot.TextMessage {
	// get github link.
	profileLink := os.Getenv("BotAuthorProfileLink")

	// create TextMessage.
	return linebot.NewTextMessage(profileLink)
}
