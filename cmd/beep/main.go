package main

import (
	"fmt"
	"github.com/sago35/go-beep"
)

func main() {
	score := `e c d a e c d a`

	fmt.Printf(score)
	beep.Play(score)
}
