# Development environment and tools for full-stack gRPC service generation.

This project aims to set up an opinionated environment for protobuf based web development, mainly to generate and prototype full-stack monitoring and control panel systems for abstract devices. Final goal would be to generate disk images of docker swarm mode Raspberry Pi cluster nodes.

## Required packages for development
* `direnv`
* `docker-compose`

## Setup the development environment
* `$ git clone https://github.com/magnuskokk/grpc-tools.git`
* `$ cd grpc-tools`
* Install `direnv` from your package manager and set up the shell hook for the terminal emulator you're using (bash, zsh, etc...): https://github.com/direnv/direnv.
* `$ direnv allow .` to load local environment variables from .envrc.
* `$ ./setup.sh` to build and install the dev environment.

## Project layout
It is possible to run commands through the root `tusk.yml` file using docker or locally by manually using `app/tusk.yml` if you have a local go installation.

All proto services are defined in `./idl/{servicename}/{servicename}{version}`. The generated go package for each service is `app/idl/{servicename}/{servicename}{version}`. A Typescript client is also generated into `./frontend/generated` and swagger doc jsons into `./swagger`

Now that you have the environment loaded, you can run some commands.

## List all commands
* `$ tusk`

## Some useful commands
* `$ tusk protolint`
* `$ tusk protoc`
* `$ tusk app.generate.go`
* `$ tusk app.test.go`
* `$ tusk app.bench.go`

## Build and run the echo server
* `$ tusk serve.echo`
* Or start locally: `$ cd app; go run cmd/echo-server/main.go`
* Make sure it's responding: `$ curl 'http://localhost:8000/echo?Message=test'` You should see the same message returned

## Swagger UI (echo-server must be running beforehand to enable the try out feature)
### TODO needs development
* `$ tusk serve.echo.swagger`

* Open a browser at `http://localhost:8080` to see the docs.

## Reset the dev environment:
* `$ tusk env.reset`

### VSCode
direnv extension: https://marketplace.visualstudio.com/items?itemName=Rubymaniac.vscode-direnv. This allows installing all go tools in the .direnv directory.
