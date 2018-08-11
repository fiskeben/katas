package main

import (
	"fmt"
)

func main() {
	g := makeGenerator()
	b := g.MakeRandomBoard()
	fmt.Println(b)
}
