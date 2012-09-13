#!/bin/sh
APP="website-admin"
ADDR=":8086"
PWD=`pwd`/..

echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."

export GOPATH=${PWD}/src/add-on:${PWD}

if [ ! -d ./add-on/src/golanger.com/framework ]; then
    echo "go get -u golanger.com/framework"
    go get -u golanger.com/framework
fi

if [ -f ${APP} ]; then
    rm ${APP}
fi

echo "Building ${APP}"
go build .

if [ -f src ]; then
    mv ./src ${APP}
    echo "Runing ${APP}"
    ./$APP -addr=${ADDR}
fi