.PHONY: build
build:
	tinygo build -wasm-abi=generic -target=wasi -gc=leaking -no-debug -o records.wasm main.go
	wasm-tools component new records.wasm -o records.wasm --adapt wasi_snapshot_preview1.wasm --wit world.wit
	python -m wasmtime.bindgen records.wasm --out-dir host

test: build
	python host.py