package estrutura_de_formato

import "fmt"

type FormatoEntidade struct {
	nome        string
	idade       uint
	programador bool
}

var request = FormatoEntidade{"Rei", 46, true}

func createEntidade(f FormatoEntidade) FormatoEntidade {
	return f
}

func Execute_EstruturaDeFormato() {

	reinaldo := createEntidade(request)

	fmt.Println(reinaldo)
}
