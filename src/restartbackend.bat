@echo off
taskkill /f /im go.exe /t
<<<<<<< HEAD

start cmd /k "cd wikiscrapper\backend && go run main.go prioqueue.go links.go ids.go safemap.go ids_colly.go"
=======
start cmd /k "cd wikiscrapper\backend && go run main.go prioqueue.go links.go ids.go bfs.go safemap.go"
>>>>>>> 7a1cff5a99f5f8765452f9c09d5c93e6a1a2cd0f
