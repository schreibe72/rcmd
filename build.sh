#!/bin/bash

mkdir build
VERSION=$(git name-rev --tags --name-only $(git rev-parse HEAD))
GITHASH=$(git log -n1 --pretty='%h')
rm build/darwin_amd64.zip
rm build/linux_amd64.zip
rm build/linux_i386.zip

rm -rf build/darwin
rm -rf build/linux

mkdir -p build/darwin/amd64
mkdir -p build/linux/amd64
mkdir -p build/linux/i386

echo VERSION: $VERSION
echo GITHASH: $GITHASH
cd build/darwin/amd64
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$VERSION -X main.githash=$GITHASH" ../../../
zip ../../darwin_amd64.zip rcmd
cd ../../linux/amd64
env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$VERSION -X main.githash=$GITHASH" ../../../
zip ../../linux_amd64.zip rcmd
cd ../i386
env GOOS=linux GOARCH=386 go build -ldflags "-X main.version=$VERSION -X main.githash=$GITHASH" ../../../
zip ../../linux_i386.zip rcmd
