#! /bin/bash -e

NAME=unik
INTERFACE=Interface
PKG=github.com/plimble

mockgen \
  -destination=mock.go \
  --self_package=$PKG/$NAME \
  -package=$NAME \
  $PKG/$NAME \
  $INTERFACE

echo "OK"
