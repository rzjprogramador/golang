package pgm

type Args_Rotativa struct {
	linguagem  string
	referencia Referencia
}

type Referencia struct {
	palavrasChaveReferencia PalavrasChaveReferencia
	exemplo                 Exemplo
}
