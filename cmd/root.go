package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	// Use:   "display",
	Short: "List the size of a local directory.",
	Long:  `This command will display the size of a directory with several different options.`,
}

func init() {
	RootCmd.AddCommand(NamespaceCmd)
}
