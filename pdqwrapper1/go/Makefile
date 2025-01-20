default: clean install

./pdq:
	CGO_ENABLED=1 go build -ldflags '-extldflags "-static"' pdq.go
	@ldd pdq 2>/dev/null && (echo "ldd pdq # this should have been 'not a dynamic executable'", halting"; exit 1") || echo "ldd pdq # this was non-executable, as intended"


	./pdq -s 0.1 -z 0.9 1 1 1

install: ./pdq
	cp ./pdq /usr/local/bin
	pdq -s 0.1 -z 0.9 1 10 1

test:
	go test .

clean:
	rm ./pdq
