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
