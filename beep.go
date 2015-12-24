package main

import (
	"fmt"
	"strings"
	"syscall"
	//"unsafe"
)

var (
	kernel32, _ = syscall.LoadLibrary("kernel32.dll")
	beep32, _   = syscall.GetProcAddress(kernel32, "Beep")
)

func beep(freq, duration int) {
	syscall.Syscall(
		uintptr(beep32),
		uintptr(2),
		uintptr(freq),
		uintptr(duration),
		0,
	)
}


var note2freq = map[string]int{
	`c`: 261,
	`d`: 293,
	`e`: 329,
	`f`: 349,
	`g`: 392,
	`a`: 440,
	`b`: 493,
}

func play(score string) {
	for _, n := range strings.Split(score, " ") {
		beep(note2freq[n] * 3, 130)
	}
}

func main() {
	defer syscall.FreeLibrary(kernel32)

	score := `e c d a e c d a`

	fmt.Printf(score)
	play(score)
}
