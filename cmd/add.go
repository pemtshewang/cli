/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/pemtshewang/cli/pkg/tasks"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		priority, _ := cmd.Flags().GetString("priority")

		if title == "" {
			fmt.Println("Task title is required for this action")
			return
		}

		err := tasks.AddTask(title, priority)
		if err != nil {
			fmt.Println("Error adding task: ", err)
			return
		}
		fmt.Println("Task added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("title", "t", "", "Title of the task")
	addCmd.Flags().StringP("priority", "p", "medium", "Priority of the task")
}
