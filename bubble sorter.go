package main

import "fmt"

func main() {
	var i, j, aux, tam, num int
	lista := []int{}

	fmt.Printf("digite o tamanho da lista:")
	fmt.Scan(&tam)
	fmt.Println()

	for i = 0; i < tam; i++ {
		fmt.Printf("%d° numero:", i+1)
		fmt.Scan(&num)
		println()
		lista = append(lista, num)
	}

	for j = 0; j < tam; j++ {

		for i = 0; i <= tam-2; i++ {
			if lista[i] > lista[i+1] {
				aux = lista[i]
				lista[i] = lista[i+1]
				lista[i+1] = aux
			}

		}
	}
	fmt.Printf("a sua lista organizada de forma crescente: \n")
	for i = 0; i < tam; i++ {
		fmt.Printf("%d ", lista[i])
	}

}
