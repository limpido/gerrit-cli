package cmd

import (
	"fmt"

	"github.com/limpido/gerrit-cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inside gerrit push")
		fmt.Printf("The date is %s\n", utils.GitRepoUrl())
	},
}
