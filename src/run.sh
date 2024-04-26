#!/bin/bash

# Call the stop script
./stop.bat

# Change the file below to main.go once it's ready
gnome-terminal --working-directory=wikiscrapper/backend -e "go run main.go prioqueue.go bfs.go links.go ids.go safemap.go"
gnome-terminal --working-directory=wikiscrapper -e "npm run dev"