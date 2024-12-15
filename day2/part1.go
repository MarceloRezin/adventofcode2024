package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("./../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	niveisSeguros := 0

	for scanner.Scan() {
		linha := scanner.Text()
		linhaArr := strings.Split(linha, " ")

		nivelAnterior := 0
		nivelSeguro := true
		sequenciaPositiva := true
		for i, nivelString := range linhaArr {
			nivel, _ := strconv.Atoi(nivelString)

			if i != 0 {
				diferenca := float64(nivelAnterior - nivel)
				diferencaAbs := math.Abs(diferenca)

				if diferencaAbs < 1 || diferencaAbs > 3 {
					nivelSeguro = false
					break
				}

				if i == 1 {
					sequenciaPositiva = math.Signbit(diferenca)
				} else {
					if sequenciaPositiva != math.Signbit(diferenca) {
						nivelSeguro = false
						break
					}
				}
			}

			nivelAnterior = nivel
		}

		if nivelSeguro {
			niveisSeguros++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(niveisSeguros)
}
