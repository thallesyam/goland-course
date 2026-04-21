# Pacotes, Módulos e Submódulos em Go

## Conceitos

### Módulo
Definido pelo arquivo `go.mod` na raiz do projeto. É a unidade de distribuição e versionamento do código. Todo projeto Go precisa de um módulo.

```
module meu-projeto   ← nome do módulo
go 1.21
```

### Pacote
É um conjunto de arquivos `.go` dentro de uma mesma pasta, todos com a mesma declaração `package` no topo. É a unidade de organização do código.

- Cada pasta = um pacote
- Todos os arquivos de uma pasta devem ter o mesmo `package`

```go
// pasta: saudacao/hello.go
package saudacao   // ← nome do pacote

func Say(name string) string { ... }
```

### Submódulo
Um módulo independente dentro de um repositório maior, com seu próprio `go.mod`. É usado quando partes do projeto precisam ser completamente independentes (versionamento, dependências próprias).

---

## Como o Go monta o caminho de import

A regra é:

```
<nome do módulo> + <caminho relativo à raiz do módulo>
```

### Exemplo

```
meu-projeto/          ← raiz do módulo
├── go.mod            → module meu-projeto
└── saudacao/
    └── hello.go      → package saudacao
```

Import correto:

```go
import "meu-projeto/saudacao"
```

---

## Múltiplos pacotes em um projeto

Um único `go.mod`, cada pasta é um pacote diferente:

```
meu-projeto/
├── go.mod              → module meu-projeto
├── saudacao/
│   └── hello.go        → package saudacao
├── matematica/
│   └── calc.go         → package matematica
└── cmd/
    └── main.go
```

```go
// cmd/main.go
import (
    "meu-projeto/saudacao"
    "meu-projeto/matematica"
)

func main() {
    saudacao.Say("Thalles")
    matematica.Soma(1, 2)
}
```

---

## Regras importantes

| Regra | Correto | Errado |
|-------|---------|--------|
| Uma pasta, um pacote | todos os `.go` da pasta com mesmo `package` | arquivos da mesma pasta com `package` diferentes |
| Import começa pelo módulo | `"meu-modulo/pasta/subpasta"` | `"pasta/subpasta"` sem o nome do módulo |
| Nome do módulo ≠ nome da pasta | `go.mod` pode ter qualquer nome | não precisa ser igual ao nome da pasta raiz |
