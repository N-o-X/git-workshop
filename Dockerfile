FROM golang:1.20-alpine

RUN apk add git vim nano

RUN git config --global user.email "demo@example.com" && git config --global user.name "Demo User"
RUN git config --global --add safe.directory /git

WORKDIR /git
