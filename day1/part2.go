package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("./../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	scanner := bufio.NewScanner(f)

	listA := []int{}
	mapB := make(map[int]int)

	for scanner.Scan() {
		linha := scanner.Text()
		linhaArr := strings.Split(linha, "   ")

		itemA, _ := strconv.Atoi(linhaArr[0])
		itemB, _ := strconv.Atoi(linhaArr[1])

		listA = append(listA, itemA)

		frequencia, exists := mapB[itemB]
		if !exists {
			frequencia = 1
		} else {
			frequencia++
		}

		mapB[itemB] = frequencia
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	somaSimilaridade := 0
	for _, v := range listA {

		frequencia, exists := mapB[v]
		if !exists {
			frequencia = 0
		}

		somaSimilaridade += v * frequencia
	}

	fmt.Println(int32(somaSimilaridade))
}
