# syntax=docker/dockerfile:1

FROM golang:1.26.1 AS setup

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p /opt/spinner

FROM setup AS build-api
RUN make build-api
RUN cp ./bin/badmoviespinner-api /opt/spinner/spinner-api
CMD ["/opt/spinner/spinner-api"]
