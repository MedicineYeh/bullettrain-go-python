package carPython

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mgutz/ansi"
)

const carPaint = "black:220"
const pythonSymbolPaint = "32:220"
const pythonSymbolIcon = ""
const virtualenvSymbolIcon = "🐍"
const virtualenvSymbolPaint = "32:220"

// Car for Python and virtualenv
type Car struct {
	paint string
}

func paintedPythonSymbol() string {
	var pythonSymbol string
	if pythonSymbol = os.Getenv("BULLETTRAIN_CAR_PYTHON_ICON"); pythonSymbol == "" {
		pythonSymbol = pythonSymbolIcon
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_PYTHON_ICON_PAINT"); symbolPaint == "" {
		symbolPaint = pythonSymbolPaint
	}

	return ansi.Color(pythonSymbol, symbolPaint)
}

func paintedVirtualenvSymbol() string {
	var virtualenvSymbol string
	if virtualenvSymbol = os.Getenv("BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SYMBOL_ICON"); virtualenvSymbol == "" {
		virtualenvSymbol = virtualenvSymbolIcon
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SYMBOL_PAINT"); symbolPaint == "" {
		symbolPaint = virtualenvSymbolPaint
	}

	return ansi.Color(virtualenvSymbol, symbolPaint)
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_PYTHON_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	if e := os.Getenv("BULLETTRAIN_CAR_PYTHON_SHOW"); e == "true" {
		return true
	}

	cmd := exec.Command("pwd", "-P")
	pwd, err := cmd.Output()
	if err != nil {
		return false
	}
	d := strings.Trim(string(pwd), "\n")

	// Show when .py files exist in current directory
	pyPattern := fmt.Sprintf("%s%s*.py", d, string(os.PathSeparator))
	pyFiles, err := filepath.Glob(pyPattern)
	if pyFiles != nil {
		return true
	}

	// Show when .python-version file exist in current directory
	versionFiles, _ := filepath.Glob(fmt.Sprintf("%s%s.python-version",
		d, string(os.PathSeparator)))
	if versionFiles != nil {
		return true
	}

	return false
}

// getPythonVersion gets the available version number for a python executable
//
// Use it to check if python2 responds, python3 responds or only python does.
func getPythonVersion(pythonExecutable string) string {
	cmdPython := exec.Command(pythonExecutable, "--version")
	resultPython, errPython := cmdPython.CombinedOutput()
	if errPython == nil {
		return strings.TrimSpace(strings.TrimLeft(
			string(resultPython), "Python "))
	} else {
		return ""
	}
}

func collectPythonVersions(o *bytes.Buffer, carPaint func(string) string) {
	pythonVersions := make([]string, 0)
	var p string
	if p = getPythonVersion("python2"); p != "" {
		pythonVersions = append(pythonVersions, p)
	}
	if p = getPythonVersion("python3"); p != "" {
		pythonVersions = append(pythonVersions, p)
	}
	if len(pythonVersions) == 0 {
		if p = getPythonVersion("python"); p != "" {
			pythonVersions = append(pythonVersions, p)
		}
	}

	if len(pythonVersions) > 0 {
		o.WriteString(paintedPythonSymbol())
		o.WriteString(carPaint(
			strings.TrimSpace(strings.Join(pythonVersions, " "))))
	}
}

func collectPythonVirtualenvs(o *bytes.Buffer, carPaint func(string) string) {
	cmdPyenv := exec.Command("pyenv", "version")
	cmdOut, errPyenv := cmdPyenv.Output()
	if errPyenv == nil {
		re := regexp.MustCompile(`(?m)^([a-zA-Z0-9_\-]+)`)
		versions := re.FindAllStringSubmatch(string(cmdOut), -1)
		var versionsInfo string
		for _, i := range versions {
			versionsInfo = fmt.Sprintf("%s %s", versionsInfo, i[1])
		}

		o.WriteString(carPaint(" "))
		o.WriteString(paintedVirtualenvSymbol())
		o.WriteString(carPaint(versionsInfo))
	}
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
	// Output collector buffer.
	var o bytes.Buffer

	if pvers := os.Getenv("BULLETTRAIN_CAR_PYTHON_VERSION_SHOW"); pvers != "false" {
		collectPythonVersions(&o, carPaint)
	}

	if pvenvs := os.Getenv("BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SHOW"); pvenvs != "false" {
		collectPythonVirtualenvs(&o, carPaint)
	}

	out <- o.String()
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
