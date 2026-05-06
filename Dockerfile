# syntax=docker/dockerfile:1

FROM golang:1.26.1 AS setup

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p /opt/spinner

FROM setup
ENV BIN_NAME_PREFIX=spinner

RUN make build-api
RUN make install-api
CMD ["/opt/spinner/spinner-api"]
