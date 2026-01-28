# DNS

## Records

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#DNSRecordListResponse">DNSRecordListResponse</a>

Methods:

- <code title="get /zones/{zone_id}/dns_records">client.DNS.Records.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#DNSRecordService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#DNSRecordListParams">DNSRecordListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#DNSRecordListResponse">DNSRecordListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Zones

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneNewResponse">ZoneNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneEditResponse">ZoneEditResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneGetResponse">ZoneGetResponse</a>

Methods:

- <code title="post /zones">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneNewParams">ZoneNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneNewResponse">ZoneNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_id}">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneEditParams">ZoneEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneEditResponse">ZoneEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneGetParams">ZoneGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#ZoneGetResponse">ZoneGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#AccountListResponse">AccountListResponse</a>

Methods:

- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#AccountListParams">AccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6#AccountListResponse">AccountListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
