package main

import "fmt"

func Tribonacci(lista []int, n int) []int {
	if n <= 3 {
		return lista[:n]
	}

	for iAtual := 3; iAtual < n; iAtual++ {
		nextI := lista[iAtual-3] + lista[iAtual-2] + lista[iAtual-1]

		lista = append(lista, nextI)
	}

	return lista
}

func main() {
	var n1 int
	fmt.Println("Digite o número de vezes para o Tribonacci rodar:")
	fmt.Scan(&n1)

	listaNumeros := []int{1, 1, 2}
	resultado := Tribonacci(listaNumeros, n1)

	fmt.Println("Resultado Tribonacci:", resultado)
}
