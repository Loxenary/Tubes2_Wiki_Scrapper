@echo off
taskkill /f /im go.exe /t
start cmd /k "cd wikiscrapper\backend && go run main.go prioqueue.go bfs.go ids.go safemap.go crawler.go links_util.go"
