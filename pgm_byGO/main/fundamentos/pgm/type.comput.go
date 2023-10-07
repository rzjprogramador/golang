package pgm

type Comput_Fixa struct {
	linguagem                 string
	dado                      Dado
	tipo                      Tipo
	funcao_gera_dado          Funcao_Gera_Dado
	ferramentasComputacionais FerramentasComputacionais
}

type Dado struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	valor          string
	parteDado      string
	dado           string
	dadoViaFuncao  Funcao_Gera_Dado
}

type PalavrasChaveReferencia struct {
	temNaLinguagem Habilitado
	mudavel        string
	naoMudavel     string
}

type Declaracao_Referencia struct {
	temNaLinguagem         Habilitado
	sintaxe_padrao         string
	referencia_mudavel     string
	referencia_nao_mudavel string
	funcao                 string
	metodo_de_novo_tipo    string
	array                  string
}

type Tipo struct {
	temNaLinguagem        Habilitado
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
	temNaLinguagem Habilitado
	conceito       Conceito
	via_Classe     Classe
	via_Struct     Struct
	exemplo        Exemplo
}

type Classe struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	construtor     string
}

type Struct struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	modelagem      string
	exemplo        Exemplo
}

type FerramentasComputacionais struct {
	temNaLinguagem    Habilitado
	atribuicao        Atribuicao
	operacoes         Operacao_em_Funcao
	comparacao        Comparacao
	controle_de_fluxo Controle_De_Fluxo
}

type Operacao_em_Funcao struct {
	temNaLinguagem    Habilitado
	atribuicao        Atribuicao
	comparacao        Comparacao
	castingConversao  string
	aritmeticos       []string
	controle_de_fluxo Controle_De_Fluxo
}

type Comparacao struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	unitaria       ComparacaoUnitaria
	mutipla        ComparacaoMultipla
}

type ComparacaoUnitaria struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	exemplo        Exemplo
}
type ComparacaoMultipla struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	exemplo        Exemplo
}

type Atribuicao struct {
	temNaLinguagem              Habilitado
	acumulativa                 string
	auto_operavel_Unitariamente string
	exemplo                     Exemplo
}

type Controle_De_Fluxo struct {
	temNaLinguagem   Habilitado
	conceito         Conceito
	checagem_if_else Checagem
	checagem_switch  Checagem
	ternaria         Checagem
	detalhes         string
}

type Checagem struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	sintaxe        string
	exemplo        Exemplo
}
