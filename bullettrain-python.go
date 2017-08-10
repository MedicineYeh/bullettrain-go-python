package carPython

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/mgutz/ansi"
)

type Car struct {
	paint string
}

func paintedSymbol() string {
	var symbol string
	if symbol = os.Getenv("BULLETTRAIN_CAR_PYTHON_ICON"); symbol == "" {
		symbol = " "
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_PYTHON_ICON_PAINT"); symbolPaint == "" {
		symbolPaint = "32:220"
	}

	return ansi.Color(symbol, symbolPaint)
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_PYTHON_PAINT"); c.paint == "" {
		c.paint = "black:220"
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := true

	// TODO check for
	// 	 *.py
	//   .python-version

	if e := os.Getenv("BULLETTRAIN_CAR_PYTHON_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
//
// Car version managers can expose multiple Python versions too.
// Python version managers analyzed first, then system Pythons are looked at.
// Empty string is returned when no interpreter could be reached.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!
	carPaint := ansi.ColorFunc(c.GetPaint())

	// ______
	// | ___ \
	// | |_/ /   _  ___ _ ____   __
	// |  __/ | | |/ _ \ '_ \ \ / /
	// | |  | |_| |  __/ | | \ V /
	// \_|   \__, |\___|_| |_|\_/
	//        __/ |
	//       |___/

	cmdPyenv := exec.Command("pyenv", "version")
	cmdOut, errPyenv := cmdPyenv.Output()
	if errPyenv == nil {
		re := regexp.MustCompile(`(?m)^([a-zA-Z0-9_\-]+)`)
		versions := re.FindAllStringSubmatch(string(cmdOut), -1)
		var versionsInfo string
		for _, i := range versions {
			versionsInfo = fmt.Sprintf("%s %s", versionsInfo, i[1])
		}

		out <- fmt.Sprintf("%s%s%s%s",
			carPaint(" "),
			paintedSymbol(),
			carPaint(versionsInfo),
			carPaint(" "))

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

	//cmdPython := exec.Command("python", "--version")
	//var stderr bytes.Buffer
	//cmdPython.Stderr = &stderr
	//errPython := cmdPython.Run()
	//if errPython == nil {
	//	out <- ansi.Color(
	//		fmt.Sprintf(" %s %s ",
	//			symbol, strings.Trim(stderr.String(), "\n")),
	//		c.GetPaint())
	//}
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_PYTHON_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_PYTHON_SEPARATOR_SYMBOL")
}
