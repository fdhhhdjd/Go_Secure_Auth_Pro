package third_party

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
)

type TelegramMessage struct {
	ChatID                string `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool   `json:"disable_notification,omitempty"`
}

// SendTelegramMessage sends a message to a Telegram chat using the Telegram Bot API.
// It takes the message content, parse mode, disable web page preview, and disable notification as parameters.
// Returns an error if there was a problem sending the message.
func SendTelegramMessage(message string, parseMode string, disableWebPagePreview bool, disableNotification bool) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", global.Cfg.Telegram.BotToken)

	telegramMessage := TelegramMessage{
		ChatID:                global.Cfg.Telegram.ChatID,
		Text:                  message,
		ParseMode:             parseMode,
		DisableWebPagePreview: disableWebPagePreview,
		DisableNotification:   disableNotification,
	}

	jsonData, err := json.Marshal(telegramMessage)
	if err != nil {
		return err
	}

	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Response from Telegram API: %s\n", string(body))
	return nil
}

// PingTelegram pings the Telegram API to check if the bot is reachable.
// It sends a GET request to the Telegram API's getMe endpoint using the provided bot token.
// If the API returns a non-OK status, it returns an error.
func PingTelegram(botToken string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", botToken)

	response, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}
