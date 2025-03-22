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
    <h1>Running Odin WebAssembly</h1>
    <script src="/static/wasm_loader.js"></script>
</body>
</html>`

	_, err := fmt.Fprint(writer, pageHTML)
	if err != nil {
		log.Fatal(err)
	}
}
