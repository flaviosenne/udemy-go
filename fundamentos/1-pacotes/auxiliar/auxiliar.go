package auxiliar

import "fmt"

// diferença entre escrever a função começando
// com a letra maiúsula e minúscula é a visibilidade.
// Caso comece em maiúsula a função está diponível fora do pacote
// como é o caso da função abaixo 'Escrever'.
// Agora se estivesse em minúscula estaria visivel apenas  dentro
// do pacote em específico: 'escrever'
func Escrever() {
	fmt.Println("escrevendo do pkg auxiliar")
	escrever2()
}
