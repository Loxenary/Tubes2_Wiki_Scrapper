@echo off
taskkill /f /im go.exe /t

start cmd /k "cd wikiscrapper\backend\test && go run test.go"