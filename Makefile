.PHONY: build

build:
	go build -ldflags '-w -s -linkmode internal -extldflags "-static"' \
	-o vector-ble \
	cmd/main.go