# Python car for the [Bullettrain](https://github.com/bullettrain-sh/bullettrain-go-core) shell prompt

## Features:

- Displaying only when used
- Virtualenv display

**Callword**: `python`

**Template variables**:

* `.VenvIcon`: the Python virtualenv icon
* `.Venv`: the Python virtualenv text

**Template colours**:

* `c`: the Python version colour
* `cs`: the Python virtualenv symbol's colour


## Car options

| Environment variable                               | Description                                                    | Default value                                                                                                               |
|:---------------------------------------------------|:---------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------|
| BULLETTRAIN_CAR_VIRTUALENV_SHOW                    | Whether the car needs to be shown all the time.                | true                                                                                                                        |
| BULLETTRAIN_CAR_VIRTUALENV_TEMPLATE                | The car's template.                                            | `{{.VersionIcon \| printf "%s " \| cs}}{{.Version \| printf "%s " \| c}}{{.VenvIcon \| printf "%s " \| cvs}}{{.Venv \| c}}` |
| BULLETTRAIN_CAR_VIRTUALENV_PAINT                   | Colour override for the car't paint.                           | 15:214                                                                                                                      |
| BULLETTRAIN_CAR_VIRTUALENV_SYMBOL_ICON             | Icon displayed on the car.                                     | `🐍`                                                                                                                        |
| BULLETTRAIN_CAR_VIRTUALENV_SYMBOL_PAINT            | Colour override for the car's symbol.                          | 32:214                                                                                                                      |
| BULLETTRAIN_CAR_VIRTUALENV_SEPARATOR_PAINT         | Colour override for the car's right hand side separator paint. | Using default painting algorythm.                                                                                           |
| BULLETTRAIN_CAR_VIRTUALENV_SEPARATOR_SYMBOL        | Override the car's right hand side separator symbol.           | Using global symbol.                                                                                                        |
| BULLETTRAIN_CAR_VIRTUALENV_SEPARATOR_TEMPLATE      | Defines the car separator's template.                          | Using global template.                                                                                                      |

# Contribute

Even reporting your use case will greatly help us to figure out/improve
this product, so feel free to reach out in the Issues section.
