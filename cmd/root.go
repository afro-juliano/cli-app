package cmd

import (
    "fmt"
	"os"

	"github.com/spf13/cobra"
    //"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cli-app",
	Short: "cli-app is a todo application",
	Long: `cli-app will help you get more done in less time.
            It's designed to be as simple as possible to Help
            you acomplish your goals`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("cli-app is good")
    },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
        fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
