package main

import (
	"net"
	"strconv"
)

const (
	SIMPLE_STRING string = "+"
	INTEGER       string = ":"
	SIMPLE_ERROR  string = "-"
	ARRAY         string = "*"
	BULK_STRING   string = "$"
)

type RESP struct {
	conn *net.TCPConn
}

func (resp *RESP) encodeSimpleString(str string) string {
	encoded := SIMPLE_STRING + str + "\r\n"
	return encoded
}

func (resp *RESP) encodeSimpleError(str string) string {
	encoded := SIMPLE_ERROR + str + "\r\n"
	return encoded
}

func (resp *RESP) encodeArray(tokens []string) string {
	encoded := ARRAY + strconv.Itoa(len(tokens)) + "\r\n"
	for _, it := range tokens {
		encoded = encoded + "$" + strconv.Itoa(len(it)) + "\r\n" + it + "\r\n"
	}
	return encoded
}
func (resp *RESP) encodeInteger(n int) string {
	encoded := INTEGER + strconv.Itoa(n) + "\r\n"
	return encoded
}

func (resp *RESP) encodeBulkString(str string) string {
	encoded := BULK_STRING + strconv.Itoa(len(str)) + "\r\n" + str + "\r\n"
	return encoded
}
