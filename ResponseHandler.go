package main

import (
	"log"
	"net"
)

type ResponseHandler struct {
}

func initResponseHandler() *ResponseHandler {
	return &ResponseHandler{}
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

func (rh *ResponseHandler) parseResponse(conn *net.TCPConn) string {
	// Handle response based on prefix character
	prefix := readByte(conn)
	switch prefix {
	case "+":
		return rh.parseSimpleString(conn)
	case "-":
		return rh.parseSimpleError(conn)
	//case ":":
	//	return rh.parseIntegers(conn)
	//case "$":
	//	return rh.parseBulkString(conn)
	//case "*":
	//	return rh.parseArray(conn)
	default:
		log.Fatal("Unknown error")
	}
	return ""
}

func (rh *ResponseHandler) parseSimpleString(conn *net.TCPConn) string {
	return readLine(conn)
}

func (rh *ResponseHandler) parseSimpleError(conn *net.TCPConn) string {
	return readLine(conn)
}

//func (rh *ResponseHandler) parseIntegers(conn *net.TCPConn) string    {}
//func (rh *ResponseHandler) parseArray(conn *net.TCPConn) string       {}
//func (rh *ResponseHandler) parseBulkString(conn *net.TCPConn) string  {}
