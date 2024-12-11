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

	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	operacoesValidas := r.FindAllString(memoriaCorrompida, -1)

	total := 0
	for _, op := range operacoesValidas {

		opDecomposta := strings.Split(op, ",")

		x, _ := strconv.Atoi(strings.Replace(opDecomposta[0], "mul(", "", -1))
		y, _ := strconv.Atoi(strings.Replace(opDecomposta[1], ")", "", -1))

		total += x * y
	}

	println(total)
}
