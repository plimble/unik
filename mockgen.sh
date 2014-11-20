#! /bin/bash -e

PKG=github.com/plimble
NAME=unik
INTERFACE=Interface

mkdir -p mock_$NAME
mockgen \
  -destination=mock_$NAME/mock.go \
  $PKG/$NAME \
  $INTERFACE

echo "OK"
