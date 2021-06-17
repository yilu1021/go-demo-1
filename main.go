package main

import (
	"fmt"

	"example.com/go-demo-1/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMoscot())
	fmt.Print(quote.Go())
}
