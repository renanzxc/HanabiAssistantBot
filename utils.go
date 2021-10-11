package main

import "strings"

func getCurrentMatch(senderID int) (*Match, bool) {
	var match, ok = matches[senderID]
	return match, ok
}

func parseTelegramData(data string) []string {
	return strings.Split(data, "|")
}
