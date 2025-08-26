package main

import (
	"log"
	"net"
)

type RedisClient struct {
	host      *string
	port      int
	conn      net.Conn
	connected bool
}

func (rc *RedisClient) connect() bool {
	IPs, _ := net.LookupIP(*rc.host)
	for i := 0; i < len(IPs); i++ {
		addr := net.TCPAddr{IP: IPs[i], Port: rc.port}
		conn, err := net.DialTCP("tcp", nil, &addr)
		if err == nil {
			rc.conn = conn
			rc.connected = true
			return true

		}
	}
	return false
}

func (rc *RedisClient) disConnect() {
	if rc.connected {
		rc.connected = false
		err := rc.conn.Close()
		rc.conn = nil
		if err != nil {
			log.Fatal("Error while closing connection")
		}
		println("Disconnected")
	}
}

func (rc *RedisClient) getConnection() *net.Conn {
	return &rc.conn
}
