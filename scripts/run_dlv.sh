#!/usr/bin/env bash
set -e

TARGET="${1:-api}"
case "$TARGET" in
    api|spinner|web) ;;
    *)
        echo "Usage: $0 [api|spinner]"
        exit 1
        ;;
esac

if [ "$TARGET" = "spinner" ]; then
    CMD="cmd/spinner/main.go"
elif [ "$TARGET" = "web" ]; then
    CMD="cmd/web/main.go"
else
    CMD="cmd/api/main.go"
fi

if [ ! -f "$CMD" ]; then
    echo "Package not found: $CMD"
    exit 1
fi

kill_process() {
    kill 0
}

trap kill_process SIGINT SIGTERM

go tool dlv debug \
    --headless \
    --listen=:2345 \
    --api-version=2 \
    $CMD
