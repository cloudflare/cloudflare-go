# Spectrum

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#DNSParam">DNSParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#EdgeIPsUnionParam">EdgeIPsUnionParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#OriginDNSParam">OriginDNSParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#OriginPortUnionParam">OriginPortUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#DNS">DNS</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#EdgeIPs">EdgeIPs</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#OriginDNS">OriginDNS</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#OriginPortUnion">OriginPortUnion</a>

## Analytics

### Aggregates

#### Currents

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsAggregateCurrentGetResponse">AnalyticsAggregateCurrentGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/spectrum/analytics/aggregate/current">client.Spectrum.Analytics.Aggregates.Currents.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsAggregateCurrentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsAggregateCurrentGetParams">AnalyticsAggregateCurrentGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsAggregateCurrentGetResponse">AnalyticsAggregateCurrentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Events

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#Dimension">Dimension</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#Dimension">Dimension</a>

#### Bytimes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsEventBytimeGetResponse">AnalyticsEventBytimeGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/spectrum/analytics/events/bytime">client.Spectrum.Analytics.Events.Bytimes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsEventBytimeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsEventBytimeGetParams">AnalyticsEventBytimeGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsEventBytimeGetResponse">AnalyticsEventBytimeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Summaries

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsEventSummaryGetResponse">AnalyticsEventSummaryGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/spectrum/analytics/events/summary">client.Spectrum.Analytics.Events.Summaries.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsEventSummaryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsEventSummaryGetParams">AnalyticsEventSummaryGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AnalyticsEventSummaryGetResponse">AnalyticsEventSummaryGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Apps

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppNewResponse">AppNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppUpdateResponse">AppUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppListResponse">AppListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppDeleteResponse">AppDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppGetResponse">AppGetResponse</a>

Methods:

- <code title="post /zones/{zone_id}/spectrum/apps">client.Spectrum.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppNewParams">AppNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppNewResponse">AppNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_id}/spectrum/apps/{app_id}">client.Spectrum.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, appID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppUpdateParams">AppUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppUpdateResponse">AppUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/spectrum/apps">client.Spectrum.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppListParams">AppListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppListResponse">AppListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/spectrum/apps/{app_id}">client.Spectrum.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, appID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppDeleteParams">AppDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppDeleteResponse">AppDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/spectrum/apps/{app_id}">client.Spectrum.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, appID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppGetParams">AppGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum">spectrum</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/spectrum#AppGetResponse">AppGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
