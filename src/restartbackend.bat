@echo off
taskkill /f /im go.exe /t

start cmd /k "cd wikiscrapper\backend && go run main.go prioqueue.go"