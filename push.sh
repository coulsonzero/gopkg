#! /bin/bash

export CUR="shell"

# shellcheck disable=SC2120
function main() {
   if [ "$2" = "-d" ]; then         # 判断第二个输入参数为'-d' ?
       bash ./shell/tag-delete.sh
   elif [ -z "$2" ]; then           # 判断第二个输入参数 ?
       bash ./shell/tag-release.sh
   else                             # 只有一个输入参数
       echo 'undefined the second argument '
   fi
   exit 0
}

# shellcheck disable=SC2164
cd $CUR
# 判断输入参数数量 > 1 ？
if [ ${#} -ge 1 ]; then
  case "$1" in
    "dev"   ) bash dev-push.sh                        ;;
    "master") bash master-push.sh                     ;;
    "both"  ) bash dev-push.sh && bash master-push.sh ;;
    "tag"   ) main                                    ;;
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

    命令示例: sh push.sh dev
    "
fi


