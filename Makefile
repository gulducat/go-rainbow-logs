bin/rainbow: bin
	go build -o ./bin/rainbow ./cmd/rainbow/

bin:
	mkdir -p bin

.PHONY: clean
clean:
	rm -rf ./bin
