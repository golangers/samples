@echo off
set APP=webserver.exe
set ADDR=:8082
set PWD=%cd%\..
echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."

set GOPATH=%PWD%\src\add-on;%PWD%

if not exist %PWD%\src\add-on\src\golanger.com\framework\web  (
    set GOCMD="go get -d golanger.com/framework/web"
    echo %GOCMD%
    %GOCMD%
)

if not exist %PWD%\src\add-on\src\golanger.com\middleware  (
    set GOCMD="go get -d golanger.com/middleware"
    echo %GOCMD%
    %GOCMD%
)

if exist %APP% del %APP%

echo "Building %APP%"
go build .

if exist src.exe (
    rename src.exe %APP%
    echo "Runing %APP%"
    %APP% -addr=%ADDR%
)