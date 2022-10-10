#! /bin/bash

# 判断输入参数数量 > 1 ？
if [ ${#} -ge 1 ]; then
   if [ "$1" = "dev" ]; then
      bash ./shell/dev-push.sh
   elif [ "$1" = "master" ]; then
      bash ./shell/master-push.sh
   elif [ "$1" = 'both' ]; then
      bash ./shell/dev-push.sh && bash ./shell/master-push.sh
   elif [ "$1" = 'tag' ]; then
       # 判断第二个输入参数为'-d' ?
       if [ "$2" = '-d' ]; then
           bash ./shell/tag-delete.sh
       # 判断第二个输入参数 ?
       elif [ -z "$2" ]; then
           bash ./shell/tag-release.sh
       # 只有一个输入参数
       else
           echo 'undefined the second argument '
       fi
   else
       echo "$1 No input"
  fi
else
    echo "
    [Bash基础命令]:
    1) dev    : 推送dev
    2) master : 推送master
    3) both   : 同时推送dev与master
    4a) tag    : 推送tag
    4b) tag -d : 删除tag

    命令示例: sh push.sh dev
    "
fi
