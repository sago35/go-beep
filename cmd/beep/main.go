package main

import (
	"fmt"
	"github.com/sago35/go-beep"
)

func main() {
	score := `c2. d4 e2. c4 e2 c e1`

	fmt.Printf("%v\n", score)
	beep.Play(score)
}
