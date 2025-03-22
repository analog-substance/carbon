#! /bin/bash
gomarkdoc --output 'docs/docs/code/{{.Dir}}/_index.md' --template-file package=.hugo/gomarkdoc/package.gotxt --template-file file=./.hugo/gomarkdoc/file.gotxt --exclude-dirs ./.hugo/... --exclude-dirs ./deployments/... ./... --repository.url https://github.com/analog-substance/carbon
