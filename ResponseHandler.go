package main

import (
	"net"
)

type ResponseHandler struct {
	resp *RESP
}

func initResponseHandler(conn *net.TCPConn) *ResponseHandler {
	return &ResponseHandler{resp: &RESP{conn: conn}}
}

func (rh *ResponseHandler) parseResponse(conn *net.TCPConn) string {
	return rh.resp.decode(conn)
}
