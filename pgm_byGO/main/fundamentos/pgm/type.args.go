package pgm

type Args_Rotativa struct {
	temNaLinguagem Habilitado
	linguagem      string
	referencia     Referencia
	metodo_Para_NovosTipos Metodo_Para_Novos_Tipos
}

type Referencia struct {
	palavrasChaveReferencia PalavrasChaveReferencia
	declaracaoReferencia Declaracao_Referencia
	exemplo                 Exemplo
}

type Metodo_Para_Novos_Tipos struct {
	acionar_via string

}
