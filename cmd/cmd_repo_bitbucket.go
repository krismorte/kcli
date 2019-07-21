package cmd

import (
	"fmt"
	"k-cli/conf"
	"strings"

	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var bitBucket = &cobra.Command{
	Use:     "bitbucket",
	Aliases: []string{"bit"},
}

var bitBucketConf = &cobra.Command{
	Use: "conf",
	Run: runAddBitBucketConf,
}

var bitBucketNew = &cobra.Command{
	Use: "new",
	Run: runNewBitbucketrRepo,
}

type BitBucketUser struct {
	Username        string
	Nickname        string
	Account         string
	Account_status  string
	Display_name    string
	Website         string
	Created_on      string
	Uuid            string
	Has_2fa_enabled bool
}

func BitBucketCommand() *cobra.Command {
	bitBucketConf.Flags().StringP("username", "u", "", "Set your username")
	bitBucketConf.Flags().StringP("password", "p", "", "Set your password")
	bitBucketNew.Flags().StringP("name", "n", "", "Set your repository name")
	bitBucket.AddCommand(bitBucketNew)
	bitBucket.AddCommand(bitBucketConf)
	return bitBucket
}

func runNewBitbucketrRepo(cmd *cobra.Command, args []string) {

	description := "Repository made by k-cli"
	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		fmt.Println("--name [-t] is required")
		return
	}

	rows := conf.RunDatabaseQuery("select value from repo_conf where repo_type='BITBUCKET' and key_value='USERNAME'")

	var username string
	var password string

	for rows.Next() {
		rows.Scan(&username)
	}
	rows.Close()

	rows = conf.RunDatabaseQuery("select value from repo_conf where repo_type='BITBUCKET' and key_value='PASSWORD'")

	for rows.Next() {
		rows.Scan(&password)
	}
	rows.Close()

	if username == "" || password == "" {
		fmt.Println("You need to configure your credentials firts")
		return
	}

	c := bitbucket.NewBasicAuth(username, password)

	user, _ := c.User.Profile()
	str := fmt.Sprintf("%v", user)

	opt := &bitbucket.RepositoryOptions{
		RepoSlug:    name,
		IsPrivate:   "true",
		Description: description,
		Owner:       getUsername(str),
	}
	fmt.Println(opt)
	repo, err := c.Repositories.Repository.Create(opt)
	if err != nil {
		err.Error()
	}
	fmt.Println(repo)
}

func runAddBitBucketConf(cmd *cobra.Command, args []string) {

	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
	if username == "" || password == "" {
		fmt.Println("Username and Password are required")
		return
	}
	insert := fmt.Sprintf("INSERT INTO repo_conf ( repo_type , key_value , value) VALUES ('%s', '%s', '%s')", "BITBUCKET", "USERNAME", username)
	insert2 := fmt.Sprintf("INSERT INTO repo_conf ( repo_type , key_value , value) VALUES ('%s', '%s', '%s')", "BITBUCKET", "PASSWORD", password)

	conf.RunDatabaseCommand(insert)
	conf.RunDatabaseCommand(insert2)

}

func getUsername(str string) string {

	ini := strings.Index(str, "username:")
	fim := strings.Index(str, "uuid:")
	username := str[ini:fim]
	username = strings.Replace(username, "username:", "", -1)
	return strings.Trim(username, " ")
}
