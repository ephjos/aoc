package main

import (
	"fmt"
	"os"
	"testing"
)

type TC struct {
	expected string
	input    string
}

func testDay(t *testing.T, d Day, a_cases []TC, b_cases []TC) {
	t.Run("a", func(t *testing.T) {
		for _, in := range a_cases {
			exp := in.expected
			res := d.a(in.input)
			if exp != res {
				t.Errorf("expected %s, got %s", exp, res)
			}
		}
	})

	t.Run("b", func(t *testing.T) {
		for _, in := range b_cases {
			exp := in.expected
			res := d.b(in.input)
			if exp != res {
				t.Errorf("expected %s, got %s", exp, res)
			}
		}
	})
}

func BenchmarkAll(b *testing.B) {
	inputs := get_inputs()

	os.Stdout, _ = os.Open(os.DevNull)
	b.ResetTimer()

	for i, day := range DAYS {
		b_a := fmt.Sprintf("day%02d a", i+1)
		b.Run(b_a, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				day.a(inputs[i])
			}
		})

		b_b := fmt.Sprintf("day%02d b", i+1)
		b.Run(b_b, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				day.b(inputs[i])
			}
		})
	}
}
