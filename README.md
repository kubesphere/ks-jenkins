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
Most of the plugins come from Jenkins community, but parts of them don't:

| Name | Git Repo |
|---|---|
| `kubesphere-token-auth` | https://github.com/kubesphere/kubesphere-token-auth-plugin |
| `kubesphere-extension` | https://github.com/jenkinsci/kubesphere-extension-plugin |
