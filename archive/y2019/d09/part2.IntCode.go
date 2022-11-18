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
	OP_ADD  int64 = 1
	OP_MULT int64 = 2

	OP_INPUT  int64 = 3
	OP_OUTPUT int64 = 4

	OP_JUMP_IF_TRUE  int64 = 5
	OP_JUMP_IF_FALSE int64 = 6

	OP_LESS_THAN int64 = 7
	OP_EQUALS    int64 = 8

	OP_REL_OFFSET int64 = 9

	OP_HALT int64 = 99

	MODE_POSITION  int64 = 0
	MODE_IMMEDIATE int64 = 1
	MODE_RELATIVE  int64 = 2
)

type IntCode struct {
	Data []int64
}

func parseOpCode(opcode int64) (op int64, m1 int64, m2 int64, m3 int64) {
	op = opcode % 100
	m1 = opcode % 1000 / 100
	m2 = opcode % 10000 / 1000
	m3 = opcode % 100000 / 10000
	return
}

type ComputeState struct {
	index, opcode, relativeBase int64
}

func (ic IntCode) getValue(state *ComputeState, mode, position, value int64) int64 {

	if ((state.opcode == 1 || state.opcode == 2 || state.opcode == 7 ||
		state.opcode == 8) && position == 3) ||
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

func (ic *IntCode) write(dest, value int64) {
	if dest >= int64(len(ic.Data)) {
		temp := make([]int64, dest*2)
		copy(temp, ic.Data)
		ic.Data = temp
	}

	ic.Data[dest] = value
}

func (ic IntCode) Compute(ch chan int64) {
	relativeBase := int64(0)
	for i := int64(0); i < int64(len(ic.Data)); {
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

			ic.write(v3, v1+v2)

			i += 4
			break
		case OP_MULT:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(state, m1, 1, p1)
			v2 := ic.getValue(state, m2, 2, p2)
			v3 := ic.getValue(state, m3, 3, p3)

			ic.write(v3, v1*v2)

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

			ic.write(v1, <-ch)

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

func New(rawData []string) *IntCode {
	var Data []int64

	for _, v := range rawData {
		t, _ := strconv.Atoi(v)
		Data = append(Data, int64(t))
	}

	return &IntCode{Data}
}
