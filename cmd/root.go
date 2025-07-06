package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "github.com/WelintonJunior/my-sentiment",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help Message for toggle")
}
