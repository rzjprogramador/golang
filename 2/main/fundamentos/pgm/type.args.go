package pgm

type Args struct {
	linguagem  string
	referencia Referencia
}

type Referencia struct {
	palavrasChaveReferencia PalavrasChaveReferencia
	exemplo                 Exemplo
}


