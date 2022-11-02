.PHONY: all
all: clean build

.PHONY: clean
clean:
	rm -f p1-eporter

.PHONY: build
build: p1-exporter

p1-exporter:
	go build -o p1-exporter main.go

.PHONY: install
install: p1-exporter
	cp p1-exporter /usr/local/bin/
