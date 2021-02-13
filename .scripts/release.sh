#! /bin/bash

if [ ! -d release ]; then
    mkdir release
else
    rm -rf release/*
fi

VERSION=`git describe --tags`

docker build --build-arg VERSION=$VERSION -t bonedaddy/circuit-breaker:$VERSION .
docker image save bonedaddy/circuit-breaker:$VERSION --output release/circuit-breaker-docker_$VERSION.tar

go build -o release/"circuit-breaker-$VERSION" -ldflags "-X main.Version=$VERSION" ./cmd

ls ./release/* > files
for i in $(cat files); do
    sha256sum "$i" > "$i.sha256"
done