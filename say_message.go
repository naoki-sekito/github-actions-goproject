package main

import (
	"fmt"

	"rsc.io/quote"
)

func SayMessage(msg string) string {

	fmt.Println(msg)

	if msg == "say" {
		return quote.Hello()
	}
	return "no msg"

}
