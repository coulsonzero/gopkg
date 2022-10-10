#! /bin/bash

if [ ${#} -ge 1 ]
then
   case ${1} in
        "dev") bash ./shell/dev-push.sh ;;
        "master") bash ./shell/master-push.sh ;;
        "both") bash ./shell/dev-push.sh && bash ./shell/master-push.sh ;;
        "tag")
            if [[ ${2} == '-d' ]]; then
                bash ./shell/tag-delete.sh
            elif [[ -n ${2} ]]; then
                echo 'undefined the second argument '
            else
                bash ./shell/tag-release.sh
            fi
        ;;
        *) echo "${1} undefined the first argument" ;;
    esac
else
    echo "
    [Bash基础命令]:
    dev    : 推送dev
    master : 推送master
    both   : 同时推送dev与master
    tag    : 推送tag
    tag -d : 删除tag

    命令示例: sh push.sh dev
    "
fi


#   if [ "$1" = "dev" ]; then
#      bash ./shell/dev-push.sh
#   elif [ "$1" = "master" ]; then
#      bash ./shell/master-push.sh
#   elif [ "$1" = 'both' ]; then
#      bash ./shell/dev-push.sh && bash ./shell/master-push.sh
#   elif [ "$1" = 'tag' ]; then
#       # 判断第二个输入参数为'-d' ?
#       if [ "$2" = '-d' ]; then
#           bash ./shell/tag-delete.sh
#       # 判断第二个输入参数 ?
#       elif [ -z "$2" ]; then
#           bash ./shell/tag-release.sh
#       # 只有一个输入参数
#       else
#           echo 'undefined the second argument '
#       fi
#   else
#       echo "$1 No input"
#  fi