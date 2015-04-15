# WGT Playground

## Extract

Visit [WGT's confirmed artists list](http://www.wave-gotik-treffen.de/english/bands.php). Copy the big blob of artists and save to `artists-2015-04-05` (or something appropriate).

## Load/Transform

``` sh
export ECHO_NEST_API_KEY=
bin/clean-artists artists-2015-04-15 | bin/generate-data > data.json
```

## View

``` sh
go get
go build
./wgt
```

Open [localhost:8080](http://localhost:8080).
