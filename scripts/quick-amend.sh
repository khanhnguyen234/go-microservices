#!/usr/bin/env bash

branch_name=$(git symbolic-ref -q HEAD)
branch_name=${branch_name##refs/heads/}
branch_name=${branch_name:-HEAD}

echo Current Branch: '['"$branch_name"']'
read -r -p "git add . && git commit --amend && git push origin HEAD -f"
git add .
git commit --amend
git push origin HEAD -f