package main

import (
	"testing"
)

func TestDay02(t *testing.T) {
	testDay(t, day02{}, []TC{
		TC{
			"1985",
			`ULL
RRDDD
LURDL
UUUUD`,
		},
	}, []TC{
		TC{
			"5DB3",
			`ULL
RRDDD
LURDL
UUUUD`,
		},
	})
}
