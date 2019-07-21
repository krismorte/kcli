package cmd

import (
	"github.com/spf13/cobra"
)

var RepoSQLCommand = "INSERT INTO repo_conf ( repo_tipe , key_value , value) VALUES (?, ?, ?)"
var RepoTableName = "repo_conf"

var repo = &cobra.Command{
	Use:     "repo",
	Aliases: []string{},
	Short:   "Help with your repositories.",
	Long:    "Help with your repositories",
	//Run:     runAddGitConf,
}

func RepoCommand() *cobra.Command {
	repo.AddCommand(GitHubCommand())
	repo.AddCommand(BitBucketCommand())
	return repo
}
