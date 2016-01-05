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

func Play(score string, bpm int) {
	re := regexp.MustCompile(`([a-gr])(is|es)?([',]?)(\d*)(\.?)`)
	octave := 1.0
	duration := 4

	for _, n := range strings.Split(score, " ") {
		var alteration float32 = 1.0
		match := re.FindAllStringSubmatch(n, -1)
		fmt.Println(match)
		if match[0][3] == `'` {
			octave = 2.0
		} else if match[0][3] == `,` {
			octave = 0.5
		} else {
			octave = 1.0
		}

		if match[0][2] == `is` {
			alteration = 0.94
		} else if match[0][2] == `es` {
			alteration = 1.06
		}

		if match[0][4] != `` {
			d, err := strconv.Atoi(match[0][4])
			if err != nil {
			} else {
				if match[0][5] == `.` {
					duration = d * 3 / 4
				} else {
					duration = d
				}
			}
		}

		timer := time.After(time.Duration(_duration(duration, bpm)) * time.Millisecond)
		if match[0][1] != `r` {
			beep(int(float32(note2freq[match[0][1]])*float32(octave)*alteration), int(_duration(duration, bpm) * 0.9))
		}
		<-timer
	}
}

func _duration(duration, bpm int) float64 {
	return 60.0 * 1000.0 / float64(bpm) * (4.0 / float64(duration))
}
