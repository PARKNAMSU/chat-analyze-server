#!/bin/bash

export TAG="$1"
export COMMIT_EXPLAIN=""

if [ "$TAG" == "patch" ] || [ "$TAG" == "fix" ] || [ "$TAG" == "add" ] || [ "$TAG" == "setting" ]; then
    COMMIT_EXPLAIN="$2"
else
    TAG="default"
    COMMIT_EXPLAIN="$1"
fi

if [ "$COMMIT_EXPLAIN" == "" ]; then
    echo "not exist commit explain"
    exit
fi

while [ $# -ne 0 ]; do
    if [ "$TAG" == "default" ]; then
        COMMIT_EXPLAIN="$COMMIT_EXPLAIN $2"
    else
        COMMIT_EXPLAIN="$COMMIT_EXPLAIN $3"
    fi
    shift
done


git add .
git commit -m "[${TAG}][${COMMIT_EXPLAIN}]"
