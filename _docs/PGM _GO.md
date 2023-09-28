# LINGUAGEM GOLANG

### TIPO
> any_qualquer_formato:
para descobrir o tipoQualquerFormato é só ver o tipo da funcao que printa no console ja que ela aceita qualquer tipo, em ts é any, no java Any, em golang tambem tem o any ou legado que era interface{}  /interface vazia/,

> texto_string:
string

> numeros:
conceito: prefira inferir , assim ao setar e passar o mouse no artefato ele mostrará o tipo inferido.
numeros_inteiros: int, quando somente positivo use uint (pega o padrao da amaquinaInUso 32 ou 64 bytes)
numeros_decimais: float64 (pega o padrao da amaquinaInUso)

> boleano:
sim: true, nao: false

> error:
lib_golang: error, membros
  - criar novo erro: error.new("mensagem texto")

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

### LOOP
loop_FOR_com_Closure : dentro do escopo de uma funcao posso usar o closure de uma variavel mutavel modificavel inicuando-a  com valor zero, acumulando novos valores a ela no for e usando closure retornar esta var que lembra onde ela se iniciou, o que lhe foi lhe foi atribuido e devolve seus novos valores.

