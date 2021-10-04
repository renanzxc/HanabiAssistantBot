package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	matches       = map[int]*Match{}
	positionsText = []string{"First", "Second", "Third", "Fourth", "Fifth"}
	numbersEmojis = []NumberType{"0️⃣", "1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣"}
	colorEmojis   = []ColorType{"⬜️", "🟨", "🟦", "🟥", "🟩"}

	//buttons
	numCardsSelector    = &tb.ReplyMarkup{ResizeReplyKeyboard: true, OneTimeKeyboard: true, ReplyKeyboardRemove: true}
	numCardsReply4      = numCardsSelector.Text("4")
	numCardsReply5      = numCardsSelector.Text("5")
	selectCardBtn       = &tb.Btn{Unique: "select_card_btn"}
	addNumberInfoBtn    = &tb.Btn{Text: "🔢 Number", Unique: "add_card_number"}
	addColorInfoBtn     = &tb.Btn{Text: "🎨 Color", Unique: "add_card_color"}
	selectNumberInfoBtn = &tb.Btn{Unique: "add_card_number_info"}
	selectColorInfoBtn  = &tb.Btn{Unique: "add_card_color_info"}
)

func setupHandlers(bot *tb.Bot) {
	var (
		menu *tb.ReplyMarkup

		btnStartMatch    = menu.Text("▶ Start Match")
		btnContinueMatch = menu.Text("▶ Continue Match")
	)

	bot.Handle("/start", func(m *tb.Message) {
		menu = &tb.ReplyMarkup{ResizeReplyKeyboard: true, OneTimeKeyboard: true}
		if _, ok := matches[m.Sender.ID]; ok {
			menu.Reply(
				menu.Row(btnStartMatch),
				menu.Row(btnContinueMatch),
			)
		} else {
			menu.Reply(
				menu.Row(btnStartMatch),
			)
		}

		bot.Send(m.Sender, "Welcome to 🎆 Hanabi Assistant 🎆!", menu)
	})

	bot.Handle(&btnStartMatch, func(m *tb.Message) {
		numCardsSelector.Reply(
			numCardsSelector.Row(numCardsReply4),
			numCardsSelector.Row(numCardsReply5),
		)

		bot.Send(m.Sender, "Select cards number", numCardsSelector)
	})

	selectNumCardsHandler := func(numCards int, sender *tb.User) {
		matches[sender.ID] = newMatch(bot, numCards)

		matches[sender.ID].showCards(sender)
	}

	bot.Handle(&numCardsReply4, func(m *tb.Message) {
		selectNumCardsHandler(4, m.Sender)
	})

	bot.Handle(&numCardsReply5, func(m *tb.Message) {
		selectNumCardsHandler(5, m.Sender)
	})

	bot.Handle(selectCardBtn, func(c *tb.Callback) {
		cardIndex, err := strconv.ParseInt(c.Data, 10, 64)
		if err != nil {
			log.Panic(err)
		}

		infoToAddSelector := bot.NewMarkup()
		infoToAddSelector.Inline(
			infoToAddSelector.Row(infoToAddSelector.Data(addNumberInfoBtn.Text, addNumberInfoBtn.Unique, c.Data)),
			infoToAddSelector.Row(infoToAddSelector.Data(addColorInfoBtn.Text, addColorInfoBtn.Unique, c.Data)),
		)
		bot.Edit(c.Message, "What information do you want to add to "+positionsText[cardIndex]+" Card?", infoToAddSelector)
		bot.Respond(c, &tb.CallbackResponse{})
	})

	bot.Handle(addNumberInfoBtn, func(c *tb.Callback) {
		cardIndex, err := strconv.ParseInt(c.Data, 10, 64)
		if err != nil {
			log.Panic(err)
		}

		numberInfoSelector := bot.NewMarkup()

		for ii := 1; ii <= 5; ii++ {
			btn := numberInfoSelector.Data(string(numbersEmojis[ii]), selectNumberInfoBtn.Unique, c.Data, string(numbersEmojis[ii])).Inline()
			numberInfoSelector.InlineKeyboard = append(numberInfoSelector.InlineKeyboard, []tb.InlineButton{*btn})
		}

		bot.Edit(c.Message, "What number to add a "+positionsText[cardIndex]+" Card?", numberInfoSelector)
		bot.Respond(c, &tb.CallbackResponse{})
	})

	bot.Handle(selectNumberInfoBtn, func(c *tb.Callback) {
		var data = parseTelegramData(c.Data)
		cardIndex, err := strconv.ParseInt(data[0], 10, 64)
		if err != nil {
			log.Panic(err)
		}

		for ii := range numbersEmojis {
			if numbersEmojis[ii] == NumberType(data[1]) {
				matches[c.Sender.ID].Cards[cardIndex].Number = numbersEmojis[ii]
				break
			}
		}

		matches[c.Sender.ID].showCards(c.Sender, c.Message)
		bot.Respond(c, &tb.CallbackResponse{})
	})

	bot.Handle(addColorInfoBtn, func(c *tb.Callback) {
		cardIndex, err := strconv.ParseInt(c.Data, 10, 64)
		if err != nil {
			log.Panic(err)
		}

		colorInfoSelector := bot.NewMarkup()

		for ii := 0; ii < len(colorEmojis); ii++ {
			btn := colorInfoSelector.Data(string(colorEmojis[ii]), selectColorInfoBtn.Unique, c.Data, string(colorEmojis[ii])).Inline()
			colorInfoSelector.InlineKeyboard = append(colorInfoSelector.InlineKeyboard, []tb.InlineButton{*btn})
		}

		bot.Edit(c.Message, "What color to add a "+positionsText[cardIndex]+" Card?", colorInfoSelector)
		bot.Respond(c, &tb.CallbackResponse{})
	})

	bot.Handle(selectColorInfoBtn, func(c *tb.Callback) {
		var data = parseTelegramData(c.Data)
		cardIndex, err := strconv.ParseInt(data[0], 10, 64)
		if err != nil {
			log.Panic(err)
		}

		for ii := range numbersEmojis {
			if colorEmojis[ii] == ColorType(data[1]) {
				matches[c.Sender.ID].Cards[cardIndex].Color = colorEmojis[ii]
				break
			}
		}

		matches[c.Sender.ID].showCards(c.Sender, c.Message)
		bot.Respond(c, &tb.CallbackResponse{})
	})

	bot.Handle(&btnContinueMatch, func(m *tb.Message) {
		matches[m.Sender.ID].showCards(m.Sender)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		fmt.Println(m.Text)
	})
}

