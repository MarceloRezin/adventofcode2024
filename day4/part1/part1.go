package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var xmasRegex, _ = regexp.Compile(`XMAS`)
var samxRegex, _ = regexp.Compile(`SAMX`)

func main() {

	f, err := os.Open("./../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	cacaPalavras := make([][]string, 0)
	contagemXmas := 0

	scanner := bufio.NewScanner(f)
	// linha
	for scanner.Scan() {
		linha := scanner.Text()

		contagemXmas += countXmas(linha)

		linhaArr := strings.Split(linha, "")
		cacaPalavras = append(cacaPalavras, [][]string{linhaArr}...)
	}

	// coluna
	for col := 0; col < len(cacaPalavras[0]); col++ {
		colunaTransposta := ""
		for _, linha := range cacaPalavras {
			colunaTransposta += linha[col]
		}

		contagemXmas += countXmas(colunaTransposta)
	}

	// diagonal cima baixo \
	//Metade de cima
	linhaMax := len(cacaPalavras)
	for col := 0; col < len(cacaPalavras[0])-3; col++ {
		colTmp := col

		linhaDiagonal := ""
		for lin := 0; lin < linhaMax; lin++ {
			linhaDiagonal += cacaPalavras[lin][colTmp]
			colTmp++
		}

		contagemXmas += countXmas(linhaDiagonal)

		linhaMax--
	}

	// Metade de baixo
	colMax := len(cacaPalavras[0]) - 1
	for lin := 1; lin < len(cacaPalavras)-3; lin++ {
		linTmp := lin

		linhaDiagonal := ""
		for col := 0; col < colMax; col++ {
			linhaDiagonal += cacaPalavras[linTmp][col]
			linTmp++
		}

		contagemXmas += countXmas(linhaDiagonal)

		colMax--
	}

	// diagonal baixo cima /
	//Metade de baixo
	linhaMin := 0
	for col := 0; col < len(cacaPalavras[0])-3; col++ {
		colTmp := col

		linhaDiagonal := ""
		for lin := len(cacaPalavras) - 1; lin >= linhaMin; lin-- {
			linhaDiagonal += cacaPalavras[lin][colTmp]
			colTmp++
		}

		contagemXmas += countXmas(linhaDiagonal)

		linhaMin++
	}

	// Metade de cima
	colMax = len(cacaPalavras[0]) - 1
	for lin := len(cacaPalavras) - 2; lin > 2; lin-- {
		linTmp := lin

		linhaDiagonal := ""
		for col := 0; col < colMax; col++ {
			linhaDiagonal += cacaPalavras[linTmp][col]
			linTmp--
		}

		contagemXmas += countXmas(linhaDiagonal)

		colMax--
	}

	println(contagemXmas)
}

func countXmas(texto string) int {
	return len(xmasRegex.FindAllString(texto, -1)) + len(samxRegex.FindAllString(texto, -1))
}
