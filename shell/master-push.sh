#! /bin/bash

git checkout master
git merge dev
git pull origin master
#git pull origin master --rebase
git push origin master
git checkout dev
echo '------------------------'