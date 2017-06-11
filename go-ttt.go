// 2017-06-06 new type elm, tbl instead of element. slice (bounds) check
// 2017-06-05

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"rose.local/yuse/go-ttt/ttt"
)

func do_ttt(args []string, opt_n bool, opt_w bool) {
	dst := ""
	len := len(args)
	for i, str := range args {
		if opt_w {
			dst += ttt.Decode_string(str)
		} else {
			dst += ttt.Decode_substring_or_at_marker(str)
		}
		if i == len-1 {
			if !opt_n {
				dst += "\n"
			}
		} else {
			dst += " "
		}
	}
	fmt.Print(dst)
}

func do_ttt_stdin(opt_n bool, opt_w bool) {
	r := bufio.NewReader(os.Stdin)
	for {
		// var str string
		str, err := r.ReadString('\n')
		// XXX: output here, for case that str does not end with '\n'
		var dst string
		if opt_w {
			dst = ttt.Decode_string(str)
		} else {
			dst = ttt.Decode_substring_or_at_marker(str)
		}
		fmt.Print(dst)
		if err == io.EOF {
			break
		}
	}
}

func unbackslash(str string) string {
	a := strings.Split(str, "")
	dst := ""
	for 0 < len(a) {
		var ch string
		ch, a = a[0], a[1:]	// shift
		switch {
		case ch == "\\" && 0 < len(a):
			ch, a = a[0], a[1:] // shift
			switch {
			case ch == "a": dst += "\a"
			case ch == "b": dst += "\b"
			case ch == "e": dst += "\033"
			case ch == "f": dst += "\f"
			case ch == "n": dst += "\n"
			case ch == "r": dst += "\r"
			case ch == "t": dst += "\t"
			case ch == "c":
				var n rune = 0	// XXX: default value
				if 0 < len(a) && a[0] <= "\x7f" {
					ch, a = a[0], a[1:]
					n = []rune(ch)[0] & 0x1f
					dst += fmt.Sprintf("%c", n)
				} else {
					// XXX
					dst += "c"
				}
			case ch == "x":
				var n int64 = 0	// XXX: default value
				if 0 < len(a) &&
					("0" <= a[0] && a[0] <="9" ||
					"a" <= strings.ToLower(a[0]) &&
					strings.ToLower(a[0]) <= "f") {
					ch, a = a[0], a[1:]
					m, _ := strconv.ParseInt(ch, 16, 8)
					n = m
				}
				if 0 < len(a) &&
					("0" <= a[0] && a[0] <="9" ||
					"a" <= strings.ToLower(a[0]) &&
					strings.ToLower(a[0]) <= "f") {
					ch, a = a[0], a[1:]
					m, _ := strconv.ParseInt(ch, 16, 8)
					n = n * 0x10 + m
				}
				dst += fmt.Sprintf("%c", n)
			case "0" <= ch && ch <= "7":
				var n int64 = 0	// XXX: default value
				m, _ := strconv.ParseInt(ch, 8, 8)
				n = m
				if 0 < len(a) && "0" <= a[0] && a[0] <= "7" {
					ch, a = a[0], a[1:]
					m, _ := strconv.ParseInt(ch, 8, 8)
					n = n * 010 + m
				}
				if 0 < len(a) && "0" <= a[0] && a[0] <= "7" && n < 040 {
					ch, a = a[0], a[1:]
					m, _ := strconv.ParseInt(ch, 8, 8)
					n = n * 010 + m
				}
				dst += fmt.Sprintf("%c", n)
			default:
				dst += ch
			}
		default:
			dst += ch
		}
	}
	return dst
}

func main() {
	opt_m := flag.String("m", "", "Decode at marker STRING")
	opt_n := flag.Bool("n", false, "no newline")
	opt_w := flag.Bool("w", false, "Decode whole string")
	flag.Parse()
	args := flag.Args()
	if (*opt_m != "") {
		ttt.Set_marker(unbackslash(*opt_m))
	}
	if len(args) == 0 {
		do_ttt_stdin(*opt_n, *opt_w)
	} else {
		do_ttt(args, *opt_n, *opt_w)
	}
}
