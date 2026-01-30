# Zaraz

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ButtonTextTranslationParam">ButtonTextTranslationParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#NeoEventParam">NeoEventParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ButtonTextTranslation">ButtonTextTranslation</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#NeoEvent">NeoEvent</a>

Methods:

- <code title="put /zones/{zone_id}/settings/zaraz/workflow">client.Zaraz.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ZarazService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ZarazUpdateParams">ZarazUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Workflow">Workflow</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Config

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Configuration">Configuration</a>

Methods:

- <code title="put /zones/{zone_id}/settings/zaraz/config">client.Zaraz.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ConfigService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ConfigUpdateParams">ConfigUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/settings/zaraz/config">client.Zaraz.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ConfigGetParams">ConfigGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Default

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/default">client.Zaraz.Default.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#DefaultService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#DefaultGetParams">DefaultGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Export

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/export">client.Zaraz.Export.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ExportService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#ExportGetParams">ExportGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## History

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryListResponse">HistoryListResponse</a>

Methods:

- <code title="put /zones/{zone_id}/settings/zaraz/history">client.Zaraz.History.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryUpdateParams">HistoryUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/settings/zaraz/history">client.Zaraz.History.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryListParams">HistoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryListResponse">HistoryListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Configs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryConfigGetResponse">HistoryConfigGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/history/configs">client.Zaraz.History.Configs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryConfigGetParams">HistoryConfigGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#HistoryConfigGetResponse">HistoryConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Publish

Methods:

- <code title="post /zones/{zone_id}/settings/zaraz/publish">client.Zaraz.Publish.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#PublishService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#PublishNewParams">PublishNewParams</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Workflow

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Workflow">Workflow</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Workflow">Workflow</a>

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/workflow">client.Zaraz.Workflow.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#WorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#WorkflowGetParams">WorkflowGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz">zaraz</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/zaraz#Workflow">Workflow</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
