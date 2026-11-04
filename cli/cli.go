// Package cli implementa o menu interativo de tarefas.
package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/victorcaio06/taskflow-api/task"
)

// CLI lê comandos de uma entrada e escreve as respostas em uma saída.
type CLI struct {
	store   *task.Store
	scanner *bufio.Scanner
	out     io.Writer
}

// New cria uma CLI. Entrada e saída são recebidas como parâmetro
// para que possam ser trocadas nos testes.
func New(store *task.Store, in io.Reader, out io.Writer) *CLI {
	return &CLI{
		store:   store,
		scanner: bufio.NewScanner(in),
		out:     out,
	}
}

// Run executa o menu até a opção de sair ou o fim da entrada.
func (c *CLI) Run() error {
	for {
		c.printMenu()

		option, err := c.read("Opção: ")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch option {
		case "1":
			err = c.addTask()
		case "2":
			c.listTasks()
		case "3":
			err = c.changeStatus()
		case "4":
			err = c.showTask()
		case "5":
			err = c.completeTask()
		case "0":
			return nil
		default:
			c.println("Opção inválida")
		}

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

func (c *CLI) printMenu() {
	c.println()
	c.println("1 - Cadastrar tarefa")
	c.println("2 - Listar tarefas")
	c.println("3 - Alterar status")
	c.println("4 - Consultar tarefa")
	c.println("5 - Concluir tarefa")
	c.println("0 - Sair")
}

// read mostra uma mensagem e devolve o que foi digitado.
// Retorna io.EOF quando a entrada termina.
func (c *CLI) read(message string) (string, error) {
	fmt.Fprint(c.out, message)

	if !c.scanner.Scan() {
		if err := c.scanner.Err(); err != nil {
			return "", fmt.Errorf("ler entrada: %w", err)
		}
		return "", io.EOF
	}

	return strings.TrimSpace(c.scanner.Text()), nil
}

// readID lê um ID. ok é false quando o valor digitado não é um número.
func (c *CLI) readID() (id int, ok bool, err error) {
	text, err := c.read("ID da tarefa: ")
	if err != nil {
		return 0, false, err
	}

	id, convErr := strconv.Atoi(text)
	if convErr != nil {
		c.println("ID inválido")
		return 0, false, nil
	}

	return id, true, nil
}

func (c *CLI) addTask() error {
	title, err := c.read("Título: ")
	if err != nil {
		return err
	}

	t, err := c.store.Add(title)
	if err != nil {
		c.println(err)
		return nil
	}

	c.printf("Tarefa %d cadastrada\n", t.ID)

	return nil
}

func (c *CLI) listTasks() {
	tasks := c.store.List()

	if len(tasks) == 0 {
		c.println("Nenhuma tarefa cadastrada")
		return
	}

	for _, t := range tasks {
		c.printTask(t)
	}
}

func (c *CLI) showTask() error {
	id, ok, err := c.readID()
	if err != nil || !ok {
		return err
	}

	t, err := c.store.Get(id)
	if err != nil {
		c.println(err)
		return nil
	}

	c.printTask(t)

	return nil
}

func (c *CLI) completeTask() error {
	id, ok, err := c.readID()
	if err != nil || !ok {
		return err
	}

	if err := c.store.Complete(id); err != nil {
		c.println(err)
		return nil
	}

	c.println("Tarefa concluída")

	return nil
}

var statusOptions = map[string]task.Status{
	"1": task.Pending,
	"2": task.InProgress,
	"3": task.Done,
}

func (c *CLI) changeStatus() error {
	id, ok, err := c.readID()
	if err != nil || !ok {
		return err
	}

	c.println("1 - Pendente")
	c.println("2 - Em andamento")
	c.println("3 - Concluída")

	option, err := c.read("Novo status: ")
	if err != nil {
		return err
	}

	status, exists := statusOptions[option]
	if !exists {
		c.println("Status inválido")
		return nil
	}

	if err := c.store.UpdateStatus(id, status); err != nil {
		c.println(err)
		return nil
	}

	c.println("Status alterado")

	return nil
}

func (c *CLI) printTask(t task.Task) {
	c.printf("[%d] %s (%s)\n", t.ID, t.Title, t.Status)
}

func (c *CLI) println(a ...any) {
	fmt.Fprintln(c.out, a...)
}

func (c *CLI) printf(format string, a ...any) {
	fmt.Fprintf(c.out, format, a...)
}
