package enderecos

import "strings"

// TipoDeEndereco verifica se o endereço é da rua, avenida ou rodovia
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "rodovia"}

	endereco = strings.ToLower(endereco)
	primeiraPalavra := strings.Split(endereco, " ")[0]

	enderecoEhValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoEhValido = true
		}
	}

	if enderecoEhValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Inválido"
}
