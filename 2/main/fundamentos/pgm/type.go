package pgm

type Computacional struct {
	linguagem        Linguagem
	dado             Dado
	referencia       Referencia
	tipo             Tipo
	ferramentaFuncao Operacao_em_Funcao
}

type Linguagem struct {
	retentora    string
	nome         string
	transpilacao any
}

const (
	COMPILADA = iota
	INTERPRETADA
)

type Dado struct {
	conceito      Conceito
	valor         string
	parteDado     string
	dado          string
	dadoViaFuncao DadoViaFuncao
}

type Conceito struct {
	conceito string
}

type Exemplo struct {
	exemplo string
}

type Referencia struct {
	conceito                Conceito
	sinonimos               []string
	palavrasChaveReferencia PalavrasChaveReferencia
	exemplo                 Exemplo
}

type PalavrasChaveReferencia struct {
	conceito   Conceito
	mudavel    string
	naoMudavel string
}

type Tipo struct {
	conceito              Conceito
	possiveis             []string
	origem                any //ENUM: primitivo - naoPrimitivo - Personalizado
	nome                  string
	representacao         string
	sintaxe_modo_de_tipar string
	exemplo               Exemplo
}

const (
	PRIMITIVO = iota
	NAO_PRIMITIVO
	PERSONALIZADO_CLASS_STRUCT
)

type DadoViaFuncao struct {
	conceito   Conceito
	via_Classe Classe
	via_Struct Struct
	exemplo    Exemplo
}

type Classe struct {
	conceito   Conceito
	construtor string
}

type Struct struct {
	conceito  Conceito
	modelagem string
	exemplo   Exemplo
}

type Operacao_em_Funcao struct {
	atribuicao        Atribuicao
	comparacao        Comparacao
	castingConversao  string
	aritmeticos       []string
	controle_de_fluxo Controle_De_Fluxo
}

type Comparacao struct {
	conceito Conceito
	unitaria ComparacaoUnitaria
	mutipla  ComparacaoMultipla
}

type ComparacaoUnitaria struct {
	conceito Conceito
	exemplo  Exemplo
}
type ComparacaoMultipla struct {
	conceito Conceito
	exemplo  Exemplo
}

type Atribuicao struct {
	acumulativa string
	exemplo     Exemplo
}

type Controle_De_Fluxo struct {
	conceito         Conceito
	checagem_if_else Checagem
	checagem_switch  Checagem
	ternaria         Checagem
}

type Checagem struct {
	conceito Conceito
	sintaxe  string
	exemplo  Exemplo
}
