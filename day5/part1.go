package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var posterioresDe = make(map[string][]string)

func main() {

	f, err := os.Open("./../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	atualizoesPaginas := [][]string{}

	sessaoRegras := true
	for scanner.Scan() {
		linha := scanner.Text()

		if len(linha) < 2 {
			sessaoRegras = false
			continue
		}

		if sessaoRegras {
			regrasArr := strings.Split(linha, "|")

			regrasExistentes := posterioresDe[regrasArr[1]]

			regrasExistentes = append(regrasExistentes, regrasArr[0])

			posterioresDe[regrasArr[1]] = regrasExistentes
		} else {
			atualizoesPaginas = append(atualizoesPaginas, strings.Split(linha, ","))
		}

	}

	somaPagsMeio := 0
	for _, att := range atualizoesPaginas {

		valido := true
		for i, pag := range att {

			for j := i + 1; j < len(att); j++ {
				if isPosteriorDe(pag, att[j]) {
					valido = false
					break
				}
			}

			if !valido {
				break
			}

		}

		if valido {
			pagMeio, _ := strconv.Atoi(att[(len(att) / 2)])
			somaPagsMeio += pagMeio
		}
	}

	fmt.Println(somaPagsMeio)
}

func isPosteriorDe(numero string, checkPosterior string) bool {
	return slices.Contains(posterioresDe[numero], checkPosterior)
}
