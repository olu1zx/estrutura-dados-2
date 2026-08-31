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

func (a *Arvore) Buscar(v int) *No {
	return buscar(a.Raiz, v)
}

func buscar(no *No, v int) *No {
	if no == nil || no.Valor == v {
		return no
	}

	if v < no.Valor {
		return buscar(no.Esquerdo, v)
	}

	return buscar(no.Direito, v)
}

func (a *Arvore) BuscarIter(v int) *No {
	atual := a.Raiz

	for atual != nil {
		if v == atual.Valor {
			return atual
		}

		if v < atual.Valor {
			atual = atual.Esquerdo
		} else {
			atual = atual.Direito
		}
	}

	return nil
}

func imprimirResultado(metodo string, valor int, no *No) {
	if no == nil {
		fmt.Printf("%s: valor %d nao encontrado\n", metodo, valor)
		return
	}

	fmt.Printf("%s: valor %d encontrado no no %d\n", metodo, valor, no.Valor)
}

func construirArvore() *Arvore {
	arvore := NovaArvore()
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		arvore.Inserir(valor)
	}
	return arvore
}

func main() {
	arvore := construirArvore()

	for _, valor := range []int{65, 45} {
		imprimirResultado("Busca recursiva", valor, arvore.Buscar(valor))
		imprimirResultado("Busca iterativa", valor, arvore.BuscarIter(valor))
	}
}
