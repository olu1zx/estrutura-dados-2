# Estrutura de Dados 2 - Atividade BST em Go

Atividade pratica sobre Arvores Binarias de Busca (BST), com exercicio em papel e implementacoes em Go.

## Estrutura

- `exercicio-1-papel.md`: resposta explicada do exercicio em papel.
- `01-criar-e-inserir-bst`: criacao da arvore, criacao de no e insercao.
- `02-busca-bst`: busca recursiva e iterativa.
- `03-percursos-bst`: percursos em pre-ordem, em-ordem, pos-ordem e largura.
- `04-consultas-bst`: altura, contagem, folhas, minimo e maximo.
- `05-remover-bst`: remocao nos tres casos.

## Como executar

Com o Go instalado, rode cada exercicio a partir da raiz do repositorio:

```powershell
go run ./01-criar-e-inserir-bst
go run ./02-busca-bst
go run ./03-percursos-bst
go run ./04-consultas-bst
go run ./05-remover-bst
```

Para verificar todos os pacotes:

```powershell
go test ./...
```

## Fluxo de entrega sugerido

1. Crie o repositorio da disciplina no GitHub.
2. Deixe a `main` com a estrutura inicial da disciplina.
3. Crie a branch da atividade:

```powershell
git checkout -b atividade-bst
```

4. Adicione os arquivos da atividade, faca o commit e envie a branch:

```powershell
git add .
git commit -m "feat: adiciona atividade de bst em go"
git push -u origin atividade-bst
```

5. No GitHub, abra um Pull Request de `atividade-bst` para `main`.

Se ainda nao configurou seu usuario do Git local, use seus dados do GitHub:

```powershell
git config --global user.name "Seu Nome"
git config --global user.email "seu-email-do-github@example.com"
```
