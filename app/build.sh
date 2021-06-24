#!/usr/bin/env bash

# This script builds omnichannel for most common platforms.

DIRECTORY=builds

if [ ! -d "$DIRECTORY" ]; then
  mkdir "$DIRECTORY"
  echo "creating directory $DIRECTORY"
else
  find $DIRECTORY -type f -delete
  echo "cleaning directory $DIRECTORY/"
fi

if [ -d "resources" ]; then
  cp -rf resources/ ./$DIRECTORY/resources
   echo "coping resources to $DIRECTORY/resources"
fi

set -ex
export CGO_ENABLED=0

GOOS=linux   GOARCH=386   go build -o ./$DIRECTORY/omnichannel_linux_386
GOOS=linux   GOARCH=amd64 go build -o ./$DIRECTORY/omnichannel_linux_amd64
GOOS=linux   GOARCH=arm   go build -o ./$DIRECTORY/omnichannel_linux_arm7
GOOS=linux   GOARCH=arm64 go build -o ./$DIRECTORY/omnichannel_linux_arm64
GOOS=darwin  GOARCH=amd64 go build -o ./$DIRECTORY/omnichannel_mac_amd64
GOOS=windows GOARCH=386   go build -o ./$DIRECTORY/omnichannel_windows_386.exe
GOOS=windows GOARCH=amd64 go build -o ./$DIRECTORY/omnichannel_windows_amd64.exe
GOOS=freebsd GOARCH=386   go build -o ./$DIRECTORY/omnichannel_freebsd_386
GOOS=freebsd GOARCH=amd64 go build -o ./$DIRECTORY/omnichannel_freebsd_amd64
GOOS=freebsd GOARCH=arm   go build -o ./$DIRECTORY/omnichannel_freebsd_arm7
GOOS=openbsd GOARCH=386   go build -o ./$DIRECTORY/omnichannel_openbsd_386
GOOS=openbsd GOARCH=amd64 go build -o ./$DIRECTORY/omnichannel_openbsd_amd64

echo "Build completed"