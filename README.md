# Harbor
## Description

Harbor is a manager for non-Docker dependencies in applications container building.

Using Harbor should be simple and Harbor should help diminish customized scripts run before a `docker build`.

## Objectives
At the time, Harbor main objectives are:

+ Manage and download configuration files that don't belong in code repositories.
 + Manage per environment configuration such as: downloading different files for dev, test or production environments.
+ Execute shell commands before a `docker build` run (such as running some `ant` or `maven` build).

## Usage
Usage: `harbor [-e KEY=VALUE]`

Harbor looks up a file named harbor.yml in the same directory where run from, harbor.yml structure is:
```
imagetag: <tag to be used on 'docker build'>
 tags: [] <tags to create and push into registry>
 downloadpath: <local root path to download files into>
 s3:
   bucket: <base bucket to download files from>
   basepath: <inside the bucket the root path for files to be downloaded>
 files:
   - s3path: <path to file in S3 after [s3.bucket]/[s3.basepath]>
     localname: <local path + name of the file, will be downloaded into [downloadpath]/[localname]>
     permission: <[optional] file permissions, default 0644>
 - commands:
   <YAML array containing shell commands (currently /bin/bash) to be run before 'docker build'>
```

### Templating in harbor.yml
You can use ${<KEY>} as a placeholder in harbor.yml to be replaced by the value passed in a -e flag