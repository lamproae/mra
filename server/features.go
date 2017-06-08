package main

type Feature struct {
	Name   string
	CCount int
	Cases  []*Cases
}

func (f *Feature) Add(c *Case) error {

}

func (f *Feature) Del(c *Case) error {

}

func (f *Feature) Get(c *Case) (*Case, error) {

}
