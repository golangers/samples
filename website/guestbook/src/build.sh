#!/bin/sh
APP="guestbook"
ADDR=":8081"

echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."

if [[ ${1} == "nb" && -f ${APP} ]]; then
    echo "Runing ${APP}"
    ./$APP -addr=${ADDR}
    exit 0
fi

PWD=`pwd`/..
GO_GET_LIST="framework/web middleware"
ADD_ON="${PWD}/src/add-on"
export GOPATH=${ADD_ON}:${PWD}

goget() {
    pkg="golanger.com/${1}"
    if [[ ! -d "${ADD_ON}/src/${pkg}" ]]; then
        GOCMD="go get -d ${pkg}"
        echo ${GOCMD}
        ${GOCMD}
    fi 
}


for pkg in ${GO_GET_LIST}; do
    goget ${pkg}
done

if [ -f ${APP} ]; then
    rm ${APP}
fi

echo "Building ${APP}"
go build .

if [[ -f src ]]; then
    mv ./src ${APP}
    echo "Runing ${APP}"
    ./$APP -addr=${ADDR}
fi