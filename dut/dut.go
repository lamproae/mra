package dut

import (
	"errors"
	"github.com/lamproae/mra/assert"
	"github.com/lamproae/mra/cli"
	"github.com/lamproae/mra/command"
	"github.com/lamproae/mra/config"
	"github.com/lamproae/mra/routine"
	"github.com/lamproae/mra/task"
	"log"
)

//DUT should be and interface
type DUT struct {
	name string
	cli  *cli.Cli
	L3   bool //We need to have a feature list. For run each case
	L2   bool
}

func NewDUT(conf *config.Config) (*DUT, error) {
	if conf == nil {
		return nil, errors.New("Invalid config")
	}

	c, err := cli.NewCli(conf)
	if err != nil {
		return nil, errors.New("Cannot create new DUT with: " + err.Error())
	}

	err = c.Init()
	if err != nil {
		return nil, errors.New("Cannot create new DUT with: " + err.Error())
	}

	return &DUT{
		name: conf.DeviceName,
		cli:  c,
	}, nil
}

func (d *DUT) RunTask(t *task.Task) error {
	if err := t.CheckPreCondition(); err != nil {
		log.Println("PreCondition check failed for task: ", t.Name, " with: ", err.Error())
		return errors.New("PreCondition check failed!: " + err.Error())
	}

	for _, r := range t.Routines {
		err := d.RunRoutine(r)
		if r != nil {
			log.Println(err)
		}
		return errors.New("Cannot run task: " + t.Name + " with: " + err.Error())
	}

	if err := t.CheckPostCondition(); err != nil {
		log.Println("PostCondition check failed for task: ", t.Name, " with: ", err.Error())
		return errors.New("PostCondition check failed!: " + err.Error())
	}

	return nil
}

func (d *DUT) RunRoutine(r *routine.Routine) error {
	log.Println("Running Routine: ", r.Name)
	for _, c := range r.CMD {
		_, err := d.RunCommand(c)
		if err != nil {
			log.Println("Error happend when run routine: ", r.Name, " with: ", err.Error())
			return errors.New("Cannot run routine: " + r.Name + " with: " + err.Error())
		}
	}

	for _, a := range r.Assert {
		success := d.Assert(a)
		if !success {
			return errors.New("Assertion failed for routine: " + r.Name + " Message: " + a.String())
		}
	}

	return nil
}

func (d *DUT) Assert(a *assert.Assert) bool {
	result, err := d.RunCommand(a.CMD)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	a.Raw = result

	return a.Do()
}

func (d *DUT) RunCommand(cmd *command.Command) (string, error) {
	data, err := d.cli.RunCommand(cmd)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (d *DUT) CreateVlan(id int) error {
	return nil

}

func (d *DUT) DestroyVlan(id int) error {

	return nil
}

func (d *DUT) DestroyAllVlan() error {

	return nil
}

func (d *DUT) CreateVlanInterface(id int, ip string) error {

	return nil
}

func (d *DUT) DestroyVlanInterface(id int) error {

	return nil
}

func (d *DUT) AddIPAddress(ifname, ip string) error {

	return nil
}

func (d *DUT) DelIPAddress(ifname, ip string) error {

	return nil
}

func (d *DUT) AddSecondaryIPAddress(ifname, ip string) error {

	return nil
}

func (d *DUT) DelSecondaryIPAddress(ifname, ip string) error {

	return nil
}

func (d *DUT) DelAllIPAddress(ifname string) error {

	return nil
}

func (d *DUT) CreateOSPFInstance(id, tag string) error {

	return nil
}

func (d *DUT) DestroyOSPFInstance(id string) error {

	return nil
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
