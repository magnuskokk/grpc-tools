# Development environment and tools for full-stack gRPC service generation.

This project aims to set up an opinionated environment for protobuf based web development, mainly to generate and prototype remote monitoring and control panel systems for abstract devices.

## Ideas

Imagine you have a custom home automation device (raspi with a temperature sensor and a radiator). It may have its own internal logic to turn on/off the heat based on indoor temperature but that's not important here. We need to remotely see if everything's working and intervene if needed.

Example proto from `idl/raspi/raspiv1/types.proto`:
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
* `docker-compose`

## Setup the development environment
* `$ git clone https://github.com/magnuskokk/grpc-tools.git`
* `$ cd grpc-tools`
* Install `direnv` from your package manager and set up the shell hook for the terminal emulator you're using (bash, zsh, etc...): https://github.com/direnv/direnv.
* `$ direnv allow .` to load local environment variables from .envrc.
* `$ ./setup.sh` to set up the local dev environment.

### Optionally install tools for code generation.
* `$ tusk env.install.generate`

## Project layout
It is possible to run commands through the root `tusk.yml` file using docker or locally by manually using `app/tusk.yml` if you have a local go installation.

All proto services are defined in `./idl/{servicename}/{servicename}{version}`. The generated go package for each service is `app/idl/{servicename}/{servicename}{version}`. A Typescript client is also generated into `./frontend/generated` and swagger doc jsons into `./swagger`

Now that you have the environment loaded, you can run some commands. There are two example services `raspi` and `echo`.

## List all tasks
* `$ tusk`
```
Tasks:
   app.bench               Run all go benchmarks.
   app.test                Run all go tests.
   docker.cleanall         Clean all all docker stuff for services.
   docker.cleancontainers  Remove all local containers, images and any anonymous volumes attached to containers.
   docker.cleanimages      Stop all containers and remove all images used by any defined services.
   docker.cleanvolumes     Stop and clean all local volumes.
   docker.down             Stop all containers.
   env.build               Build the docker containers for dev tools.
   env.clean               Clean the environment including docker containers but keep images and named volumes.
   env.install.tools       Install tools and dependencies for generating code.
   env.reset               Reset and rebuild the dev environment.
   gen.clean               Remove all generated files.
   gen.go                  Run all //go:generate directives.
   gen.protoc              Generate gRPC server, client, gateway, typescript and swagger for all services.
   gen.protolint           Lint protobuf definitions using prototool.
   monitor.grafana         Start Grafana.
   stack.build             Build stack.
   stack.down              Stop the stack.
   stack.up                Start the stack.
 ```

## Build and run the stack
* `$ tusk stack.build`
* `$ tusk stack.up`
* Make sure it's responding: `$ curl http://localhost:8000/echo?message=test` 

## Swagger UI
* `$ tusk serve.echo.swagger`
* Open a browser at `http://localhost:8080` to see the swagger docs.

## Reset the dev environment:
* `$ tusk env.reset`

### VSCode
direnv extension: https://marketplace.visualstudio.com/items?itemName=Rubymaniac.vscode-direnv. This allows installing all go tools in the .direnv directory.
