#!/bin/bash

set -e

echo "Building Gremlin Benchmark Binaries..."
echo ""

mkdir -p binaries

# FreeBSD AMD64
echo "Building for FreeBSD (amd64)..."
GOOS=freebsd GOARCH=amd64 go build -o binaries/gremlin-bench-freebsd-amd64 ./cmd/benchmark
echo "✅ binaries/gremlin-bench-freebsd-amd64"

# Linux AMD64
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o binaries/gremlin-bench-linux-amd64 ./cmd/benchmark
echo "✅ binaries/gremlin-bench-linux-amd64"

# Linux ARM64
echo "Building for Linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -o binaries/gremlin-bench-linux-arm64 ./cmd/benchmark
echo "✅ binaries/gremlin-bench-linux-arm64"

# macOS ARM64
echo "Building for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o binaries/gremlin-bench-darwin-arm64 ./cmd/benchmark
echo "✅ binaries/gremlin-bench-darwin-arm64"

echo ""
echo "All binaries built successfully!"
echo ""
echo "Directory contents:"
ls -lh binaries/
echo ""
echo "To run on target device:"
echo "  FreeBSD (amd64):    ./gremlin-bench-freebsd-amd64"
echo "  Linux (amd64):      ./gremlin-bench-linux-amd64"
echo "  Linux (arm64):      ./gremlin-bench-linux-arm64"
echo "  macOS (arm64):      ./gremlin-bench-darwin-arm64"
echo ""
echo "Default: 10,000,000 iterations (~30-60 seconds total)"
echo "Custom:  ./gremlin-bench-darwin-arm64 -n 1000000"
