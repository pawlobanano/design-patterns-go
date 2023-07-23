abstract-factory:
	go build -o abstract-factory ./creational/abstract-factory && ./abstract-factory

builder:
	go build -o builder ./creational/builder && ./builder

simple-factory:
	go build -o simple-factory ./creational/simple-factory && ./simple-factory

.PHONY: abstract-factory builder simple-factory