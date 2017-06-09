package ccase

import (
	"errors"
	"log"
	"net/url"
	"strings"
)

type DUT struct {
	Name     string
	Device   string
	UserName string
	Password string
}

type Condition struct {
	Assertions  []*Assertion
	Description string
}

type Routine struct {
	Assertions  []*Assertion
	Description string
}

type Command struct {
	Mode string
	Cli  string
}

type Assertion struct {
	DUT      DUT
	Command  Command
	Expected string
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
