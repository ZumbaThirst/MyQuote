package main

import ("rsc.io/quote"
    "fmt")

func main() {
    fmt.Print(glass())
    fmt.Print(Go())
    fmt.Print(hello())
    fmt.Print(opt())
}

func glass() string {
    return quote.Glass()
}

func Go() string {
    return quote.Go()
}

func hello() string {
    return quote.Hello()
}

func opt() string {
    return quote.Opt()
}

