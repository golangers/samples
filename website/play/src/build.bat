@echo off
set APP=play.exe
set ADDR=:8087
set PWD=%cd%\..
echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."

set GOPATH=%PWD%\src\add-on;%PWD%
cd %PWD%\src

if exist %APP% del %APP%

echo "Building %APP%"
go build .

if exist src.exe (
    rename src.exe %APP%
    echo "Runing %APP%"
    %APP% -addr=%ADDR%
)
