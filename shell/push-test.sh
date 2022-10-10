#! /bin/bash

if [ ${#} -ge 1 ]
then
   case ${1} in
        "dev")
            bash ./shell/dev-push.sh
        ;;
        "master")
            bash ./shell/master-push.sh
        ;;
        "both")
            bash ./shell/dev-push.sh && bash ./shell/master-push.sh
        ;;
        "tag")
            if [[ ${2} == '-d' ]]; then
                bash ./shell/tag-delete.sh
            elif [[ -n ${2} ]]; then
                echo 'undefined the second argument '
            else
                bash ./shell/tag-release.sh
            fi
        ;;
        *)
            echo "${1} undefined the first argument"
        ;;
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
