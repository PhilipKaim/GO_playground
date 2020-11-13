package main

import (
        "fmt"

        "rsc.io/quote"

        "hello/morestrings"

        "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(quote.Go())
    fmt.Println(4 + 5)
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
