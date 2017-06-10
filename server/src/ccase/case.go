package ccase

import (
	"errors"
	"log"
	"net/url"
	//"strconv"
	"strings"
)

type Case struct {
	Group    string
	SubGroup string
	Feature  string
	Name     string
	Tasks    []*Task
	DUTs     []*DUT
	TCount   int
}

func IsValidCaseParas(in url.Values) bool {
	if v, ok := in["case_group"]; !ok {
		log.Println(v)
		return false
	} else if v[0] == "" {
		return false
	}

	if v, ok := in["case_sub_group"]; !ok {
		return false
	} else if v[0] == "" {
		return false
	}

	if v, ok := in["case_feature"]; !ok {
		return false
	} else if v[0] == "" {
		return false
	}

	if v, ok := in["case_name"]; !ok {
		return false
	} else if v[0] == "" {
		return false
	}

	for k, _ := range in {
		if strings.HasPrefix(k, "device") {
			//if _, err := strconv.ParseInt(k[3:], 10, 64); err != nil {
			return true
			//	}
		}
	}

	return false
}

func createNewCase(in url.Values) *Case {
	group, _ := in["case_group"]
	sgroup, _ := in["case_sub_group"]
	feature, _ := in["case_feature"]
	name, _ := in["case_name"]

	duts, _ := GetAllDutFromRequest(in)

	return &Case{
		Group:    group[0],
		SubGroup: sgroup[0],
		Feature:  feature[0],
		Name:     name[0],
		DUTs:     duts,
	}
}

func CreateNewCase(in url.Values) (*Case, error) {
	for k, v := range in {
		log.Println(k, v)
	}

	if !IsValidCaseParas(in) {
		return nil, errors.New("Invalid parameter for Create a new Case")
	}

	return createNewCase(in), nil
}

func GetAllDutFromRequest(in url.Values) ([]*DUT, error) {
	dutmap := make(map[string]*DUT, 1)
	duts := make([]*DUT, 0, 1)
	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "device" {
				if _, ok := dutmap[fields[1]]; ok {
					log.Println("Save DUT alread exist: ", k)
					continue
				}
				dutmap[fields[1]] = &DUT{Name: "DUT" + fields[1], Device: v[0]}
			}
		}
	}

	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "username" {
				dutmap[fields[1]].UserName = v[0]

			} else if fields[0] == "password" {
				dutmap[fields[1]].Password = v[0]
			}
		}
	}

	if len(dutmap) == 0 {
		return nil, errors.New("Get no DUT from input request")
	}

	for _, d := range dutmap {
		log.Printf("Find New DuT: %q in requst", d)
		duts = append(duts, d)
	}

	return duts, nil
}

func (c *Case) AddTask(t *Task) error {
	if c.Tasks == nil {
		c.Tasks = make([]*Task, 0, 1)
	}

	if c.IsTaskExist(t) {
		return errors.New("Same task :" + t.Name + " already exist in case: " + c.Name)
	}

	c.Tasks = append(c.Tasks, t)
	c.TCount++

	return nil
}

func (c *Case) DelTask(t *Task) error {
	if !c.IsTaskExist(t) {
		return errors.New("Cannot find task :" + t.Name + " in case: " + c.Name)
	}

	for i, v := range c.Tasks {
		if v.Name == t.Name {
			c.Tasks = append(c.Tasks[:i], c.Tasks[i+1:]...)
		}
	}

	c.TCount--
	return nil
}

func (c *Case) GetTask(name string) *Task {
	for _, v := range c.Tasks {
		if v.Name == name {
			return v
		}
	}

	return nil
}

func (c *Case) IsTaskExist(t *Task) bool {
	for _, v := range c.Tasks {
		if v.Name == t.Name {
			return true
		}
	}

	return false
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
