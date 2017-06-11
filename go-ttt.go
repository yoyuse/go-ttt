// 2017-06-06 new type elm, tbl instead of element. slice (bounds) check
// 2017-06-05

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"rose.local/yuse/go-ttt/ttt"
)

func do_ttt(args []string, opt_n bool, opt_Z bool) {
	dst := ""
	len := len(args)
	for i, str := range args {
		if opt_Z {
			dst += ttt.Decode_substring(str)
		} else {
			dst += ttt.Decode_at_marker(str)
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

func do_ttt_stdin(opt_n bool, opt_Z bool) {
	r := bufio.NewReader(os.Stdin)
	for {
		// var str string
		str, err := r.ReadString('\n')
		// XXX: output here, for case that str does not end with '\n'
		var dst string
		if opt_Z {
			dst = ttt.Decode_substring(str)
		} else {
			dst = ttt.Decode_at_marker(str)
		}
		fmt.Print(dst)
		if err == io.EOF {
			break
		}
	}
}

func main() {
	opt_n := flag.Bool("n", false, "no newline")
	opt_Z := flag.Bool("Z", false, "implicit ttt at end of each str or line")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		do_ttt_stdin(*opt_n, *opt_Z)
	} else {
		do_ttt(args, *opt_n, *opt_Z)
	}
}
