package main

type CLI struct {
	redisClient *RedisClient
}

func MakeCLI(host *string, port int) *CLI {
	return &CLI{
		redisClient: &RedisClient{host: host, port: port, conn: nil},
	}
}

func (cli *CLI) Run() {
	if !cli.redisClient.connect() {
		return
	}
	println("Connected to redis server")
}
