package main

type CaseDBInMem struct {
	Device string
	GCount int
	FCount int
	CCount int
	Groups []*Group
}

func (cdbim *CaseDBInMem) Add(c *Case) error {

}

func (cdbim *CaseDBInMem) Del(c *Case) error {

}

func (cdbim *CaseDBInMem) Get(c *Case) (*Case, error) {

}
