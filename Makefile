simple-factory:
	go build -o simple-factory ./creational/simple-factory && ./simple-factory

abstract-factory:
	go build -o abstract-factory ./creational/abstract-factory && ./abstract-factory

.PHONY: simple-factory abstract-factory