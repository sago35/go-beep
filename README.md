# go-beep
beeper for windows

PCのbeepを鳴らすためのライブラリを作成していく予定

# build

```
$ go build ./cmd/beep
```

# usage

```
usage: beep [<flags>] [<score>...]

Flags:
  -h, --help     Show context-sensitive help (also try --help-long and
                 --help-man).
      --version  Print version information and quit
      --bpm=120  Change the tempo of the music

Args:
  [<score>]  Input score
```

