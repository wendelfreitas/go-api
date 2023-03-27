package main

import "github.com/wendelfreitas/go-api/api/configs"

func main() {
	config,_ := configs.LoadConfig(".")

	println(config.DBHost)
}