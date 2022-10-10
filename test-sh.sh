#! /bin/bash

if [ $# -eq 0 ]; then
  echo "no arg"
fi
if [ $1 == '-dev' ]; then
  echo 'dev'
elif [ $1 == '-del' ]; then
  echo 'del'
else
  echo 'no things'
fi