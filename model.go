package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

type ColorType string
type NumberType string

const unknown = "❓"

type ButtonData struct {
	CardIndex          *int
	SelectedNumberData *NumberType
	SelectedColorData  *ColorType
}

func newButton(text, unique string, data ButtonData) *tb.Btn {
	return &tb.Btn{
		Text:   text,
		Unique: unique,
		Data:   data.ToString(),
	}
}

func (btn *ButtonData) ToString() (btnDataString string) {
	syncBtnValues(btn, &btnDataString)

	return
}

func (btn *ButtonData) Parse(btnDataString string) {
	syncBtnValues(btn, &btnDataString)
}

func syncBtnValues(btnData *ButtonData, btnDataString *string) {
	var (
		btnDataArray []string
		err          error
	)

	if *btnDataString != "" {
		btnDataArray = strings.Split(*btnDataString, "|")
	} else {
		btnDataArray = make([]string, 3)
	}

	if btnData.CardIndex != nil {
		btnDataArray[0] = ValueToString(btnData.CardIndex)
	} else if btnDataArray[0] != "" {
		var cardIndex int64

		if cardIndex, err = strconv.ParseInt(btnDataArray[0], 10, 0); err != nil {
			log.Fatal(err)
		}

		btnData.CardIndex = GetIntPointer(int(cardIndex))
	}

	if btnData.SelectedColorData != nil {
		btnDataArray[1] = ValueToString(btnData.SelectedColorData)
	} else if btnDataArray[1] != "" {
		var colorData = ColorType(btnDataArray[1])

		btnData.SelectedColorData = &colorData
	}

	if btnData.SelectedNumberData != nil {
		btnDataArray[2] = ValueToString(btnData.SelectedNumberData)
	} else if btnDataArray[2] != "" {
		var numberData = NumberType(btnDataArray[2])

		btnData.SelectedNumberData = &numberData
	}

	*btnDataString = strings.Join(btnDataArray, "|")
}

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

func getCurrentMatch(senderID int) (*Match, bool) {
	var match, ok = matches[senderID]
	return match, ok
}

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
