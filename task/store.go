package task

import "fmt"

type Store struct {
	tasks  map[int]Task
	nextID int
}

func NewStore() *Store {
	return &Store{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (s *Store) Add(title string) (Task, error) {
	title, err := normalizeTitle(title)
	if err != nil {
		return Task{}, err
	}

	t := Task{
		ID:     s.nextID,
		Title:  title,
		Status: Pending,
	}

	s.tasks[t.ID] = t
	s.nextID++

	return t, nil
}

func (s *Store) Get(id int) (Task, error) {
	t, exists := s.tasks[id]
	if !exists {
		return Task{}, fmt.Errorf("%w: id %d", ErrNotFound, id)
	}

	return t, nil
}

func (s *Store) List() []Task {
	list := []Task{}

	for id := 1; id < s.nextID; id++ {
		list = append(list, s.tasks[id])
	}

	return list
}

func (s *Store) UpdateStatus(id int, status Status) error {
	if !status.Valid() {
		return fmt.Errorf("%w: %q", ErrInvalidStatus, status)
	}

	t, err := s.Get(id)
	if err != nil {
		return err
	}

	t.Status = status
	s.tasks[id] = t

	return nil
}

func (s *Store) Complete(id int) error {
	t, err := s.Get(id)
	if err != nil {
		return err
	}

	if err := t.Complete(); err != nil {
		return fmt.Errorf("concluir tarefa %d: %w", id, err)
	}

	s.tasks[id] = t

	return nil
}
