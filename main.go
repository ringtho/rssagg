package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ringtho/rssagg/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}


func main() {
	fmt.Println("hello world")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port:", portString)
}