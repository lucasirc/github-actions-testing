#!/bin/bash
echo "initing entrypoint"
set -e

ls -lash

if [ -z "$GITHUB_URL" ] || [ -z "$RUNNER_TOKEN" ] || [ -z "$RUNNER_NAME" ]; then
  echo "The required variables were not defined: GITHUB_URL, RUNNER_TOKEN e RUNNER_NAME"
  exit 1
fi

LABELS_ARG=""
if [ -n "$RUNNER_LABELS" ]; then
  LABELS_ARG="--labels $RUNNER_LABELS"
fi

whoami
nginx

gosu runner whoami

echo "executin github actions config.sh"

gosu runner ./config.sh \
  --url "$GITHUB_URL" \
  --token "$RUNNER_TOKEN" \
  --name "$RUNNER_NAME" \
  --work "_work" \
  --unattended \
  --replace \
  $LABELS_ARG

echo "executing run.sh "

exec gosu runner ./run.sh