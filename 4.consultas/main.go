package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func NovoNo(v int) *No {
	return &No{Valor: v}
}

func NovaArvore() *Arvore {
	return &Arvore{}
}

func (a *Arvore) Inserir(v int) {
	a.Raiz = inserir(a.Raiz, v)
}

func inserir(no *No, v int) *No {
	if no == nil {
		return NovoNo(v)
	}

	if v < no.Valor {
		no.Esquerdo = inserir(no.Esquerdo, v)
	} else if v > no.Valor {
		no.Direito = inserir(no.Direito, v)
	}

	return no
}

func (a *Arvore) Altura() int {
	return altura(a.Raiz)
}

func altura(no *No) int {
	if no == nil {
		return -1
	}

	alturaEsquerda := altura(no.Esquerdo)
	alturaDireita := altura(no.Direito)

	if alturaEsquerda > alturaDireita {
		return alturaEsquerda + 1
	}

	return alturaDireita + 1
}

func (a *Arvore) Contar() int {
	return contar(a.Raiz)
}

func contar(no *No) int {
	if no == nil {
		return 0
	}

	return 1 + contar(no.Esquerdo) + contar(no.Direito)
}

func (a *Arvore) ContarFolhas() int {
	return contarFolhas(a.Raiz)
}

func contarFolhas(no *No) int {
	if no == nil {
		return 0
	}

	if no.Esquerdo == nil && no.Direito == nil {
		return 1
	}

	return contarFolhas(no.Esquerdo) + contarFolhas(no.Direito)
}

func (a *Arvore) Minimo() (int, bool) {
	if a.Raiz == nil {
		return 0, false
	}

	atual := a.Raiz
	for atual.Esquerdo != nil {
		atual = atual.Esquerdo
	}

	return atual.Valor, true
}

func (a *Arvore) Maximo() (int, bool) {
	if a.Raiz == nil {
		return 0, false
	}

	atual := a.Raiz
	for atual.Direito != nil {
		atual = atual.Direito
	}

	return atual.Valor, true
}

func imprimirConsultas(nome string, arvore *Arvore) {
	fmt.Println(nome)
	fmt.Println("Altura:", arvore.Altura())
	fmt.Println("Quantidade de nos:", arvore.Contar())
	fmt.Println("Quantidade de folhas:", arvore.ContarFolhas())

	if minimo, ok := arvore.Minimo(); ok {
		fmt.Println("Minimo:", minimo)
	} else {
		fmt.Println("Minimo: arvore vazia")
	}

	if maximo, ok := arvore.Maximo(); ok {
		fmt.Println("Maximo:", maximo)
	} else {
		fmt.Println("Maximo: arvore vazia")
	}

	fmt.Println()
}

func construirArvore() *Arvore {
	arvore := NovaArvore()
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		arvore.Inserir(valor)
	}
	return arvore
}

func main() {
	arvoreVazia := NovaArvore()
	imprimirConsultas("Arvore vazia", arvoreVazia)

	arvoreComUmNo := NovaArvore()
	arvoreComUmNo.Inserir(50)
	imprimirConsultas("Arvore com um no", arvoreComUmNo)

	arvoreCompleta := construirArvore()
	imprimirConsultas("Arvore do exercicio", arvoreCompleta)
}
