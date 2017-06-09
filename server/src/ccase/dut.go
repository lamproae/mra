package ccase

type DUT struct {
	Name     string
	Device   string
	UserName string
	Password string
}

func (d *DUT) String() string {
	return "DUT{" + d.Device + ":" + d.Name + ":" + d.UserName + ":" + d.Password + "}"
}
