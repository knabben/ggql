#!/bin/bash -x

WD=/app
FILE_NAME=sqlite3

if [ ! -d /github/workflow ]; then
   mkdir -p /github/workflow
fi

${WD}/gql $@
cp ${WD}/${FILE_NAME} /github/workspace/

echo ::set-output name=file::${FILE_NAME}
