#! /bin/bash

# shellcheck disable=SC2162
read -p "[dev] Enter commit >>> " message
if [[ $message != "" ]]; then
  echo "[dev] Success added a commit: ${message}"
else
  echo "[dev] No commit input !"
fi

git checkout dev
git add .
git commit -m "[update]: ${message}"
git pull origin dev --rebase
git push origin dev
