package main

import (
	"testing"
)

func TestDay03(t *testing.T) {
	testDay(t, day03{}, []TC{
		TC{
			"0",
			"5 10 25\n",
		},
	}, []TC{})
}
