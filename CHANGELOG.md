## 0.47.0 (Unreleased)

## 0.46.0 (3rd August, 2022)

NOTES:

* docs: add release notes ([#1001](https://github.com/cloudflare/cloudflare-go/issues/1001))

ENHANCEMENTS:

* filter: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1004](https://github.com/cloudflare/cloudflare-go/issues/1004))
* firewall_rule: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1004](https://github.com/cloudflare/cloudflare-go/issues/1004))
* rulesets: add support for `http_custom_errors` phase ([#998](https://github.com/cloudflare/cloudflare-go/issues/998))
* rulesets: add support for `serve_error` action ([#998](https://github.com/cloudflare/cloudflare-go/issues/998))

BUG FIXES:

* access_application: fix inability to set bool values to false ([#1006](https://github.com/cloudflare/cloudflare-go/issues/1006))
* rulesets: fix sni action parameter ([#1002](https://github.com/cloudflare/cloudflare-go/issues/1002))

DEPENDENCIES:

* provider: bumps github.com/golangci/golangci-lint from 1.47.1 to 1.47.2 ([#1005](https://github.com/cloudflare/cloudflare-go/issues/1005))
* provider: bumps github.com/golangci/golangci-lint from 1.47.2 to 1.47.3 ([#1008](https://github.com/cloudflare/cloudflare-go/issues/1008))
* provider: bumps github.com/urfave/cli/v2 from 2.11.0 to 2.11.1 ([#1003](https://github.com/cloudflare/cloudflare-go/issues/1003))

## 0.45.0 (July 20th, 2022)
