package main

import (
	"fmt"
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Match contains all data about the match
type Match struct {
	CardsNumber int
	Cards       []Card
	Bot         *tb.Bot
}

// ShowCards show all cards on match to user
func (match *Match) ShowCards(sender *tb.User, messageToEdit ...*tb.Message) {
	selectorCards := match.Bot.NewMarkup()

	for ii := range match.Cards {
		var btn = newButton(fmt.Sprintf(`🎆 Number: %s Color: %s`, match.Cards[ii].Number, match.Cards[ii].Color), selectCardBtn.Unique, ButtonData{
			CardIndex: &ii,
		})
		selectorCards.InlineKeyboard = append(selectorCards.InlineKeyboard, []tb.InlineButton{*btn.Inline()})
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

// newMatch initialize a new match with the initial values
func newMatch(b *tb.Bot, cardsNumber int) (match *Match) {
	match = &Match{Bot: b, CardsNumber: cardsNumber}

	for ii := 0; ii < cardsNumber; ii++ {
		match.Cards = append(match.Cards, newCard())
	}

	return
}

// getCurrentMatch get a current match of an user
func getCurrentMatch(senderID int) (*Match, bool) {
	var match, ok = matches[senderID]
	return match, ok
}

// Handlers

func startMatchHandler(m *tb.Message) {
	var (
		numCardsSelector = &tb.ReplyMarkup{ResizeReplyKeyboard: true, OneTimeKeyboard: true, ReplyKeyboardRemove: true}
		err              error
	)

	numCardsSelector.Reply(
		numCardsSelector.Row(*numCardsBtn4),
		numCardsSelector.Row(*numCardsBtn5),
	)

	if _, err = bot.Send(m.Sender, "Select cards number", numCardsSelector); err != nil {
		log.Fatal(err)
	}
}

func selectCardHandler(c *tb.Callback) {
	var (
		match             *Match
		matchOK           bool
		infoToAddSelector = bot.NewMarkup()
		actionsRows       = []tb.Row{
			infoToAddSelector.Row(infoToAddSelector.Data(addNumberCardBtn.Text, addNumberCardBtn.Unique, c.Data)),
			infoToAddSelector.Row(infoToAddSelector.Data(addColorCardBtn.Text, addColorCardBtn.Unique, c.Data)),
		}
		btnData ButtonData
		err     error
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

	if _, err = bot.Edit(c.Message, "What information do you want to add to "+positionsText[*btnData.CardIndex]+" Card?", infoToAddSelector); err != nil {
		log.Fatal(err)
	}
	if err = bot.Respond(c, &tb.CallbackResponse{}); err != nil {
		log.Fatal(err)
	}
}
