
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

