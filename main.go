package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	bot *tb.Bot

	matches = map[int]*Match{}

	positionsText = []string{"First", "Second", "Third", "Fourth", "Fifth"}
	numbersEmojis = []NumberType{"0️⃣", "1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣"}
	colorEmojis   = []ColorType{"⬜️", "🟨", "🟦", "🟥", "🟩"}
)

func main() {
	var (
		token string
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

	setupHandlers()

	fmt.Println("Run bot")
	bot.Start()
}
