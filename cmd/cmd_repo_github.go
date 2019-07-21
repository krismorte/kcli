package cmd

import (
	"context"
	"fmt"
	"k-cli/conf"

	"github.com/google/go-github/v27/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var gitHub = &cobra.Command{
	Use:     "github",
	Aliases: []string{"gh"},
}

var gitHubConf = &cobra.Command{
	Use: "conf",
	Run: runAddGitHubConf,
}

var gitHubNew = &cobra.Command{
	Use: "new",
	Run: runNewGitRepo,
}

func GitHubCommand() *cobra.Command {
	gitHubConf.Flags().StringP("token", "t", "", "Set your token")
	gitHubNew.Flags().StringP("name", "n", "", "Set your repository name")
	gitHub.AddCommand(gitHubConf)
	gitHub.AddCommand(gitHubNew)
	return gitHub
}

func runNewGitRepo(cmd *cobra.Command, args []string) {
	private := false
	description := "Repository made by k-cli"
	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		fmt.Println("--name [-t] is required")
		return
	}

	rows := conf.RunDatabaseQuery("select value from repo_conf where repo_type='GITHUB' and key_value='TOKEN'")

	var token string

	for rows.Next() {
		rows.Scan(&token)
	}
	rows.Close()

	if token == "" {
		fmt.Println("You need to configure your token firts")
		return
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	r := &github.Repository{Name: &name, Private: &private, Description: &description}
	client.Repositories.Create(ctx, "", r)
	fmt.Println("Create the repo")
}

func runAddGitHubConf(cmd *cobra.Command, args []string) {

	token, _ := cmd.Flags().GetString("token")

	if token == "" {
		fmt.Println("--token [-t] is required")
		return
	}

	insert := fmt.Sprintf("INSERT INTO repo_conf ( repo_type , key_value , value) VALUES ('%s', '%s', '%s')", "GITHUB", "TOKEN", token)
	conf.RunDatabaseCommand(insert)
}
