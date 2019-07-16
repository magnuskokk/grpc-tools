# GRPC tools

Playground repo for gRPC. Example roughly based on https://github.com/gogo/grpc-example


## Required packages for development
* `go`
* `npm` (for downloading swagger-ui-dist)
* `direnv`
* `wget`
* `bsdtar`

## Setup the development environment
Do not open any IDE-s yet. Let's first create a local development environment something akin to virtualenv for python.

* Install direnv from your package manager and setup direnv shell hook for the terminal you're using (bash, zsh, etc...): https://github.com/direnv/direnv.
* `$ direnv allow .` to allow loading environment variables from .envrc.
* `$ setup` to install tools necessary for code generation.

### VSCode environment
Install direnv extension: https://marketplace.visualstudio.com/items?itemName=Rubymaniac.vscode-direnv. Then open the directory. Click allow when asked to allow direnv. Do not click install anything before allowing direnv, this way you get the right GOPATH set to .direnv  with all Go related plugins. Now when asked about missing go tools, press Install all in the right corner. When it says "All tools successfully installed. You're ready to Go :)." then restart VSCode.

## Generate API, test and view swagger docs:
* `$ make`
* `$ make test` to run all tests.
* `$ cd serverapp; make docserver`
* Open browser at `http://localhost:8000/openapi-ui`
