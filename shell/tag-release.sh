#! /bin/bash

main() {
  # shellcheck disable=SC2086
  git tag -a ${version} -m "[tag]: add a new tag-${version}"
  # shellcheck disable=SC2086
  git push origin ${version}
  git fetch -p
  git tag -l | tail -n 5
}


git checkout master
git tag -l | tail -n 4
# shellcheck disable=SC2162
read -t 20 -p "[tag] release version >>> " version
if [ -z "$version" ]; then
  echo  '[tag] No input tag version !'
else
    main
    echo "[tag] Success added a new tag: ${version}"
fi

git checkout dev
echo '------------------------'
