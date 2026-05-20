# Billing

## Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileGetResponse">ProfileGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/billing/profile">client.Billing.Profiles.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileGetParams">ProfileGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileGetResponse">ProfileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Usage

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetResponse">UsageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoResponse">UsagePaygoResponse</a>

Methods:

- <code title="get /accounts/{account_id}/billable/usage">client.Billing.Usage.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetParams">UsageGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetResponse">UsageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/paygo-usage">client.Billing.Usage.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageService.Paygo">Paygo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoParams">UsagePaygoParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoResponse">UsagePaygoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
