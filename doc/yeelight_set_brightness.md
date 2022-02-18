## yeelight set brightness

Change brightness

### Synopsis

Change brightness

Arguments:
  [host] - IP address or hostname of a device. If a port is missed then by default will be used 55443.
  [bright] - Brightness value [0 - 100]

```
yeelight set brightness [host] [bright] [flags]
```

### Examples

```
  yeelight set brightness 192.168.1.79:55443 10
```

### Options

```
  -b, --background          apply for background
  -d, --duration duration   Effect duration (default 30ms)
      --effect string       Effect name (sudden, smooth) (default "smooth")
  -h, --help                help for brightness
```

### SEE ALSO

* [yeelight set](yeelight_set.md)	 - Specify the light parameter

###### Auto generated by spf13/cobra on 18-Feb-2022