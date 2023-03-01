# ks-jenkins
Jenkins distribution for [kubesphere](https://github.com/kubesphere/kubesphere)

`ks-jenkins` is an out-of-the-box solution which base on [custom-war-packager](https://github.com/jenkinsci/custom-war-packager).

# Get started
The docker images are below:

| | |
|---|---|
| official | `kubesphere/ks-jenkins:v3.4.0-2.319.3` |
| experimental | `kubespheredev/ks-jenkins:master` |

## Build from source

[jcli](https://github.com/jenkins-zh/jenkins-cli) is a handy tool which can generate jenkins.war and docker image by one command line.

> For some reasons, you need to config your Maven settings file. Please [see also](https://github.com/kubesphere/ks-jenkins/issues/16).

Build it via: `make build`

Run it via: `make run`

# Plugins
Please pay attention to these plugins, we still need to keep use a special version of them:

| Name | Version | Description |
|---|---|---|
| `pipeline-input-step` | `2.12-rc390.24ce2a334298` | Depends on [PR-33](https://github.com/jenkinsci/pipeline-input-step-plugin/pull/33) |
