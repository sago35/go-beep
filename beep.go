package beep

import (
	"fmt"
	"regexp"
	"strings"
	"syscall"
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

func Play(score string) {
	re := regexp.MustCompile(`([a-g])([',]?)`)
	octave := 1

	for _, n := range strings.Split(score, " ") {
		match := re.FindAllStringSubmatch(n, -1)
		fmt.Println(match)
		if match[0][2] == `'` {
			octave = octave * 2
		} else if match[0][2] == `,` {
			octave = octave / 2
		}
		beep(note2freq[match[0][1]] * octave, 130)
	}
}

