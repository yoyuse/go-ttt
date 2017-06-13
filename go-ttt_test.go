package main

import (
	"testing"
)

func ExampleDecode_string() {
	do_ttt([]string{"kdjd;sjdkd;s", "lgkg;gjslghdhghdkshdkd"})
	// Output:
	// の、が、のが なにをしないでいたいの
}

type stringTest struct {
	in, out string
}

var Tests0 = []stringTest{
	{"abefnrt", "abefnrt"},
	{"\\a\\b\\e\\f\\n\\r\\t", "\a\b\033\f\n\r\t"},
	{"\\\\", "\\"},
	{"[\\033j]", "[\x1bj]"},
	{"\\400", "\040" + "0"},
	{"\\xa", "\x0a"},
	{"\\cM", "\r"},
	{"\\0", "\x00"},
	{"bad \\", "bad \\"},
	{"bad \\z", "bad z"},
	{"bad \\x", "bad \x00"},
	{"bad \\c", "bad c"},
}

func Test_Unbackslash(t *testing.T) {
	for _, ts := range Tests0 {
		unbs := unbackslash(ts.in)
		if unbs != ts.out {
			t.Errorf("unbackslash: %s == %q, want %q", ts.in, unbs, ts.out)
		}
	}
}
