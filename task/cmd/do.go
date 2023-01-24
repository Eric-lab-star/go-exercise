/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Eric-lab-star/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as complete",

	Run: func(cmd *cobra.Command, args []string) {
		ids := []int{}
		for _, v := range args {
			id, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(v, "is not valid arg")
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
		tasks, err := db.AllTasks()
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTask(task.ID)
			if err != nil {
				fmt.Printf("Faild to mark \"%d\" as completed. Error: %v\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed\n", id)
			}
		}
		if err != nil {
			fmt.Printf("something went wrong")
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
