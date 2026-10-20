package task

// Status guarda a situação de uma tarefa.
type Status string

const (
	Pending    Status = "pendente"
	InProgress Status = "em andamento"
	Done       Status = "concluída"
)

// Task representa uma tarefa.
type Task struct {
	ID     int
	Title  string
	Status Status
}
