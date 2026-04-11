#!/usr/bin/env bash
set -e

BIN=bin/badmoviespinner

kill_process() {
    kill 0
}

trap kill_process SIGINT SIGTERM

go tool dlv exec \
    --headless \
    --listen=:2345 \
    --api-version=2 \
    $BIN
