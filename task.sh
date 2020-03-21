#!/bin/sh

set -e

date +"%D %T" >> output

git add .

git commit -m "docs(output) new record"

git push origin master