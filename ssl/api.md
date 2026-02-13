# SSL

## Analyze

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#AnalyzeNewResponse">AnalyzeNewResponse</a>

Methods:

- <code title="post /zones/{zone_id}/ssl/analyze">client.SSL.Analyze.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#AnalyzeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#AnalyzeNewParams">AnalyzeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#AnalyzeNewResponse">AnalyzeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CertificatePacks

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#HostParam">HostParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#RequestValidity">RequestValidity</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#Host">Host</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#RequestValidity">RequestValidity</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#Status">Status</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#ValidationMethod">ValidationMethod</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackNewResponse">CertificatePackNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackListResponse">CertificatePackListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackDeleteResponse">CertificatePackDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackEditResponse">CertificatePackEditResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackGetResponse">CertificatePackGetResponse</a>

Methods:

- <code title="post /zones/{zone_id}/ssl/certificate_packs/order">client.SSL.CertificatePacks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackNewParams">CertificatePackNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackNewResponse">CertificatePackNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/ssl/certificate_packs">client.SSL.CertificatePacks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackListParams">CertificatePackListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackListResponse">CertificatePackListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}">client.SSL.CertificatePacks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, certificatePackID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackDeleteParams">CertificatePackDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackDeleteResponse">CertificatePackDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}">client.SSL.CertificatePacks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, certificatePackID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackEditParams">CertificatePackEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackEditResponse">CertificatePackEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}">client.SSL.CertificatePacks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, certificatePackID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackGetParams">CertificatePackGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackGetResponse">CertificatePackGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Quota

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackQuotaGetResponse">CertificatePackQuotaGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/ssl/certificate_packs/quota">client.SSL.CertificatePacks.Quota.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackQuotaService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackQuotaGetParams">CertificatePackQuotaGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#CertificatePackQuotaGetResponse">CertificatePackQuotaGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recommendations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#RecommendationGetResponse">RecommendationGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/ssl/recommendation">client.SSL.Recommendations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#RecommendationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#RecommendationGetParams">RecommendationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#RecommendationGetResponse">RecommendationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Universal

### Settings

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#UniversalSSLSettingsParam">UniversalSSLSettingsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#UniversalSSLSettings">UniversalSSLSettings</a>

Methods:

- <code title="patch /zones/{zone_id}/ssl/universal/settings">client.SSL.Universal.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#UniversalSettingService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#UniversalSettingEditParams">UniversalSettingEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#UniversalSSLSettings">UniversalSSLSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/ssl/universal/settings">client.SSL.Universal.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#UniversalSettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#UniversalSettingGetParams">UniversalSettingGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#UniversalSSLSettings">UniversalSSLSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Verification

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#Verification">Verification</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#VerificationEditResponse">VerificationEditResponse</a>

Methods:

- <code title="patch /zones/{zone_id}/ssl/verification/{certificate_pack_id}">client.SSL.Verification.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#VerificationService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, certificatePackID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#VerificationEditParams">VerificationEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#VerificationEditResponse">VerificationEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/ssl/verification">client.SSL.Verification.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#VerificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#VerificationGetParams">VerificationGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl">ssl</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/ssl#Verification">Verification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
