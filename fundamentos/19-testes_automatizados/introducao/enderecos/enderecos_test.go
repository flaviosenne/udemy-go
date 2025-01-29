package enderecos_test

import (
	. "introducao/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		// {"Praça das Rosas", "Tipo Inválido"},
		{"Estrada qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA PAULISTA", "Avenida"},
		// {"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido, cenario.enderecoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}

// dentro do terminal para ver qual linha nâo está coberta basta usar esses comando
// go tool coverprofile cobertura.txt
// go tool cover --html=cobertura.txt
// para rodar varios testes a partir do modulo, basta colocar go test ./...
