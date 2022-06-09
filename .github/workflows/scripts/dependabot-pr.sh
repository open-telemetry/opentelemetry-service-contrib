#!/bin/bash -ex

git config user.name $GITHUB_ACTOR
git config user.email $GITHUB_ACTOR@users.noreply.github.com

PR_NAME=dependabot-prs/`date +'%Y-%m-%dT%H%M%S'`
git checkout -b $PR_NAME

IFS=$'\n'
requests=$(gh pr list --search "author:app/dependabot" --json number,title --template '{{range .}}{{tablerow .title}}{{end}}')
message=""

for line in $requests; do
    if [[ $line != Bump* ]]; then 
        continue
    fi

    module=$(echo $line | cut -f 2 -d " ")
    if [[ $module == go.opentelemetry.io/collecto* ]]; then
        continue
    fi
    version=$(echo $line | cut -f 6 -d " ")
    make for-all CMD="$GITHUB_WORKSPACE/internal/buildscripts/update-dep" MODULE=$module VERSION=v$version
    message+=$line
    message+=$'\n'
done

make gotidy
make otelcontribcol

git add go.sum go.mod
git add "**/go.sum" "**/go.mod"
git commit -m "dependabot updates `date`
$message"
git push origin $PR_NAME

gh pr create --title "dependabot updates `date`" --body "$message" -l "Skip Changelog"
