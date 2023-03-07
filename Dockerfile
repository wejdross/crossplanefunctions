# FROM golang:latest AS build
# WORKDIR /build
# COPY go.mod /build/
# COPY go.sum /build/
# RUN go mod download
# COPY main.go /build/
# RUN go build -o /build/fnio


# FROM ubuntu:latest
# COPY --from=build /build/fnio /fnio
# ENTRYPOINT [ "/fnio" ]

FROM python:3.9-slim-buster AS build
WORKDIR /build
COPY requirements.txt /build/requirements.txt
COPY main2.py /fnio
RUN pip3 install -r requirements.txt

ENTRYPOINT [ "python3", "/fnio" ]