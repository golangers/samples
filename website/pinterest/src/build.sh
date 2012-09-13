#!/bin/sh
APP="pinterest"
ADDR=":8083"
PWD=`pwd`/..

echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."
export GOPATH=${PWD}/src/add-on:${PWD}
cd ${PWD}/src

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