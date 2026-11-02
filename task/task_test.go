package task

import (
	"errors"
	"testing"
)

func TestStatusValid(t *testing.T) {
	tests := []struct {
		name   string
		status Status
		want   bool
	}{
		{name: "pendente", status: Pending, want: true},
		{name: "em andamento", status: InProgress, want: true},
		{name: "concluída", status: Done, want: true},
		{name: "vazio", status: "", want: false},
		{name: "desconhecido", status: "cancelada", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskComplete(t *testing.T) {
	tests := []struct {
		name    string
		status  Status
		wantErr error
	}{
		{name: "pendente", status: Pending},
		{name: "em andamento", status: InProgress},
		{name: "já concluída", status: Done, wantErr: ErrAlreadyDone},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := Task{ID: 1, Title: "Estudar Go", Status: tt.status}

			err := task.Complete()

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Complete() error = %v, want %v", err, tt.wantErr)
			}
			if task.Status != Done {
				t.Errorf("Status = %q, want %q", task.Status, Done)
			}
		})
	}
}
