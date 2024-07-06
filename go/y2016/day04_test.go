package main

import (
	"testing"
)

func TestDay04(t *testing.T) {
	testDay(t, day04{}, []TC{
		TC{
			"1514",
			`aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`,
		},
	}, []TC{})
}
