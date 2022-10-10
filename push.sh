#! /bin/bash

if [ ${#} -ge 1 ]; then
   if [[ ${1} == 'dev' ]]; then
      bash ./shell/dev-push.sh
   elif [[ ${1} == 'master' ]]; then
      bash ./shell/master-push.sh
   elif [[ ${1} == 'both' ]]; then
      bash ./shell/dev-push.sh && bash ./shell/master-push.sh
   elif [[ ${1} == 'tag' ]]; then
       if [[ ${2} == '-d' ]]; then
           bash ./shell/tag-delete.sh
       elif [[ -n ${2} ]]; then
           echo 'undefined the second argument '
       else
           bash ./shell/tag-release.sh
       fi
   else
       echo "${1}No input"
  fi
else
    echo "
    command如下命令:
    dev    : 推送dev
    master : 推送master
    tag    : 推送tag
    tag-d  : 删除远程tag

    命令示例: sh push.sh dev
    "
fi
