# syntax=docker/dockerfile:1

FROM golang:1.26.1 AS setup

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p /opt/spinner

ENV BIN_NAME_PREFIX=spinner

# TARGET: API
FROM setup AS api
RUN make build-api
RUN make install-api
CMD ["/opt/spinner/api/spinner-api"]

# TARGET: WEB
FROM setup AS web

CMD ["/opt/spinner/web/spinner-web"]
