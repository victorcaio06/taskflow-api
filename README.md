# taskflow-api

CLI simples para cadastrar, listar e alterar o status de tarefas.
As tarefas ficam em memória, então são perdidas ao fechar o programa.

## Como rodar

```bash
go run .
```

## Estrutura

- `main.go`: menu da CLI
- `task/task.go`: tipos `Task` e `Status`
- `task/store.go`: operações de cadastro, listagem e alteração de status

## Verificações

```bash
go fmt ./...
go vet ./...
```
