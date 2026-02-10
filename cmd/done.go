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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "To mark the task as done",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println("Error while reading the parameter")
		}
		num, err := strconv.Atoi(id)
		if err == nil {
			tasks.CompleteTask(num)
			fmt.Println("Updated successfully")
			return
		}
		fmt.Println(err)
		fmt.Println("Not task exists for it now as of now")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	doneCmd.Flags().StringP("id", "i", "", "to mark particular task as done")

}
