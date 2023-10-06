package pgm

type Comput_Fixa struct {
	linguagem                 string
	dado                      Dado
	tipo                      Tipo
	funcao_gera_dado          Funcao_Gera_Dado
	ferramentasComputacionais FerramentasComputacionais
}

type Dado struct {
	conceito      Conceito
	valor         string
	parteDado     string
	dado          string
	dadoViaFuncao Funcao_Gera_Dado
}

type PalavrasChaveReferencia struct {
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

type Funcao_Gera_Dado struct {
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

type FerramentasComputacionais struct {
	atribuicao        Atribuicao
	operacoes         Operacao_em_Funcao
	comparacao        Comparacao
	controle_de_fluxo Controle_De_Fluxo
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
	detalhes         string
}

type Checagem struct {
	conceito Conceito
	sintaxe  string
	exemplo  Exemplo
}
