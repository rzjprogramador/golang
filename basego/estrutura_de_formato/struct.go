package estrutura_de_formato

import "fmt"

type FormatoEntidade struct {
	nome        string
	idade       uint
	programador bool
}

var request = FormatoEntidade{"Rei", 46, true}

// request := FormatoEntidade{"Rei", 46, true} // declaracao_fora_de_escopo: somente sendo a declaracaoExplicita

func createEntidade(f FormatoEntidade) FormatoEntidade {
	return f
}

func Execute_EstruturaDeFormato() {

	reinaldo := createEntidade(request)

	fmt.Println(reinaldo)
}
