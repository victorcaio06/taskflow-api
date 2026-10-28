package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		fmt.Println("3 - Alterar status")
		fmt.Println("0 - Sair")

		option := read(scanner, "Opção: ")

		switch option {
		case "1":
			addTask(store, scanner)
		case "2":
			listTasks(store)
		case "3":
			changeStatus(store, scanner)
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
	t, err := store.Add(read(scanner, "Título: "))
	if err != nil {
		fmt.Println(err)
		return
	}

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

func changeStatus(store *task.Store, scanner *bufio.Scanner) {
	id, err := strconv.Atoi(read(scanner, "ID da tarefa: "))
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	statusOptions := map[string]task.Status{
		"1": task.Pending,
		"2": task.InProgress,
		"3": task.Done,
	}

	fmt.Println("1 - Pendente")
	fmt.Println("2 - Em andamento")
	fmt.Println("3 - Concluída")

	status, exists := statusOptions[read(scanner, "Novo status: ")]
	if !exists {
		fmt.Println("Status inválido")
		return
	}

	if err := store.UpdateStatus(id, status); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Status alterado")
}
