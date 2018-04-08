build:
	go build -o myhead src/*.go

clean:
	rm myhead

test: setup
	richgo test -v -race ./...

setup:
	go get github.com/kyoh86/richgo
