package main

import tb "gopkg.in/tucnak/telebot.v2"

func startCommand(m *tb.Message) {
	var (
		menuStart = &tb.ReplyMarkup{ResizeReplyKeyboard: true, OneTimeKeyboard: true}
		menuRows  = []tb.Row{menuStart.Row(*startMatchBtn)}
	)

	if _, ok := getCurrentMatch(m.Sender.ID); ok {
		menuRows = append(menuRows, menuStart.Row(*continueMatchBtn))
	}

	menuStart.Reply(menuRows...)

	bot.Send(m.Sender, "Welcome to 🎆 Hanabi Assistant 🎆!", menuStart)
}
