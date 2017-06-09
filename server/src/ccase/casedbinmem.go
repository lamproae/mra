package ccase

import (
	"errors"
	"sort"
)

type CaseDBInMem struct {
	Device string
	GCount int
	FCount int
	CCount int
	Groups map[string]*Group
}

func (cdbim *CaseDBInMem) Add(c *Case) error {
	g, ok := cdbim.Groups[c.Group]
	if !ok {
		cdbim.Groups[c.Group] = &Group{
			Name:     c.Group,
			Features: make(map[string]*Feature, 1),
		}
		cdbim.GCount++
		g, _ = cdbim.Groups[c.Group]
	}

	return g.Add(c)
}

func (cdbim *CaseDBInMem) Del(c *Case) error {
	g, ok := cdbim.Groups[c.Group]
	if !ok {
		return errors.New("Cannot find Group: " + c.Group + " for delete case: " + c.Name)
	}

	err := g.Del(c)
	if err != nil {
		return err
	}

	if len(g.Features) == 0 {
		delete(cdbim.Groups, c.Group)
		cdbim.GCount--
	}

	return nil
}

func (cdbim *CaseDBInMem) Get(c *Case) (*Case, error) {
	g, ok := cdbim.Groups[c.Group]
	if !ok {
		return nil, errors.New("Cannot find Group: " + c.Group + " for Get case: " + c.Name)
	}

	return g.Get(c)
}

func (cdbim *CaseDBInMem) Dump() []*Case {
	result := make([]*Case, 0, 10)
	gs := make([]*Group, 0, len(cdbim.Groups))

	for _, g := range cdbim.Groups {
		gs = append(gs, g)
	}

	//sort.Slice(gs, func(i, j int) bool { return gs[i].Name < gs[j].Name })
	sort.Stable(GroupSlice(gs))
	for _, g := range gs {
		result = append(result, g.Dump()...)
	}

	return result
}

func (cdbim *CaseDBInMem) DumpGroup(group string) ([]*Case, error) {
	g, ok := cdbim.Groups[group]
	if !ok {
		return nil, errors.New("Cannot find Group: " + group + " for dump")
	}

	return g.Dump(), nil
}

func (cdbim *CaseDBInMem) DumpFeature(group, feature string) ([]*Case, error) {
	g, ok := cdbim.Groups[group]
	if !ok {
		return nil, errors.New("Cannot find Group: " + group + " for dump feature")
	}

	return g.DumpFeature(feature)
}

type GroupSlice []*Group

func (s GroupSlice) Len() int           { return len(s) }
func (s GroupSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s GroupSlice) Less(i, j int) bool { return s[i].Name < s[j].Name }
