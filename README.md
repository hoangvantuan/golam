# Auto deploy lambda@edge for cloudfront

## Install

```
go get github.com/mdshun/golam
```

## Usage

```
NAME:
   Auto deploy lambda@edge function cmd - A new cli application

USAGE:
   golam [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     deploy   publish new lambda@edge version and connect to cloudfront distribution
     update   update lambda function, publish new version from source
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --distribution value, -d value      cloudfront dítribution id
   --even-type value, --et value       event type of cloudfront distribution
   --lambda-version value, --lv value  version of lambda function
   --name value, -n value              name of lambda function
   --path value, -p value              lambda source code path directory
   --path-pattern value, --pt value    path pattern of cloudfront distribution
   --publish-new-version, --pnv        event type of cloudfront distribution
   --region value, -r value            aws region
   --help, -h                          show help
   --version, -v                       print the version
```
