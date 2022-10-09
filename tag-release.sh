version=v0.5.2-beta.1

git checkout master
git tag
git tag -a ${version} -m "[tag]: modify module version-${version}"
git push origin ${version}
git tag
git checkout dev