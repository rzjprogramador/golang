package main

import "fmt"

// Tipo_Personal_Novo

// estrutura { criar fabrica : [struct, ou classe] }
type User struct {
	nome      string
	sobrenome string
	idade     uint
}

// acoes { vincular a estrutura }
func (u User) nomeCompleto() string {
	return fmt.Sprintf("%s %s TEXTO_LIVRE onde Quiser entre as variaveis ", u.nome, u.sobrenome)
}

func (u *User) fazAniversario() {
	u.idade++
}

func (u *User) ehMaiorDeIdade() bool {
	if u.idade > 18 {
		return true
	}
	return false
}

// uso
func main() {
	reinaldo := User{"Reinaldo", "Zachars", 46} // instanciando
	fmt.Println(reinaldo.nomeCompleto())
	reinaldo.fazAniversario()              // fez 47
	reinaldo.fazAniversario()              // fez 48
	fmt.Println(reinaldo)                  // mostra que a idade é 48
	fmt.Println(reinaldo.ehMaiorDeIdade()) // true
}

/*
Tipo_Personal_Novo

ciclo_Logicos_Novo_TIPO :
-- estrutura: [
	- criar_Struct , exemplo: struct User
	]
-- acoes: [
	criar metodo vinculado a estrutura , exemplo: func nomeCompleto()

	- ferramentas :
	ferramentas_de_operacoes: [
		- autoIncremente ex: idade++ nao retorna nada aqui na funcao de uso porque age diretamente no campo da estrutura de origem.
	]
	ferramenta_atribuicao: [
		ponteiro: ponteiro_de_olho_na_propriedade_dentro_do_tipo_e_nao_em_copias: insere o asteristico no tipoAlvo, exemplo (u *User), detalhe: o asteristico é sempre no tipo pois fica de olho na propriedade dentro do tipo gravado e nao na copia
	]
	ferramenta_devolve_texto: [
		-- texto_Interpolado: return fmt.Sprintf("%s %s TEXTO_LIVRE onde Quiser entre as variaveis ", u.nome, u.sobrenome) // com fmt.Sprintf("entre aspasDuplas o %TIPO_DA_VARIAVEL Qualquer_Texto_Opcional", separados por virgula variaveis corresponte aos que estao nas aspas por justa posicao ex: variavel1, variavel2)
	]
]

- Uso_em_escopo_de_uso: [
	-- instanciacao : instanciar a estrutura , ex: instancia := Tipo{arg1, arg2, arg3 }, não preciso criar uma funcao apra instanciar o Tipo ex: func criarUser(u User)
	-- receber entradas,
	-- usar os membros da novaInstancia
	---> exemplo: func main{ criarUser( input)}, usar as acoes ex: instancia1.metodo()
]

*/
