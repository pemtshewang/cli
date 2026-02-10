/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/pemtshewang/cli/pkg/tasks"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the task added in json",
	Run: func(cmd *cobra.Command, args []string) {
		TaskList, err := tasks.LoadTasks()
		if err != nil {
			fmt.Println("error loading the task")
		}
		if len(TaskList) == 0 {
			fmt.Println("No task has been added")
			return
		}
		fmt.Printf("Title\t\tPriority\n")
		for i := range TaskList {
			fmt.Printf("%s\t\t%s\n", TaskList[i].Title, TaskList[i].Priority)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
