# compile Go code into executable brokerApp
# given stage name of builder
# FROM golang:1.18-alpine as builder

# create directory called /app
# RUN mkdir /app

# copy everything from broker service to /app
# COPY . /app

# following commands set to /app
# WORKDIR /app

# not running C libraries, complies code with go tool, output is brokerApp, points to main.go
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# makes output file executable
# RUN chmod +x /app/brokerApp

# build tiny docker image, copy only compiled executable to smaller image (alpine)
FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]