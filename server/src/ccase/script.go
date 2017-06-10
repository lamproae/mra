package ccase

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Script struct {
	Commands []*Command
}

func RunUserScript(in url.Values) (*Script, error) {
	commands, err := GetAllCommandFromRequest(in)
	if err != nil {
		return nil, err
	}

	log.Printf("Get Script: %#q", commands)

	return &Script{
		Commands: commands,
	}, nil
}

func GetAllCommandFromRequest(in url.Values) ([]*Command, error) {
	commandmap := make(map[string]*Command, 1)
	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "command_mode" {
				if _, ok := commandmap[fields[1]]; ok {
					log.Println("Same Command alread exist: ", k)
					continue
				}
				commandmap[fields[1]] = &Command{Mode: v[0], Seq: fields[1]}
			}
		}
	}

	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "command_cli" {
				commandmap[fields[1]].Cli = v[0]
			} else if fields[0] == "command_delay" {
				commandmap[fields[1]].Delay = v[0]
			} else if fields[0] == "command_expected" {
				commandmap[fields[1]].Expected = v[0]
			}
		}
	}

	if len(commandmap) == 0 {
		return nil, errors.New("Get no Command from input request")
	}

	commands := make([]*Command, len(commandmap), len(commandmap))
	for k, c := range commandmap {
		log.Printf("Find New Command: %q in requst", c)

		j, err := strconv.Atoi(k)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		commands[j] = c
	}

	return commands, nil
}
