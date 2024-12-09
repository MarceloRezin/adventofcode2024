package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
	listB := []int{}

	for scanner.Scan() {
		linha := scanner.Text()
		linhaArr := strings.Split(linha, "   ")

		itemA, _ := strconv.Atoi(linhaArr[0])
		itemB, _ := strconv.Atoi(linhaArr[1])

		listA = append(listA, itemA)
		listB = append(listB, itemB)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(listA)
	sort.Ints(listB)

	var somaDistancias float64 = 0
	for i := 0; i < 1000; i++ {
		somaDistancias += math.Abs(float64(listA[i]) - float64(listB[i]))
	}

	fmt.Println(int32(somaDistancias))
}
