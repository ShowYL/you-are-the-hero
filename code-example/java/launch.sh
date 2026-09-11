#!/usr/bin/env bash

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"

javac "$SCRIPT_DIR/Main.java" && java -cp "$SCRIPT_DIR" Main
