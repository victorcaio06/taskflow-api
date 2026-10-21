package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/victorcaio06/taskflow-api/task"
)

func main() {
	store := task.NewStore()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("1 - Cadastrar tarefa")
		fmt.Println("2 - Listar tarefas")
		fmt.Println("0 - Sair")

		option := read(scanner, "Opção: ")

		switch option {
		case "1":
			addTask(store, scanner)
		case "2":
			listTasks(store)
		case "0":
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}

// read mostra uma mensagem e devolve o que foi digitado.
func read(scanner *bufio.Scanner, message string) string {
	fmt.Print(message)

	// Sem mais entrada (Ctrl+D), encerra o programa.
	if !scanner.Scan() {
		os.Exit(0)
	}

	return strings.TrimSpace(scanner.Text())
}

func addTask(store *task.Store, scanner *bufio.Scanner) {
	title := read(scanner, "Título: ")

	if title == "" {
		fmt.Println("O título não pode ser vazio")
		return
	}

	t := store.Add(title)
	fmt.Printf("Tarefa %d cadastrada\n", t.ID)
}

func listTasks(store *task.Store) {
	tasks := store.List()

	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa cadastrada")
		return
	}

	for _, t := range tasks {
		fmt.Printf("[%d] %s (%s)\n", t.ID, t.Title, t.Status)
	}
}
