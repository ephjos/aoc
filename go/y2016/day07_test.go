package main

import (
	"testing"
)

func TestDay07(t *testing.T) {
	testDay(t, day07{}, []TC{
		TC{
			"2",
			`abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyui
ioxxoj[asdfgh]zxcvbn`,
		},
	}, []TC{
		TC{
			"3",
			`aba[bab]xyz
xyx[xyx]xyx
aaa[kek]eke
zazbz[bzb]cdb`,
		},
	})
}
