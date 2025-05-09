package cmd

import (
	"fmt"

	"github.com/limpido/gerrit-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(branchCmd)
}

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		origin := util.GitOrigin()
		upstream := util.GitUpstream()
		util.Execute(fmt.Sprintf("git checkout %s", upstream))
		util.Execute(fmt.Sprintf("git checkout -b %s", branch))
		util.Execute(fmt.Sprintf("git branch -u %s/%s", origin, upstream))
	},
}
