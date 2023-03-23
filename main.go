package main

import (
	"challange-7/routers"
)

func main () {
	var port = ":8080"

	router.StartServer().Run(port)
}