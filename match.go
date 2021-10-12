package main

import (
	"log"

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
		btnData ButtonData
	)
	btnData.Parse(c.Data)

	if match, matchOK = getCurrentMatch(c.Sender.ID); !matchOK {
		log.Panic("No match")
	}

	if match.Cards[*btnData.CardIndex].HasAnyInfo() {
		actionsRows = append(actionsRows, infoToAddSelector.Row(infoToAddSelector.Data(removeInfoBtn.Text, removeInfoBtn.Unique, c.Data)))
	}

	infoToAddSelector.Inline(
		actionsRows...,
	)

	bot.Edit(c.Message, "What information do you want to add to "+positionsText[*btnData.CardIndex]+" Card?", infoToAddSelector)
	bot.Respond(c, &tb.CallbackResponse{})
}

func addNumberCard(c *tb.Callback) {
	var (
		btnData            ButtonData
		numberInfoSelector = bot.NewMarkup()
	)
	btnData.Parse(c.Data)

	for ii := 1; ii <= 5; ii++ {
		btnData.SelectedNumberData = &numbersEmojis[ii]
		btn := newButton(string(numbersEmojis[ii]), selectNumberCardBtn.Unique, btnData).Inline()

		numberInfoSelector.InlineKeyboard = append(numberInfoSelector.InlineKeyboard, []tb.InlineButton{*btn})
	}

	bot.Edit(c.Message, "What number to add a "+positionsText[*btnData.CardIndex]+" Card?", numberInfoSelector)
	bot.Respond(c, &tb.CallbackResponse{})
}

func selectNumberCard(c *tb.Callback) {
	var (
		btnData ButtonData
	)
	btnData.Parse(c.Data)

	for ii := range numbersEmojis {
		if numbersEmojis[ii] == NumberType(*btnData.SelectedNumberData) {
			matches[c.Sender.ID].Cards[*btnData.CardIndex].Number = numbersEmojis[ii]
			break
		}
	}

	matches[c.Sender.ID].ShowCards(c.Sender, c.Message)
	bot.Respond(c, &tb.CallbackResponse{})
}

func addColorCard(c *tb.Callback) {
	var (
		btnData ButtonData
	)
	btnData.Parse(c.Data)

	colorInfoSelector := bot.NewMarkup()

	for ii := 0; ii < len(colorEmojis); ii++ {
		btnData.SelectedColorData = &colorEmojis[ii]
		btn := newButton(string(colorEmojis[ii]), selectColorCardBtn.Unique, btnData).Inline()

		colorInfoSelector.InlineKeyboard = append(colorInfoSelector.InlineKeyboard, []tb.InlineButton{*btn})
	}

	bot.Edit(c.Message, "What color to add a "+positionsText[*btnData.CardIndex]+" Card?", colorInfoSelector)
	bot.Respond(c, &tb.CallbackResponse{})
}

func selectColorCard(c *tb.Callback) {
	var (
		btnData ButtonData
	)
	btnData.Parse(c.Data)

	for ii := range numbersEmojis {
		if colorEmojis[ii] == ColorType(*btnData.SelectedColorData) {
			matches[c.Sender.ID].Cards[*btnData.CardIndex].Color = colorEmojis[ii]
			break
		}
	}

	matches[c.Sender.ID].ShowCards(c.Sender, c.Message)
	bot.Respond(c, &tb.CallbackResponse{})
}

func removeInfoCard(c *tb.Callback) {
	var btnData ButtonData
	btnData.Parse(c.Data)

	var match = matches[c.Sender.ID]
	match.Cards[*btnData.CardIndex] = newCard()

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
