@echo off
set APP=webserver.exe
set ADDR=:8082
set PWD=%cd%\..
echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."

set GOPATH=%PWD%\src\add-on;%PWD%

if not exist %PWD%\src\add-on\src\golanger.com\framework  (
    echo "go get golanger.com/framework"
    go get golanger.com/framework
)

if not exist %PWD%\src\add-on\src\golanger.com\i18n  (
    echo "go get golanger.com/i18n"
    go get golanger.com/i18n
)

if not exist %PWD%\src\add-on\src\golanger.com\log  (
    echo "go get golanger.com/log"
    go get golanger.com/log
)

if not exist %PWD%\src\add-on\src\golanger.com\session  (
    echo "go get golanger.com/session"
    go get golanger.com/session
)

if not exist %PWD%\src\add-on\src\golanger.com\urlmanage  (
    echo "go get golanger.com/urlmanage"
    go get golanger.com/urlmanage
)

if not exist %PWD%\src\add-on\src\golanger.com\validate  (
    echo "go get golanger.com/validate"
    go get golanger.com/validate
)

if not exist %PWD%\src\add-on\src\golanger.com\utils  (
    echo "go get golanger.com/utils"
    go get golanger.com/utils
)

if not exist %PWD%\src\add-on\src\golanger.com\middleware  (
    echo "go get golanger.com/middleware"
    go get golanger.com/middleware
)

if not exist %PWD%\src\add-on\src\golanger.com\database  (
    echo "go get golanger.com/database"
    go get golanger.com/database
)

if exist %APP% del %APP%

echo "Building %APP%"
go build .

if exist src.exe (
    rename src.exe %APP%
    echo "Runing %APP%"
    %APP% -addr=%ADDR%
)