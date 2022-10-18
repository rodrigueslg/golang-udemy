package enderecos_test

import (
	. "instroducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	tipoEsperado     string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"RODOVIA BR 381", "Rodovia"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.tipoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.tipoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	t.Error("Esse teste sempre vai falhar")
}

// Comandos para rodar o teste:
// 1:	teste do arquivo local								 	| go test
// 2:	teste de todos os arquivos em cascata				 	| go test ./...
// 3:	teste do arquivo local verboso						 	| go test -v
// 4:	teste do arquivo local com % de cobertura			 	| go test --cover
// 5.1: teste do arquivo local com % de cobertura em arquivo 	| go test ./... --coverprofile cobertura.txt
// 5.2:	ler o arquivo de cobertura criado no passo anterior	 	| go tool cover --func=cobertura.txt
// 5.3:	let o de cobertura criado no passo anterior como html	| go tool cover --html=cobertura.txt
