# Simple echo telegram bot

## Prerequisites

This example uses webhooks to interact with Telegram API.
It requires active domain with webserver serving this domain using TLS protocol.
I set it up it using personal subdomain and `Nginx` webserver.
TLS support is enabled using Let's Encrypt certificate.
Webserver proxies all requests from public IP
to localhost where telegram bot is listens to incoming requests.
It is actully deployed and can be tested by searching for `nal_test_bot` in Telegram
and chatting to them.

## Initialization of module dependecies

### Follow 12factor app rules and explicitly manage dependecies

* go mod init
* go get -u go.uber.org/zap@v1

## Running app

* run `go run main.go types.go`

## Building app

* run `make` (Works for Mac OSX/Linux)
