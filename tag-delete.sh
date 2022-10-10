#! /bin/bash

git checkout master

git tag
# shellcheck disable=SC2162
read -p "[tag] delete version >>> " version
if [[ $version != "" ]]; then
  # shellcheck disable=SC2086
  git tag -d ${version}
  # shellcheck disable=SC2086
  git push origin :refs/tags/${version}
  echo "[tag] Success deleted a tag: ${version}"
else
  echo  'No input delete version !'
fi

git checkout dev
