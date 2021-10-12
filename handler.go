package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

// setupHandlers setup all bot handlers
func setupHandlers() {
	// Handle commands
	bot.Handle("/start", startCommandHandler)

	// Handle buttons
	bot.Handle(startMatchBtn, startMatchHandler)
	bot.Handle(numCardsBtn4, func(m *tb.Message) {
		selectNumCardsHandler(4, m.Sender)
	})
	bot.Handle(numCardsBtn5, func(m *tb.Message) {
		selectNumCardsHandler(5, m.Sender)
	})
	bot.Handle(selectCardBtn, selectCardHandler)
	bot.Handle(addNumberCardBtn, addNumberCardHandler)
	bot.Handle(selectNumberCardBtn, selectNumberCardHandler)
	bot.Handle(addColorCardBtn, addColorCardHandler)
	bot.Handle(selectColorCardBtn, selectColorCardHandler)
	bot.Handle(removeInfoBtn, removeInfoCardHandler)
	bot.Handle(continueMatchBtn, func(m *tb.Message) {
		matches[m.Sender.ID].ShowCards(m.Sender)
	})
}

// selectNumCardsHandler setup match cards according to the number of cards entered
func selectNumCardsHandler(numCards int, sender *tb.User) {
	matches[sender.ID] = newMatch(bot, numCards)

	matches[sender.ID].ShowCards(sender)
}
