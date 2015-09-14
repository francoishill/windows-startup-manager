@echo off
cls

go build -o windows_startup_manager.exe -ldflags "-H windowsgui"