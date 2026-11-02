package task

import (
	"errors"
	"strings"
	"testing"
)

func TestStoreAdd(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		wantTitle string
		wantErr   error
	}{
		{name: "título simples", title: "Estudar Go", wantTitle: "Estudar Go"},
		{name: "remove espaços nas pontas", title: "  Estudar Go  ", wantTitle: "Estudar Go"},
		{name: "um caractere", title: "a", wantTitle: "a"},
		{
			name:      "no limite de caracteres",
			title:     strings.Repeat("a", MaxTitleLength),
			wantTitle: strings.Repeat("a", MaxTitleLength),
		},
		{
			name:      "limite conta caracteres e não bytes",
			title:     strings.Repeat("ç", MaxTitleLength),
			wantTitle: strings.Repeat("ç", MaxTitleLength),
		},
		{name: "vazio", title: "", wantErr: ErrInvalidTitle},
		{name: "somente espaços", title: " \t ", wantErr: ErrInvalidTitle},
		{
			name:    "acima do limite",
			title:   strings.Repeat("a", MaxTitleLength+1),
			wantErr: ErrInvalidTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStore()

			got, err := s.Add(tt.title)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Add() error = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				if n := len(s.List()); n != 0 {
					t.Errorf("tarefa inválida foi salva: List() tem %d itens", n)
				}
				return
			}

			want := Task{ID: 1, Title: tt.wantTitle, Status: Pending}
			if got != want {
				t.Errorf("Add() = %+v, want %+v", got, want)
			}
		})
	}
}

func TestStoreAddIDs(t *testing.T) {
	s := NewStore()

	mustAdd(t, s, "primeira")
	if _, err := s.Add(""); err == nil {
		t.Fatal("Add(\"\") deveria falhar")
	}
	second := mustAdd(t, s, "segunda")

	if second.ID != 2 {
		t.Errorf("ID = %d, want 2 (falha não deve consumir ID)", second.ID)
	}
}

func TestStoreList(t *testing.T) {
	t.Run("vazia", func(t *testing.T) {
		got := NewStore().List()

		if got == nil || len(got) != 0 {
			t.Errorf("List() = %#v, want slice vazio", got)
		}
	})

	t.Run("mantém ordem de cadastro", func(t *testing.T) {
		s := NewStore()
		titles := []string{"a", "b", "c"}
		for _, title := range titles {
			mustAdd(t, s, title)
		}

		got := s.List()

		if len(got) != len(titles) {
			t.Fatalf("List() tem %d itens, want %d", len(got), len(titles))
		}
		for i, task := range got {
			if task.ID != i+1 || task.Title != titles[i] {
				t.Errorf("List()[%d] = %+v", i, task)
			}
		}
	})
}

func TestStoreGet(t *testing.T) {
	s := NewStore()
	created := mustAdd(t, s, "Estudar Go")

	tests := []struct {
		name    string
		id      int
		want    Task
		wantErr error
	}{
		{name: "existente", id: created.ID, want: created},
		{name: "inexistente", id: 99, wantErr: ErrNotFound},
		{name: "zero", id: 0, wantErr: ErrNotFound},
		{name: "negativo", id: -1, wantErr: ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Get(tt.id)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Get() error = %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Get() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestStoreUpdateStatus(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		status  Status
		wantErr error
	}{
		{name: "para em andamento", id: 1, status: InProgress},
		{name: "para concluída", id: 1, status: Done},
		{name: "para pendente", id: 1, status: Pending},
		{name: "status inválido", id: 1, status: "cancelada", wantErr: ErrInvalidStatus},
		{name: "status vazio", id: 1, status: "", wantErr: ErrInvalidStatus},
		{name: "tarefa inexistente", id: 2, status: Done, wantErr: ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStore()
			mustAdd(t, s, "Estudar Go")

			err := s.UpdateStatus(tt.id, tt.status)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("UpdateStatus() error = %v, want %v", err, tt.wantErr)
			}

			want := tt.status
			if tt.wantErr != nil {
				want = Pending
			}
			if got, _ := s.Get(1); got.Status != want {
				t.Errorf("Status = %q, want %q", got.Status, want)
			}
		})
	}
}

func TestStoreComplete(t *testing.T) {
	t.Run("conclui tarefa pendente", func(t *testing.T) {
		s := NewStore()
		created := mustAdd(t, s, "Estudar Go")

		if err := s.Complete(created.ID); err != nil {
			t.Fatalf("Complete() error = %v", err)
		}

		got, _ := s.Get(created.ID)
		if !got.IsDone() {
			t.Errorf("Status = %q, want %q", got.Status, Done)
		}
	})

	t.Run("tarefa já concluída", func(t *testing.T) {
		s := NewStore()
		created := mustAdd(t, s, "Estudar Go")
		if err := s.Complete(created.ID); err != nil {
			t.Fatal(err)
		}

		err := s.Complete(created.ID)

		if !errors.Is(err, ErrAlreadyDone) {
			t.Errorf("Complete() error = %v, want %v", err, ErrAlreadyDone)
		}
	})

	t.Run("tarefa inexistente", func(t *testing.T) {
		err := NewStore().Complete(1)

		if !errors.Is(err, ErrNotFound) {
			t.Errorf("Complete() error = %v, want %v", err, ErrNotFound)
		}
	})
}

func mustAdd(t *testing.T, s *Store, title string) Task {
	t.Helper()

	task, err := s.Add(title)
	if err != nil {
		t.Fatalf("Add(%q) error = %v", title, err)
	}

	return task
}
