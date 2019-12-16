// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reactionMap = make(map[string][]Chemical, 0)
var excess = make(map[string]int, 0)

type Chemical struct {
	amount int
	name   string
}

type Reaction struct {
	inputs []Chemical
	output Chemical
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

func getOreCount(name string, amount int) int {

	if name == "ORE" {
		return amount
	}

	count := int(math.Min(float64(excess[name]), float64(amount)))
	amount -= count
	excess[name] -= count

	sub := 0
	for _, ingredient := range reactionMap[name][1:] {
		sub += getOreCount(ingredient.name, 1)
	}

	fmt.Println(name, count, amount, sub)

	return sub
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	reactions := make([]Reaction, 0)

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

	fmt.Println(getOreCount("FUEL", 1))
}
