# `/examples`

In order to use use the examples in this directory you will need to provide a `config.json`. Copy the supplied `config_template.json`, rename it and add your Exchange creds and SSL cert/key paths.

The examples can be compiled using `go build`. Each example takes a single argument of `-config` where you can supply the path to your `config.json`.


```
.\stream-sync.exe -config ..\config.json
```