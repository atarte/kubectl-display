package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var NamespaceCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "L size of a local directory.",
	Long:  `Twill display the size of a directory with several different options.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Namespace subcommande")
	},
}
