#!/bin/bash

# Build the WASM binary
GOOS=js GOARCH=wasm go build -o web/main.wasm

# Ensure wasm_exec.js is present
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
