# Dockerized Development Environment and Tools for full-stack gRPC service generation.

This project aims to set up an opinionated environment for protobuf based web development, mainly to generate and prototype remote monitoring and control panel systems for abstract devices.

## Ideas
This project is inspired by https://github.com/gogo/grpc-example and uses https://github.com/uber/prototool to lint and generate protos.

Imagine you have a custom home automation device (raspi with a temperature sensor and a radiator). It may have its own internal logic to turn on/off the heat based on indoor temperature but that's not important here. We need to remotely see if everything's working and intervene if needed.

The next idea would be to research how to generate Prometheus and Grafana provisioning scripts using a custom protoc plugin and custom annotations.

Example proto proto based on a real service from `idl/raspi/raspiv1`
```protobuf
message Temperature {
  sint32 reading = 1;
}

message Radiator {
  bool enabled = 1;
  uint32 level = 2;
}

message Status {
  Temperature temperature = 1;
  Radiator radiator = 2;
}
```

## Required packages for development
* `direnv`
* `docker`
* `docker-compose`

## Setup the development environment
* `$ git clone https://github.com/magnuskokk/grpc-tools.git`
* `$ cd grpc-tools`
* Install `direnv` from your package manager and set up the shell hook for the terminal emulator you're using (bash, zsh, etc...): https://github.com/direnv/direnv.
* `$ direnv allow .` to load local environment variables from .envrc.
* `$ ./install-tusk.sh` to install the task runner.
* `$ tusk` to list all tasks.

## Project layout
It is possible to run commands through the root `tusk.yml` file using docker or locally by manually using `app/tusk.yml` if you have a local go installation.

All proto services are defined in `./idl/{name}/{name}{version}`. The generated go package for each service is `app/idl/{name}/{name}{version}`.

A Typescript client is also generated into `./frontend/generated` and swagger doc jsons into `./swagger`.

Servers are defined in `docker-compose.yml` and `app/cmd/{name}-server`. The implementations are in `app/api/{name}`

Now that you have the environment loaded, you can run some commands. There are two example services `raspi` and `echo`.

## List all tasks
* `$ tusk`
```
Tasks:
   app.bench                Run all go benchmarks in ./app.
   app.test                 Run all go tests in ./app.
   docker.clean             Stop and clean everything but keep system cache for fast rebuilds.
   docker.clean.containers  Stop and remove all containers, images and any anonymous volumes attached to containers.
   docker.clean.images      Stop all containers and remove all images.
   docker.clean.volumes     Stop all containers and remove all volumes.
   docker.down              Stop all containers. All docker.* commands include only services defined in docker-compose.* files.
   docker.prune.system      WARNING! Prunes the whole system.
   env.build                Build containers for dev tools.
   env.clean                Reset an existing functional environment and install tusk.
   env.clean.sudo           Reset any environment in this directory.
   gen.app                  Run all //go:generate directives in ./app.
   gen.app.dockerfiles      Generate Dockerfiles for servers in ./app/cmd.
   gen.clean                Remove all generated files.
   gen.install.tools        Install tools and dependencies necessary for using protolint and protoc.
   gen.protoc               Generate gRPC server, client, gateway, typescript and swagger for all services.
   gen.protolint            Lint all protobuf definitions using prototool.
   stack.build              Build the whole stack.
   stack.down               Stop the stack.
   stack.up                 Start the stack.
 ```

## Build and run the stack
* `$ tusk stack.build`
* `$ tusk stack.up`
* Make sure it's responding: `$ curl http://localhost:8000/echo?message=test` 

## Swagger UI with "Try out"
* echo-server: `http://localhost:8080`
* raspi-server: `http://localhost:8081`

## Reset the dev environment:
* `$ tusk env.clean.sudo`
* `$ tusk env.build`

### VSCode
direnv extension: https://marketplace.visualstudio.com/items?itemName=Rubymaniac.vscode-direnv. This allows installing all go tools in the .direnv directory.
