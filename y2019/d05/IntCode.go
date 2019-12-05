// Part 1

package main

import (
	"fmt"
	"strconv"
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
	POSITION := 0
	IMMEDIATE := 1

	switch mode {
	case POSITION:
		return data[p]
	case IMMEDIATE:
		return p
	default:
		panic(fmt.Sprintf("Mode: %d,Param: %d\n", mode, p))
	}

}

func (ic IntCode) Compute(input int) {
	for i := 0; i < len(ic.data); {
		opcode, m1, m2, m3 := parseOpCode(ic.data[i])
		m1 = m1
		m2 = m2
		m3 = m3
		if opcode == 99 {
			break
		}

		switch opcode {
		case 1:
			p1 := ic.data[i+1]
			p2 := ic.data[i+2]
			p3 := ic.data[i+3]

			v1 := getValue(ic.data, m1, p1)
			v2 := getValue(ic.data, m2, p2)
			ic.data[p3] = v1 + v2

			i += 4
			break
		case 2:
			p1 := ic.data[i+1]
			p2 := ic.data[i+2]
			p3 := ic.data[i+3]

			v1 := getValue(ic.data, m1, p1)
			v2 := getValue(ic.data, m2, p2)

			ic.data[p3] = v1 * v2

			i += 4
			break
		case 3:
			p1 := ic.data[i+1]

			ic.data[p1] = input

			i += 2
			break
		case 4:
			p1 := ic.data[i+1]

			v1 := getValue(ic.data, m1, p1)
			fmt.Println(v1)

			i += 2
			break
		case 5:
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
		case 6:
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
		case 7:
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
		case 8:
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
