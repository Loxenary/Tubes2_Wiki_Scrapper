#!/bin/bash

# Run the Go backend
cd wikiscrapper/backend
go run main.go prioqueue.go bfs.go ids.go safemap.go crawler.go links_util.go &

# Run the Node.js frontend
cd ../
npm run dev