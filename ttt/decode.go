package ttt

import (
	"strings"
)

func DecodeString(str string) string {
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

func DecodeSubstring(str string) string {
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
	return head + DecodeString(body) + tail
}

func DecodeAtMarker(str, marker string) string {
	for marker != "" { // if marker != "" {for {...}}
		i := strings.Index(str, marker)
		if i < 0 {
			break
		}
		src := str[:i]
		str = DecodeSubstring(src) + str[(i+len(marker)):]
	}
	return str
}
