# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A Go program that reads a coded `input.txt` file and prints animated text to a terminal. The program compiles to WebAssembly (WASM) and runs in a browser via a Next.js + xterm.js frontend.

## Commands

### Go (CLI)
```bash
# Run directly
go run .

# Compile to WebAssembly (required after any main.go/main_wasm.go changes)
GOOS=js GOARCH=wasm go build -o web/public/go/main.wasm .

# If wasm_exec.js needs to be refreshed (after a Go version upgrade):
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/public/wasm_exec.js
```

### Next.js frontend (run from `web/` directory)
```bash
npm install       # Install dependencies (first time)
npm run dev       # Dev server at localhost:3000
npm run build     # Production build
npm run start     # Run production server
```

> No Makefile, linter config, or test suite exists.

## Architecture

### Go (3 source files + `go.mod`)
- `main.go` — shared logic: `//go:embed` for all 3 txt files, `var printFn func(string)` output abstraction, `typePrint()`, `dance()`, `typeDelete()`, `checkErr()`
- `main_cli.go` (`//go:build !js`) — sets `printFn = fmt.Print`, contains `main()` for CLI use
- `main_wasm.go` (`//go:build js && wasm`) — sets `printFn` to call JS `__goOutput`, contains `main()` that calls `__goDone()` then `select{}` to keep runtime alive

`typePrint()` special characters:
- `$` → recurse with `goodJobTxt` at 2ms delay
- `#` → HTTP GET California tours API, print response
- `@` → sleep 500ms
- `^` → backspace (`\b \b`)
- `&` → dance animation (splits `danceTxt` by line, renders frames with ANSI clear)
- default → print char after random 0–100ms delay

### Browser integration (`web/`)
- `app/page.tsx`: loads `wasm_exec.js` via `<Script strategy="afterInteractive" onLoad={initWasm}>`, dynamically imports xterm, fetches/instantiates `main.wasm`, renders xterm terminal (neon green `#39ff14` on black)
- Two JS globals the WASM binary calls:
  - `__goOutput(text)` — converts `\n` → `\r\n` then writes to xterm
  - `__goDone()` — sets `isDone` state, reveals email + LinkedIn link bar
- `web/public/go/main.wasm` — compiled binary; rebuild after any Go source changes
- `web/public/wasm_exec.js` — Go's WASM runtime; must match the Go version used to compile

### Configuration
- `next.config.js`: `reactStrictMode: false` — required; strict mode double-invokes effects, causing WASM to run twice
- `tsconfig.json`: target ES2022, strict mode, Next.js plugin
