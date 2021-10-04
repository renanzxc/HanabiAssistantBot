package main

import tb "gopkg.in/tucnak/telebot.v2"

type ColorType string
type NumberType string

const unknown = "❓"

type Card struct {
	Color  ColorType
	Number NumberType
}

type Match struct {
	CardsNumber int
	Cards       []Card
	Bot         *tb.Bot
}
