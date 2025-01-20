./pdq:
	CGO_ENABLED=1 go build -ldflags '-extldflags "-static"' pdq.go
	-ldd pdq

install: ./pdq
	cp ./pdq /usr/local/bin


