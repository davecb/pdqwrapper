default: clean install

./pdq:
	CGO_ENABLED=1 go build -ldflags '-extldflags "-static"' pdq.go
	@ldd pdq 2>/dev/null && (echo "ldd pdq # this should have been 'not a dynamic executable'", halting"; exit 1") || echo "ldd pdq # this was non-executable, as intended"
	./pdq -s 0.1 -z 0.9 1 1 1

install: ./pdq
	cp ./pdq /usr/local/bin
	pdq -s 0.1 -z 0.9 1 10 1

test:
	go test -v -run TestPdq

single: ./pdq # run just one case. for debugging
	go test -v -run TestSingleton -args -z 1.7 -s 1.7976931348623157e+308 1 10000 1

clean:
	rm ./pdq

default: # for intellij
	CGO_ENABLED=1 go build -ldflags '-extldflags "-static"' pdq.go
	go test -v -run TestSingleton -args -z 1.7 -s -1.7976931348623157e+308 1 10000 1

