abstract-factory:
	go build -o abstract-factory ./creational/abstract-factory && ./abstract-factory

simple-factory:
	go build -o simple-factory ./creational/simple-factory && ./simple-factory

.PHONY: abstract-factory simple-factory