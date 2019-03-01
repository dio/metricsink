.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build \
		-a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo \
		-o build/sink main.go

.PHONY: clean
clean:
	rm -fr build
