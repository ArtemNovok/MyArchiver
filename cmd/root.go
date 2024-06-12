package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Simple archiver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		HandleErr(err)
	}
}

func HandleErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
