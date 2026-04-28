# Aula 02 — Basic Types (Go)

## O que foi estudado

O código implementa um **programa de média aritmética** com leitura contínua de valores via `stdin`, explorando:

- Tipos básicos: `float64`, `int`
- Declaração de variáveis com `var`
- Loop infinito com `for {}` e quebra por erro (`break`)
- Ponteiros: `*float64` para modificar valor no endereço original
- Múltiplos retornos: `(int, error)` — padrão idiomático Go
- Blank identifier `_` para ignorar retornos desnecessários
- Conversão explícita de tipos: `float64(n)`
- Separação de responsabilidades em pacotes internos (`logger`)

---

## Principais tipos básicos de Go

| Tipo | Descrição |
|------|-----------|
| `bool` | `true` ou `false` |
| `string` | Sequência de bytes UTF-8 imutável |
| `int` | Inteiro com tamanho dependente da plataforma (32 ou 64 bits) |
| `int8/16/32/64` | Inteiros com tamanho fixo |
| `uint`, `uint8/16/32/64` | Inteiros sem sinal |
| `float32` | Ponto flutuante de 32 bits |
| `float64` | Ponto flutuante de 64 bits (mais comum) |
| `complex64/128` | Números complexos |
| `byte` | Alias para `uint8` |
| `rune` | Alias para `int32`, representa um Unicode code point |

---

## Diferença entre `=` e `:=`

### `var` + `=` — declaração explícita
```go
var x int = 10
var y float64       // zero value: 0.0
var name string     // zero value: ""
```
- Pode ser usado fora e dentro de funções
- Tipo pode ser omitido se o valor deixar claro
- Inicializa com **zero value** se nenhum valor for atribuído

### `:=` — declaração curta (short variable declaration)
```go
x := 10          // tipo inferido: int
name := "Go"     // tipo inferido: string
```
- Só pode ser usado **dentro de funções**
- Declara **e** inicializa ao mesmo tempo
- O tipo é inferido automaticamente
- Não funciona para redeclarar uma variável já existente no mesmo escopo (exceto se ao menos uma variável do lado esquerdo for nova)

---

## Curiosidades

- **Zero values**: toda variável declarada sem valor inicial recebe um zero value (`0`, `false`, `""`, `nil`) — nunca há valor indefinido em Go.
- **Sem conversão implícita**: `int` e `float64` não se misturam sem conversão explícita. `float64(n)` é obrigatório.
- **`byte` e `rune`**: são apenas aliases — `byte` = `uint8`, `rune` = `int32`. Usados para clareza semântica com texto.
- **`fmt.Fscanln` com ponteiro**: o scanner precisa do endereço da variável (`&val`) para escrever nela — mesmo comportamento do C.
- **`os.Stderr` vs `os.Stdout`**: erros são escritos no stderr por convenção, permitindo redirecionar saída e erro separadamente no terminal.
- **Inteiros têm tamanho de plataforma**: `int` pode ser 32 ou 64 bits dependendo do sistema. Para tamanho fixo, prefira `int32`/`int64`.
- **`complex` existe nativamente**: Go suporta `complex64` e `complex128` como tipos built-in, sem biblioteca externa.
