package main

import "apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}
