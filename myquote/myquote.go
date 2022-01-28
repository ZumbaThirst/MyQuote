package myquote

import (
	"fmt"

	"rsc.io/quote"
)

func PrintQuote() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Hello())
}
