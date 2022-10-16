echo 'make dirs and files...'
mkdir -p cmd/api config models server
touch cmd/api/main.go config/config.yml config/init.go server/app.go Dockerfile docker-compose.yml Makefile README.md

echo 'package config' > config/init.go
echo 'package server' > server/app.go
echo 'package main' > cmd/api/main.go