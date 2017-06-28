package ttt

import (
	"testing"
)

type stringTest struct {
	in, out string
}

var Tests1 = []stringTest{
	{"kdjd;sjdkd;s", "の、が、のが"},
	{"ABCDyfhdyd", "ABCDあいう"},
	{"abcd yfhdyd", "abcd あいう"},
	{"abcd:yfhdyd", "abcdあいう"},
	{"ysksjsks/ajgjdjfkdt8p3;gpzjshdjtighdiakslghdhgiajd", "わたしたちは、氷砂糖をほしいくらいもたないでも、"},
	{"あの Iha-Tovo kd,fhrjaoajrks風、", "あの Iha-Tovo のすきとおった風、"},
	{"うつくしい森で飾られた Morio:/vjd", "うつくしい森で飾られた Morio市、"},
	{"(またAladdin  lyfjlk[ラムプ]とり)", "(またAladdin  洋燈[ラムプ]とり)"},
	{";d;fha", "123"},
}

func TestDecodeSubstring(t *testing.T) {
	for _, ts := range Tests1 {
		dec := DecodeSubstring(ts.in)
		if dec != ts.out {
			t.Errorf("DecodeSubstring(%q) == %q, want %q", ts.in, dec, ts.out)
		}
	}
}

var Tests2 = []stringTest{
	{"kdjd;sjdkd;s\x1bj", "の、が、のが"},
	{"ABCDyfhdyd\x1bj", "ABCDあいう"},
	{"abcd yfhdyd\x1bj", "abcd あいう"},
	{"abcd:yfhdyd\x1bj", "abcdあいう"},
	{"ysksjsks/ajgjdjfkdt8p3;gpzjshdjtighdiakslghdhgiajd\x1bj", "わたしたちは、氷砂糖をほしいくらいもたないでも、"},
	{"yfkd\x1bj Iha-Tovo kd,fhrjaoajrks風、\x1bj", "あの Iha-Tovo のすきとおった風、"},
	{"うつくしい森で飾られた Morio:/vjd\x1bj", "うつくしい森で飾られた Morio市、"},
	{"(またAladdin  lyfjlk[usubmw]jajc)\x1bj\x1bj\x1bj", "(またAladdin  洋燈[ラムプ]とり)"},
	{";d\x1bj;f\x1bjha", "岳ha"},
}

func TestDecodeAtMarker(t *testing.T) {
	var marker = "\x1bj"
	for _, ts := range Tests2 {
		dec := DecodeAtMarker(ts.in, marker)
		if dec != ts.out {
			t.Errorf("DecodeAtMarker(%q) == %q, want %q", ts.in, dec, ts.out)
		}
	}
}
