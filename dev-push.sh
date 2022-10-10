#!/usr/bin/env bash

message=$*

git checkout dev
git add .
git commit -m "update: ${message}"
git push origin dev