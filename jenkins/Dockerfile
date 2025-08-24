FROM docker:24-dind

ARG KUBECTL_VERSION=v1.28.0
ARG KUSTOMIZE_VERSION=v5.3.0


RUN apk add --no-cache curl unzip make build-base \
    \
    # kubectl 
    && curl -LO https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl \
    && chmod +x kubectl \
    && mv kubectl /usr/local/bin/ \
    \
    # kustomize 
    && curl -s "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F${KUSTOMIZE_VERSION}/kustomize_${KUSTOMIZE_VERSION}_linux_amd64.tar.gz" -L -o kustomize.tar.gz \
    && tar -xzf kustomize.tar.gz \
    && chmod +x kustomize \
    && mv kustomize /usr/local/bin/ \
    && rm -f kustomize.tar.gz

ENTRYPOINT ["dockerd-entrypoint.sh"]
CMD ["--host=tcp://0.0.0.0:2375", "--host=unix:///var/run/docker.sock"]
