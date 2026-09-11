#!/usr/bin/env bash

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"

case "$1" in
    "1")
        gcc -o "$SCRIPT_DIR/build/mainV1" "$SCRIPT_DIR/mainV1.c" && "$SCRIPT_DIR/build/mainV1"
    ;;
    "2")
        gcc -o "$SCRIPT_DIR/build/mainV2" "$SCRIPT_DIR/mainV2.c" && "$SCRIPT_DIR/build/mainV2"
    ;;
    "3")
        gcc -o "$SCRIPT_DIR/build/mainV3" "$SCRIPT_DIR/mainV3.c" && "$SCRIPT_DIR/build/mainV3"
    ;;
    *)
        echo -e "No script with that number.\n"
    ;;
esac
