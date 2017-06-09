package ccase

import (
	"errors"
	"sort"
)

type Group struct {
	Name     string
	FCount   int
	CCount   int
	Features map[string]*Feature
}

func (g *Group) Add(c *Case) error {
	f, ok := g.Features[c.Feature]
	if !ok {
		g.Features[c.Feature] = &Feature{
			Name:  c.Feature,
			Cases: make(map[string]*Case, 1),
		}
		g.FCount++
		f, _ = g.Features[c.Feature]
	}

	return f.Add(c)
}

func (g *Group) Del(c *Case) error {
	f, ok := g.Features[c.Feature]
	if !ok {
		return errors.New("Cannot find Feature: " + c.Feature + " in Group: " + c.Group + " for delete case: " + c.Name)
	}

	err := f.Del(c)
	if err != nil {
		return err
	}

	if len(f.Cases) == 0 {
		delete(g.Features, c.Feature)
		g.FCount--
	}

	return nil
}

func (g *Group) Get(c *Case) (*Case, error) {
	f, ok := g.Features[c.Feature]
	if !ok {
		return nil, errors.New("Cannot find Feature: " + c.Feature + " in Group: " + c.Group + " for Get case: " + c.Name)
	}

	return f.Get(c)
}

func (g *Group) Dump() []*Case {
	result := make([]*Case, 0, 10)
	fs := make([]*Feature, 0, len(g.Features))

	for _, g := range g.Features {
		fs = append(fs, g)
	}

	//sort.Slice(fs, func(i, j int) bool { return fs[i].Name < fs[j].Name })
	sort.Stable(FeatureSlice(fs))
	for _, f := range fs {
		result = append(result, f.Dump()...)
	}

	return result
}

func (g *Group) DumpFeature(feature string) ([]*Case, error) {
	f, ok := g.Features[feature]
	if !ok {
		return nil, errors.New("Cannot find Group: " + feature + " for dump")
	}

	return f.Dump(), nil
}

type FeatureSlice []*Feature

func (s FeatureSlice) Len() int           { return len(s) }
func (s FeatureSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s FeatureSlice) Less(i, j int) bool { return s[i].Name < s[j].Name }
