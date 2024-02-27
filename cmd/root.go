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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
