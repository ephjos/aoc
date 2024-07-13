#!/bin/sh

if [ "$#" -ne 1 ]; then
  echo "Error, must provide argument NN like 03"
  exit 1
fi

# Write file
cat > "day$1.go" << EOF
package main

/*
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
*/

import (
)

type day$1 struct{}

func (_d day$1) a(input string) string {
  return "0"
}

func (_d day$1) b(input string) string {
  return "0"
}
EOF

# Write test file
cat > "day$1_test.go" << EOF
package main

import (
	"testing"
)

func TestDay$1(t *testing.T) {
	testDay(t, day$1{}, []TC{
	}, []TC{
	})
}
EOF
