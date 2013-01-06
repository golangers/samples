#!/bin/sh
APP="chatroom"
ADDR=":8080"
PWD=`pwd`/..

echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."

export GOPATH=${PWD}/src/add-on:${PWD}

if [ ! -d ./add-on/src/golanger.com/framework ]; then
    echo "go get golanger.com/framework"
    go get golanger.com/framework
fi

if [ ! -d ./add-on/src/golanger.com/i18n ]; then
    echo "go get golanger.com/i18n"
    go get golanger.com/i18n
fi

if [ ! -d ./add-on/src/golanger.com/log ]; then
    echo "go get golanger.com/log"
    go get golanger.com/log
fi

if [ ! -d ./add-on/src/golanger.com/session ]; then
    echo "go get golanger.com/session"
    go get golanger.com/session
fi

if [ ! -d ./add-on/src/golanger.com/urlmanage ]; then
    echo "go get golanger.com/urlmanage"
    go get golanger.com/urlmanage
fi

if [ ! -d ./add-on/src/golanger.com/validate ]; then
    echo "go get golanger.com/validate"
    go get golanger.com/validate
fi

if [ ! -d ./add-on/src/golanger.com/utils ]; then
    echo "go get golanger.com/utils"
    go get golanger.com/utils
fi

if [ ! -d ./add-on/src/golanger.com/middleware ]; then
    echo "go get golanger.com/middleware"
    go get golanger.com/middleware
fi

if [ ! -d ./add-on/src/golanger.com/database ]; then
    echo "go get golanger.com/database"
    go get golanger.com/database
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