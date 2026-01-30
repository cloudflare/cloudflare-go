# Logs

## Control

### Retention

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlRetentionNewResponse">ControlRetentionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlRetentionGetResponse">ControlRetentionGetResponse</a>

Methods:

- <code title="post /zones/{zone_id}/logs/control/retention/flag">client.Logs.Control.Retention.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlRetentionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlRetentionNewParams">ControlRetentionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlRetentionNewResponse">ControlRetentionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/logs/control/retention/flag">client.Logs.Control.Retention.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlRetentionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlRetentionGetParams">ControlRetentionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlRetentionGetResponse">ControlRetentionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Cmb

#### Config

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#CmbConfigParam">CmbConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#CmbConfig">CmbConfig</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlCmbConfigDeleteResponse">ControlCmbConfigDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_id}/logs/control/cmb/config">client.Logs.Control.Cmb.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlCmbConfigService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlCmbConfigNewParams">ControlCmbConfigNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#CmbConfig">CmbConfig</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/logs/control/cmb/config">client.Logs.Control.Cmb.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlCmbConfigService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlCmbConfigDeleteParams">ControlCmbConfigDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlCmbConfigDeleteResponse">ControlCmbConfigDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/logs/control/cmb/config">client.Logs.Control.Cmb.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlCmbConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ControlCmbConfigGetParams">ControlCmbConfigGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#CmbConfig">CmbConfig</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RayID

Methods:

- <code title="get /zones/{zone_id}/logs/rayids/{ray_id}">client.Logs.RayID.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#RayIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, RayID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#RayIDGetParams">RayIDGetParams</a>) (\*interface{}, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Received

Methods:

- <code title="get /zones/{zone_id}/logs/received">client.Logs.Received.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ReceivedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ReceivedGetParams">ReceivedGetParams</a>) (\*interface{}, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Fields

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ReceivedFieldGetResponse">ReceivedFieldGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/logs/received/fields">client.Logs.Received.Fields.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ReceivedFieldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ReceivedFieldGetParams">ReceivedFieldGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/logs#ReceivedFieldGetResponse">ReceivedFieldGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
