package main

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

// ColorType represents a type of card color
type ColorType string

// NumberType represents a type of card number
type NumberType string

// unknown represents unknown data about card
const unknown = "❓"

// Card contains all data about card
type Card struct {
	Color  ColorType
	Number NumberType
}

// HasAnyInfo verify if exists any info about card
func (card *Card) HasAnyInfo() bool {
	return card.Color != unknown || card.Number != unknown
}

// newCard initialize a new card with the initial values
func newCard() Card {
	return Card{Color: unknown, Number: unknown}
}

// Handlers

func addNumberCardHandler(c *tb.Callback) {
	var (
		btnData            ButtonData
		numberInfoSelector = bot.NewMarkup()
		err                error
	)
	btnData.Parse(c.Data)

	for ii := 1; ii <= 5; ii++ {
		btnData.SelectedNumberData = &numbersEmojis[ii]
		btn := newButton(string(numbersEmojis[ii]), selectNumberCardBtn.Unique, btnData).Inline()

		numberInfoSelector.InlineKeyboard = append(numberInfoSelector.InlineKeyboard, []tb.InlineButton{*btn})
	}

	if _, err = bot.Edit(c.Message, "What number to add a "+positionsText[*btnData.CardIndex]+" Card?", numberInfoSelector); err != nil {
		log.Fatal(err)
	}
	if err = bot.Respond(c, &tb.CallbackResponse{}); err != nil {
		log.Fatal(err)
	}
}

func selectNumberCardHandler(c *tb.Callback) {
	var (
		btnData ButtonData
		err     error
	)
	btnData.Parse(c.Data)

	for ii := range numbersEmojis {
		if numbersEmojis[ii] == NumberType(*btnData.SelectedNumberData) {
			matches[c.Sender.ID].Cards[*btnData.CardIndex].Number = numbersEmojis[ii]
			break
		}
	}

	matches[c.Sender.ID].ShowCards(c.Sender, c.Message)
	if err = bot.Respond(c, &tb.CallbackResponse{}); err != nil {
		log.Fatal(err)
	}
}

func addColorCardHandler(c *tb.Callback) {
	var (
		btnData ButtonData
		err     error
	)
	btnData.Parse(c.Data)

	colorInfoSelector := bot.NewMarkup()

	for ii := 0; ii < len(colorEmojis); ii++ {
		btnData.SelectedColorData = &colorEmojis[ii]
		btn := newButton(string(colorEmojis[ii]), selectColorCardBtn.Unique, btnData).Inline()

		colorInfoSelector.InlineKeyboard = append(colorInfoSelector.InlineKeyboard, []tb.InlineButton{*btn})
	}

	if _, err = bot.Edit(c.Message, "What color to add a "+positionsText[*btnData.CardIndex]+" Card?", colorInfoSelector); err != nil {
		log.Fatal(err)
	}
	if err = bot.Respond(c, &tb.CallbackResponse{}); err != nil {
		log.Fatal(err)
	}
}

func selectColorCardHandler(c *tb.Callback) {
	var (
		btnData ButtonData
		err     error
	)
	btnData.Parse(c.Data)

	for ii := range numbersEmojis {
		if colorEmojis[ii] == ColorType(*btnData.SelectedColorData) {
			matches[c.Sender.ID].Cards[*btnData.CardIndex].Color = colorEmojis[ii]
			break
		}
	}

	matches[c.Sender.ID].ShowCards(c.Sender, c.Message)
	if err = bot.Respond(c, &tb.CallbackResponse{}); err != nil {
		log.Fatal(err)
	}
}

func removeInfoCardHandler(c *tb.Callback) {
	var (
		btnData ButtonData
		err     error
	)
	btnData.Parse(c.Data)

	var match = matches[c.Sender.ID]
	match.Cards[*btnData.CardIndex] = newCard()

	match.ShowCards(c.Sender, c.Message)
	if err = bot.Respond(c, &tb.CallbackResponse{}); err != nil {
		log.Fatal(err)
	}
}
