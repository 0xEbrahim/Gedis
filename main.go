package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	host := "127.0.0.1"
	port := 6379
	args := os.Args
	i := 1
	for i < len(args) {
		if args[i] == "-h" {
			if i+1 < len(args) {
				i += 1
				host = args[i]
			} else {
				log.Fatal("you have to provide a host if you used -h")
			}
		}
		if args[i] == "-p" {
			if i+1 < len(args) {
				i += 1
				port, _ = strconv.Atoi(args[i])
			} else {
				log.Fatal("you have to provide a port if you used -p")
			}
		}
		i++
	}
	cli := MakeCLI(&host, port)
	cli.Run()
	defer cli.redisClient.disConnect()
}
