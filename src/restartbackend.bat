@echo off
taskkill /f /im go.exe /t

start cmd /k "cd wikiscrapper\backend && go run main.go prioqueue.go links.go ids.go ids_node.go node.go safemap.go ids_colly.go"
