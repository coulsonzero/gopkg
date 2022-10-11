#! /bin/bash

set -e

# shellcheck disable=SC2162
read -t 30 -p "[dev] Enter commit >>> " message
if [ "$message" != "" ]; then
  echo "[dev] Success added a commit: $message"
  git checkout dev
  git add .
  git commit -m "[update]: $message"
  git pull origin dev --rebase
  git push origin dev
else
  echo "[dev] No commit input !"
  exit 1
fi


echo '------------------------'
