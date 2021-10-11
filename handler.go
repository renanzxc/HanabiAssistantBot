package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	//buttons
	startMatchBtn    = &tb.Btn{Text: "▶ Start Match"}
	continueMatchBtn = &tb.Btn{Text: "▶ Continue Match"}

	numCardsBtn4        = &tb.Btn{Text: "4"}
	numCardsBtn5        = &tb.Btn{Text: "5"}
	selectCardBtn       = &tb.Btn{Unique: "select_card_btn"}
	addNumberCardBtn    = &tb.Btn{Text: "🔢 Number", Unique: "add_card_number"}
	addColorCardBtn     = &tb.Btn{Text: "🎨 Color", Unique: "add_card_color"}
	removeInfoBtn       = &tb.Btn{Text: "❌ Remove Info", Unique: "remove_info"}
	selectNumberCardBtn = &tb.Btn{Unique: "add_card_number_info"}
	selectColorCardBtn  = &tb.Btn{Unique: "add_card_color_info"}
)

func setupHandlers() {
	// Handle commands
	bot.Handle("/start", startCommand)

	// Handle buttons
	bot.Handle(startMatchBtn, startMatch)
	bot.Handle(numCardsBtn4, func(m *tb.Message) {
		selectNumCardsHandler(4, m.Sender)
	})
	bot.Handle(numCardsBtn5, func(m *tb.Message) {
		selectNumCardsHandler(5, m.Sender)
	})
	bot.Handle(selectCardBtn, selectCard)
	bot.Handle(addNumberCardBtn, addNumberCard)
	bot.Handle(selectNumberCardBtn, selectNumberCard)
	bot.Handle(addColorCardBtn, addColorCard)
	bot.Handle(selectColorCardBtn, selectColorCard)
	bot.Handle(removeInfoBtn, removeInfoCard)
	bot.Handle(continueMatchBtn, func(m *tb.Message) {
		matches[m.Sender.ID].ShowCards(m.Sender)
	})
}

func selectNumCardsHandler(numCards int, sender *tb.User) {
	matches[sender.ID] = newMatch(bot, numCards)

	matches[sender.ID].ShowCards(sender)
}
