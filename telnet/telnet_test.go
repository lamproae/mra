package telnet_test

import (
	"fmt"
	"github.com/lamproae/mra/telnet"
	"github.com/lamproae/mra/util"
	"strings"
	"testing"
)

var host = "10.71.20.198:23"

func TestTelnetConnection(t *testing.T) {
	c, err := telnet.NewClient(host)
	if err != nil {
		fmt.Println("Error happend when connect to DUT: ", err.Error())
	}
	c.SetUnixWriteMode(true)
	data, err := c.ReadUntil("login")
	if err != nil {
		fmt.Println("Error happend when get login: ", err.Error())
		return
	}
	fmt.Println(string(data))
}

func TestLoginFunction(t *testing.T) {
	c, err := telnet.NewClient("10.71.20.198:23")
	if err != nil {
		fmt.Println("Error happend when connect to DUT: ", err.Error())
	}
	c.SetUnixWriteMode(true)
	data, err := c.ReadUntil("login")
	if err != nil {
		fmt.Println("Error happend when get login: ", err.Error())
		return
	}
	fmt.Println(string(data))
	c.Write([]byte("admin" + "\n"))
	data, err = c.ReadUntil("Password: ")
	if err != nil {
		fmt.Println("Error happend when get login prompt: ", err.Error())
		return
	}
	fmt.Println(string(data))
	c.Write([]byte("\n"))
	data, err = c.ReadUntil(">")
	if err != nil {
		fmt.Println("Error happend when login: ", err.Error())
		return
	}
	fmt.Println(string(data))

	c.Write([]byte("enable" + "\n"))
	data, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when goto enable mode: ", err.Error())
		return
	}
	fmt.Println(string(data))

	c.Write([]byte("terminal length 0" + "\n"))
	data, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when SetTerminalLength: ", err.Error())
		return
	}
	fmt.Println(string(data))
	c.Write([]byte("show running-config" + "\n"))
	data, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when do show running: ", err.Error())
		return
	}
	fmt.Println(string(data))
}

func TestShowAll(t *testing.T) {
	c, err := telnet.NewClient("10.71.20.198:23")
	if err != nil {
		fmt.Println("Error happend when connect to DUT: ", err.Error())
	}
	c.SetUnixWriteMode(true)
	data, err := c.ReadUntil("login")
	if err != nil {
		fmt.Println("Error happend when get login: ", err.Error())
		return
	}
	fmt.Println(string(data))
	c.Write([]byte("admin" + "\n"))
	data, err = c.ReadUntil("Password: ")
	if err != nil {
		fmt.Println("Error happend when get login prompt: ", err.Error())
		return
	}
	fmt.Println(string(data))
	c.Write([]byte("\n"))
	data, err = c.ReadUntil(">")
	if err != nil {
		fmt.Println("Error happend when login: ", err.Error())
		return
	}
	fmt.Println(string(data))

	c.Write([]byte("enable" + "\n"))
	data, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when goto enable mode: ", err.Error())
		return
	}
	fmt.Println(string(data))

	c.Write([]byte("terminal length 0" + "\n"))
	data, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when SetTerminalLength: ", err.Error())
		return
	}
	fmt.Println(string(data))
	c.WriteLine("configure terminal")
	data, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when do show running: ", err.Error())
		return
	}
	fmt.Println(string(data))

	c.WriteLine("@show all")
	data, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when do show running: ", err.Error())
		return
	}
	fmt.Println(string(data))
	util.SaveToFile("commandList.txt", data)

	result := strings.Split(string(data), "\n")
	fmt.Println(result)
	for _, c := range result {
		c = strings.TrimSpace(c)
		if strings.HasPrefix(c, "show") ||
			strings.HasPrefix(c, "help") ||
			strings.HasPrefix(c, "no ") ||
			//strings.HasPrefix(c, "write") ||
			len(c) == 0 {
			continue
		}
		fmt.Println(c)
	}
}
