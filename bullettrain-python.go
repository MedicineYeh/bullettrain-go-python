package bullettrain_go_python

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type PythonSegment struct {
	fg, bg color.Attribute
}

func (p *PythonSegment) SetFg(fg color.Attribute) {
	p.fg = fg
}

func (p *PythonSegment) SetBg(bg color.Attribute) {
	p.bg = bg
}

//// Builds the version string of the currently available Python interpreter(s).
//// Python version managers can expose multiple versions too.
//// Version managers analyzed first, then system Pythons.
//// Empty string is returned when no interpreter could be reached.
func (p *PythonSegment) Render(ch chan<- string) {
	const python_symbol string = "🐍"
	col := color.New(p.fg, p.bg)

	// ______
	// | ___ \
	// | |_/ /   _  ___ _ ____   __
	// |  __/ | | |/ _ \ '_ \ \ / /
	// | |  | |_| |  __/ | | \ V /
	// \_|   \__, |\___|_| |_|\_/
	//        __/ |
	//       |___/

	pyenvCmd := exec.Command("pyenv", "version")
	pyenvOut, err := pyenvCmd.Output()
	if err == nil {
		re := regexp.MustCompile(`(?m)^([a-zA-Z0-9_\-]+)`)
		versions := re.FindAllStringSubmatch(string(pyenvOut), -1)
		var versions_info string
		for _, i := range versions {
			versions_info = fmt.Sprintf("%s %s", versions_info, i[1])
		}

		ch <- col.Sprintf(" %s%s ", python_symbol, versions_info)
		close(ch)
		return
	}

	// ______      _   _
	// | ___ \    | | | |
	// | |_/ /   _| |_| |__   ___  _ __
	// |  __/ | | | __| '_ \ / _ \| '_ \
	// | |  | |_| | |_| | | | (_) | | | |
	// \_|   \__, |\__|_| |_|\___/|_| |_|
	//        __/ |
	//       |___/

	// TODO python 2 and python 3 version info!

	pythonCmd := exec.Command("python", "--version")
	var stderr bytes.Buffer
	pythonCmd.Stderr = &stderr
	pyErr := pythonCmd.Run()
	if pyErr == nil {
		ch <- col.Sprintf(" %s %s ",
			python_symbol, strings.Trim(stderr.String(), "\n"))
		close(ch)
	}

	// Always close the channel!
	close(ch)
	return
}
