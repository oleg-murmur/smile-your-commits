#!/bin/bash

set -e

# Перейти в корень проекта независимо от места запуска скрипта
cd "$(dirname "$0")/.."

SRC="./main"        # <-- Путь до main.go, относительно корня проекта
OUT="./build"

mkdir -p "$OUT"

echo "🔨 Building for Windows (x64)..."
GOOS=windows GOARCH=amd64 go build -o "$OUT/smile-your-commits.exe" "$SRC"

echo "🐧 Building for Linux (x64)..."
GOOS=linux GOARCH=amd64 go build -o "$OUT/smile-your-commits-linux" "$SRC"

echo "🍏 Building for macOS (x64)..."
GOOS=darwin GOARCH=amd64 go build -o "$OUT/smile-your-commits-macos" "$SRC"

echo "✅ Done. Binaries are in the ./build folder."
