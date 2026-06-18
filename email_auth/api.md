# EmailAuth

## DMARCReports

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#DMARCReportEditResponse">DMARCReportEditResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#DMARCReportGetResponse">DMARCReportGetResponse</a>

Methods:

- <code title="patch /zones/{zone_id}/email/auth/dmarc-reports">client.EmailAuth.DMARCReports.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#DMARCReportService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#DMARCReportEditParams">DMARCReportEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#DMARCReportEditResponse">DMARCReportEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/email/auth/dmarc-reports">client.EmailAuth.DMARCReports.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#DMARCReportService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#DMARCReportGetParams">DMARCReportGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#DMARCReportGetResponse">DMARCReportGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SPF

### Inspect

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#SPFInspectGetResponse">SPFInspectGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/email/auth/spf/inspect">client.EmailAuth.SPF.Inspect.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#SPFInspectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#SPFInspectGetParams">SPFInspectGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth">email_auth</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/email_auth#SPFInspectGetResponse">SPFInspectGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
