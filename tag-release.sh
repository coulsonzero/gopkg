#! /bin/bash

git checkout master
git tag

# shellcheck disable=SC2162
read -p "[tag] release version >>> " version
if [[ $version != "" ]]; then
  main
  echo "[tag] Success added a new tag: ${version}"
else
  echo  'No input tag version !'
fi

#version=v0.5.2-beta.1

main() {
  # shellcheck disable=SC2086
  git tag -a ${version} -m "[tag]: add a new tag-${version}"
  # shellcheck disable=SC2086
  git push origin ${version}
  git tag
  git checkout dev
}

