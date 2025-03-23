package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	address := "127.0.0.1:3030"

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveWasm)
	http.ListenAndServe(address, nil)
}

func serveWasm(writer http.ResponseWriter, request *http.Request) {

	pageHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Odin WebGPU Example</title>
</head>
<body>
<canvas id="wgpu-canvas" style="height: 100%; width: 100%;"></canvas>
	<script type="text/javascript" src="/static/odin.js"></script>
	<script type="text/javascript" src="/static/wgpu.js"></script>
	<script type="text/javascript">
		const mem = new WebAssembly.Memory({ initial: 2000, maximum: 65536, shared: false });
		const memInterface = new odin.WasmMemoryInterface();
		memInterface.setMemory(mem);

		const wgpuInterface = new odin.WebGPUInterface(memInterface);

		odin.runWasm("/static/ljosalf.wasm", null, { wgpu: wgpuInterface.getInterface() }, memInterface, /*intSize=8*/);
		//odin.runWasm("/static/odin_stripped.wasm", null, { wgpu: wgpuInterface.getInterface() }, memInterface, /*intSize=8*/);
	</script>
</body>
</html>`

	_, err := fmt.Fprint(writer, pageHTML)
	if err != nil {
		log.Fatal(err)
	}
}
