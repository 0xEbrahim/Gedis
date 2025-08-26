package main

import (
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
				// Error
			}
		}
		if args[i] == "-p" {
			if i+1 < len(args) {
				i += 1
				port, err := strconv.Atoi(args[i])
				if err != nil {
					// Error
				}
				println(port)
			} else {
				// Error
			}
		}

		i++
	}
}
