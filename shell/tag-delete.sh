#! /bin/bash

function main() {
    git fetch
    # shellcheck disable=SC2086
    git tag -d ${version}
    # shellcheck disable=SC2086
    git push origin :refs/tags/${version}
}

git checkout master

git tag -l | tail -n 5
# shellcheck disable=SC2162
read -t 20 -p "[tag] delete version >>> " version
if [[ $version != "" ]]; then
  main
  echo "[tag] Success deleted a tag: ${version}"
else
  echo  'No input delete version !'
fi

git checkout dev
echo '------------------------'
