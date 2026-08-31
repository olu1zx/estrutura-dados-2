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

func (a *Arvore) Remover(v int) {
	a.Raiz = a.Raiz.Remover(v)
}

func (no *No) Remover(v int) *No {
	if no == nil {
		return nil
	}

	if v < no.Valor {
		no.Esquerdo = no.Esquerdo.Remover(v)
		return no
	}

	if v > no.Valor {
		no.Direito = no.Direito.Remover(v)
		return no
	}

	if no.Esquerdo == nil && no.Direito == nil {
		return nil
	}

	if no.Esquerdo == nil {
		return no.Direito
	}

	if no.Direito == nil {
		return no.Esquerdo
	}

	sucessor := minimoNo(no.Direito)
	no.Valor = sucessor.Valor
	no.Direito = no.Direito.Remover(sucessor.Valor)
	return no
}

func minimoNo(no *No) *No {
	atual := no
	for atual.Esquerdo != nil {
		atual = atual.Esquerdo
	}
	return atual
}

func (a *Arvore) EmOrdem() []int {
	valores := []int{}
	emOrdem(a.Raiz, &valores)
	return valores
}

func emOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}

	emOrdem(no.Esquerdo, valores)
	*valores = append(*valores, no.Valor)
	emOrdem(no.Direito, valores)
}

func construirArvore() *Arvore {
	arvore := NovaArvore()
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		arvore.Inserir(valor)
	}
	return arvore
}

func testarRemocao(descricao string, valor int) {
	arvore := construirArvore()

	fmt.Println(descricao)
	fmt.Println("Antes:", arvore.EmOrdem())
	arvore.Remover(valor)
	fmt.Println("Depois:", arvore.EmOrdem())
	fmt.Println()
}

func main() {
	testarRemocao("Remocao de folha: 20", 20)
	testarRemocao("Remocao de no com um filho: 60", 60)
	testarRemocao("Remocao de no com dois filhos: 50", 50)
}
