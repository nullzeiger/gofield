build: main.go
	go build -o build/gofield main.go

run: build/gofield
	go run main.go

clean: build/gofield
	rm -rf build
