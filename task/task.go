package task

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// MaxTitleLength é o número máximo de caracteres de um título.
const MaxTitleLength = 100

// Status guarda a situação de uma tarefa.
type Status string

const (
	Pending    Status = "pendente"
	InProgress Status = "em andamento"
	Done       Status = "concluída"
)

// Valid informa se o status é um dos valores conhecidos.
func (s Status) Valid() bool {
	switch s {
	case Pending, InProgress, Done:
		return true
	}

	return false
}

// Task representa uma tarefa.
type Task struct {
	ID     int
	Title  string
	Status Status
}

// IsDone informa se a tarefa já foi concluída.
func (t Task) IsDone() bool {
	return t.Status == Done
}

// Complete marca a tarefa como concluída.
// Usa receiver ponteiro porque altera a própria tarefa.
func (t *Task) Complete() error {
	if t.IsDone() {
		return ErrAlreadyDone
	}

	t.Status = Done

	return nil
}

// normalizeTitle remove espaços nas pontas e valida o título.
func normalizeTitle(title string) (string, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return "", fmt.Errorf("%w: não pode ser vazio", ErrInvalidTitle)
	}

	if utf8.RuneCountInString(title) > MaxTitleLength {
		return "", fmt.Errorf("%w: máximo de %d caracteres", ErrInvalidTitle, MaxTitleLength)
	}

	return title, nil
}
