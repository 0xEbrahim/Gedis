package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type CLI struct {
	redisClient *RedisClient
}

func MakeCLI(host *string, port int) *CLI {
	println("PORT", port)
	return &CLI{
		redisClient: &RedisClient{host: host, port: port, conn: nil},
	}
}

func (cli *CLI) Run() {
	if !cli.redisClient.connect() {
		return
	}
	cmdHandler := initHandler(cli.redisClient.getConnection())
	responseHandler := initResponseHandler(cli.redisClient.getConnection())
	host := *cli.redisClient.host
	port := cli.redisClient.port
	println("Connected to Gedis server at", host, ":", port)
	buf := bufio.NewScanner(os.Stdin)
	for true {
		print(host, ":", port, "> ")
		if buf.Scan() {
			line := buf.Text()
			line = strings.TrimSpace(line)
			if len(line) == 0 {
				continue
			}
			if line == "exit" || line == "quit" {
				println("Goodbye :(")
				break
			}
			if line == "help" {
				println("go run main.go -h <host> -p <port>")
				continue
			}
			tokens := cmdHandler.tokenizeArgs(line)
			if len(tokens) == 0 {
				continue
			}
			cmd := cmdHandler.buildRESP(tokens)
			if !cli.redisClient.sendCommand(cmd) {
				log.Fatal("Error: Failed to send command")
			}
			resp := responseHandler.parseResponse(cli.redisClient.getConnection())
			println(resp)
		} else {
			return
		}
	}
}
