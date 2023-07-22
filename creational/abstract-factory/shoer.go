package main

type Shoer interface {
	setLogo(logo string)
	setSize(size int)
	logo() string
	size() int
}

type Shoe struct {
	logo_ string
	size_ int
}

func (s *Shoe) setLogo(brandName string) {
	s.logo_ = brandName
}

func (s *Shoe) logo() string {
	return s.logo_
}

func (s *Shoe) setSize(size int) {
	s.size_ = size
}

func (s *Shoe) size() int {
	return s.size_
}
