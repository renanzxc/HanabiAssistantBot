package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		token string
		bot   *tb.Bot
		err   error
	)

	if token = os.Getenv("HANABI_TOKEN_BOT"); token == "" {
		log.Fatal("Invalid Bot token")
	}
	if bot, err = tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}); err != nil {
		log.Fatal(err)
		return
	}

	setupHandlers(bot)

	fmt.Println("Run bot")
	bot.Start()
}
