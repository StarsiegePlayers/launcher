@echo off
rem windows makefile????

set version=1.0.0
set ldflags=-H windowsgui -X 'main.VERSION=%version%' -X 'main.TIME=%TIME%' -X 'main.DATE=%DATE%'
set GOOS=windows
set GOARCH=386

if "%1"=="" call :release
if "%1"=="release" call :release
if "%1"=="debug" call :debug
if "%1"=="generate" call :generate
if "%1"=="clean" call :clean
goto end

:generate
echo ---generate---
echo go generate -x
echo.
go generate -x
goto :eof

:debug
echo =====Debug Build started %DATE% %TIME%=====
echo.
call :clean
call :generate
echo ---debug---
echo go build -x -ldflags="%ldflags%"
echo.
go build -x -ldflags="%ldflags%"
goto :eof

:release
echo =====Release Build started %DATE% %TIME%=====
echo.
call :clean
call :generate
set ldflags=%ldflags% -s -w
echo ---release---
echo go build -x -ldflags="%ldflags%"
echo.
go build -x -ldflags="%ldflags%"
goto :eof

:clean
echo ---clean---
echo del *.exe *.syso
echo.
del *.exe *.syso >nul 2>nul
goto :eof

:end
