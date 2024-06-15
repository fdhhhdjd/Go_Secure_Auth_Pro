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
