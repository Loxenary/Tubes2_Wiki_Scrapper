@echo off
call ./stop.bat

@REM ganti file dibawah pake main.go kalo udah kelar
start cmd /k "cd wikiscrapper\backend && go run main.go prioqueue.go bfs.go ids.go safemap.go crawler.go links_util.go"
start cmd /k "cd wikiscrapper && npm run dev"
