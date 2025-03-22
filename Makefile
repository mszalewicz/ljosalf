.PHONY: run bootstrap  build

run:
	mkdir -p server/public
	odin build . -target:js_wasm32 -extra-linker-flags:"--export-table" -out:server/public/ljosalf.wasm
	cd ./server && go run .

bootstrap:
	cp $(HOME)/odin/dist/core/sys/wasm/js/odin.js server/public
	cp $(HOME)/odin/dist/vendor/wgpu/wgpu.js server/public

build: bootstrap
	mkdir -p bin
	odin build . -target:js_wasm32 -extra-linker-flags:"--export-table" -out:server/public/ljosalf.wasm
	cp -r server/public bin
	cd server && go build -o server.bin .
	cp server/server.bin bin
