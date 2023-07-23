package main

type Gunner interface {
	setName(name string)
	setPower(power int)
	Name() string
	Power() int
}
