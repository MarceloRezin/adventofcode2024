package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("./../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	memoriaCorrompida := string(content)

	mulRegex, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	dontParts := strings.Split(memoriaCorrompida, "don't()")

	var opsToExecute []string
	for i, p := range dontParts {
		if i == 0 {
			opsToExecute = append(opsToExecute, mulRegex.FindAllString(p, -1)...)
			continue
		}

		doParts := strings.Split(p, "do()")
		for j := 1; j < len(doParts); j++ {
			opsToExecute = append(opsToExecute, mulRegex.FindAllString(doParts[j], -1)...)
		}

	}

	total := 0
	for _, op := range opsToExecute {

		opDecomposta := strings.Split(op, ",")

		x, _ := strconv.Atoi(strings.Replace(opDecomposta[0], "mul(", "", -1))
		y, _ := strconv.Atoi(strings.Replace(opDecomposta[1], ")", "", -1))

		total += x * y
	}

	println(total)
}
