package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of azure-terraform-generator",
	Long:  `All software has versions. This is azure-terraform-generator`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Azure Terraform Generator v0.1.0 -- HEAD")
	},
}
