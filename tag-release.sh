#! /bin/bash

git checkout master
git tag

# shellcheck disable=SC2162
read -p "[tag] release version >>> " version
if [[ $version != "" ]]; then
  echo "[tag] Success added a new tag: ${version}"
else
  echo  'No commit input !'
fi

#version=v0.5.2-beta.1

git tag -a "${version}" -m "[tag]: add a new tag-${version}"
git push origin "${version}"
git tag
git checkout dev
