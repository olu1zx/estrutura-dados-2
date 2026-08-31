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

func (a *Arvore) PreOrdem() []int {
	valores := []int{}
	preOrdem(a.Raiz, &valores)
	return valores
}

func preOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}

	*valores = append(*valores, no.Valor)
	preOrdem(no.Esquerdo, valores)
	preOrdem(no.Direito, valores)
}

func (a *Arvore) EmOrdem() []int {
	valores := []int{}
	emOrdem(a.Raiz, &valores)
	return valores
}

// EmOrdem produz valores crescentes em uma BST porque visita primeiro
// todos os menores valores da subarvore esquerda, depois a raiz, e por fim
// os maiores valores da subarvore direita.
func emOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}

	emOrdem(no.Esquerdo, valores)
	*valores = append(*valores, no.Valor)
	emOrdem(no.Direito, valores)
}

func (a *Arvore) PosOrdem() []int {
	valores := []int{}
	posOrdem(a.Raiz, &valores)
	return valores
}

func posOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}

	posOrdem(no.Esquerdo, valores)
	posOrdem(no.Direito, valores)
	*valores = append(*valores, no.Valor)
}

func (a *Arvore) EmLargura() []int {
	valores := []int{}
	if a.Raiz == nil {
		return valores
	}

	fila := []*No{a.Raiz}
	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]
		valores = append(valores, atual.Valor)

		if atual.Esquerdo != nil {
			fila = append(fila, atual.Esquerdo)
		}
		if atual.Direito != nil {
			fila = append(fila, atual.Direito)
		}
	}

	return valores
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

	fmt.Println("Pre-ordem:", arvore.PreOrdem())
	fmt.Println("Em-ordem:", arvore.EmOrdem())
	fmt.Println("Pos-ordem:", arvore.PosOrdem())
	fmt.Println("Em largura:", arvore.EmLargura())
}
