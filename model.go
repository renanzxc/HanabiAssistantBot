package main

import (
	"fmt"
	"log"
	"strconv"

	tb "gopkg.in/tucnak/telebot.v2"
)

type ColorType string
type NumberType string

const unknown = "❓"

type Card struct {
	Color  ColorType
	Number NumberType
}

func (card *Card) HasAnyInfo() bool {
	return card.Color != unknown || card.Number != unknown
}

type Match struct {
	CardsNumber int
	Cards       []Card
	Bot         *tb.Bot
}

func (match *Match) ShowCards(sender *tb.User, messageToEdit ...*tb.Message) {
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
