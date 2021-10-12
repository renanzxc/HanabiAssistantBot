package main

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

// startCommandHandler handle start command
func startCommandHandler(m *tb.Message) {
	var (
		menuStart = &tb.ReplyMarkup{ResizeReplyKeyboard: true, OneTimeKeyboard: true}
		menuRows  = []tb.Row{menuStart.Row(*startMatchBtn)}
		err       error
	)

	// If there is a match in progress the option "Continue" is added
	if _, ok := getCurrentMatch(m.Sender.ID); ok {
		menuRows = append(menuRows, menuStart.Row(*continueMatchBtn))
	}

	menuStart.Reply(menuRows...)

	if _, err = bot.Send(m.Sender, "Welcome to 🎆 Hanabi Assistant 🎆!", menuStart); err != nil {
		log.Fatal(err)
	}
}
