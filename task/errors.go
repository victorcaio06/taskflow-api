package task

import "errors"

// ErrNotFound indica que não existe tarefa com o ID informado.
var ErrNotFound = errors.New("tarefa não encontrada")
