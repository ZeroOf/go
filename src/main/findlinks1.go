package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main()  {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprint(os.Stderr, "findlinks1: %v \n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
