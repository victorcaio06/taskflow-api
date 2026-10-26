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

func (s *Store) List() []Task {
	list := []Task{}

	for id := 1; id < s.nextID; id++ {
		list = append(list, s.tasks[id])
	}

	return list
}

func (s *Store) UpdateStatus(id int, status Status) error {
	t, exists := s.tasks[id]
	if !exists {
		return fmt.Errorf("%w: id %d", ErrNotFound, id)
	}

	t.Status = status
	s.tasks[id] = t

	return nil
}
