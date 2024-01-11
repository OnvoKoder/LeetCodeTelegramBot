package main

import (
	"log"
)

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	startBot()
}
