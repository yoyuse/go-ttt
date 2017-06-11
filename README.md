# go-ttt

`go-ttt` is a Tiny TT-Code Translation program for CLI, written in Go.

## Usage

``` console
$ go-ttt [options] [--] [string ...]
```

`go-ttt` decodes each T-Code strings given as command line arguments to Japanese text.
If no argument is provided, `go-ttt` decodes standard input.

Some examples:

``` console
$ go-ttt 'ysksjsks/ajgjdjfkdt8p3;gpzjshdjtighdiakslghdhgia'
わたしたちは、氷砂糖をほしいくらいもたないでも

$ echo 'yfkd% Iha-Tovo kd,fhrjaoajrksqr%' | go-ttt -m %
あの Iha-Tovo のすきとおった風

$ go-ttt -m % <<< 'yd.djtjshdjfoxhgw7ig;eks% Morio:/v%'
うつくしい森で飾られた Morio市
```

Decode is done as well as [ttt](https://github.com/yoyuse/ttt);
that is, `go-ttt` scans string from its tail to head,
finds first valid T-Code string and decode it to Japanese.

When `-m` *marker* option is specified,
`go-ttt` does ttt decode at each point where *marker* appears in given strings.

### Command line options

| Option | Meaning |
|--------|---------|
| `-m` *marker* | Decode at each occurrence of *marker* |
| `-n` | Do not output newline (when echo-mode) |
| `-w` | Decode whole string rather than à la ttt |

### Marker string

Marker string can include following sequence:

| Sequence | Meaning |
|----------|---------|
| `\a` | Bell |
| `\b` | BackSpace|
| `\e` | Escape |
| `\f` | Form feed |
| `\n` | Newline |
| `\r` | Carriage return |
| `\t` | Tab |
| `\\` | Backslash |
| `\`*ooo* | Octal character |
| `\x`*hh* | Hexadecimal character |
| `\c`*C* | Control character |

For example, `$ go-ttt -m '\ej'` sets marker to `Esc j`,
which can be input by <kbd>Alt</kbd>+<kbd>j</kbd> in some terminal.
Either `'\033j'`, `'\x1bj'` or `'\c[j'` works as well.

## Sample configurations

Following settings enable ttt Japanse input with <kbd>Alt</kbd>+<kbd>j</kbd>
on command line of interactive shell.

### Bash

Bash 4 or later is required.

``` shell
function bash-ttt() {
    local src=$(printf "%s" "$READLINE_LINE" | cut -b 1-$READLINE_POINT)
    local rest=$(printf "%s" "$READLINE_LINE" | cut -b $(expr $READLINE_POINT + 1)-)
    local dst=$(go-ttt <<< "$src")
    READLINE_LINE="$dst$rest"
    READLINE_POINT=$(expr $(printf "%s" "$dst" | wc -c))
}

bind -x '"\ej": bash-ttt'
```

### Zsh

``` shell
function zsh-ttt() {
    local rbuf="${RBUFFER}"
    BUFFER=$(go-ttt <<< "${LBUFFER}")
    CURSOR=$#BUFFER
    BUFFER="$BUFFER${rbuf}"
}

zle -N zsh-ttt
bindkey '\ej' zsh-ttt
```

## License

UNDEFINED © yoyuse
