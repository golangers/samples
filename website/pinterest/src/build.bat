@echo off
set APP=pinterest.exe
set ADDR=:8083
set PWD=%cd%\..
echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."

set GOPATH=%PWD%\src\add-on;%PWD%

if not exist %PWD%\src\add-on\src\golanger.com\framework  (
    echo "go get -u golanger.com/framework"
    go get -u golanger.com/framework
)

if exist %APP% del %APP%

echo "Building %APP%"
go build .

if exist src.exe (
    rename src.exe %APP%
    echo "Runing %APP%"
    %APP% -addr=%ADDR%
)
