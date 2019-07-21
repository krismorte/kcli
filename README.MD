# k-cli

The main object here is build a personal CLI to help the developer's life :D

## Build the CLI

`git clone https://github.com/krismorte/kcli`

`go build -o kcli main.go`

`sudo mv kcli  /usr/bin/`

## Hot to use

Creating repositories in Github and Bitbucket

`kcli repo github cont -t <YOUR_GITHUB_TOKEN>`

`kcli repo github new -n <MY_REPO>`

`kcli repo bitbucket cont -u <YOUR_BITBUCKET_USER> -p <YOUR_BITBUCKET_PASS>`

`kcli repo bitbucket new -n <MY_REPO>`

##TODO

| Provider        | Action           | Status  |
| ------------- |:-------------:| -----:|
| github      | Create Repo | OK |
| bitbucket      | Create Repo      |   OK |
| AWS | Create Bucket      |    --- |
| AWS | Upalod to  Bucket      |    --- |
| AWS | Start/Stop RDS      |    --- |