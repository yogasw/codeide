package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Simplifying VS Code for Open Specific Programming Languages",
	Run:   detectLanguage,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func detectLanguage(cmd *cobra.Command, args []string) {

}
