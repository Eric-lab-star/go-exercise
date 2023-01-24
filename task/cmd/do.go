/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

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
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
