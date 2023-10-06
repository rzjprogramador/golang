package pgm

import "fmt"

var linguagemGo = Linguagem{
	retentora:    "Google",
	nome:         "Golang",
	transpilacao: "Compilada",
}

var linguagemTs = Linguagem{
	retentora:    "Microsoft",
	nome:         "Typescript",
	transpilacao: "Interpretada",
}

func linguagem(l Linguagem) Linguagem {
	return l
}

func ExecuteLinguagem() {
	fmt.Println(linguagem(linguagemGo))
	fmt.Println(linguagem(linguagemTs))
}
