#!/bin/sh -l

WD=/app

if [ ! -d /github/workflow ]; then
   mkdir -p /github/workflow
fi

${WD}/gql scrape --url https://api.graph.cool/simple/v1/ciyz901en4j590185wkmexyex
cp ${WD}/sqlite3 /github/workspace/

echo ::set-output name=file::sqlite3
