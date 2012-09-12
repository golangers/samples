@echo off
set APP=chatroom.exe
set ADDR=:8080
set PWD=%cd%\..
set GO_PWD=%PWD%\..\..
echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."

if not exist %PWD%\src\add-on\src\golanger.com  (
    md %PWD%\src\add-on\src\golanger.com
    xcopy /K /S %GO_PWD%\framework\*  %PWD%\src\add-on\src\golanger.com
)

set GOPATH=%GO_PWD%\framework;%PWD%\src\add-on;%PWD%
cd %PWD%\src

if exist %APP% del %APP%

echo "Building %APP%"
go build .

if exist src.exe (
    rename src.exe %APP%
    echo "Runing %APP%"
    %APP% -addr=%ADDR%
)
