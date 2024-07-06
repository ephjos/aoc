package main

import (
	"testing"
)

func TestDay01(t *testing.T) {
	testDay(t, day01{}, []TC{
		TC{"5", "R2, L3"},
		TC{"2", "R2, R2, R2"},
		TC{"12", "R5, L5, R5, R3"},
	}, []TC{
		TC{"4", "R8, R4, R4, R8"},
	})
}
