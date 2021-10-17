package main

import (
	"bytes"
	"go/build"
	"net/http"
	"os"
	"os/exec"
)

var (
	Outdir = "build/"
	Srcdir = "src/"

	GOROOT = build.Default.GOROOT
)

const (
	web_dir  = "build-web/"
	htmlText = `
	<!DOCTYPE html>
	<script src="wasm_exec.js"></script>
	<script>
	// Polyfill
	if (!WebAssembly.instantiateStreaming) {
	  WebAssembly.instantiateStreaming = async (resp, importObject) => {
		const source = await (await resp).arrayBuffer();
		return await WebAssembly.instantiate(source, importObject);
	  };
	}
	
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(result => {
	  go.run(result.instance);
	});
	</script>
	`
)

func main() {

	BuildWeb(true)
	test()
	// test()
}

func test() {

	// println(a.a)
}

func BuildDir(path string) {

	// go file
	os.Mkdir("build/"+path, 0777)
	cmd := exec.Command("go", "build", "-o", Outdir+path+"main.wasm", Srcdir+path+"main.go")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	cmd.Start()

	//html file
	os.WriteFile("build/"+path+"index.html", []byte(htmlText), 0755)

}
func BuildWeb(serve bool) {

	os.Mkdir(web_dir, 0777)

	// copy the static resources on build
	hcp := exec.Command("cp", "-r", "assets/", web_dir)
	hcp.Start()
	wecp := exec.Command("cp", GOROOT+"/misc/wasm/wasm_exec.js", web_dir)
	wecp.Start()
	cmd := exec.Command("go", "build", "-o", web_dir+"main.wasm", Srcdir+"main.go")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	cmd.Run()
	println("out:", outb.String(), "err:", errb.String())

	//html file
	os.WriteFile(web_dir+"index.html", []byte(htmlText), 0755)

	/// serve the build on browser
	if serve {
		http.ListenAndServe(":3000", http.FileServer(http.Dir(web_dir)))
	}
}
