# Harbor
## Description

Harbor is a manager for non-Docker dependencies in applications container building.

Using Harbor should be simple and Harbor should help diminish customized scripts run before a `docker build`.

## Objectives
At the time, Harbor main objectives are:

+ Manage and download configuration files that don't belong in code repositories.
 + Manage per environment configuration such as: downloading different files for dev, test or production environments.
+ Execute shell commands before a `docker build` run (such as running some `ant` or `maven` build).