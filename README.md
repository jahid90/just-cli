# just-cli - A command runner

## Description

`just` is a command line tool to execute arbitrary commands.

`just` uses a project specific config file to discover available commands and allows executing them via their defined aliases.
`just` will look for a config file named `Justfile` by default.

An example of a config file is as below:

```
$ cd /a/project/directory
$ cat Justfile
{
  "version": "1",
  commands: {
    "build": "npm run build",
    "docker:build": "docker build -t image:tag .",
    "docker:start": "docker-compose up -d",
    "clean": "rm -rf ./dist/"
  }
}
```
The commands can be run using the `do` sub-command

```
$ just do build
npm run build
...
BUILD SUCCESSFUL
```

A custom file can be passed with the `--config-file` flag.

```
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
$ git clone git@github.com:jahid90/just-cli.git just
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
