package cmd

import (
    //"fmt"

	"github.com/spf13/cobra"
    "github.com/afro-juliano/cli-app/todo"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add will create a new todo item to the list`,
    Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
    var items = []todo.Item{}
    for _, x := range args {
        items = append(items, todo.Item{Text:x})
    }
    todo.SaveItems("x", items)
}

func init() {
	rootCmd.AddCommand(addCmd)

}
