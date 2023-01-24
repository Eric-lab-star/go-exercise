package cmd

import (
	"fmt"
	"strings"

	"github.com/Eric-lab-star/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Printf("Could not Add task: %s", task)
		}
		fmt.Printf("Added \"%s\" to list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
