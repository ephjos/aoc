// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Point struct {
	X, Y int
}

type DistanceOnWire = int

type Direction = string
type Length = int

func parseInstruction(instruction string) (Direction, Length) {
	dir := string(instruction[0])
	length, _ := strconv.Atoi(instruction[1:])
	return dir, length
}

func parseWire(instructions []string) map[Point]DistanceOnWire {
	output := make(map[Point]DistanceOnWire)
	currentPoint := Point{0, 0}
	totalDistance := 0

	for _, instruction := range instructions {
		direction, length := parseInstruction(instruction)
		for i := 0; i < length; i++ {
			switch direction {
			case "U":
				currentPoint.Y += 1
				break
			case "D":
				currentPoint.Y -= 1
				break
			case "L":
				currentPoint.X -= 1
				break
			case "R":
				currentPoint.X += 1
				break
			default:
				panic("Instruction parse error")
				break
			}

			totalDistance += 1
			output[currentPoint] = totalDistance
		}
	}

	return output
}

func getKeys(m map[Point]DistanceOnWire) []Point {
	keys := make([]Point, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func contains(a []Point, p Point) bool {
	for _, v := range a {
		if p == v {
			return true
		}
	}

	return false
}

func worker(wg *sync.WaitGroup, keys1, keys2 []Point, a, b int, channel chan Point) {
	defer wg.Done()
	for i := a; i < b; i++ {
		key := keys1[i]
		if contains(keys2, key) {
			channel <- key
		}
	}
}

// Using gorountines provided a ~3x speed up
func handleWires(wires *[2]map[Point]DistanceOnWire) {
	wire1 := (*wires)[0]
	keys1 := getKeys(wire1)
	wire2 := (*wires)[1]
	keys2 := getKeys(wire2)

	crosses := make(chan Point)
	wg := &sync.WaitGroup{}
	batch := 5000
	l := len(keys1)

	for i := 0; i < l; i += batch {
		wg.Add(1)
		if i+batch < l {
			go worker(wg, keys1, keys2, i, i+batch, crosses)
		} else {
			go worker(wg, keys1, keys2, i, l-1, crosses)
		}
	}

	go func() {
		wg.Wait()
		close(crosses)
	}()

	min := math.MaxInt32
	for point := range crosses {
		s1 := wire1[point]
		s2 := wire2[point]
		steps := s1 + s2
		if steps < min {
			min = steps
		}
	}
	fmt.Println(min)
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	wires := [2]map[Point]DistanceOnWire{}
	count := 0

	for input.Scan() {
		line := input.Text()
		if line != "" {
			tokens := strings.Split(line, ",")
			wires[count] = parseWire(tokens)
		}
		count += 1
	}

	handleWires(&wires)

}
