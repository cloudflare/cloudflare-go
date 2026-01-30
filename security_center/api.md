# SecurityCenter

## Insights

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightListResponse">InsightListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightDismissResponse">InsightDismissResponse</a>

Methods:

- <code title="get /{accounts_or_zones}/{account_or_zone_id}/security-center/insights">client.SecurityCenter.Insights.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightListParams">InsightListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePagination">V4PagePagination</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightListResponse">InsightListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{accounts_or_zones}/{account_or_zone_id}/security-center/insights/{issue_id}/dismiss">client.SecurityCenter.Insights.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightService.Dismiss">Dismiss</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, issueID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightDismissParams">InsightDismissParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightDismissResponse">InsightDismissResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Class

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightClassGetResponse">InsightClassGetResponse</a>

Methods:

- <code title="get /{accounts_or_zones}/{account_or_zone_id}/security-center/insights/class">client.SecurityCenter.Insights.Class.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightClassService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightClassGetParams">InsightClassGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightClassGetResponse">InsightClassGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Severity

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightSeverityGetResponse">InsightSeverityGetResponse</a>

Methods:

- <code title="get /{accounts_or_zones}/{account_or_zone_id}/security-center/insights/severity">client.SecurityCenter.Insights.Severity.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightSeverityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightSeverityGetParams">InsightSeverityGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightSeverityGetResponse">InsightSeverityGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Type

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightTypeGetResponse">InsightTypeGetResponse</a>

Methods:

- <code title="get /{accounts_or_zones}/{account_or_zone_id}/security-center/insights/type">client.SecurityCenter.Insights.Type.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightTypeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightTypeGetParams">InsightTypeGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center">security_center</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/security_center#InsightTypeGetResponse">InsightTypeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
