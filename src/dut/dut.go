package dut

//DUT should be and interface
type DUT struct {
	name string
	L3   bool //We need to have a feature list. For run each case
	L2   bool
}

func (d *DUT) CreateVlan(id int) error {

}

func (d *DUT) DestroyVlan(id int) error {

}

func (d *DUT) DestroyAllVlan() error {

}

func (d *DUT) CreateVlanInterface(id int, ip string) error {

}

func (d *DUT) DestroyVlanInterface(id) {

}

func (d *DUT) AddIPAddress(ifname, ip string) error {

}

func (d *DUT) DelIPAddress(ifname, ip string) error {

}

func (d *DUT) AddSecondaryIPAddress(ifname, ip string) error {

}

func (d *DUT) DelSecondaryIPAddress(ifname, ip string) error {

}

func (d *DUT) DelAllIPAddress(ifname string) error {

}

func (d *DUT) CreateOSPFInstance(id, tag string) error {

}

func (d *DUT) DestroyOSPFInstance(id string) error {

}
