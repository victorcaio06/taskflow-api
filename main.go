package main

import (
	"fmt"
	"os"

	"github.com/victorcaio06/taskflow-api/cli"
	"github.com/victorcaio06/taskflow-api/task"
)

func main() {
	app := cli.New(task.NewStore(), os.Stdin, os.Stdout)

	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
