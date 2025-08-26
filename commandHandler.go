package main

import (
	"net"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type CommandHandler struct {
	resp *RESP
}

func initHandler(conn *net.TCPConn) *CommandHandler {
	return &CommandHandler{resp: &RESP{conn: conn}}
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

func (ch *CommandHandler) buildRESP(tokens []string) string {
	resp := "*" + strconv.Itoa(len(tokens)) + "\r\n"
	for _, it := range tokens {
		resp = resp + "$" + strconv.Itoa(len(it)) + "\r\n" + it + "\r\n"
	}
	return resp
}
