package main

import (
	"fmt"
	"github.com/Xanonymous-GitHub/HiYaYa-Bot/utils"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const GetConfirmedAmountCmd = "@today"

type CovidCurrentStatus struct {
	TotalConfirmed    string `json:"確診"`
	ReleaseQuarantine string `json:"解除隔離"`
	TotalDeath        int    `json:"死亡"`
	TotalInspection   string `json:"送驗"`
	TotalExclusion    string `json:"排除"`
	NewConfirmedCase  int    `json:"昨日確診"`
	NewExclusion      string `json:"昨日排除"`
	NewInspection     string `json:"昨日送驗"`
}

func fetchConfirmedAmountFromCdc() (int, error) {
	const dataSourceUrl = "https://covid19dashboard.cdc.gov.tw/dash3"

	resp, err := http.Get(dataSourceUrl)
	if err != nil {
		return -1, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// fix "0" key at top level of our response body.
	bodyString, err := utils.ReaderToString(resp.Body)
	if err != nil {
		return -1, err
	}
	fixBody := utils.StringToReadCloser(bodyString[5 : len(bodyString)-1])

	currentStatus := &CovidCurrentStatus{}
	err = utils.ParseJSONBody(fixBody, currentStatus)
	if err != nil {
		return -1, err
	}

	return currentStatus.NewConfirmedCase, nil
}

func (app *HiYaYaBot) replyPictureFromText(text string) {
	rapidAPIKey := os.Getenv("RapidApiKey")
	url := fmt.Sprintf("https://img4me.p.rapidapi.com/?text=%s&type=png&bcolor=FFFFFF&fcolor=000000&size=35&font=trebuchet", text)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", rapidAPIKey)
	req.Header.Add("x-rapidapi-host", "img4me.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	body, _ := ioutil.ReadAll(res.Body)

	app.bot.ReplyMessage(string(body))
}

func (app *HiYaYaBot) GetConfirmedAmount() *linebot.TextMessage {
	confirmedAmount, err := fetchConfirmedAmountFromCdc()
	if err != nil {
		log.Println(err)
		return nil
	}

	result := fmt.Sprintf("New confirmed case today in Taiwan: %d", confirmedAmount)
	//app.replyPictureFromText(string(rune(confirmedAmount)))
	return linebot.NewTextMessage(result)
}
