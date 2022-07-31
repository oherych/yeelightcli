[![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/banner-direct-single.svg)](https://stand-with-ukraine.pp.ua)

`yeelightcli` is a command-line tool for managing Yeelight devices. 
Right now it supports 95% of all features presented in the official documentation.

## Installation

### Go
`go install github.com/oherych/yeelightcli`

### Brew
```shell
  brew tap oherych/yeelightcli
  brew install yeelightcli
```



## Usage
`yeelightcli help`

```
Usage:
  yeelightcli [command]

Available Commands:
  discovery   Discovers all the devices in the local network
  power       Switches the light on/off
  set         Specify the light parameter
  adjust      Adjust the light parameter
  default     Declare current light settings as default
  cron        Schedule light turn on/off
  music       Start or stop music mode on a device
  name        Change device name
  read        Fetch current device settings
  help        Help about any command
  
Flags:
  -h, --help   help for main

Use "main [command] --help" for more information about a command.

```

### TODO
- SetAdjust
- SetBackgoundAdjust
- StartColorFlow
- StartBackgroundColorFlow
- StopColorFlow
- StopBackgroundColorFlow
- SetMusic