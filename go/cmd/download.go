package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/limpido/gerrit-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}

var downloadCmd = &cobra.Command{
	Use:   "download <change> [<change>...]",
	Short: "Download the specified change(s) from remote to current branch",
	Args:  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		repoUrl := util.GitRepoUrl()
		for _, change := range args {
			var resp map[string]json.RawMessage
			var curPatchSet map[string]json.RawMessage
			json.Unmarshal([]byte(util.Query(change)), &resp)
			json.Unmarshal([]byte(resp["currentPatchSet"]), &curPatchSet)
			ref := curPatchSet["ref"]
			util.Execute(fmt.Sprintf("git fetch %s %s", repoUrl, ref))
			util.Execute("git cherry-pick FETCH_HEAD")
		}
	},
}
