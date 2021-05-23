package main

import "github.com/line/line-bot-sdk-go/v7/linebot"

const GetHelpCmd = "@help"

const availableFeatures = "There is all available features:\n" +
	GetBotAuthorProfileCmd

func (app *HiYaYaBot) ShowHelpMsg(event *linebot.Event) error {
	// create TextMessage.
	replyMessage := linebot.NewTextMessage(availableFeatures)

	// create Task.
	task := app.bot.ReplyMessage(event.ReplyToken, replyMessage)

	// execute the Task.
	_, err := task.Do()
	if err != nil {
		return err
	}

	return nil
}
