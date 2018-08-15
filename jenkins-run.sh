#!/bin/bash

set -e

title() { echo -e "\e[1;34m==> $1\e[0m"; }

docker_login() {
  title "gcloud docker login"
  ACCESS_TOKEN=$(curl -s -H 'Metadata-Flavor: Google' http://metadata.google.internal./computeMetadata/v1/instance/service-accounts/default/token | cut -d'"' -f 4)
  docker login -e not@val.id -u _token -p ${ACCESS_TOKEN} http://gcr.io
  gcloud docker -a
}

docker_login

IMG=gcr.io/px_docker_repo/pxlogstash_exporter
BUILDTAG=$(git describe --tags --abbrev=1)
echo ${BUILDTAG} | tee buildtag
title "docker build ${BUILDTAG}"
docker build -t ${IMG} .

BUILD_ENV=""
if [[ ${BRANCH} == 'master' ]]; then
  BUILD_ENV=prod
elif [[ ${BRANCH} == 'dev' ]]; then
  BUILD_ENV=stg
fi

if [ -n "${BUILD_ENV}" ]; then
  title "docker tag ${BUILDTAG} and ${BUILD_ENV}-latest"
  docker tag ${IMG} ${IMG}:${BUILDTAG}
  docker tag ${IMG} ${IMG}:${BUILD_ENV}-latest

  title "docker push"
  docker push ${IMG}:${BUILDTAG}
  docker push ${IMG}:${BUILD_ENV}-latest
fi
