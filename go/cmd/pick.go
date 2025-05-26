package cmd

import (
	"fmt"

	"github.com/limpido/gerrit-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pickCmd)
}

var pickCmd = &cobra.Command{
	Use:   "pick <branch> [<branch>...]",
	Short: "Cherrypick HEAD commit to the specified branch(es) on remote server",
	Args:  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		head := util.GitHead()
		origin := util.GitOrigin()
		curBranch := util.GitBranch()
		for _, branch := range args {
			tmpBranch := fmt.Sprintf("pick/%s", head)
			util.Execute(fmt.Sprintf("git checkout %s", branch))
			util.Execute(fmt.Sprintf("git checkout -b %s", tmpBranch))
			util.Execute(fmt.Sprintf("git cherry-pick %s", head))
			util.Execute(fmt.Sprintf("git push %s HEAD:refs/for/%s", origin, branch))
			util.Execute(fmt.Sprintf("git checkout %s", curBranch))
			util.Execute(fmt.Sprintf("git branch -D %s", tmpBranch))
		}
	},
}
