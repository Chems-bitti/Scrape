package main

import (
	"bufio"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("insert bot token")
	if err != nil {
		panic(err) // Don't judge me using panic please
	}
	var chatID int64 = insertChatId
	bot.Debug = true
	scanner := bufio.NewScanner(os.Stdin) // Reading from stdin, you should pipe what you want to send to linkedIn bot, try echo Hi | linkedInbot
	for scanner.Scan() {
		jobs := fmt.Sprintf("Nouvelles offre dispo :\n %s", scanner.Text())
		msg := tgbotapi.NewMessage(chatID, jobs)
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
