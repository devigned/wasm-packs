#!/usr/bin/env bash
set -e

DEFAULT_PREFIX=wasm/demo-base
DEFAULT_PLATFORM="linux/arm64,linux/amd64"

REPO_PREFIX=${DEFAULT_PREFIX}
PLATFORM=${DEFAULT_PLATFORM}

usage() {
  echo "Usage: "
  echo "  $0 [-f <prefix>] [-p <platform>] <dir>"
  echo "   <dir>  directory to build"
  exit 1;
}

while getopts "v:p:" o; do
  case "${o}" in
    f)
      REPO_PREFIX=${OPTARG}
      ;;
    p)
      PLATFORM=${OPTARG}
      ;;
    \?)
      echo "Invalid option: -$OPTARG" 1>&2
      usage
      ;;
    :)
      usage
      ;;
  esac
done

BASE_DIR=${@:$OPTIND:1}

if [[ -z ${REPO_PREFIX} ]]; then
  echo "Prefix cannot be empty"
  echo
  usage
  exit 1
fi

if [[ -z ${BASE_DIR} ]]; then
  echo "Must specify directory"
  echo
  usage
  exit 1
fi

cd $(dirname $0)

IMAGE_DIR=$(realpath "${BASE_DIR}")
TAG=$(basename "${IMAGE_DIR}")
BASE_IMAGE=heroku/heroku:24-build
RUN_IMAGE=${REPO_PREFIX}-run:${TAG}
BUILD_IMAGE=${REPO_PREFIX}-build:${TAG}

echo "BUILDING ${BUILD_IMAGE}..."
docker buildx build --load \
	--platform "${PLATFORM}" \
	--build-arg "BASE_IMAGE=${BASE_IMAGE}" \
	-t "${BUILD_IMAGE}" \
	--output type=image \
	"${IMAGE_DIR}/build"

echo "BUILDING ${RUN_IMAGE}..."
docker buildx build --load \
	--platform "${PLATFORM}" \
	--build-arg "BASE_IMAGE=${BUILD_IMAGE}" \
	-t "${RUN_IMAGE}" \
	--output type=image \
	"${IMAGE_DIR}/run"

echo
echo "BASE IMAGES BUILT!"
echo
echo "Images:"
for IMAGE in "${BASE_IMAGE}" "${BUILD_IMAGE}" "${RUN_IMAGE}"; do
  echo "    ${IMAGE}"
done
