.PHONY: run bootstrap

run:
	odin build . -target:js_wasm32 -extra-linker-flags:"--export-table" -out:server/public/ljosalf.wasm
	cd ./server && go run .

bootstrap:
	cp $(HOME)/odin/dist/core/sys/wasm/js/odin.js server/public
	cp $(HOME)/odin/dist/vendor/wgpu/wgpu.js server/public
