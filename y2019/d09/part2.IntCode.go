// Part 1

package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

var m *sync.Mutex = &sync.Mutex{}

const (
	OP_ADD  int = 1
	OP_MULT int = 2

	OP_INPUT  int = 3
	OP_OUTPUT int = 4

	OP_JUMP_IF_TRUE  int = 5
	OP_JUMP_IF_FALSE int = 6

	OP_LESS_THAN int = 7
	OP_EQUALS    int = 8

	OP_REL_OFFSET int = 9

	OP_HALT int = 99

	MODE_POSITION  int = 0
	MODE_IMMEDIATE int = 1
	MODE_RELATIVE  int = 2
)

type IntCode struct {
	Data []int
}

func parseOpCode(opcode int) (op int, m1 int, m2 int, m3 int) {
	op = opcode % 100
	m1 = opcode % 1000 / 100
	m2 = opcode % 10000 / 1000
	m3 = opcode % 100000 / 10000
	return
}

type ComputeState struct {
	index, opcode, relativeBase int
}

func (ic IntCode) getValue(state *ComputeState, mode, position, value int) int {

	if ((state.opcode == 1 || state.opcode == 2 || state.opcode == 7 || state.opcode == 8) && position == 3) ||
		(state.opcode == 3 && position == 1) {
		switch mode {
		case MODE_POSITION, MODE_IMMEDIATE:
			return value
		case MODE_RELATIVE:
			return state.relativeBase + value
		default:
			panic(fmt.Sprintf("Mode: %d,Param: %d\n", mode, value))
		}
	}

	switch mode {
	case MODE_POSITION:
		return ic.Data[value]
	case MODE_IMMEDIATE:
		return value
	case MODE_RELATIVE:
		return ic.Data[state.relativeBase+value]
	default:
		panic(fmt.Sprintf("Mode: %d,Param: %d\n", mode, value))
	}

}

func (ic IntCode) Compute(ch chan int) {
	// Pad memory
	PAD_LENGTH := 2000000
	PADDING := make([]int, PAD_LENGTH)
	ic.Data = append(ic.Data, PADDING...)

	relativeBase := 0
	for i := 0; i < len(ic.Data); {
		opcode, m1, m2, m3 := parseOpCode(ic.Data[i])
		state := &ComputeState{
			i, opcode, relativeBase,
		}

		if opcode == OP_HALT {
			break
		}

		switch opcode {
		case OP_ADD:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(state, m1, 1, p1)
			v2 := ic.getValue(state, m2, 2, p2)
			v3 := ic.getValue(state, m3, 3, p3)

			ic.Data[v3] = v1 + v2

			i += 4
			break
		case OP_MULT:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(state, m1, 1, p1)
			v2 := ic.getValue(state, m2, 2, p2)
			v3 := ic.getValue(state, m3, 3, p3)

			ic.Data[v3] = v1 * v2

			i += 4
			break
		case OP_INPUT:
			p1 := ic.Data[i+1]
			v1 := ic.getValue(state, m1, 1, p1)
			log.Print("Waiting for input...")

			// I don't think this should fix the race condition.
			// I mean, it should just toggle the lock...
			// But without it go -race thinks there is a race
			// condition, so it stays.
			m.Lock()
			m.Unlock()

			ic.Data[v1] = <-ch

			log.Print(fmt.Sprintf("Input %d received.\n", ic.Data[v1]))

			i += 2
			break
		case OP_OUTPUT:
			p1 := ic.Data[i+1]
			v1 := ic.getValue(state, m1, 1, p1)

			ch <- v1

			i += 2
			break
		case OP_JUMP_IF_TRUE:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]

			v1 := ic.getValue(state, m1, 1, p1)
			v2 := ic.getValue(state, m2, 2, p2)

			if v1 != 0 {
				i = v2
			} else {
				i += 3
			}
			break
		case OP_JUMP_IF_FALSE:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]

			v1 := ic.getValue(state, m1, 1, p1)
			v2 := ic.getValue(state, m2, 2, p2)

			if v1 == 0 {
				i = v2
			} else {
				i += 3
			}
			break
		case OP_LESS_THAN:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(state, m1, 1, p1)
			v2 := ic.getValue(state, m2, 2, p2)
			v3 := ic.getValue(state, m3, 3, p3)

			if v1 < v2 {
				ic.Data[v3] = 1
			} else {
				ic.Data[v3] = 0
			}

			i += 4
			break
		case OP_EQUALS:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(state, m1, 1, p1)
			v2 := ic.getValue(state, m2, 2, p2)
			v3 := ic.getValue(state, m3, 3, p3)

			if v1 == v2 {
				ic.Data[v3] = 1
			} else {
				ic.Data[v3] = 0
			}

			i += 4
			break
		case OP_REL_OFFSET:
			p1 := ic.Data[i+1]
			v1 := ic.getValue(state, m1, 1, p1)

			relativeBase += v1

			i += 2
			break
		default:
			break
		}
	}
	close(ch)
}

func MakeIntCode(rawData []string) *IntCode {
	var Data []int

	for _, v := range rawData {
		t, _ := strconv.Atoi(v)
		Data = append(Data, t)
	}

	return &IntCode{Data}
}
