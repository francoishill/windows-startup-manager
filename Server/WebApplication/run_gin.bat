@echo off
cls

cd %GOPATH%\src\github.com\francoishill\windows-startup-manager\Server\WebApplication

gin -p 12345 -a 54321 -i run "config/server.gcfg"
pause