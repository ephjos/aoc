// Part 2

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
var store = make(map[string]int, 0)

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

func getOreCount(name string, needed int) int {
	if name == "ORE" {
		return needed
	}

	stored := store[name]
	if stored > 0 {
		if stored >= needed {
			store[name] -= needed
			return 0
		} else {
			needed -= stored
			store[name] = 0
			return getOreCount(name, needed)
		}
	}

	canMake := reactionMap[name][0].amount

	batches := int(math.Ceil(float64(needed) / float64(canMake)))
	excess := (canMake * batches) - needed

	store[name] = excess

	oreNeeded := 0
	for _, ingredient := range reactionMap[name][1:] {
		oreNeeded += getOreCount(ingredient.name, batches*ingredient.amount)
	}

	return oreNeeded
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

	ore := 0
	i := 3840000
	tril := 1000000000000
	for ore < tril {
		store = make(map[string]int, 0)
		ore = getOreCount("FUEL", i)
		fmt.Printf("i=%d ore=%d\n", i, ore)

		i += 1
	}

	i -= 2
	fmt.Printf("\ni=%d\n", i)
}
