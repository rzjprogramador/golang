package main

import "fmt"

type User struct {
	nome      string
	sobrenome string
	idade     uint
}

func (u User) nomeCompleto() string {
	return fmt.Sprintf("%s %s TEXTO_LIVRE onde Quiser entre as variaveis ", u.nome, u.sobrenome)
}

func (u *User) fazAniversario() {
	u.idade++
}

func criarUser(u User) User {
	return u
}

var user1 = User{
	nome:      "Reinaldo",
	sobrenome: "Zachars",
	idade:     46,
}

func main() {
	reinaldo := criarUser(user1)
	fmt.Println(reinaldo.nomeCompleto())
	reinaldo.fazAniversario() // fez 47
	reinaldo.fazAniversario() // fez 48
	fmt.Println(reinaldo)     // mostra que a idade é 48
}

/*
passosLogicos_CriarTipo_com_AcaoMetodo :
- criar estrutura , exemplo: struct User
- criar metodo vinculado a estrutura , exemplo: func nomeCompleto()
- configurar poder de instanciar a estrutura, exemplo: func criarUser(u User)
- em_escopo_de_uso: [
	-- instanciar a estrutura,
	-- receber entradas,
	-- usar os membros da novaInstancia
	---> exemplo: func main{ criarUser( input)}, usar as acoes ex: instancia1.metodo()
]

:: ponteiro_de_olho_na_instancia_certa: insere o asteristico na instancia, exemplo (u* User)
:: autoIncremente ex: idade++ nao retorna nada aqui na funcao de uso porque age diretamente no campo da estrutura de origem.


*/
