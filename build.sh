#!/bin/bash

mkdir build
VERSION=$(git name-rev --tags --name-only $(git rev-parse HEAD))
GITHASH=$(git log -n1 --pretty='%h')
BINNAME="rcmd"
rm build/*.zip

rm -rf build/darwin
rm -rf build/linux

mkdir -p build/darwin/amd64
mkdir -p build/linux/amd64
mkdir -p build/linux/i386

echo VERSION: $VERSION
echo GITHASH: $GITHASH
cd build/darwin/amd64
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$VERSION -X main.githash=$GITHASH" ../../../
zip ../../${BINNAME}_darwin_amd64.zip $BINNAME
cd ../../linux/amd64
env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$VERSION -X main.githash=$GITHASH" ../../../
zip ../../${BINNAME}_linux_amd64.zip $BINNAME
cd ../i386
env GOOS=linux GOARCH=386 go build -ldflags "-X main.version=$VERSION -X main.githash=$GITHASH" ../../../
zip ../../${BINNAME}_linux_i386.zip $BINNAME
