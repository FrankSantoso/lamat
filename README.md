# lamat
Small CLI to Geocode / Reverse Geocode an Address / Coordinate. 


### Usage
Get your API keys for nominatim openstreetmap & Google Geocoding

Copy the provided `config.toml.sample` to `config.toml`

Insert your keys to the config.toml.

### Reverse Geocoding
```
    lamat rev -- <longitude> <latitude>
``` 

### Forward Geocoding

```
    lamat find "<YOUR ADDRESS STRING HERE>"
```

Note that this tool currently doesn't support partial matches yet.
