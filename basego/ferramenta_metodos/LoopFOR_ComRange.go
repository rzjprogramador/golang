package ferramenta_metodos

import "fmt"

func Execute_FOR_com_Range() {
	nomes := [3]string{"Rei", "Guga", "Leo"}

	for _, valor := range nomes {
		valor += " :: concatenando em todos"
		fmt.Println(valor)
		// fmt.Println(indice, valor)
	}
}

/*
FOR_com_Range,
conceito: só tem o for para iterar fazer loop no golang, a clausula range é para iterar/percorrer sobre iteraveis [string, map dicionario, array slice lista, ]

sintaxe: for <variaveis de captura indice, valor> := range alvo { acao }

passes_logicos: for variavelCaputuradora_INDICE, variavelCaputuradora_VALOR := range onde quero iterar { O QUE QUERO REPETIR usando as variaveis a cada rodagem do loop }

obs: senão quer utilizar uma variavel substitua por _ anderline na captura, a cao que fizer com cada valor capturado será com efeitoColateral ou seja vai modificar o alvo original.

*/
