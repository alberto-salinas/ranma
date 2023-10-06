package internals

import "github.com/spf13/cobra"
import "os"
import "fmt"

var rootCmd = cobra.Command{
	Use: "randma",
	Short: "randma is a tool to solve leet code problems in GO",
	Run: func(cmd *cobra.Command, args [] string) {
		fmt.Println("here we go")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}