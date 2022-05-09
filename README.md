![Go](https://github.com/jahid90/just-cli/workflows/Go/badge.svg)

# JustCLI - A command runner

## Description

`just` is a command line tool to execute arbitrary commands.

`just` uses a project specific config file to discover available commands and allows executing them via their defined aliases. `just` will look for a config file named `just.yaml` or `just.json` by default.

### Usage

```sh
$ just --help

NAME:
   just - A command runner

USAGE:
   just [global options] command [command options] [arguments...]

VERSION:
   1.0.0 - 34b01060ca5275766c80e16f3370f8f5e0358020

DESCRIPTION:
   Runs commands defined by aliases in a config file.
   Looks for a config file named just.json/just.yaml in the current directory.
   A different config file can be provided using the `--config-file` switch

   Usage examples:
     To list the available commands, run `just --list`
     To execute a command, run `just <alias>`

   If no sub-command is passed, `do` is inferred.

COMMANDS:
   do       Runs a command
   hello    Says hello
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config-file value, -c value  the config file to use
   --list, -l                     list the available commands (default: false)
   --short, -s                    list a short version of the available commands (default: false)
   --output value, -o value       Print config as json/yaml
   --convert                      Convert config files between different versions (default: false)
   --command value                Prints the command for a given alias
   --vars value                   Uses provided vars to interpolate into runs  (accepts multiple inputs)
   --skip-failures, -k            Keep going even if some steps fail (default: false)
   --help, -h                     show help (default: false)
   --version, -v                  print the version (default: false)
```

### Sample config file

```sh
$ cat just.yaml

version: 6
variables:
   environ: production
commands:
  build:
    description: Builds the app
    steps:
      - name: Invoke webpack
        env:
          - NODE_ENV={{ .environ }}
        run: webpack-cli
  test:
    description: Tests the app
    needs:
      - build
    steps:
      - name: Invoke the test target
        run: ./gradlew test
  multi:step:
    description: Run a multi-step alias
    steps:
      - name: list dir contents
        run: ls
      - name: fail step
        run: false
      - name: say bye
        run: echo bye
```

### The `do` sub-command
The `do` sub-command can be used to run the commands listed in a config file

(**Note**: As of version `1.0.0`, the `do` sub-command is no longer needed to be specified explicitly. Any arguments to `just` is forwarded to the `do` sub-command if it does not match any other sub-commands. So for e.g., `just do build` can be replaced with `just build`)

### List available commands:

```shell
$ just --list # or just do --list

Available commands are:

  ALIAS         COMMAND
  -----         -------
  build         npm run build
  docker:build  docker build -t image:tag
  docker:start  docker-compose up -d
  clean         rm -rf ./dist/
```

### To run a command, pass the alias to `just do`

```shell
$ just build # or just do build
npm run build
...
BUILD SUCCESSFUL
```

### A custom config file can be passed with the `--config-file` flag.

```shell
$ cat my-config-file
{
  "build": "mvn package",
  ...
}
$ just --config-file=my-config-file do build
mvn package
...
BUILD SUCCESSFUL
```

## Development

### Checkout the package locally from github
```
$ cd /workspace
$ git clone https://github.com/jahid90/just-cli.git just
```
### Run a local build
```
$ go build
$ ./just help
```
This will fetch the needed dependencies and create an executable in the local directory

### Install it locally
```
$ go install
$ just help
```
This will install the executable to `$GOBIN`. Adding `$GOBIN` to `$PATH` will allow executing the command from anywhere.

```
$ export PATH=$PATH:$GOBIN
```
