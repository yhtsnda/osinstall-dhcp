#!/usr/bin/env bash
__SCRIPTDIR="${0%/*}"
#echo $__SCRIPTDIR
__REAL_SCRIPTDIR=$( cd -P -- "$(dirname -- "$(command -v -- "$0")")" && pwd -P )
#echo $__REAL_SCRIPTDIR

sudo docker  run  --rm  -e GOOS=linux -e GOARCH=386 -v $__REAL_SCRIPTDIR:/go/src/github.com/JinLinGan/osinstall-dhcp  golang sh -c  "env && cd /go/src/github.com/JinLinGan/osinstall-dhcp && go build "

sudo docker build -t jinlingan/osinstall-dhcp:latest -f ./Dockerfile.scratch .

rm -rf $__REAL_SCRIPTDIR/osinstall-dhcp
