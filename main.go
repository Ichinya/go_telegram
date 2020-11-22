package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type GetMeT struct {
	Ok     bool         `json:"ok"`
	Result GetMeResultT `json:"result"`
}

type GetMeResultT struct {
	Id        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type SendMessageT struct {
	OK     bool     `json:"ok"`
	Result MessageT `json:"result"`
}

type MessageT struct {
	MessageID int                          `json:"message_id"`
	From      GetUpdatesResultMessageFromT `json:"from"`
	Chat      GetUpdatesResultMessageChatT `json:"chat"`
	Date      int                          `json:"date"`
	Text      string                       `json:"text"`
}

type GetUpdatesResultMessageFromT struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type GetUpdatesResultMessageChatT struct {
	Id        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type GetUpdatesT struct {
	Ok     bool                `json:"ok"`
	Result []GetUpdatesResultT `json:"result"`
}

type GetUpdatesResultT struct {
	UpdateID int                `json:"ipdate_id"`
	Message  GetUpdatesMessageT `json:"message,omitemply"`
}

type GetUpdatesMessageT struct {
	MessageID int                          `json:"message_id"`
	From      GetUpdatesResultMessageFromT `json:"from"`
	Chat      GetUpdatesResultMessageChatT `json:"chat"`
	Date      int                          `json:"date"`
	Text      string                       `json:"text"`
}

const telegramBaseUrl = "https://api.telegram.org/bot"
const telegramToken = "1459181557:AAEvS2w8ItmgpEu0i93LGGOV-YGS-vfekoE"

const methodGetMe = "getMe"
const methodGetUpdates = "getUpdates"
const methodSendMessage = "sendMessage"

func main() {
	body := getBodyByUrl(getUrlByMethod(methodGetUpdates))
	getUpdates := GetUpdatesT{}
	err := json.Unmarshal(body, &getUpdates)
	if err != nil {
		fmt.Printf("Error in Unmarshal: %s", err.Error())
		return
	}
	for _, update := range getUpdates.Result {
		if strings.ToLower(update.Message.Text) == "go" {
			url := getUrlByMethod(methodSendMessage) + "?chat_id=" + strconv.Itoa(update.Message.Chat.Id) + "&text=go go"
			body := getBodyByUrl(url)

			fmt.Printf("%s", body)
		}
		if strings.Contains(update.Message.Text, "load") {
			url := getUrlByMethod(methodSendMessage) + "?chat_id=" + strconv.Itoa(update.Message.Chat.Id) + "&text=load is found"
			body := getBodyByUrl(url)

			fmt.Printf("%s", body)
		}
		// fix нужен offset
	}

}

func getUrlByMethod(methodName string) string {
	return telegramBaseUrl + telegramToken + "/" + methodName
}

func getBodyByUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return body
}
