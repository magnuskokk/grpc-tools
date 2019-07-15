# GRPC tools

Playground repo for gRPC. Example roughly based on https://github.com/gogo/grpc-example


## Required packages for development
* `go`
* `npm` (for downloading swagger-ui-dist)
* `direnv`
* `wget`
* `bsdtar`

## Setup the development environment
The environment variables are described in .envrc. This enables us to
create a local development setup. Its purpose modify global env variables to download and keep all dependencies in the .direnv directory.
* Install direnv from your package manager and setup direnv shell hook for the terminal you're using (bash, zsh, etc...): https://github.com/direnv/direnv
* `$ direnv allow .` to allow loading environment variables from .envrc
* `$ setup` to install local development tools

Do not open the directory in VSCode yet. Make sure to run `setup` in a terminal before opening any IDE-s. Install VSCode direnv extension: https://marketplace.visualstudio.com/items?itemName=Rubymaniac.vscode-direnv. Then open the directory. Click allow when asked to allow direnv. When VSCode asks you to install some go tools, press Install all in the right corner. When it says "All tools successfully installed. You're ready to Go :)." then restart VSCode. That way you have the right GOPATH set to .direnv directory.

## Testing:
* `$ make test`

## View generated swagger docs:
* `$ make`
* `$ make docserver`
* Open browser at `http://localhost:8000/openapi-ui`
