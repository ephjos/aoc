package main

import (
	"testing"
)

func TestDay06(t *testing.T) {
	testDay(t, day06{}, []TC{
		TC{
			"easter",
			`eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar
`,
		},
	}, []TC{
		TC{
			"advent",
			`eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar
`,
		},
	})
}
