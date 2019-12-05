// Part 1

package main

import (
	"fmt"
	"strconv"
)

const (
	OP_ADD  int = 1
	OP_MULT int = 2

	OP_INPUT  int = 3
	OP_OUTPUT int = 4

	OP_JUMP_IF_TRUE  int = 5
	OP_JUMP_IF_FALSE int = 6

	OP_LESS_THAN int = 7
	OP_EQUALS    int = 8

	OP_HALT int = 99

	MODE_POSITION  int = 0
	MODE_IMMEDIATE int = 1
)

type IntCode struct {
	data []int
}

func parseOpCode(opcode int) (op int, m1 int, m2 int, m3 int) {
	op = opcode % 100
	m1 = opcode % 1000 / 100
	m2 = opcode % 10000 / 1000
	m3 = opcode % 100000 / 10000
	return
}

func getValue(data []int, mode, p int) int {
	switch mode {
	case MODE_POSITION:
		return data[p]
	case MODE_IMMEDIATE:
		return p
	default:
		panic(fmt.Sprintf("Mode: %d,Param: %d\n", mode, p))
	}

}

func (ic IntCode) Compute(input int) {
	for i := 0; i < len(ic.data); {
		opcode, m1, m2, m3 := parseOpCode(ic.data[i])
		m3 = m3

		if opcode == OP_HALT {
			break
		}

		switch opcode {
		case OP_ADD:
			p1 := ic.data[i+1]
			p2 := ic.data[i+2]
			p3 := ic.data[i+3]

			v1 := getValue(ic.data, m1, p1)
			v2 := getValue(ic.data, m2, p2)
			ic.data[p3] = v1 + v2

			i += 4
			break
		case OP_MULT:
			p1 := ic.data[i+1]
			p2 := ic.data[i+2]
			p3 := ic.data[i+3]

			v1 := getValue(ic.data, m1, p1)
			v2 := getValue(ic.data, m2, p2)

			ic.data[p3] = v1 * v2

			i += 4
			break
		case OP_INPUT:
			p1 := ic.data[i+1]

			ic.data[p1] = input

			i += 2
			break
		case OP_OUTPUT:
			p1 := ic.data[i+1]

			v1 := getValue(ic.data, m1, p1)
			fmt.Println(v1)

			i += 2
			break
		case OP_JUMP_IF_TRUE:
			p1 := ic.data[i+1]
			p2 := ic.data[i+2]

			v1 := getValue(ic.data, m1, p1)
			v2 := getValue(ic.data, m2, p2)

			if v1 != 0 {
				i = v2
			} else {
				i += 3
			}
			break
		case OP_JUMP_IF_FALSE:
			p1 := ic.data[i+1]
			p2 := ic.data[i+2]

			v1 := getValue(ic.data, m1, p1)
			v2 := getValue(ic.data, m2, p2)

			if v1 == 0 {
				i = v2
			} else {
				i += 3
			}
			break
		case OP_LESS_THAN:
			p1 := ic.data[i+1]
			p2 := ic.data[i+2]
			p3 := ic.data[i+3]

			v1 := getValue(ic.data, m1, p1)
			v2 := getValue(ic.data, m2, p2)

			if v1 < v2 {
				ic.data[p3] = 1
			} else {
				ic.data[p3] = 0
			}

			i += 4
			break
		case OP_EQUALS:
			p1 := ic.data[i+1]
			p2 := ic.data[i+2]
			p3 := ic.data[i+3]

			v1 := getValue(ic.data, m1, p1)
			v2 := getValue(ic.data, m2, p2)

			if v1 == v2 {
				ic.data[p3] = 1
			} else {
				ic.data[p3] = 0
			}

			i += 4
			break
		default:
			break
		}
	}
}

func MakeIntCode(rawData []string) *IntCode {
	var data []int

	for _, v := range rawData {
		t, _ := strconv.Atoi(v)
		data = append(data, t)
	}

	return &IntCode{data}
}
