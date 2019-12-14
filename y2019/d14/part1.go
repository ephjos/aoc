// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Chemical struct {
	amount int
	name   string
}

type Reaction struct {
	inputs []Chemical
	output Chemical
}

type Queue struct {
	arr []string
}

func (q *Queue) Enqueue(s string) {
	q.arr = append(q.arr, s)
}

func (q *Queue) Dequeue() string {
	temp := q.arr[0]
	q.arr = q.arr[1:]
	return temp
}

func (q *Queue) IsEmpty() bool {
	if len(q.arr) == 0 {
		return true
	}
	return false
}

func parseChemical(obj string) Chemical {
	tokens := strings.Split(obj, " ")
	amount, err := strconv.Atoi(tokens[0])

	if err != nil {
		panic(err)
	}

	name := tokens[1]

	chemical := Chemical{
		amount,
		name,
	}

	return chemical
}

func parseChemicalList(list string) []Chemical {
	tokens := strings.Split(list, ", ")

	chemicals := make([]Chemical, 0)
	for _, tok := range tokens {
		chemical := parseChemical(tok)
		chemicals = append(chemicals, chemical)
	}

	return chemicals
}

func parseReaction(line string) Reaction {
	tokens := strings.Split(line, "=> ")
	inputs := parseChemicalList(tokens[0])
	output := parseChemical(tokens[1])

	r := Reaction{inputs, output}
	return r
}

// https://gitlab.com/FrankEndriss/aoc2019/blob/master/src/d14.cc

func getOreCount(chemicalName string,
	reactionMap map[string][]Chemical) int {
	Q := &Queue{make([]string, 0)}

	discovered := make(map[string]bool)
	discovered["FUEL"] = true

	Q.Enqueue("FUEL")

	for !Q.IsEmpty() {
		v := Q.Dequeue()

		if len(discovered) == len(reactionMap) {
			break
		}

		for _, c := range reactionMap[v][1:] {
			fmt.Println(c)
			discovered[c.name] = true
			Q.Enqueue(c.name)
		}

	}

	return 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	reactions := make([]Reaction, 0)
	reactionMap := make(map[string][]Chemical, 0)

	for input.Scan() {
		line := input.Text()

		reaction := parseReaction(line)
		reactions = append(reactions, reaction)

		reactionMap[reaction.output.name] = append(
			reactionMap[reaction.output.name],
			append([]Chemical{reaction.output},
				reaction.inputs...)...,
		)
	}

	fmt.Println(getOreCount("FUEL", reactionMap))
}
