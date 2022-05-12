#!/bin/sh

dir="$PWD"
packages="$(grep -Ril "package main" . | grep -E ".*\.go" | xargs dirname | grep -P "$FILTER" | sort | uniq)"

fail="false"
echo "" > coverage.txt

for package in $packages; do
    echo "$package"
    sh -c "cd $package; $@"
    ret="$?"
    if [ "$ret" != "0" ]; then
        fail="true"
    fi

    if [ -f $package/profile.out ]; then
        cat $package/profile.out >> coverage.txt
        rm $package/profile.out
    fi
done

if [ "$fail" = "true" ]; then
    echo "Some tests failed"
    exit 1
fi