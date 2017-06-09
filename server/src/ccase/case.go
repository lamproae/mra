package ccase

import (
	"errors"
	"log"
	"net/url"
	//"strconv"
	"strings"
)

type Case struct {
	Group     string
	SubGroup  string
	Feature   string
	Name      string
	Tasks     []*Task
	DUTs      []*DUT
	TaskCount int
}

type Task struct {
	Name          string
	Seq           int
	PreCondition  Condition
	Routine       Routine
	PostCondition Condition
	Description   string
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

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
