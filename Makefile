
# make help
help:
	@bash push.sh

# make dev
dev:
#	@git checkout dev
#	@git add .
#	@git status
#	@-if [ "$(message)" != "" ]; then git commit -m "[update]: $(message)" & git pull origin dev & git push origin dev; else git reset & echo "[dev] You haven't entered any comments !"; fi
#	@git checkout dev
#	@echo '------------------------'
	@bash push.sh dev

# make master
master:
#	@git checkout master
#	@git merge dev
#	@git pull origin master
#	@git push origin master
#	@git checkout dev
#	@echo '------------------------'
	@bash push.sh master

# make both
both:
#	make dev
#	make master
	@bash push.sh dev
	@bash push.sh master

tag-d:
#	@git checkout master
#	@git tag -l | tail -n 5
#	@-if [[ $version != "" ]]; then git fetch & git tag -d ${version} & git push origin :refs/tags/${version} & exit 0 & echo "[tag] Success deleted a tag: ${version}"; else echo  'No input delete version !'; fi
#	@git checkout dev
#	@echo '------------------------'
	@bash push.sh tag -d

# make tag version=v0.9.2
tag:
#	@git checkout master
#	@-if [ -z ${version} ]; then echo '[tag] No input tag version !'; else git tag -a ${version} -m "[tag]: add a new tag-${version}" & git push origin ${version} & git fetch -p & git tag -l | tail -n 5 & echo "[tag] Success added a new tag: ${version}"; fi
#	@git checkout dev
#	@echo '------------------------'
	@bash push.sh tag
