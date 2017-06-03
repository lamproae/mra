package assert

import (
	"github.com/lamproae/mra/command"
	"log"
	"regexp"
)

type Assert struct {
	Raw        string
	Expected   string
	UnExpected string
	CMD        *command.Command
	Done       bool
}

func (a *Assert) Do() bool {
	a.Done = true

	var re *regexp.Regexp
	if a.Expected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		if len(match) == 0 {
			return false
		} else {
			return true
		}
	} else if a.UnExpected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		if len(match) == 0 {
			return true
		} else {
			return false
		}
	}
	log.Println("Invlaid assertion, Both expcted and unexpected are empty!")
	return false
}

func (a *Assert) String() string {
	var re *regexp.Regexp
	if a.Expected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		log.Println("LEN: +++++++++++++++++++++++: ", len(match))
		if len(match) != 2 {
			return "FAILED: Command: " + a.CMD.CMD + " EXPECTED: " + a.Expected + "GET: " + a.Raw
		} else {
			return "SUCCESS: Command: " + a.CMD.CMD + " EXPECTED: " + a.Expected
		}
	} else if a.UnExpected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		log.Println("LEN: +++++++++++++++++++++++: ", len(match))
		if len(match) != 2 {
			return "SUCCESS: Command: " + a.CMD.CMD + " UnEXPECTED: " + a.UnExpected
		} else {
			return "FAILED: Command: " + a.CMD.CMD + " UnEXPECTED: " + a.UnExpected + "GET: " + a.Raw
		}
	}
	log.Println("Invlaid assertion, Both expcted and unexpected are empty!")
	return "FAILED: Command: " + a.CMD.CMD + " EXPECTED: " + a.Expected + "GET: " + a.Raw

}
