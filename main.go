package main

import (
	"ecommerce-product/cmd"
	"ecommerce-product/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupMySQL()

	// load redis
	helpers.SetupRedis()

	// run kafka consumer
	// cmd.ServeKafkaConsumer()

	// run http
	cmd.ServeHTTP()
}