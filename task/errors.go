package task

import "errors"

var (
	// ErrNotFound indica que não existe tarefa com o ID informado.
	ErrNotFound = errors.New("tarefa não encontrada")

	// ErrInvalidTitle indica um título vazio ou longo demais.
	ErrInvalidTitle = errors.New("título inválido")

	// ErrInvalidStatus indica um status fora dos valores conhecidos.
	ErrInvalidStatus = errors.New("status inválido")

	// ErrAlreadyDone indica que a tarefa já estava concluída.
	ErrAlreadyDone = errors.New("tarefa já concluída")
)
