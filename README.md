![Go](https://github.com/jahid90/just-cli/workflows/Go/badge.svg)

# JustCLI - A command runner

## Description

`just` is a command line tool to execute arbitrary commands.

`just` uses a project specific config file to discover available commands and allows executing them via their defined aliases. `just` will look for a config file named `just.json` by default.

A `v1` config file example is shown below:

```shell
$ cd /a/project/directory
$ cat just.json
{
  "version": "1",
  "commands": {
    "build": "npm run build",
    "docker:build": "docker build -t image:tag .",
    "docker:start": "docker-compose up -d",
    "clean": "rm -rf ./dist/"
  }
}
```

`v2` and `v3` are currently work-in-progress. They parse the config file w.r.t. some grammar rules before executing the commands.

`v4` executes the command using the underlying OS shell and supports environment variables and sub-command expansions.

Any version above `v4` is currently unsupported.

A `v4` config file example is presented below

```json
{
    "version": "4",
    "commands": {
        "dev": "NODE_ENV=development,DEBUG=app:* yarn start",
        "build": "NODE_ENV=production yarn build",
        "test": "PROFILE=dev,PORT=9000,SECRET=password,USER=$USER ./mvnw test",
        "docker:build": "docker build -t docker-image:local .",
        "docker:run": "docker-compose up -d",
        "k8s:generate": "VERSION=$(METADATA_FILE_NAME=.app-metadata.json get-version) envsubst < k8s/template.yaml > k8s/deployment.yaml",
        "k8s:deploy": "kubectl apply -f k8s/deployment.yaml",
        "done": "echo done",
        "ls": "ls -lh",
        "k8s:redeploy": "docker build -t $(app-name):$(get-version) . && kubectl apply -f k8s/deployment.yaml"
    }
}

```

## The `do` sub-command
The `do` sub-command can be used to run the commands listed in a config file

(**Note**: As of version `1.0.0`, the `do` sub-command is no longer needed to be specified explicitly. Any arguments to `just` is forwarded to the `do` sub-command if it does not match any other sub-commands. So for e.g., `just do build` can be replaced with `just build`)

### Show help
```shell
$ just --help # or just help
NAME:
   just - A command runner

USAGE:
   just [global options] command [command options] [arguments...]

...

GLOBAL OPTIONS:
   --config-file value, -c value  the config file to use (default: "just.json")
   --list, -l                     list the available commands (default: false)
   --help, -h                     show help (default: false)
   --version, -v                  print the version (default: false)

$ just do --help
NAME:
   just do - Runs a command

USAGE:
   just do [command options] [arguments...]

OPTIONS:
   --list, -l  list the available commands (default: false)
   --help, -h  show help (default: false)

```

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
