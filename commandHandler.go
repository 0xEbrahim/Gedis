package main

import (
	"regexp"
	"unicode/utf8"
)

type CommandHandler struct {
}

func initHandler() *CommandHandler {
	return &CommandHandler{}
}

func (ch *CommandHandler) tokenizeArgs(args string) []string {
	var tokens []string
	pattern := regexp.MustCompile(`"[^"]+"|\S+`)
	strs := pattern.FindAllString(args, -1)
	for _, it := range strs {
		token := it
		sz := utf8.RuneCountInString(it)
		if sz > 2 && it[0] == '"' && it[sz-1] == '"' {
			token = it[1 : sz-1]
		}
		tokens = append(tokens, token)
	}
	return tokens
}
