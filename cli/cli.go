package cli

import (
	"errors"
	"github.com/lamproae/mra/command"
	"github.com/lamproae/mra/config"
	"github.com/lamproae/mra/telnet"
	"log"
)

type Cli struct {
	client *telnet.Client
	conf   *config.Config
	mode   string
}

func (c *Cli) RunCommand(cmd *command.Command) (result []byte, err error) {
	if cmd.RequiredMode() != c.mode {
		return nil, errors.New("Cannot run : " + cmd.CMD() + " under " + c.mode + " mode!")
	}

	c.WriteLine(cmd.CMD())
	data, err = c.ReadUntil(c.conf.Prompt())
	if err != nil {
		fmt.Println("Error happend when get login prompt: ", err.Error())
		return nil, err
	}
	return data, nil
}

func NewCli(conf *config.Config) (c *Cli, err error) {
	tc, err := telnet.NewClient(conf.IP + ":" + conf.Port)
	if err != nil {
		log.Println("error happend when connect to: ", conf.IP, " with: ", err.Error())
		return nil, errors.New("Cannot connect to host")
	}

	return &Cli{
		client: tc,
		conf:   conf,
	}, nil
}

func (c *Cli) CurrentMode() string {
	return c.mode
}

func (c *Cli) Init() error {
	err := c.login()
	if err != nil {
		log.Println("Error happened when login: ", err.Error())
		return err
	}

	c.WriteLine("enable")
	data, err = c.ReadUntil(c.conf.Prompt)
	if err != nil {
		fmt.Println("Error happend when goto enable mode: ", err.Error())
		return err
	}
	fmt.Println(string(data))

	c.WriteLine("terminal length 0")
	data, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when SetTerminalLength: ", err.Error())
		return err
	}
	fmt.Println(string(data))

	return nil
}

func (c *Cli) login() error {
	c.client.SetUnixWriteMode(true)
	data, err := c.ReadUntil(c.conf.LoginPrompt)
	if err != nil {
		fmt.Println("Error happend when get login: ", err.Error())
		return err
	}
	fmt.Println(string(data))
	c.WriteLine(c.conf.UserName)
	data, err = c.ReadUntil(c.conf.PasswordPrompt)
	if err != nil {
		fmt.Println("Error happend when get login prompt: ", err.Error())
		return err
	}
	fmt.Println(string(data))
	c.WriteLine(c.conf.Password)
	data, err = c.ReadUntil(c.conf.EnablePrompt)
	if err != nil {
		fmt.Println("Error happend when login: ", err.Error())
		return err
	}
	fmt.Println(string(data))

	return nil
}