func newMatch(b *tb.Bot, cardsNumber int) (match *Match) {
	match = &Match{Bot: b, CardsNumber: cardsNumber}

	for ii := 0; ii < cardsNumber; ii++ {
		match.Cards = append(match.Cards, Card{Color: unknown, Number: unknown})
	}

	return
}

func (match *Match) showCards(sender *tb.User, messageToEdit ...*tb.Message) {
	selectorCards := match.Bot.NewMarkup()

	for ii := range match.Cards {
		btn := selectorCards.Data(fmt.Sprintf(`🎆 Number: %s Color: %s`, match.Cards[ii].Number, match.Cards[ii].Color), selectCardBtn.Unique, strconv.FormatInt(int64(ii), 10)).Inline()
		selectorCards.InlineKeyboard = append(selectorCards.InlineKeyboard, []tb.InlineButton{*btn})
	}

	if len(messageToEdit) > 0 {
		if _, err := match.Bot.Edit(messageToEdit[0], "Cards Info:", selectorCards); err != nil {
			log.Panic(err)
		}
	} else {
		if _, err := match.Bot.Send(sender, "Cards Info:", selectorCards); err != nil {
			log.Panic(err)
		}
	}
}

func parseTelegramData(data string) []string {
	return strings.Split(data, "|")
}
