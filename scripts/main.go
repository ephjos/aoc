package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

var tmpl string = `
| -    | *Part 1*     | *Part 2*      |
| :--: | :--:         | :--:          |
| Time | {{ .Time1 }} | {{ .Time2 }}  |
| Rank | {{ .Rank1 }} | {{ .Rank2 }}  |
`

type Stats struct {
	Time1, Time2, Rank1, Rank2 string
}

func main() {
	data, err := ioutil.ReadFile("./input")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	line := strings.Split(strings.TrimSpace(lines[2]), " ")
	var values []string

	for _, v := range line {
		clean := strings.TrimSpace(v)
		if clean != "" {
			values = append(values, clean)
		}
	}

	var stats Stats = Stats{
		values[1], values[4], values[2], values[5],
	}

	var buffer bytes.Buffer

	tpl := template.Must(template.New("").Parse(tmpl))
	tpl.Execute(&buffer, stats)

	fmt.Println(buffer.String())
}
