# BotnetFeed

## ASN

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ASNDayReportResponse">ASNDayReportResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ASNFullReportResponse">ASNFullReportResponse</a>

Methods:

- <code title="get /accounts/{account_id}/botnet_feed/asn/{asn_id}/day_report">client.BotnetFeed.ASN.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ASNService.DayReport">DayReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, asnID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ASNDayReportParams">ASNDayReportParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ASNDayReportResponse">ASNDayReportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/botnet_feed/asn/{asn_id}/full_report">client.BotnetFeed.ASN.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ASNService.FullReport">FullReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, asnID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ASNFullReportParams">ASNFullReportParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ASNFullReportResponse">ASNFullReportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Configs

### ASN

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ConfigASNDeleteResponse">ConfigASNDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ConfigASNGetResponse">ConfigASNGetResponse</a>

Methods:

- <code title="delete /accounts/{account_id}/botnet_feed/configs/asn/{asn_id}">client.BotnetFeed.Configs.ASN.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ConfigASNService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, asnID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ConfigASNDeleteParams">ConfigASNDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ConfigASNDeleteResponse">ConfigASNDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/botnet_feed/configs/asn">client.BotnetFeed.Configs.ASN.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ConfigASNService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ConfigASNGetParams">ConfigASNGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed">botnet_feed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/botnet_feed#ConfigASNGetResponse">ConfigASNGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
