// Part 1

package main

import (
	"fmt"
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

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{[]int{}}
}

func (s *Queue) Push(i int) {
	s.data = append(s.data, i)
	return
}

func (s *Queue) Pop() int {
	temp := s.data[0]
	s.data = s.data[1:]
	return temp
}

type IntCode struct {
	Data      []int64
	IP        int64
	RelBase   int64
	IsRunning bool
	Inputs    *Queue
}

func (ic *IntCode) Copy() *IntCode {
	temp := &IntCode{
		Data:      make([]int64, len(ic.Data)),
		IP:        ic.IP,
		RelBase:   ic.RelBase,
		IsRunning: ic.IsRunning,
	}
	copy(temp.Data, ic.Data)

	return temp
}

func parseOpCode(opcode int64) (op int64, m1 int64, m2 int64, m3 int64) {
	op = opcode % 100
	m1 = opcode % 1000 / 100
	m2 = opcode % 10000 / 1000
	m3 = opcode % 100000 / 10000
	return
}

func (ic IntCode) getValue(opcode, mode, position, value int64) int64 {

	if ((opcode == 1 || opcode == 2 || opcode == 7 ||
		opcode == 8) && position == 3) ||
		(opcode == 3 && position == 1) {
		switch mode {
		case MODE_POSITION, MODE_IMMEDIATE:
			return value
		case MODE_RELATIVE:
			return ic.RelBase + value
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
		return ic.Data[ic.RelBase+value]
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

func (ic *IntCode) AddInput(i int) {
	ic.Inputs.Push(i)
	return
}

func (ic *IntCode) Run() int64 {
	for {
		i := ic.IP
		opcode, m1, m2, m3 := parseOpCode(ic.Data[i])

		if opcode == OP_HALT {
			ic.IsRunning = false
			break
		}

		switch opcode {
		case OP_ADD:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(opcode, m1, 1, p1)
			v2 := ic.getValue(opcode, m2, 2, p2)
			v3 := ic.getValue(opcode, m3, 3, p3)

			ic.write(v3, v1+v2)

			ic.IP += 4
			break
		case OP_MULT:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(opcode, m1, 1, p1)
			v2 := ic.getValue(opcode, m2, 2, p2)
			v3 := ic.getValue(opcode, m3, 3, p3)

			ic.write(v3, v1*v2)

			ic.IP += 4
			break
		case OP_INPUT:
			p1 := ic.Data[i+1]
			v1 := ic.getValue(opcode, m1, 1, p1)

			ic.write(v1, int64(ic.Inputs.Pop()))

			ic.IP += 2
			break
		case OP_OUTPUT:
			p1 := ic.Data[i+1]
			v1 := ic.getValue(opcode, m1, 1, p1)

			ic.IP += 2
			return v1
		case OP_JUMP_IF_TRUE:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]

			v1 := ic.getValue(opcode, m1, 1, p1)
			v2 := ic.getValue(opcode, m2, 2, p2)

			if v1 != 0 {
				ic.IP = v2
			} else {
				ic.IP += 3
			}
			break
		case OP_JUMP_IF_FALSE:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]

			v1 := ic.getValue(opcode, m1, 1, p1)
			v2 := ic.getValue(opcode, m2, 2, p2)

			if v1 == 0 {
				ic.IP = v2
			} else {
				ic.IP += 3
			}
			break
		case OP_LESS_THAN:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(opcode, m1, 1, p1)
			v2 := ic.getValue(opcode, m2, 2, p2)
			v3 := ic.getValue(opcode, m3, 3, p3)

			if v1 < v2 {
				ic.Data[v3] = 1
			} else {
				ic.Data[v3] = 0
			}

			ic.IP += 4
			break
		case OP_EQUALS:
			p1 := ic.Data[i+1]
			p2 := ic.Data[i+2]
			p3 := ic.Data[i+3]

			v1 := ic.getValue(opcode, m1, 1, p1)
			v2 := ic.getValue(opcode, m2, 2, p2)
			v3 := ic.getValue(opcode, m3, 3, p3)

			if v1 == v2 {
				ic.Data[v3] = 1
			} else {
				ic.Data[v3] = 0
			}

			ic.IP += 4
			break
		case OP_REL_OFFSET:
			p1 := ic.Data[i+1]
			v1 := ic.getValue(opcode, m1, 1, p1)

			ic.RelBase += v1
			ic.IP += 2
			break
		default:
			break
		}
	}

	return 0
}

func NewIntCode(rawData []string) *IntCode {
	var Data []int64

	for _, v := range rawData {
		t, _ := strconv.Atoi(v)
		Data = append(Data, int64(t))
	}

	return &IntCode{
		Data:      Data,
		IP:        int64(0),
		RelBase:   int64(0),
		IsRunning: true,
		Inputs:    NewQueue(),
	}
}
