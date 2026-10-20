package task

// Store guarda as tarefas em memória.
type Store struct {
	tasks  map[int]Task
	nextID int
}

// NewStore cria um Store vazio.
func NewStore() *Store {
	return &Store{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

// Add cadastra uma nova tarefa com status pendente.
func (s *Store) Add(title string) Task {
	t := Task{
		ID:     s.nextID,
		Title:  title,
		Status: Pending,
	}

	s.tasks[t.ID] = t
	s.nextID++

	return t
}

// List retorna as tarefas em ordem de ID.
func (s *Store) List() []Task {
	list := []Task{}

	for id := 1; id < s.nextID; id++ {
		list = append(list, s.tasks[id])
	}

	return list
}
