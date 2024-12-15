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

		if isNivelSafe(linhaArr) || isNivelSafe(reverse(linhaArr)) {
			niveisSeguros++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(niveisSeguros)
}

func isNivelSafe(linha []string) bool {
	nivelAnterior := 0
	nivelSeguro := true
	sequenciaPositiva := true
	existeErro := false
	for i, nivelString := range linha {
		nivel, _ := strconv.Atoi(nivelString)

		if i != 0 {
			diferenca := float64(nivelAnterior - nivel)
			diferencaAbs := math.Abs(diferenca)

			if diferencaAbs < 1 || diferencaAbs > 3 {
				if existeErro {
					nivelSeguro = false
					break
				} else {
					existeErro = true
					continue
				}

			}

			if i == 1 {
				sequenciaPositiva = math.Signbit(diferenca)
			} else {
				if sequenciaPositiva != math.Signbit(diferenca) {
					if existeErro {
						nivelSeguro = false
						break
					} else {
						existeErro = true
						continue
					}
				}
			}
		}

		nivelAnterior = nivel
	}

	return nivelSeguro
}

func reverse(list []string) []string {
	var reverseList []string

	for i := len(list) - 1; i >= 0; i-- {
		reverseList = append(reverseList, list[i])
	}

	return reverseList
}
