
dev:
	git checkout dev
	git add .
	git status
	read -t 40 -p "[dev] Enter commit >>> " message
	if [ "$message" != "" ]; then git commit -m "[update]: $message" & git pull origin dev & git push origin dev; else git reset & echo "[dev] You haven't entered any comments !"; fi
	git checkout dev
	echo '------------------------'