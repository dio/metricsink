example-up: build/sink
	make -C example up

example-down:
	make -C example down

build/sink:
	CGO_ENABLED=0 GOOS=linux go build \
		-a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo \
		-o build/sink main.go

.PHONY: clean
clean:
	rm -fr build

.PHONY: clean example-up example-down
