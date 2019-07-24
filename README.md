# GRPC tools

Playground repo for gRPC. Example roughly based on https://github.com/gogo/grpc-example


## Required packages for development
* `direnv`
* `docker-compose`

## Setup the development environment
Do not open any IDE-s yet. Let's first create a local development environment something akin to virtualenv for python.

* Install direnv from your package manager and setup direnv shell hook for the terminal you're using (bash, zsh, etc...): https://github.com/direnv/direnv.
* `$ direnv allow .` to setup a local environment based on .envrc.
* `$ ./setup.sh` to build and install dev dependencies.

### VSCode
direnv extension: https://marketplace.visualstudio.com/items?itemName=Rubymaniac.vscode-direnv. This allows installing all tools to .direnv directory.

## List all commands
* `$ tusk`

## Some useful commands
* `$ tusk protolint`
* `$ tusk protoc`
* `$ tusk test.go`

## Run server
* `$ tusk serve-echo`
* Make sure it's responding: `$ curl 'http://localhost:8081/echo?Message=test'` You should see the same message returned

## Swagger UI (echo-server must be running beforehand)
* `$ tusk serve-doc`

* Open a browser at `http://localhost:8000`.
