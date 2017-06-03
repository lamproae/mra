package cli_test

import (
	"cli"
	"config"
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	c, err := cli.NewCli(&config.Config{
		DeviceName:     "V8500",
		IP:             "10.71.20.198",
		Port:           "23",
		UserName:       "admin",
		Password:       "",
		EnablePrompt:   ">",
		LoginPrompt:    "login",
		PasswordPrompt: "Password",
		Prompt:         "#",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}
