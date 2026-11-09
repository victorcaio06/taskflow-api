# taskflow-api

CLI simples para cadastrar, consultar, listar, alterar o status e concluir tarefas.
As tarefas ficam em memória, então são perdidas ao fechar o programa.

## Requisitos

- Go 1.27 ou superior

## Como rodar

```bash
go run .
```

Ou gerando o binário:

```bash
go build -o taskflow-api .
./taskflow-api
```

O menu aceita as opções:

| Opção | Ação                |
| ----- | ------------------- |
| 1     | Cadastrar tarefa    |
| 2     | Listar tarefas      |
| 3     | Alterar status      |
| 4     | Consultar tarefa    |
| 5     | Concluir tarefa     |
| 0     | Sair                |

O programa também encerra quando a entrada termina (Ctrl+D).

## Estrutura

```
.
├── main.go          # monta as dependências e inicia a CLI
├── cli/             # entrada e saída: menu, leitura e mensagens
│   ├── cli.go
│   └── cli_test.go
└── task/            # núcleo: regras de negócio, sem entrada e saída
    ├── errors.go    # erros de domínio
    ├── task.go      # tipos Task e Status e validações
    ├── store.go     # operações de cadastro, consulta, listagem e conclusão
    └── *_test.go
```

O pacote `task` não conhece terminal nem `fmt.Print`: ele só recebe dados e devolve
valores ou erros. O pacote `cli` recebe um `io.Reader` e um `io.Writer`, o que
permite testá-lo com `strings.Reader` e `strings.Builder` no lugar de `os.Stdin` e
`os.Stdout`.

Não há interface para o armazenamento: existe uma única implementação em memória,
então uma abstração ainda não tem uso concreto.

## Regras de negócio

- O título é obrigatório, tem espaços das pontas removidos e aceita até 100 caracteres.
- Toda tarefa nova começa como `pendente`.
- Os status válidos são `pendente`, `em andamento` e `concluída`.
- Concluir uma tarefa já concluída é um erro.

Os erros de domínio ficam em `task/errors.go` (`ErrNotFound`, `ErrInvalidTitle`,
`ErrInvalidStatus`, `ErrAlreadyDone`). Eles são embrulhados com `fmt.Errorf` e `%w`
para adicionar contexto, e podem ser identificados com `errors.Is`:

```go
if errors.Is(err, task.ErrNotFound) {
	// ...
}
```

## Testes

```bash
go test ./...
```

Com cobertura e saída detalhada:

```bash
go test -cover ./...
go test -v ./...
```

Para rodar um teste ou subteste específico:

```bash
go test ./task -run '^TestStoreAdd$/acima_do_limite'
```

## Verificações

```bash
go fmt ./...
go vet ./...
```
