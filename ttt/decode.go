// 2017-06-07

package ttt

import (
	"strings"
)

func Decode_string(str string) string {
	code := strings.Split(str, "")
	dst := ""
	var t elm = table
	for _, ch := range code {
		k := strings.Index(keys, ch)
		switch k {
		case -1:
			t = table
			dst += ch
		default:
			u, ok := t.(tbl)
			// slice (bounds) check
			if !ok || len(u) <= k {
				t = table
				dst += ch
				break
			}
			t = u[k]
			switch val := t.(type) {
			case nil:
				t = table
			case string:
				dst += val
				t = table
			}
		}
	}
	return dst
}

func Decode_substring(str string) string {
	code := strings.Split(str, "")
	var ch, tail, body, head string
	i := len(code) - 1
	for 0 <= i {
		ch = code[i]
		if 0 <= strings.Index(keys, ch) {
			break
		}
		tail = ch + tail
		i -= 1
	}
	for 0 <= i {
		ch = code[i]
		if strings.Index(keys, ch) < 0 {
			break
		}
		body = ch + body
		i -= 1
	}
	if ch == delimiter {
		i -= 1
	}
	for 0 <= i {
		ch = code[i]
		head = ch + head
		i -= 1
	}
	return head + Decode_string(body) + tail
}

func Decode_at_marker(str string) string {
	// ;dMARKER;fMARKERhaMARKER -> 岳3
	for marker != "" {			// if marker != "" {for {...}}
		i := strings.Index(str, marker)
		if i < 0 {
			break
		}
		src := str[:i]
		str = Decode_substring(src) + str[(i+len(marker)):]
	}
	return str
}
