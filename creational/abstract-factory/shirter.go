package main

type Shirter interface {
	setLogo(logo string)
	setSize(size int)
	logo() string
	size() int
}

type Shirt struct {
	logo_ string
	size_ int
}

func (s *Shirt) setLogo(brandName string) {
	s.logo_ = brandName
}

func (s *Shirt) logo() string {
	return s.logo_
}

func (s *Shirt) setSize(size int) {
	s.size_ = size
}

func (s *Shirt) size() int {
	return s.size_
}
