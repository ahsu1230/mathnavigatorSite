#!/bin/bash

OUTPUT="$(test -z `gofmt -l ./`)"
STATUS=$?
if [ $STATUS != "0" ]; then
    echo "Go files are not correctly formatted. Please run gofmt."
else
    echo "Go files correctly formatted!"
fi
exit $STATUS
