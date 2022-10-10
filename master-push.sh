#! /bin/bash

git checkout master
git merge dev
#git pull origin master --rebase
git pull origin master
git push origin master
git checkout dev