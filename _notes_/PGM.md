# PGM

### linguagens:
- Golang, aplicabilidade: Backend, top: 1, performance : Pc: bom, servidor: bom, alcance: amplo, desvantagem: comunidade baixa, vantagem: [ utilizado em Big Techs]
-
- Kotlin, aplicabilidade: [Backend, Interface ], top: 2, performance: Pc: médio, servidor: bom, alcance: amplo, desvantagem: [comunidade baixa, pc tem que ser potente, ], vantagem: [ utilizado em Big Techs]
-
- Typescript ,aplicabilidade: [Backend, Interface ], top: 2, performance: Pc: bom, servidor: bom,  alcance: web, server, desvantagem: [ não utilizado em Big Techs para backend, ], vantagem: [ utilizado em Start ups, médios negócios, ]
-

### ferramentas_programacao_universal:
1. baseDominada( declaracoes( escrita mutável, leituraSomenteImutavel), funções, POO[ abstração, encapsulamento, polimorfismo, ],  Lopps, coleções, Api(resto, graphql ))
2. opcional_recursos_da_linguagem(  )

### nomeacoes :
variáveis_e_objeto : substantivo,
array : substantivo no plural do que agrupa
função: verbo_Imperativo+Substantivo,  pense: QUANDO FOR CHAMADA ELA VAI ???... a resposta sera o nome um verboImperativo: exemplo : somar(), calcular(), reduzir(),
tutorial : https://fb.watch/nfhSF12MzR/?mibextid=Nif5oz

> nomeacao:
  1. nome_arquivo: SnackCase : porque pode ser o nome do artefato principal dele



# paradigma_codificao:
  possiveis: [ declarativo, imperativo ]

3. imperativo: um legado verboso que diz como fazer passo a passo,
4.
5. declarativa: é moderno usa -se funções prontas da linguagem que já abstrairam e encapsularam o legado imperativo..
6. tutoriais : https://youtube.com/shorts/WCpIbFOs0VQ?si=spr_LHkG49JNS1tS

codificacao_conforme_comunicacao_IO_de_entrada_e_saida_do_dado:
possíveis: [ code_Interno_Memoria_IO, code_Externo_Servidor_IO ],
conceito:
temos que ter o código natural para máquina entender e junto o código que faz o provedor inOut entender ex: Rest,  Graphql, então para cada instrução natural mapear reproduzir também a instrução para o provedorInOut entender,

# IMPLEMENTACAO

### principais_conceitos:
- programe respostas resilientes de forma assincrona garantindo dar a resposta em segundos, minutos, horas ou dias depois mas entrega.

### contratos:
conceito: toda declaracao é um contrato. narração: VAI TER ESTE FORMATO.
possíveis: [
 de_valor_variavel,
 de_,função,
 de_metodo: é uma função atrelada a um objeto.
 de_colecao_de_campos_objeto,
]

### TIPO
> conceito:
todo tipo é representado por valor e todos tem seu valor padraoPorDefault exemplo_universal : texto="", número = 0, lógico = false, nulos = null, inesistente = indefinido,

> possíveis:
Primitivos da linguagem, e os criados por Gerador de Objeto ex: Class, struct, FactoryFunctions

> any_qualquer_formato:
para descobrir o tipoQualquerFormato é só ver o tipo da funcao que printa no console ja que ela aceita qualquer tipo, em ts é any, no java Any, em golang tambem tem o any ou legado que era interface{}  /interface vazia/,

> texto_colecao_de_string:
string

> numeros:
inteiro, decimal

> boleano:
sim: true, nao: false

> error:
cada linguagem tem seu objeto error. temos que usar o objeto error da linguem e dele criar nosso novo error personalizado.

### ERROS_EXCECOES
matar_execucao_programa: na maioria usam throw, em GO usa a funcao panic ( "message") recupera com a keyword recovery


### VALOR
temos_sempre_tipos: seu valor a esquerda sempre é de tipo então use os membros ferramentas deste tipo.

Valor_Zero_por_Default :  todo tipo é inicializado por um valor zero ou seja DEFAULT se nao for passado inicializado com nada ele devolve este valor ZeroPorDefault

atribuicao:
conceito: atribuicao é inicializar valor dar preenchimento aos artefato que vem de contrato,
possíveis: [ atribuicaoDireta com =,

atribuicaoInferida no Go é com :=,

atribuicaoAcumulativa com +=, -=, *=, %= ,

atribuicao_de_funcao: é invocar chamar e argumentar se preciso.,

atribuir_por_Composicao: é no contrato assinalar um campo  com valor de um outro Tipo personalizado e no uso dar como valor o objeto deste tipoPsrsonalizado como valor ao campo sob contrato.

 ]

 ## referencia e copia
 referencia : é a variavel original o que modificar vai modificar direto nela é marcavel como um ponteiro em algumas linguagens como [ Golang, ]

 copia: é quando atrinuimos uma variavel a outra , então esta nova atribuida é uma copia e qualquer mudanca de valor nela nao atinge a original e simm a copia somente, a original só se modifica se for um ponteiro.


## SOMADORES E SUBTRAIDORES
adiciona_mais_um: variavel++ , cada vez que for chamado sera adicionado +1 na variavel que tem que ser um numero, este operador ja torna a variavel um ponteiro por padrao isto é ja modifica diretamente na variavel original e nao em copia.

diminui_menos_um: variavel-- , mesma coisa da adicionadora ja é um ponteiro age direto na referencia.

## ACOES DE OBJETO ENTIDADE
metodos_prototype: Funcoes ligadas ao estruturador de entidade ( classe, ou struct ) São acoes que manipulam os campos da entidade.   Os campos da entidade sao acessados para configuracao com this. ou em outras linguagens com uma variavel que referencia um ponteiro da estrutura em questao (Em golang fazemso uma funcao normal e antes do nome da funcao entre parenteses referenciamos a variavel que sera o this ao ponteiro da struct)



# FERRAMENTA PARA FUNCAO

### DECISAO LIMITADA MAIS QUE 1 OPCAO
> switch
  Decisao_Interruptor_Switch
  conceito: ao inves de fazer mais que 1 : if , if else - fazer então um switch quando tenho opções limitadas para tomar uma decisao.

  passes_logicos: avalie o parametro e caso for o que passei, retorne isto. ..obs: naMelhorForma: posso analizar o parametro direto no caso.

  default: tem que estar no mesmo escopo dos casos, e o tipo do default tem que ser o mesmo prometido na decalracao do metodo.


### LOOP
loop_FOR_com_Closure : dentro do escopo de uma funcao posso usar o closure de uma variavel mutavel modificavel inicuando-a  com valor zero, acumulando novos valores a ela no for e usando closure retornar esta var que lembra onde ela se iniciou, o que lhe foi lhe foi atribuido e devolve seus novos valores.



