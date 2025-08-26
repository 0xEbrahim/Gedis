package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
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

func readByte(conn *net.TCPConn) string {
	b := make([]byte, 1)
	_, err := conn.Read(b)
	prefix := string(b)
	if err != nil {

	}
	return prefix
}
func readLine(conn *net.TCPConn) string {
	line := ""
	for {
		c := readByte(conn)
		if c == "\r" {
			readByte(conn)
			break
		}
		line = line + c
	}
	return line
}

func ReadLength(conn *net.TCPConn) int {
	length := readLine(conn)
	n, _ := strconv.Atoi(length)
	return n
}

func readBulkString(conn *net.TCPConn) string {
	ReadLength(conn)
	return readLine(conn)
}

func (resp *RESP) decode(conn *net.TCPConn) string {
	prefix := readByte(conn)
	switch prefix {
	case "+":
		return resp.parseSimpleString(conn)
	case "-":
		return resp.parseSimpleError(conn)
	case ":":
		return resp.parseIntegers(conn)
	case "$":
		return resp.parseBulkString(conn)
	case "*":
		return resp.parseArray(conn)
	case "_":
		return resp.parseNull(conn)
	case "#":
		return resp.parseBoolean(conn)
	case ",":
		return resp.parseDoubles(conn)
	case "(":
		return resp.parseBigNumbers(conn)
	case "!":
		return resp.parseBulkError(conn)
	case "=":
		return resp.parseVerbatimStrings(conn)
	case "%":
		return resp.parseMap(conn)
	default:
		log.Fatal("Unknown error")
	}
	return ""
}

func (resp *RESP) parseMap(conn *net.TCPConn) string {
	n := ReadLength(conn)
	parts := make([]string, 0, n)
	for i := 0; i < n; i++ {
		key := resp.decode(conn)
		value := resp.decode(conn)

		parts = append(parts, fmt.Sprintf("%s:%s", key, value))
	}
	return "{" + strings.Join(parts, ", ") + "}"
}

func (resp *RESP) parseSimpleString(conn *net.TCPConn) string {
	return readLine(conn)
}

func (resp *RESP) parseSimpleError(conn *net.TCPConn) string {
	return readLine(conn)
}

func (resp *RESP) parseIntegers(conn *net.TCPConn) string {
	return readLine(conn)
}

func (resp *RESP) parseArray(conn *net.TCPConn) string {
	n := ReadLength(conn)
	elems := make([]string, 0, n)
	for i := 0; i < n; i++ {
		val := resp.decode(conn)
		elems = append(elems, val)
	}
	return "[" + strings.Join(elems, ", ") + "]"
}
func (resp *RESP) parseBulkString(conn *net.TCPConn) string {
	return readBulkString(conn)
}
func (resp *RESP) parseBoolean(conn *net.TCPConn) string {
	return readLine(conn)
}

func (resp *RESP) parseDoubles(conn *net.TCPConn) string {
	return readLine(conn)
}

func (resp *RESP) parseBigNumbers(conn *net.TCPConn) string {
	return readLine(conn)
}

func (resp *RESP) parseBulkError(conn *net.TCPConn) string {
	return readBulkString(conn)
}

func (resp *RESP) parseVerbatimStrings(conn *net.TCPConn) string {
	return readBulkString(conn)
}

func (resp *RESP) parseNull(conn *net.TCPConn) string {
	readByte(conn)
	readLine(conn)
	return ""
}
