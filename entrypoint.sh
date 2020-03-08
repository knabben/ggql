#!/bin/sh -l

WORKDIR=${PWD}

if [ ! -d /github/workflow ]; then
   mkdir -p /github/workflow
fi

${WORKDIR}/gql scrape --url https://api.graph.cool/simple/v1/ciyz901en4j590185wkmexyex
cp ${WORKDIR}/sqlite3 /github/workflow/
echo ::set-output name=file::sqlite3
