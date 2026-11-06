package cli

import (
	"errors"
	"strings"
	"testing"

	"github.com/victorcaio06/taskflow-api/task"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "cadastra tarefa",
			input: []string{"1", "Estudar Go"},
			want:  []string{"Tarefa 1 cadastrada"},
		},
		{
			name:  "rejeita título vazio",
			input: []string{"1", "   "},
			want:  []string{"título inválido: não pode ser vazio"},
		},
		{
			name:  "lista sem tarefas",
			input: []string{"2"},
			want:  []string{"Nenhuma tarefa cadastrada"},
		},
		{
			name:  "lista tarefas cadastradas",
			input: []string{"1", "a", "1", "b", "2"},
			want:  []string{"[1] a (pendente)", "[2] b (pendente)"},
		},
		{
			name:  "consulta tarefa",
			input: []string{"1", "Estudar Go", "4", "1"},
			want:  []string{"[1] Estudar Go (pendente)"},
		},
		{
			name:  "consulta tarefa inexistente",
			input: []string{"4", "7"},
			want:  []string{"tarefa não encontrada: id 7"},
		},
		{
			name:  "ID não numérico",
			input: []string{"4", "abc"},
			want:  []string{"ID inválido"},
		},
		{
			name:  "conclui tarefa",
			input: []string{"1", "Estudar Go", "5", "1", "4", "1"},
			want:  []string{"Tarefa concluída", "[1] Estudar Go (concluída)"},
		},
		{
			name:  "conclui tarefa já concluída",
			input: []string{"1", "Estudar Go", "5", "1", "5", "1"},
			want:  []string{"concluir tarefa 1: tarefa já concluída"},
		},
		{
			name:  "altera status",
			input: []string{"1", "Estudar Go", "3", "1", "2", "2"},
			want:  []string{"Status alterado", "[1] Estudar Go (em andamento)"},
		},
		{
			name:  "rejeita opção de status inválida",
			input: []string{"1", "Estudar Go", "3", "1", "9"},
			want:  []string{"Status inválido"},
		},
		{
			name:  "opção de menu inválida",
			input: []string{"x"},
			want:  []string{"Opção inválida"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := run(t, tt.input...)

			for _, want := range tt.want {
				if !strings.Contains(out, want) {
					t.Errorf("saída não contém %q\nsaída:\n%s", want, out)
				}
			}
		})
	}
}

func TestRunStops(t *testing.T) {
	t.Run("opção sair ignora o resto da entrada", func(t *testing.T) {
		out := run(t, "0", "1", "não deve ser cadastrada")

		if strings.Contains(out, "cadastrada") {
			t.Errorf("continuou lendo após sair:\n%s", out)
		}
	})

	t.Run("entrada vazia", func(t *testing.T) {
		run(t)
	})

	t.Run("entrada termina no meio de um comando", func(t *testing.T) {
		run(t, "1")
	})
}

func TestRunReadError(t *testing.T) {
	app := New(task.NewStore(), errReader{}, &strings.Builder{})

	if err := app.Run(); err == nil {
		t.Fatal("Run() deveria retornar erro de leitura")
	}
}

// run executa a CLI com uma linha de entrada por argumento e devolve a saída.
func run(t *testing.T, lines ...string) string {
	t.Helper()

	in := strings.NewReader(strings.Join(lines, "\n"))
	var out strings.Builder

	if err := New(task.NewStore(), in, &out).Run(); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	return out.String()
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, errors.New("falha de leitura")
}
