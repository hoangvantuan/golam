# Auto deploy lambda@edge for cloudfront

## Install

```
go get github.com/mdshun/golam
```

## Usage:

golam [global options] command [command options][arguments...]

## Commands:

help, h Shows a list of commands or help for one command

GLOBAL OPTIONS:

1. update lambda function and configure with cloudfront

`--all, -a`

```
args[0]: lambda function name
args[1]: source code path
args[2]: distribution id
args[3]: path pattern
args[4]: event type
args[5]: aws region
```

2. update lambda function

`--update, -u`

```
args[0]: lambda function name
args[1]: source code path
args[2]: aws region
```

3. update and publish new version

`--publish, -p`

```
args[0]: lambda function name
args[1]: source code path
args[2]: aws region
```

4. configure latest version of lambda for cloudfront

`--setup, -s`

```
args[0]: distribution id
args[1]: path pattern
args[2]: event type
args[3]: lambda function name
args[4]: setting version
args[5]: aws region
```

5. show help

`--help, -h`

6. print the version

`--version, -v`
