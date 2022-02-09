build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r server make run
