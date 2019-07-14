#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

DIR="$( cd -P "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

GIT_HEAD="$(git rev-parse --short=7 HEAD)"
GIT_DATE=$(git log HEAD -n1 --pretty='format:%cd' --date=format:'%Y%m%d-%H%M')

export SHORTEN_URL_TAG="shorten_url:${GIT_HEAD}-${GIT_DATE}"
export DISPATCHER_TAG="dispatcher:${GIT_HEAD}-${GIT_DATE}"

# Stop docker images
docker-compose -f ${DIR}/docker-compose-shorten-url.yml down
