@echo off
call ./stop.bat

@REM ganti file dibawah pake main.go kalo udah kelar
start cmd /k "cd wikiscrapper\backend && go run main.go prioqueue.go links.go ids.go ids_node.go node.go safemap.go ids_colly.go"
start cmd /k "cd wikiscrapper && npm run dev"
