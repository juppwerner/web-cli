package cmt

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "web-cli",
	Short: "Read Atom feeds",
	Long: `Just a small CLI application.
		  Read Atom feeds.`,
}

// Exec ...
func Exec() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
