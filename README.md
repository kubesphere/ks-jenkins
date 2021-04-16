# ks-jenkins

Jenkins distribution for [kubesphere](https://github.com/kubesphere/kubesphere)

`ks-jenkins` is an out-of-the-box solution which base on [custom-war-packager](https://github.com/jenkinsci/custom-war-packager).

# Get started

The docker images are below:

| | |
|---|---|
| official | `kubesphere/ks-jenkins:2.249.1` |
| experimental | `kubespheredev/ks-jenkins:2.249.1` |

## Build from source

[jcli](https://github.com/jenkins-zh/jenkins-cli) is a handy tool which can generate jenkins.war and docker image by one command line.

`jcli cwp --install-artifacts --config-path formula.yaml`

# Plugins

Please pay attention to these plugins, we still need to keep use a special version of them:

| Name | Version | Description |
|---|---|---|
| `pipeline-input-step` | `2.12-rc390.24ce2a334298` | Depends on [PR-33](https://github.com/jenkinsci/pipeline-input-step-plugin/pull/33) |
