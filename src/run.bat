@echo off
call ./stop.bat

start cmd /k "cd wikiscrapper && npm run dev"
@REM ganti file dibawah pake main.go kalo udah kelar
start cmd /k "cd wikiscrapper\backend\test && go run test.go"