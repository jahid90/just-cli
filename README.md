# just-cli - A command runner

## Description

`just` is a command line tool to execute arbitrary commands.

`just` uses a project specific config file to discover available commands and allows executing them via their defined aliases.
`just` will look for a config file named `just.json` by default.

A `v1` config file example is as below:

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

The `v2` config file allows specifying environment variables to be passed to the commands

```shell
$ cat just.json
{
  "version": "2",
  "commands": [
    {
      "alias": "dev",
      "action": "yarn start dev".
      "env": {
        "NODE_ENV": "development"
      }
    },
    {
      "alias": "build",
      "action": "./mvnw package"
    }
  ]
}
```

### The `do` sub-command
The `do` sub-command can be used to run the commands listed in a config file

#### List available commands:

```shell
$ just do --list
Available commands are:
  build         npm run build
  docker:build  docker build -t image:tag
  docker:start  docker-compose up -d
  clean         rm -rf ./dist/
```

#### To run a command, pass the alias to `just do`

```shell
$ just do build
npm run build
...
BUILD SUCCESSFUL
```

#### A custom file can be passed with the `--config-file` flag.

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
