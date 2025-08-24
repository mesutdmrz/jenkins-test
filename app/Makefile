APP_NAME := pod-watcher
REGISTRY := mesut4730
IMAGE := $(REGISTRY)/$(APP_NAME)
TAG := $(shell git rev-parse --short HEAD)-$(shell date +%H%M)

OVERLAY_DIR := /home/jenkins/agent/workspace/docker-dind-pipeline

.PHONY: all docker push update-kustomize deploy clean

all: docker push

## Docker image build
docker:
	docker build -t $(IMAGE):$(TAG) .

## Docker image push
push:
	docker push $(IMAGE):$(TAG)

## overlay seçilerek kustomize image güncelle
update-kustomize:
ifndef ENV
	$(error ENV not set. Use 'make update-kustomize ENV=test' or ENV=prod)
endif
	cd $(OVERLAY_DIR)/$(ENV) && kustomize edit set image $(IMAGE)=$(IMAGE):$(TAG)

## Deploy (örn: test veya prod)
deploy: all update-kustomize
ifndef ENV
	$(error ENV not set. Use 'make deploy ENV=test' or ENV=prod)
endif
	kubectl apply -k $(OVERLAY_DIR)/$(ENV)

## Temizlik (Docker local image’ı silmek istersen)
clean:
	docker rmi $(IMAGE):$(TAG) || true
