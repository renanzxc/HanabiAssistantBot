package main

import (
	"log"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	// Buttons
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

// ButtonData contains all data about button
type ButtonData struct {
	CardIndex          *int
	SelectedNumberData *NumberType
	SelectedColorData  *ColorType
}

// newButton initialize a new button with the initial values
func newButton(text, unique string, data ButtonData) *tb.Btn {
	return &tb.Btn{
		Text:   text,
		Unique: unique,
		Data:   data.ToString(),
	}
}

// ToString convert all button data to string
func (btn *ButtonData) ToString() (btnDataString string) {
	syncBtnValues(btn, &btnDataString)

	return
}

// Parse parse string data to button struct
func (btn *ButtonData) Parse(btnDataString string) {
	syncBtnValues(btn, &btnDataString)
}

// syncBtnValues sync a button data to struct or string
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
