build:
	jcli cwp --install-artifacts --config-path formula.yaml

run:
	jcli config gen -i=false > /home/gitpod/.jenkins-cli.yaml
	jcli center start --image kubesphere/ks-jenkins  --version 2.249.1 --setup-wizard=false
