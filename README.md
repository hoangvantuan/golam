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

## Samples

### Update lambda function publish new version
```
golam update -n name-lambda-fuction -r [region] -p path-to-source-code --pnv
```
### Update lambda function no publish new version
```
golam update -n name-lambda-fuction -r [region] -p path-to-source-code
```

### Update Cloudfront with specify lambda@edge version
```
golam deploy -n lambda-function-name -r [region]  -d distribution-id --pt path-pattern --et event-type -lv lambda-function-version
```

### Pushish new lambda@edge version and deploy to cloudfront
```
golam deploy -n lambda-function-name -r [region]  -d distribution-id --pt path-pattern --et event-type -p lambda-function-path
```

