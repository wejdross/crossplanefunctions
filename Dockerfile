FROM golang:latest AS build
WORKDIR /build
COPY go.mod /build/
COPY go.sum /build/
RUN go mod download
COPY main.go /build/
RUN go build -o /build/fnio


FROM ubuntu:latest
COPY --from=build /build/fnio /fnio
ENTRYPOINT [ "/fnio" ]