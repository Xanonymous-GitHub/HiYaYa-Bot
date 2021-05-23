package main

import (
	"fmt"
	"github.com/Xanonymous-GitHub/HiYaYa-Bot/utils"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"io"
	"log"
	"net/http"
)

const GetConfirmedAmountCmd = "@today"

type CovidCurrentStatus struct {
	TotalConfirmed    int    `json:"確診"`
	ReleaseQuarantine int    `json:"解除隔離"`
	TotalDeath        int    `json:"死亡"`
	TotalInspection   string `json:"送驗"`
	TotalExclusion    string `json:"排除"`
	NewConfirmedCase  int    `json:"昨日確診"`
	NewExclusion      string `json:"昨日排除"`
	NewInspection     string `json:"昨日送驗"`
}

func fetchConfirmedAmountFromCdc() (string, error) {
	const dataSourceUrl = "https://covid19dashboard.cdc.gov.tw/dash3"

	resp, err := http.Get(dataSourceUrl)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	currentStatus := &CovidCurrentStatus{}
	err = utils.ParseJSONBody(resp.Body, currentStatus)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("New confirmed case today in Taiwan: %d", currentStatus.NewConfirmedCase)
	return result, nil
}

func (app *HiYaYaBot) GetConfirmedAmount() *linebot.TextMessage {
	confirmedAmountMsg, err := fetchConfirmedAmountFromCdc()
	if err != nil {
		log.Println(err)
		return nil
	}

	return linebot.NewTextMessage(confirmedAmountMsg)
}
