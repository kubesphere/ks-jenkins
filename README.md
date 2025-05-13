# ks-jenkins
Jenkins distribution for [kubesphere](https://github.com/kubesphere/kubesphere), but you may use it for other purposes. It's just a standard jenkins plus some plugins.

# Get started
The docker images are below:

| Type         | Image                                                                      |
|--------------|----------------------------------------------------------------------------|
| official     | [latest-release](https://github.com/kubesphere/ks-jenkins/releases/latest) |
| experimental | `kubesphere/ks-jenkins:master`                                             |

## Build your own jenkins image
1. Clone this repo
2. Modify the Dockerfile and plugins.txt as you need.
3. Build the image
```bash
docker build -t <your-image-name> .
```