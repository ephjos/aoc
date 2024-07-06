package main

import (
	//"fmt"
	"testing"
)

func TestDay05(t *testing.T) {
	testDay(t, day05{}, []TC{
		TC{
			"18f47a30",
			"abc",
		},
	}, []TC{
		TC{
			"05ace8e3",
			"abc",
		},
	})
}

/*
func BenchmarkDay05aBlockSize(b *testing.B) {
	inputs := get_inputs()

	b.ResetTimer()

  for i := 8; i <= 36; i+=2 {
		name := fmt.Sprintf("i=%d", i)
		b.Run(name, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				day05_a_core(inputs[4], 1 << i)
			}
		})
  }
}

func BenchmarkDay05bBlockSize(b *testing.B) {
	inputs := get_inputs()

	b.ResetTimer()

  for i := 8; i <= 36; i+=2 {
		name := fmt.Sprintf("i=%d", i)
		b.Run(name, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				day05_b_core(inputs[4], 1 << i)
			}
		})
  }
}
*/
