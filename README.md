# GRPC tools

Playground repo for gRPC. Example roughly based on https://github.com/gogo/grpc-example


## Required packages for development
* `go`
* `direnv`
* `wget`
* `bsdtar`

## Setup the development environment
The environment variables are described in .envrc. This enables us to
create a local development setup. Its purpose modify global env variables to download and keep all dependencies in the .direnv directory.
* Install direnv from your package manager and setup direnv shell hook for the terminal you're using (bash, zsh, etc...): https://github.com/direnv/direnv
* run `$ direnv allow .` to allow loading environment variables from .envrc
* run `$ sh setup.sh` to install local development tools

## Resetting the development env:
* Make sure direnv works by `$ echo $GOPATH`. The GOPATH needs to be set to .direnv directory.
* `$ rm -rf .direnv`
* `$ sh setup.sh`

Install VSCode direnv extension: https://marketplace.visualstudio.com/items?itemName=Rubymaniac.vscode-direnv


## Generating the application:
* `$ cd application`
* `$ make generate`

## Testing:
* `$ cd application`
* `$ make test`