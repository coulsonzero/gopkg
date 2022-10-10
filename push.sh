#! /bin/bash

if [ ${#} -ge 1 ]
then
    case ${1} in
        "dev")
            bash dev-push.sh
        ;;
        "master")
            bash master-push.sh
        ;;
        "both")
            bash dev-push.sh && bash master-push.sh
        ;;
        "tag")
            bash tag-release.sh
        ;;
        "tag-d")
            bash tag-delete.sh
        ;;
        *)
            echo "${1}No input"
        ;;
    esac
else
    echo "
    command如下命令:
    dev    : 推送dev
    master : 推送master
    tag    : 推送tag

    命令示例: sh push.sh dev
    "
fi