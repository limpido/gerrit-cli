package cmd

import (
	"fmt"

	"github.com/limpido/gerrit-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Amend the HEAD commit and push.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		origin := util.GitOrigin()
		upstream := util.GitUpstream()
		util.Execute("git commit --amend --no-edit -s")
		c := fmt.Sprintf("git push %s HEAD:refs/for/%s --no-thin", origin, upstream)
		util.Execute(c)
	},
}
