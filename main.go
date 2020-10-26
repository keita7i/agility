package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	di := DI{}
	panic(di.Router().Run(fmt.Sprintf(":%s", port)))
}
