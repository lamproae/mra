package ccase

import (
	"errors"
	"sort"
)

type Group struct {
	Name      string
	SGCount   int
	CCount    int
	SubGroups map[string]*SubGroup
}

func (g *Group) Add(c *Case) error {
	sg, ok := g.SubGroups[c.SubGroup]
	if !ok {
		g.SubGroups[c.SubGroup] = &SubGroup{
			Name:     c.SubGroup,
			Features: make(map[string]*Feature, 1),
		}
		g.SGCount++
		sg, _ = g.SubGroups[c.SubGroup]
	}

	err := sg.Add(c)
	if err != nil {
		return err
	}

	g.CCount++

	return nil
}

func (g *Group) Del(c *Case) error {
	sg, ok := g.SubGroups[c.SubGroup]
	if !ok {
		return errors.New("Cannot find Feature: " + c.Feature + " in Group: " + c.Group + " for delete case: " + c.Name)
	}

	err := sg.Del(c)
	if err != nil {
		return err
	}

	if len(sg.Features) == 0 {
		delete(g.SubGroups, c.SubGroup)
		g.SGCount--
	}

	g.CCount--
	return nil
}

func (g *Group) Get(c *Case) (*Case, error) {
	sg, ok := g.SubGroups[c.SubGroup]
	if !ok {
		return nil, errors.New("Cannot find Feature: " + c.Feature + " in Group: " + c.Group + " for Get case: " + c.Name)
	}

	return sg.Get(c)
}

func (g *Group) Dump() []*Case {
	result := make([]*Case, 0, 10)
	sgs := make([]*SubGroup, 0, len(g.SubGroups))

	for _, sg := range g.SubGroups {
		sgs = append(sgs, sg)
	}

	//sort.Slice(fs, func(i, j int) bool { return fs[i].Name < fs[j].Name })
	sort.Stable(SubGroupSlice(sgs))
	for _, sg := range sgs {
		result = append(result, sg.Dump()...)
	}

	return result
}

func (g *Group) DumpSubGroup(sgroup string) ([]*Case, error) {
	sg, ok := g.SubGroups[sgroup]
	if !ok {
		return nil, errors.New("Cannot find SubGroup: " + sgroup + " for dump")
	}

	return sg.Dump(), nil
}

func (g *Group) DumpFeature(sgroup, feature string) ([]*Case, error) {
	sg, ok := g.SubGroups[sgroup]
	if !ok {
		return nil, errors.New("Cannot find SubGroup: " + sgroup + " for dump")
	}

	return sg.DumpFeature(feature)
}

type SubGroupSlice []*SubGroup

func (s SubGroupSlice) Len() int           { return len(s) }
func (s SubGroupSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SubGroupSlice) Less(i, j int) bool { return s[i].Name < s[j].Name }
