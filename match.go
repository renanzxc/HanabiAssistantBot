package main

import (
	"log"
	"strconv"

	tb "gopkg.in/tucnak/telebot.v2"
)

func startMatch(m *tb.Message) {
	var (
		numCardsSelector = &tb.ReplyMarkup{ResizeReplyKeyboard: true, OneTimeKeyboard: true, ReplyKeyboardRemove: true}
	)

	numCardsSelector.Reply(
		numCardsSelector.Row(*numCardsBtn4),
		numCardsSelector.Row(*numCardsBtn5),
	)

	bot.Send(m.Sender, "Select cards number", numCardsSelector)
}

func selectCard(c *tb.Callback) {
	var (
		match             *Match
		matchOK           bool
		infoToAddSelector = bot.NewMarkup()
		actionsRows       = []tb.Row{
			infoToAddSelector.Row(infoToAddSelector.Data(addNumberCardBtn.Text, addNumberCardBtn.Unique, c.Data)),
			infoToAddSelector.Row(infoToAddSelector.Data(addColorCardBtn.Text, addColorCardBtn.Unique, c.Data)),
		}
	)

	cardIndex, err := strconv.ParseInt(c.Data, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	if match, matchOK = getCurrentMatch(c.Sender.ID); !matchOK {
		log.Panic("No match")
	}

	if match.Cards[cardIndex].HasAnyInfo() {
		actionsRows = append(actionsRows, infoToAddSelector.Row(infoToAddSelector.Data(removeInfoBtn.Text, removeInfoBtn.Unique, c.Data)))
	}

	infoToAddSelector.Inline(
		actionsRows...,
	)

	bot.Edit(c.Message, "What information do you want to add to "+positionsText[cardIndex]+" Card?", infoToAddSelector)
	bot.Respond(c, &tb.CallbackResponse{})
}

func addNumberCard(c *tb.Callback) {
	cardIndex, err := strconv.ParseInt(c.Data, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	numberInfoSelector := bot.NewMarkup()

	for ii := 1; ii <= 5; ii++ {
		btn := numberInfoSelector.Data(string(numbersEmojis[ii]), selectNumberCardBtn.Unique, c.Data, string(numbersEmojis[ii])).Inline()
		numberInfoSelector.InlineKeyboard = append(numberInfoSelector.InlineKeyboard, []tb.InlineButton{*btn})
	}

	bot.Edit(c.Message, "What number to add a "+positionsText[cardIndex]+" Card?", numberInfoSelector)
	bot.Respond(c, &tb.CallbackResponse{})
}

func selectNumberCard(c *tb.Callback) {
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

	matches[c.Sender.ID].ShowCards(c.Sender, c.Message)
	bot.Respond(c, &tb.CallbackResponse{})
}

func addColorCard(c *tb.Callback) {
	cardIndex, err := strconv.ParseInt(c.Data, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	colorInfoSelector := bot.NewMarkup()

	for ii := 0; ii < len(colorEmojis); ii++ {
		btn := colorInfoSelector.Data(string(colorEmojis[ii]), selectColorCardBtn.Unique, c.Data, string(colorEmojis[ii])).Inline()
		colorInfoSelector.InlineKeyboard = append(colorInfoSelector.InlineKeyboard, []tb.InlineButton{*btn})
	}

	bot.Edit(c.Message, "What color to add a "+positionsText[cardIndex]+" Card?", colorInfoSelector)
	bot.Respond(c, &tb.CallbackResponse{})
}

func selectColorCard(c *tb.Callback) {
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

	matches[c.Sender.ID].ShowCards(c.Sender, c.Message)
	bot.Respond(c, &tb.CallbackResponse{})
}

func removeInfoCard(c *tb.Callback) {
	cardIndex, err := strconv.ParseInt(c.Data, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	var match = matches[c.Sender.ID]
	match.Cards[cardIndex] = newCard()

	match.ShowCards(c.Sender, c.Message)
	bot.Respond(c, &tb.CallbackResponse{})
}

func newMatch(b *tb.Bot, cardsNumber int) (match *Match) {
	match = &Match{Bot: b, CardsNumber: cardsNumber}

	for ii := 0; ii < cardsNumber; ii++ {
		match.Cards = append(match.Cards, newCard())
	}

	return
}

func newCard() Card {
	return Card{Color: unknown, Number: unknown}
}
