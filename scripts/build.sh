#!/bin/bash

APP="DockerfileGen"

pushd ..
go build -o $APP
mv $APP build/package/.
popd
