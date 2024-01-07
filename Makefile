
.SILENT:

build:
	go build -o lib_plugin.so -buildmode=plugin lib.go
	file lib_plugin.so
	go build -o app main.go