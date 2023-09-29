package main

import (
	"fmt"

	"github.com/rzjprogramador/golang/basego/estrutura_de_formato"
	"github.com/rzjprogramador/golang/basego/ferramenta_metodos"
	"github.com/rzjprogramador/golang/basego/valor"
	"github.com/rzjprogramador/golang/basego/valor_agrupado"
)

func main() {
	valor.ExecuteDeclaracao()
	valor.ExecuteTipos()
	fmt.Print("Hello Golang")

	estrutura_de_formato.Execute_EstruturaDeFormato()

	fmt.Println(valor_agrupado.ListaLiteral())
	fmt.Println(valor_agrupado.AddElementoALista())
	fmt.Println(valor_agrupado.FatiaSlicePorPosicaoInicioeFim())

	ferramenta_metodos.ExecuteDecisao_Interruptor_Switch()
	ferramenta_metodos.Execute_FOR_com_Range()
}
