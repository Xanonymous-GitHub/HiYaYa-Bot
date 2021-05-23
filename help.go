package main

import "github.com/line/line-bot-sdk-go/v7/linebot"

const GetHelpCmd = "@help"

const availableFeatures = "There is all available features:\n" +
	GetBotAuthorProfileCmd

func (app *HiYaYaBot) GetHelpMsg() *linebot.TextMessage {
	// create TextMessage.
	return linebot.NewTextMessage(availableFeatures)
}
