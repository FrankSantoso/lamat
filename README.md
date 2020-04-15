# lamat
Small CLI to Geocode / Reverse Geocode an Address / Coordinate. 


### Usage
Get your API keys for nominatim openstreetmap & Google Geocoding

Copy the provided `config.toml.sample` to `config.toml`

Insert your keys to the config.toml.

```
Small utility to find geocode / reverse geocode

Usage:
  lamat [command]

Available Commands:
  find        Try finding address of inputted coordinates / geocode
  help        Help about any command
  rev         Try finding geocode information from specified input address

Flags:
  -c, --cfg string   config files containing nominatim & google api keys (default "./config")
  -h, --help         help for lamat
  -j, --json         outputs json instead of strings

Use "lamat [command] --help" for more information about a command.
```

### Reverse Geocoding
```
    lamat rev -- <lat> <long>
``` 

### Forward Geocoding

```
    lamat find "<YOUR ADDRESS STRING HERE>"
```

Note that this tool currently doesn't support partial matches yet.
