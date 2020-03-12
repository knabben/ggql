#!/bin/bash -x

WD=/app

if [ ! -d /github/workflow ]; then
   mkdir -p /github/workflow
fi

RESULT=`${WD}/gql $@`
echo ::set-output name=output::${RESULT}
