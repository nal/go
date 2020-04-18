# Simple echo telegram bot

## Prerequisites

This example uses webhooks to interact with Telegram API.
It requires active domain with webserver serving this domain using TLS protocol.
I set up it using personal subdomain <https://demo.nal.kiev.ua/tgbot>.
TLS support is enabled using Let's Encrypt certificate.
Nginx webserver proxies all requests from <https://demo.nal.kiev.ua/tgbot>
to localhost where telegram bot is listens to incoming requests.

## Initialization of module dependecies

* go mod init
* go get -u go.uber.org/zap@v1

## Building module

* run `make`
