FROM ghcr.io/linuxsuren/ks-jenkins-build-cache:sha-988638c as cache

# Cache the maven local repository
RUN git clone https://github.com/kubesphere/ks-jenkins && \
    cd ks-jenkins && \
    jcli cwp --install-artifacts --config-path formula.yaml || true && \
    cd .. && rm -rf ks-jenkins

FROM gitpod/workspace-full

USER root
COPY --from=cache /usr/local/bin/hd /usr/local/bin/hd
COPY --from=cache /root/.m2/repository /workspace/m2-repository/

USER gitpod
# More information: https://www.gitpod.io/docs/config-docker/
RUN sudo rm -rf /usr/bin/hd && \
    hd install cli/cli && \
    hd install jenkins-zh/jenkins-cli/jcli && \
    sudo chown gitpod:gitpod -R /home/gitpod/.m2
