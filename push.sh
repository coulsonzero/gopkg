#! /bin/bash

set -e

export CUR="shell"

# shellcheck disable=SC2164
cd $CUR
# 判断输入参数数量 > 1 ？
if [ ${#} -ge 1 ]; then
  case "$1" in
    "dev"   ) bash dev-push.sh                        ;;
    "master") bash master-push.sh                     ;;
    "both"  ) bash dev-push.sh && bash master-push.sh ;;
    "tag"   ) case $2 in
              "-d") bash tag-delete.sh                ;;
              ""  ) bash tag-release.sh               ;;
              *   ) echo 'undefined second argument!' ;;
              esac                                    ;;
    *) echo   "not input argument"                    ;;
  esac
else
    echo "
    [Bash基础命令]:
    1) dev    : 推送dev
    2) master : 推送master
    3) both   : 同时推送dev与master
    4) tag    : 推送tag
    5) tag -d : 删除tag

    命令示例: bash push.sh dev
    "
fi


