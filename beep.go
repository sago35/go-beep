package beep

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
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
	re := regexp.MustCompile(`([a-g])([',]?)(\d*)(\.?)`)
	octave := 1
	duration := 4
	bpm := 120

	for _, n := range strings.Split(score, " ") {
		match := re.FindAllStringSubmatch(n, -1)
		fmt.Println(match)
		if match[0][2] == `'` {
			octave = octave * 2
		} else if match[0][2] == `,` {
			octave = octave / 2
		}

		if match[0][3] != `` {
			d, err := strconv.Atoi(match[0][3])
			if err != nil {
			} else {
				if match[0][4] == `.` {
					duration = d * 3 / 4
				} else {
					duration = d
				}
			}
		}
		beep(note2freq[match[0][1]]*octave, int(_duration(duration, bpm)))
	}
}

func _duration(duration, bpm int) float64 {
	return 60.0 * 1000.0 / float64(bpm) * (4.0 / float64(duration))
}
