package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Informacoes do voo necessarias para o codigo
type infoVoos struct {
	NumVoo           string
	LocalDeSaida     string
	Destino          string
	horarioDeSaida   time.Time
	horarioDeChegada time.Time
	status           string
	localizacao      string
}

func SimulacaoDeBd(voos *[]infoVoos) {
	// criando as horas dos voos de saida
	saida_voo1 := time.Date(2024, time.April, 12, 12, 00, 00, 0, time.UTC)
	saida_voo2 := time.Date(2024, time.June, 22, 12, 00, 00, 0, time.UTC)
	saida_voo3 := time.Date(2024, time.May, 12, 12, 00, 00, 0, time.UTC)

	// criando a hora de chegada
	chegada_voo1 := time.Date(2024, time.April, 12, 15, 00, 00, 0, time.UTC)
	chegada_voo2 := time.Date(2024, time.June, 22, 19, 00, 00, 0, time.UTC)
	chegada_voo3 := time.Date(2024, time.May, 12, 14, 00, 00, 0, time.UTC)

	// Adicionando instâncias à slice
	*voos = append(*voos, infoVoos{"LA365", "Salvador", "Brasília", saida_voo1, chegada_voo1, "No horario", "Portao 10"})
	*voos = append(*voos, infoVoos{"AA123", "New York", "Los Angeles", saida_voo2, chegada_voo2, "Atrasado", "Gate 5"})
	*voos = append(*voos, infoVoos{"UA456", "Chicago", "Houston", saida_voo3, chegada_voo3, "No horario", "Gate 12"})
}

// Função para achar um voo
func AcharVoo(voos []infoVoos) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite o número do seu voo: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	// Criando uma slice para armazenar várias instâncias de infoVoos
	var voos []infoVoos

	// Simulando a inserção de dados no "banco de dados"
	SimulacaoDeBd(&voos)

	// Obtendo o número do voo desejado do usuário
	numeroDesejado := AcharVoo(voos)

	// Procurando a instância na slice
	var vooEncontrado *infoVoos
	for i := range voos {
		if voos[i].NumVoo == numeroDesejado {
			vooEncontrado = &voos[i]
			break
		}
	}

	// Verificando se o voo foi encontrado
	if vooEncontrado != nil {
		fmt.Printf("Número do Voo: %s\n", vooEncontrado.NumVoo)
		fmt.Printf("Local de Saída: %s\n", vooEncontrado.LocalDeSaida)
		fmt.Printf("Destino: %s\n", vooEncontrado.Destino)
		fmt.Printf("Horário de Saída: %v\n", vooEncontrado.horarioDeSaida)
		fmt.Printf("Horário de Chegada: %v\n", vooEncontrado.horarioDeChegada)
		fmt.Printf("Status: %s\n", vooEncontrado.status)
		fmt.Printf("Localização: %s\n", vooEncontrado.localizacao)
	} else {
		fmt.Println("Voo não encontrado.")
	}
}
