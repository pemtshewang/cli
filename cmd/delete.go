/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/pemtshewang/cli/pkg/tasks"
	"github.com/spf13/cobra"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specific task by an id",
	Run: func(cmd *cobra.Command, args []string) {
		Inputid, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println("There was an error while parsing the arg id")
			return
		}
		id, err := strconv.Atoi(Inputid)
		if err != nil {
			fmt.Printf("Error while trying to parse the given input")
			return
		}
		tasks, err := tasks.DeleteTask(id)
		if err != nil {
			fmt.Println("The task with given id doesn't exist")
			return
		}
		fmt.Println(tasks)
		fmt.Println("The task has been deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("id", "i", "", "Delete task by an id")
}
