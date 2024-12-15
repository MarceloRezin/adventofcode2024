package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var cacaPalavras = make([][]string, 0)

func main() {

	f, err := os.Open("./../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		linha := scanner.Text()

		linhaArr := strings.Split(linha, "")
		cacaPalavras = append(cacaPalavras, [][]string{linhaArr}...)
	}

	contagemXmas := 0

	for lin := 1; lin < len(cacaPalavras)-1; lin++ {
		linha := cacaPalavras[lin]

		for col := 1; col < len(linha)-1; col++ {
			if linha[col] == "A" && isXMas(lin, col) {
				contagemXmas++
			}
		}
	}

	println(contagemXmas)
}

func isXMas(linha int, coluna int) bool {
	linhaSuperior := linha - 1
	linhaInferior := linha + 1

	colunaEsquerda := coluna - 1
	colunaDireita := coluna + 1

	superiorEsquerda := cacaPalavras[linhaSuperior][colunaEsquerda]
	superiorDireita := cacaPalavras[linhaSuperior][colunaDireita]

	inferiorEsquerda := cacaPalavras[linhaInferior][colunaEsquerda]
	inferiorDireita := cacaPalavras[linhaInferior][colunaDireita]

	return ((superiorEsquerda == "M" && inferiorDireita == "S") || (superiorEsquerda == "S" && inferiorDireita == "M")) &&
		((inferiorEsquerda == "M" && superiorDireita == "S") || (inferiorEsquerda == "S" && superiorDireita == "M"))
}
