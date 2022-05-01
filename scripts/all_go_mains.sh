#!/bin/sh

dir="$PWD"
packages="$(grep -Ril "package main" . | grep -E ".*\.go" | xargs dirname | grep -P "$FILTER" | sort | uniq)"

fail="false"

for package in $packages; do
    echo "$package"
    sh -c "cd $package; $@"
    ret="$?"
    if [ "$ret" != "0" ]; then
        fail="true"
    fi
done

if [ "$fail" = "true" ]; then
    echo "Some tests failed"
    exit 1
fi