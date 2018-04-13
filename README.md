# Python car for the [Bullettrain](https://github.com/bullettrain-sh/bullettrain-go-core) shell prompt

## Features:

- Displaying only when needed
- Python version display
- Virtualenv display
* Support for [Pyenv](https://github.com/pyenv/pyenv) and it's
  [virtualenv](https://github.com/pyenv/pyenv-virtualenv) plugin

**Callword**: `python`

**Template variables**:

* `.VersionIcon`: the Python version icon
* `.Version`: the Python version text
* `.VenvIcon`: the Python virtualenv icon
* `.Venv`: the Python virtualenv text

**Template colours**:

* `c`: the Python version colour
* `cs`: the Python version symbol's colour
* `cvs`: the Python virtualenv symbol's colour


## Car options

| Environment variable                           | Description                                                    | Default value                                                                                                               |
|:-----------------------------------------------|:---------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------|
| BULLETTRAIN_CAR_PYTHON_SHOW                    | Whether the car needs to be shown all the time.                | false                                                                                                                       |
| BULLETTRAIN_CAR_PYTHON_TEMPLATE                | The car's template.                                            | `{{.VersionIcon \| printf "%s " \| cs}}{{.Version \| printf "%s " \| c}}{{.VenvIcon \| printf "%s " \| cvs}}{{.Venv \| c}}` |
| BULLETTRAIN_CAR_PYTHON_PAINT                   | Colour override for the car't paint.                           | black:220                                                                                                                   |
| BULLETTRAIN_CAR_PYTHON_SYMBOL_ICON             | Icon displayed on the car.                                     | ``                                                                                                                         |
| BULLETTRAIN_CAR_PYTHON_SYMBOL_PAINT            | Colour override for the car's symbol.                          | 32:220                                                                                                                      |
| BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SYMBOL_ICON  | Icon displayed on the car.                                     | `🐍`                                                                                                                        |
| BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SYMBOL_PAINT | Colour override for the car's symbol.                          | 32:220                                                                                                                      |
| BULLETTRAIN_CAR_PYTHON_SEPARATOR_PAINT         | Colour override for the car's right hand side separator paint. | Using default painting algorythm.                                                                                           |
| BULLETTRAIN_CAR_PYTHON_SEPARATOR_SYMBOL        | Override the car's right hand side separator symbol.           | Using global symbol.                                                                                                        |
| BULLETTRAIN_CAR_PYTHON_SEPARATOR_TEMPLATE      | Defines the car separator's template.                          | Using global template.                                                                                                      |

# Contribute

Even reporting your use case will greatly help us to figure out/improve
this product, so feel free to reach out in the Issues section.
