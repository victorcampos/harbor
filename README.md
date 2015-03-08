# Harbor
## Description

Harbor is a wrapper for running commands and downloading file dependencies (currently only from AWS S3) before Docker image building.

Using Harbor should be simple and Harbor should help to stop usage of customized scripts run before a `docker build`.

## Objectives
At the time, Harbor main objectives are:

+ Manage and download configuration files that don't belong in code repositories.
 + Manage per environment configuration such as: downloading different files for dev, test or production environments.
+ Execute shell commands before a `docker build` run (such as running some `ant` or `maven` build).
+ Execute `docker build`, `docker tag` and `docker push` to repository

## Usage
Harbor looks up a file named harbor.yml in the same directory where run from, harbor.yml structure is:
```
 imagetag: <tag to be used on 'docker build'>
 tags:
   - <YAML array of custom tags to create and push into registry>
 downloadpath: <local root path to download files into>
 s3:
   bucket: <base bucket to download files from>
   basepath: <inside the bucket the root path for files to be downloaded>
 files:
   - s3path: <path to file in S3 after [s3.bucket]/[s3.basepath]>
     filename: <local path + name of the file, will be downloaded into [downloadpath]/[localname]>
     permission: <[optional] file permissions, default 0644>
 commands:
   - <YAML array containing shell commands (currently /bin/bash) to be run before 'docker build'>

 You can use ${<KEY>} as a placeholder in harbor.yml to be replaced by the value passed in a -e flag

Usage:
  harbor -h | --help
  harbor --version
  harbor --list-variables
  harbor [-e KEY=VALUE]... [options]
  harbor [options]

Options:
  -h, --help                    Show this screen.
  -v, --version                 Show version.
  --list-variables              Parses harbor.yml and prints out every ${KEY} found.
  -e KEY=VALUE                  Replaces every ${KEY} in harbor.yml with VALUE
  --debug                       Dry-run and print command executions.
  --no-download                 Prevents downloading files from S3.
  --no-commands                 Prevents commands in harbor.yml from being run.
  --no-docker                   Do not run 'docker build', 'docker tag' and 'docker push' commands.
  --no-docker-push              Do not run 'docker push' after building and tagging images.
  --docker-opts=<DOCKER_OPTS>   Will be appended to the base docker command (ex.: 'docker <docker-opts> command').
  --no-latest-tag               Do not build image tagging as 'latest',
                                will use the first tag in 'tags' list from harbor.yml or
                                generate a timestamped tag if no 'tags' list exists.
```

### Templating in harbor.yml
You can use ${<KEY>} as a placeholder in harbor.yml to be replaced by the value passed in a -e flag