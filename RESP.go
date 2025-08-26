package main

import (
	"log"
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
	default:
		log.Fatal("Unknown error")
	}
	return ""
}

func readArray(conn *net.TCPConn) string {
	str := ""
	arrLen := ReadLength(conn)
	for arrLen > 0 {
		readByte(conn)
		ReadLength(conn)
		str = str + readLine(conn) + " "
		arrLen = arrLen - 1
	}
	return str
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
	return readArray(conn)
}

func (resp *RESP) parseBulkString(conn *net.TCPConn) string {
	return readBulkString(conn)
}

func (resp *RESP) parseNull(conn *net.TCPConn) string {
	readByte(conn)
	readLine(conn)
	return ""
}
