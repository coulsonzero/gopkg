#! /bin/bash

read -p "[dev] Enter commit >>> " message
if [[ $message != "" ]]; then
  echo "add a commit: ${message}"
else
  echo  'No commit input !'
fi

git checkout dev
git add .
git commit -m "update: ${message}"
git push origin dev