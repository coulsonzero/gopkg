#! /bin/bash

main() {
  git checkout dev
  git add .
  git commit -m "[update]: ${message}"
  #git pull origin dev --rebase
  git push origin dev
}

# shellcheck disable=SC2162
read -t 30 -p "[dev] Enter commit >>> " message
if [[ $message != "" ]]; then
  echo "[dev] Success added a commit: ${message}"
  main
else
  echo "[dev] No commit input !"
fi

echo '------------------------'
