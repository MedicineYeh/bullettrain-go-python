package carVirtualenv

import (
    "bytes"
    "log"
    "os"
    "text/template"
    "path"

    "github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
    carPaint              = "15:214"
    virtualenvSymbolIcon  = "🐍"
    virtualenvSymbolPaint = "32:214"
    carTemplate           = `{{.VenvIcon | printf "%s " | cs}}{{.Venv | c}}`
)

// Car for Python and virtualenv
type Car struct {
    paint string
    Pwd string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
    if c.paint = os.Getenv("BULLETTRAIN_CAR_VIRTUALENV_PAINT"); c.paint == "" {
        c.paint = carPaint
    }

    return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
    if e := os.Getenv("BULLETTRAIN_CAR_VIRTUALENV_SHOW"); e == "false" {
        return false
    }

    // Show when "VIRTUAL_ENV" exist in environment variables
    if e := os.Getenv("VIRTUAL_ENV"); e != "" {
        return true
    }

    return false
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
//
func (c *Car) Render(out chan<- string) {
    defer close(out) // Always close the channel!

    var vs string
    if vs = os.Getenv("BULLETTRAIN_CAR_VIRTUALENV_SYMBOL_ICON"); vs == "" {
        vs = virtualenvSymbolIcon
    }

    var vsp string
    if vsp = os.Getenv("BULLETTRAIN_CAR_VIRTUALENV_SYMBOL_PAINT"); vsp == "" {
        vsp = virtualenvSymbolPaint
    }

    var s string
    if s = os.Getenv("BULLETTRAIN_CAR_VIRTUALENV_TEMPLATE"); s == "" {
        s = carTemplate
    }

    funcMap := template.FuncMap{
        // Pipeline functions for colouring.
        "c":   func(t string) string { return ansi.Color(t, c.GetPaint()) },
        "cs":  func(t string) string { return ansi.Color(t, vsp) },
    }

    tpl := template.Must(template.New("python").Funcs(funcMap).Parse(s))
    data := struct {
        VenvIcon    string
        Venv        string
    }{
        VenvIcon: virtualenvSymbolIcon,
        Venv:     path.Base(os.Getenv("VIRTUAL_ENV")),
    }
    fromTpl := new(bytes.Buffer)
    err := tpl.Execute(fromTpl, data)
    if err != nil {
        log.Fatalf("Can't generate the python template: %s", err.Error())
    }

    out <- fromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
    return os.Getenv("BULLETTRAIN_CAR_VIRTUALENV_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
    return os.Getenv("BULLETTRAIN_CAR_VIRTUALENV_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
    return os.Getenv("BULLETTRAIN_CAR_VIRTUALENV_SEPARATOR_TEMPLATE")
}
