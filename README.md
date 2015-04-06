# WGT Playground

## Extract/Load/Transform

Clean up [WGT's confirmed artists list](http://www.wave-gotik-treffen.de/english/bands.php).

``` sh
./clean-artists artists-2015-04-05 | head -n 1 | dotenv ./pull-artist-info
```

## View

``` sh
go get
go build
./wgt
```

Open [localhost:8080](http://localhost:8080).
