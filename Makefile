# ----------------------- #
# file to build your code #
# ----------------------- #

# build for static services like github pages
build-web:
	- GOOS=js GOARCH=wasm go build -o web/yourgame.wasm
	- cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./web


# build for your desktop
build-desk: 
	- go build main.go

run: 
	- wasmserve ./main.go

# VALITD

# run-web: build-web
# 	- if chromium -v /dev/null