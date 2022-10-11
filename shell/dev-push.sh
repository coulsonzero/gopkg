#! /bin/bash

set -e

git checkout dev
git add .
git status
# shellcheck disable=SC2162
read -t 30 -p "[dev] Enter commit >>> " message
if [ "$message" != "" ]; then
  git commit -m "[update]: $message"
  echo "[dev] Success added a commit: $message"
  git pull origin dev
  git push origin dev
  exit 0
else
  git reset
  echo "[dev] No commit input !"
  exit 1
fi
git checkout dev

echo '------------------------'
