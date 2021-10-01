# gofair

Lightweight golang wrapper for [Betfair API-NG](http://docs.developer.betfair.com/docs/display/1smk3cen4v3lu3yomq5qye0ni)

# setup

In order to connect to the Betfair API you will need an App Key, SSL Certificates and a username/password.

### App Key
Follow [these](https://docs.developer.betfair.com/display/1smk3cen4v3lu3yomq5qye0ni/Application+Keys) instructions to get your app key, you can either use a delayed or live key.

### SSL certificates
Follow [these](https://docs.developer.betfair.com/display/1smk3cen4v3lu3yomq5qye0ni/Non-Interactive+%28bot%29+login) instructions to set up your SSL certificates. Save your .ctr and .key files to a local directory. The default directory where the library is looking for the keys is '/certs' but you can specify any other directory.

# examples

A set of examples on how to use this library are available in the `examples` directory. You will need to supply a valid `config.json` in order to interact with the Exchange see `examples/config_template.json` for an example configuration.

# streaming

gofair makes extensive use of [channels](https://tour.golang.org/concurrency/2) for handling data returned by the Stream API. This provides the user with a reasonable amount of flexibility as you have both synchronous and asynchronous options (see `examples/stream-sync` and `examples/stream-async`).

# use

```golang
config := &gofair.Config{
		"username",
		"password",
		"appKey",
		"/certs/client-2048.crt",
		"/certs/client-2048.key",
		"",
}

trading, err := gofair.NewClient(config)
if err != nil {
    panic(err)
}


fmt.Println(trading.Login())
fmt.Println(trading.KeepAlive())
fmt.Println(trading.Logout())

filter := new(gofair.MarketFilter)
event_types, err := trading.Betting.ListEventTypes(filter)

fmt.Println(event_types)
```
