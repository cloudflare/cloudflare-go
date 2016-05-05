# cloudflare-go
[![GoDoc](https://godoc.org/github.com/cloudflare/cloudflare-go?status.svg)](https://godoc.org/github.com/cloudflare/cloudflare-go)

> **Note**: This library is under active development as we expand it to cover our (expanding!) API.
Consider the public API of this package a little unstable as we work towards a v1.0.

A Go library for interacting with [CloudFlare's API v4](https://api.cloudflare.com/). This library
allows you to:

* Manage and automate changes to your DNS records within CloudFlare
* Manage and automate changes to your zones (domains) on CloudFlare, including adding new zones to
  your account
* List and modify the status of WAF (Web Application Firewall) rules for your zones
* Fetch CloudFlare's IP ranges for automating your firewall whitelisting

A command-line client, [flarectl](cmd/flarectl), is also available as part of this project.

## Features

The current feature list includes:

- [x] DNS Records
- [x] Zones
- [x] Web Application Firewall (WAF)
- [x] CloudFlare IPs
- [x] User Administration (partial)
- [ ] Organization Administration
- [ ] [Railgun](https://www.cloudflare.com/railgun/) administration
- [ ] [Keyless SSL](https://blog.cloudflare.com/keyless-ssl-the-nitty-gritty-technical-details/)
- [ ] [Origin CA](https://blog.cloudflare.com/universal-ssl-encryption-all-the-way-to-the-origin-for-free/)

Pull Requests are welcome, but please open an issue (or comment in an existing issue) to discuss any
non-trivial changes before submitting code.

## Installation

You need a working Go environment.

```
go get github.com/cloudflare/cloudflare-go
```

## Getting Started

```
package main

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
)

var api *cloudflare.API

func main() {
	// Construct a new API object
	api = cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))

	// Fetch the list of zones on the account
	zones, err := api.ListZones()
	if err != nil {
		fmt.Println(err)
	}
	// Print the zone names
	for _, z := range zones {
		fmt.Println(z.Name)
	}
}
```

Also refer to the [API documentation](https://godoc.org/github.com/cloudflare/cloudflare-go) for how
to use this package in-depth.

# License

BSD licensed. See the [LICENSE](LICENSE) file for details.
