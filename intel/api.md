# Intel

## ASN

Methods:

- <code title="get /accounts/{account_id}/intel/asn/{asn}">client.Intel.ASN.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#ASNService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, asn <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#ASNParam">ASNParam</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#ASNGetParams">ASNGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#ASN">ASN</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Subnets

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#ASNSubnetGetResponse">ASNSubnetGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/intel/asn/{asn}/subnets">client.Intel.ASN.Subnets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#ASNSubnetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, asn <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#ASNParam">ASNParam</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#ASNSubnetGetParams">ASNSubnetGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#ASNSubnetGetResponse">ASNSubnetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DNS

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DNS">DNS</a>

Methods:

- <code title="get /accounts/{account_id}/intel/dns">client.Intel.DNS.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DNSService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DNSListParams">DNSListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePagination">V4PagePagination</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DNS">DNS</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Domains

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#Domain">Domain</a>

Methods:

- <code title="get /accounts/{account_id}/intel/domain">client.Intel.Domains.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainGetParams">DomainGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#Domain">Domain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Bulks

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainBulkGetResponse">DomainBulkGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/intel/domain/bulk">client.Intel.Domains.Bulks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainBulkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainBulkGetParams">DomainBulkGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainBulkGetResponse">DomainBulkGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DomainHistory

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainHistory">DomainHistory</a>

Methods:

- <code title="get /accounts/{account_id}/intel/domain-history">client.Intel.DomainHistory.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainHistoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainHistoryGetParams">DomainHistoryGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#DomainHistory">DomainHistory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## IPs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IP">IP</a>

Methods:

- <code title="get /accounts/{account_id}/intel/ip">client.Intel.IPs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IPService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IPGetParams">IPGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IP">IP</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## IPLists

## Miscategorizations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#MiscategorizationNewResponse">MiscategorizationNewResponse</a>

Methods:

- <code title="post /accounts/{account_id}/intel/miscategorization">client.Intel.Miscategorizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#MiscategorizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#MiscategorizationNewParams">MiscategorizationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#MiscategorizationNewResponse">MiscategorizationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Whois

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#WhoisGetResponse">WhoisGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/intel/whois">client.Intel.Whois.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#WhoisService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#WhoisGetParams">WhoisGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#WhoisGetResponse">WhoisGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## IndicatorFeeds

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedNewResponse">IndicatorFeedNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedUpdateResponse">IndicatorFeedUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedListResponse">IndicatorFeedListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedGetResponse">IndicatorFeedGetResponse</a>

Methods:

- <code title="post /accounts/{account_id}/intel/indicator-feeds">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedNewParams">IndicatorFeedNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedNewResponse">IndicatorFeedNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_id}/intel/indicator-feeds/{feed_id}">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, feedID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedUpdateParams">IndicatorFeedUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedUpdateResponse">IndicatorFeedUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/intel/indicator-feeds">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedListParams">IndicatorFeedListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedListResponse">IndicatorFeedListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/intel/indicator-feeds/{feed_id}/data">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedService.Data">Data</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, feedID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedDataParams">IndicatorFeedDataParams</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/intel/indicator-feeds/{feed_id}">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, feedID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedGetParams">IndicatorFeedGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedGetResponse">IndicatorFeedGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Snapshots

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedSnapshotUpdateResponse">IndicatorFeedSnapshotUpdateResponse</a>

Methods:

- <code title="put /accounts/{account_id}/intel/indicator-feeds/{feed_id}/snapshot">client.Intel.IndicatorFeeds.Snapshots.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedSnapshotService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, feedID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedSnapshotUpdateParams">IndicatorFeedSnapshotUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedSnapshotUpdateResponse">IndicatorFeedSnapshotUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Permissions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionNewResponse">IndicatorFeedPermissionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionListResponse">IndicatorFeedPermissionListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionDeleteResponse">IndicatorFeedPermissionDeleteResponse</a>

Methods:

- <code title="put /accounts/{account_id}/intel/indicator-feeds/permissions/add">client.Intel.IndicatorFeeds.Permissions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionNewParams">IndicatorFeedPermissionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionNewResponse">IndicatorFeedPermissionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/intel/indicator-feeds/permissions/view">client.Intel.IndicatorFeeds.Permissions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionListParams">IndicatorFeedPermissionListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionListResponse">IndicatorFeedPermissionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_id}/intel/indicator-feeds/permissions/remove">client.Intel.IndicatorFeeds.Permissions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionDeleteParams">IndicatorFeedPermissionDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IndicatorFeedPermissionDeleteResponse">IndicatorFeedPermissionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Downloads

## Sinkholes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#Sinkhole">Sinkhole</a>

Methods:

- <code title="get /accounts/{account_id}/intel/sinkholes">client.Intel.Sinkholes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#SinkholeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#SinkholeListParams">SinkholeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#Sinkhole">Sinkhole</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AttackSurfaceReport

### IssueTypes

Methods:

- <code title="get /accounts/{account_id}/intel/attack-surface-report/issue-types">client.Intel.AttackSurfaceReport.IssueTypes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueTypeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueTypeGetParams">AttackSurfaceReportIssueTypeGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/builtin#string">string</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Issues

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IssueType">IssueType</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#SeverityQueryParam">SeverityQueryParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#IssueType">IssueType</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueListResponse">AttackSurfaceReportIssueListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueClassResponse">AttackSurfaceReportIssueClassResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueDismissResponse">AttackSurfaceReportIssueDismissResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueSeverityResponse">AttackSurfaceReportIssueSeverityResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueTypeResponse">AttackSurfaceReportIssueTypeResponse</a>

Methods:

- <code title="get /accounts/{account_id}/intel/attack-surface-report/issues">client.Intel.AttackSurfaceReport.Issues.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueListParams">AttackSurfaceReportIssueListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePagination">V4PagePagination</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueListResponse">AttackSurfaceReportIssueListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/intel/attack-surface-report/issues/class">client.Intel.AttackSurfaceReport.Issues.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueService.Class">Class</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueClassParams">AttackSurfaceReportIssueClassParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueClassResponse">AttackSurfaceReportIssueClassResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_id}/intel/attack-surface-report/{issue_id}/dismiss">client.Intel.AttackSurfaceReport.Issues.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueService.Dismiss">Dismiss</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, issueID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueDismissParams">AttackSurfaceReportIssueDismissParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueDismissResponse">AttackSurfaceReportIssueDismissResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/intel/attack-surface-report/issues/severity">client.Intel.AttackSurfaceReport.Issues.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueService.Severity">Severity</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueSeverityParams">AttackSurfaceReportIssueSeverityParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueSeverityResponse">AttackSurfaceReportIssueSeverityResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/intel/attack-surface-report/issues/type">client.Intel.AttackSurfaceReport.Issues.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueService.Type">Type</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueTypeParams">AttackSurfaceReportIssueTypeParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel">intel</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/intel#AttackSurfaceReportIssueTypeResponse">AttackSurfaceReportIssueTypeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
