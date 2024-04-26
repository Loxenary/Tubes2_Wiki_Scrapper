@echo off
call ./stop.bat

@REM ganti file dibawah pake main.go kalo udah kelar
start cmd /k "cd wikiscrapper\backend && go run main.go prioqueue.go bfs.go links.go ids.go"
start cmd /k "cd wikiscrapper && npm run dev"
