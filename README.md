# Python car for the [Bullettrain](https://github.com/bullettrain-sh/bullettrain-go-core) shell prompt

## Features:

- Displaying only when needed
- Python version display
- Virtualenv display
- Support for [Pyenv](https://github.com/pyenv/pyenv) and it's [virtualenv](https://github.com/pyenv/pyenv-virtualenv) plugin

## Car options

| Environment variable                           | Description                                           | Default value |
| :---                                           | :---                                                  | :---          |
| BULLETTRAIN_CAR_PYTHON_SHOW                    | Whether the car needs to be shown all the time.       | false         |
| BULLETTRAIN_CAR_PYTHON_VERSION_SHOW            | Whether to display the available versions of pythons. | true          |
| BULLETTRAIN_CAR_PYTHON_PAINT                   | Colour override for the car't paint.                  | black:220     |
| BULLETTRAIN_CAR_PYTHON_SYMBOL_ICON             | Icon displayed on the car.                            | ``           |
| BULLETTRAIN_CAR_PYTHON_SYMBOL_PAINT            | Colour override for the car's symbol.                 | 32:220        |
| BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SHOW         | Whether to display virtualenv information.            | true          |
| BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SYMBOL_ICON  | Icon displayed on the car.                            | `🐍`          |
| BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SYMBOL_PAINT | Colour override for the car's symbol.                 | 32:220        |
