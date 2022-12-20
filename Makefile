
dev:
	@git checkout dev
	@git add .
	@git status
	@read -t 40 -p "[dev] Enter commit >>> " message
	@if [ "$message" != "" ]; then git commit -m "[update]: $message" & git pull origin dev & git push origin dev; else git reset & echo "[dev] You haven't entered any comments !"; fi
	@git checkout dev
	@echo '------------------------'

master:
	@git checkout master
	@git merge dev
	@git pull origin master
	@git push origin master
	@git checkout dev
	@echo '------------------------'

both:
	make dev
	make master

tag-d:
	@git checkout master
	@git tag -l | tail -n 5
	@read -t 20 -p "[tag] delete version >>> " version
	@if [[ $version != "" ]]; then git fetch & git tag -d ${version} & git push origin :refs/tags/${version} & exit 0 & echo "[tag] Success deleted a tag: ${version}"; else echo  'No input delete version !'; fi
	@git checkout dev
	@echo '------------------------'

