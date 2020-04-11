# k-cli

The main goal here is build a personal CLI to help the developer's life :D

## Build the CLI

`git clone https://github.com/krismorte/kcli`

`go build -o kcli main.go`

`sudo chmod +x  kcli && mv kcli  /usr/bin/`

## How to use

Creating repositories in Github and Bitbucket

`kcli repo github conf -t <YOUR_GITHUB_TOKEN>`

`kcli repo github new -n <MY_REPO>`

`kcli repo bitbucket conf -u <YOUR_BITBUCKET_USER> -p <YOUR_BITBUCKET_PASS>`

`kcli repo bitbucket new -n <MY_REPO>`

## TODO

| Provider        | Action           | Status  |
| ------------- |:-------------:| -----:|
| github      | Create Repo | OK |
| bitbucket      | Create Repo      |   OK |
| AWS | Create Bucket      |    --- |
| AWS | Upalod to  Bucket      |    --- |
| AWS | Start/Stop RDS      |    --- |
