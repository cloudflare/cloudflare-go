# Zones

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneNewResponse">ZoneNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneGetResponse">ZoneGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneUpdateResponse">ZoneUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneListResponse">ZoneListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneDeleteResponse">ZoneDeleteResponse</a>

Methods:

- <code title="post /zones">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneNewParams">ZoneNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneNewResponse">ZoneNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{identifier}">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneGetResponse">ZoneGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{identifier}">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneUpdateParams">ZoneUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneUpdateResponse">ZoneUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneListParams">ZoneListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneListResponse">ZoneListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{identifier}">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneDeleteResponse">ZoneDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LoadBalancers

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerNewResponse">ZoneLoadBalancerNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerGetResponse">ZoneLoadBalancerGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerUpdateResponse">ZoneLoadBalancerUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerListResponse">ZoneLoadBalancerListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerDeleteResponse">ZoneLoadBalancerDeleteResponse</a>

Methods:

- <code title="post /zones/{identifier}/load_balancers">client.Zones.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerNewParams">ZoneLoadBalancerNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerNewResponse">ZoneLoadBalancerNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{identifier1}/load_balancers/{identifier}">client.Zones.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier1 <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerGetResponse">ZoneLoadBalancerGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{identifier1}/load_balancers/{identifier}">client.Zones.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier1 <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerUpdateParams">ZoneLoadBalancerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerUpdateResponse">ZoneLoadBalancerUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{identifier}/load_balancers">client.Zones.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerListResponse">ZoneLoadBalancerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{identifier1}/load_balancers/{identifier}">client.Zones.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier1 <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneLoadBalancerDeleteResponse">ZoneLoadBalancerDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Dnssecs

## RateLimits

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneRateLimitGetResponse">ZoneRateLimitGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneRateLimitListResponse">ZoneRateLimitListResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/rate_limits/{id}">client.Zones.RateLimits.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneRateLimitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneRateLimitGetResponse">ZoneRateLimitGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/rate_limits">client.Zones.RateLimits.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneRateLimitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneRateLimitListParams">ZoneRateLimitListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneRateLimitListResponse">ZoneRateLimitListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Settings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingListResponse">ZoneSettingListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEditResponse">ZoneSettingEditResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/settings">client.Zones.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingListResponse">ZoneSettingListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_identifier}/settings">client.Zones.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEditParams">ZoneSettingEditParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEditResponse">ZoneSettingEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### AdvancedDDOS

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAdvancedDDOSListResponse">ZoneSettingAdvancedDDOSListResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/settings/advanced_ddos">client.Zones.Settings.AdvancedDDOS.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAdvancedDDOSService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAdvancedDDOSListResponse">ZoneSettingAdvancedDDOSListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### AlwaysOnline

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysOnlineGetResponse">ZoneSettingAlwaysOnlineGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysOnlineUpdateResponse">ZoneSettingAlwaysOnlineUpdateResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/settings/always_online">client.Zones.Settings.AlwaysOnline.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysOnlineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysOnlineGetResponse">ZoneSettingAlwaysOnlineGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_identifier}/settings/always_online">client.Zones.Settings.AlwaysOnline.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysOnlineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysOnlineUpdateParams">ZoneSettingAlwaysOnlineUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysOnlineUpdateResponse">ZoneSettingAlwaysOnlineUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### AlwaysUseHTTPs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysUseHTTPUpdateResponse">ZoneSettingAlwaysUseHTTPUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysUseHTTPListResponse">ZoneSettingAlwaysUseHTTPListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/always_use_https">client.Zones.Settings.AlwaysUseHTTPs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysUseHTTPService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysUseHTTPUpdateParams">ZoneSettingAlwaysUseHTTPUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysUseHTTPUpdateResponse">ZoneSettingAlwaysUseHTTPUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/always_use_https">client.Zones.Settings.AlwaysUseHTTPs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysUseHTTPService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAlwaysUseHTTPListResponse">ZoneSettingAlwaysUseHTTPListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### AutomaticHTTPsRewrites

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticHTTPsRewriteUpdateResponse">ZoneSettingAutomaticHTTPsRewriteUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticHTTPsRewriteListResponse">ZoneSettingAutomaticHTTPsRewriteListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/automatic_https_rewrites">client.Zones.Settings.AutomaticHTTPsRewrites.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticHTTPsRewriteService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticHTTPsRewriteUpdateParams">ZoneSettingAutomaticHTTPsRewriteUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticHTTPsRewriteUpdateResponse">ZoneSettingAutomaticHTTPsRewriteUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/automatic_https_rewrites">client.Zones.Settings.AutomaticHTTPsRewrites.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticHTTPsRewriteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticHTTPsRewriteListResponse">ZoneSettingAutomaticHTTPsRewriteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### AutomaticPlatformOptimization

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticPlatformOptimizationUpdateResponse">ZoneSettingAutomaticPlatformOptimizationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticPlatformOptimizationListResponse">ZoneSettingAutomaticPlatformOptimizationListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/automatic_platform_optimization">client.Zones.Settings.AutomaticPlatformOptimization.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticPlatformOptimizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticPlatformOptimizationUpdateParams">ZoneSettingAutomaticPlatformOptimizationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticPlatformOptimizationUpdateResponse">ZoneSettingAutomaticPlatformOptimizationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/automatic_platform_optimization">client.Zones.Settings.AutomaticPlatformOptimization.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticPlatformOptimizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingAutomaticPlatformOptimizationListResponse">ZoneSettingAutomaticPlatformOptimizationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Brotli

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrotliUpdateResponse">ZoneSettingBrotliUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrotliListResponse">ZoneSettingBrotliListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/brotli">client.Zones.Settings.Brotli.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrotliService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrotliUpdateParams">ZoneSettingBrotliUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrotliUpdateResponse">ZoneSettingBrotliUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/brotli">client.Zones.Settings.Brotli.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrotliService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrotliListResponse">ZoneSettingBrotliListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### BrowserCacheTTLs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCacheTTLUpdateResponse">ZoneSettingBrowserCacheTTLUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCacheTTLListResponse">ZoneSettingBrowserCacheTTLListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/browser_cache_ttl">client.Zones.Settings.BrowserCacheTTLs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCacheTTLService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCacheTTLUpdateParams">ZoneSettingBrowserCacheTTLUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCacheTTLUpdateResponse">ZoneSettingBrowserCacheTTLUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/browser_cache_ttl">client.Zones.Settings.BrowserCacheTTLs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCacheTTLService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCacheTTLListResponse">ZoneSettingBrowserCacheTTLListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### BrowserChecks

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCheckUpdateResponse">ZoneSettingBrowserCheckUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCheckListResponse">ZoneSettingBrowserCheckListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/browser_check">client.Zones.Settings.BrowserChecks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCheckService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCheckUpdateParams">ZoneSettingBrowserCheckUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCheckUpdateResponse">ZoneSettingBrowserCheckUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/browser_check">client.Zones.Settings.BrowserChecks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCheckService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingBrowserCheckListResponse">ZoneSettingBrowserCheckListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### CacheLevels

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCacheLevelUpdateResponse">ZoneSettingCacheLevelUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCacheLevelListResponse">ZoneSettingCacheLevelListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/cache_level">client.Zones.Settings.CacheLevels.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCacheLevelService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCacheLevelUpdateParams">ZoneSettingCacheLevelUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCacheLevelUpdateResponse">ZoneSettingCacheLevelUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/cache_level">client.Zones.Settings.CacheLevels.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCacheLevelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCacheLevelListResponse">ZoneSettingCacheLevelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ChallengeTTLs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingChallengeTTLUpdateResponse">ZoneSettingChallengeTTLUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingChallengeTTLListResponse">ZoneSettingChallengeTTLListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/challenge_ttl">client.Zones.Settings.ChallengeTTLs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingChallengeTTLService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingChallengeTTLUpdateParams">ZoneSettingChallengeTTLUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingChallengeTTLUpdateResponse">ZoneSettingChallengeTTLUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/challenge_ttl">client.Zones.Settings.ChallengeTTLs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingChallengeTTLService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingChallengeTTLListResponse">ZoneSettingChallengeTTLListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Ciphers

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCipherUpdateResponse">ZoneSettingCipherUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCipherListResponse">ZoneSettingCipherListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/ciphers">client.Zones.Settings.Ciphers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCipherService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCipherUpdateParams">ZoneSettingCipherUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCipherUpdateResponse">ZoneSettingCipherUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/ciphers">client.Zones.Settings.Ciphers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCipherService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingCipherListResponse">ZoneSettingCipherListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### DevelopmentModes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingDevelopmentModeUpdateResponse">ZoneSettingDevelopmentModeUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingDevelopmentModeListResponse">ZoneSettingDevelopmentModeListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/development_mode">client.Zones.Settings.DevelopmentModes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingDevelopmentModeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingDevelopmentModeUpdateParams">ZoneSettingDevelopmentModeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingDevelopmentModeUpdateResponse">ZoneSettingDevelopmentModeUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/development_mode">client.Zones.Settings.DevelopmentModes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingDevelopmentModeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingDevelopmentModeListResponse">ZoneSettingDevelopmentModeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### EarlyHints

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEarlyHintUpdateResponse">ZoneSettingEarlyHintUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEarlyHintListResponse">ZoneSettingEarlyHintListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/early_hints">client.Zones.Settings.EarlyHints.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEarlyHintService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEarlyHintUpdateParams">ZoneSettingEarlyHintUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEarlyHintUpdateResponse">ZoneSettingEarlyHintUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/early_hints">client.Zones.Settings.EarlyHints.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEarlyHintService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEarlyHintListResponse">ZoneSettingEarlyHintListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### EmailObfuscations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEmailObfuscationUpdateResponse">ZoneSettingEmailObfuscationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEmailObfuscationListResponse">ZoneSettingEmailObfuscationListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/email_obfuscation">client.Zones.Settings.EmailObfuscations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEmailObfuscationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEmailObfuscationUpdateParams">ZoneSettingEmailObfuscationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEmailObfuscationUpdateResponse">ZoneSettingEmailObfuscationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/email_obfuscation">client.Zones.Settings.EmailObfuscations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEmailObfuscationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingEmailObfuscationListResponse">ZoneSettingEmailObfuscationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### H2Prioritizations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingH2PrioritizationUpdateResponse">ZoneSettingH2PrioritizationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingH2PrioritizationListResponse">ZoneSettingH2PrioritizationListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/h2_prioritization">client.Zones.Settings.H2Prioritizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingH2PrioritizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingH2PrioritizationUpdateParams">ZoneSettingH2PrioritizationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingH2PrioritizationUpdateResponse">ZoneSettingH2PrioritizationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/h2_prioritization">client.Zones.Settings.H2Prioritizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingH2PrioritizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingH2PrioritizationListResponse">ZoneSettingH2PrioritizationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### HotlinkProtections

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHotlinkProtectionUpdateResponse">ZoneSettingHotlinkProtectionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHotlinkProtectionListResponse">ZoneSettingHotlinkProtectionListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/hotlink_protection">client.Zones.Settings.HotlinkProtections.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHotlinkProtectionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHotlinkProtectionUpdateParams">ZoneSettingHotlinkProtectionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHotlinkProtectionUpdateResponse">ZoneSettingHotlinkProtectionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/hotlink_protection">client.Zones.Settings.HotlinkProtections.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHotlinkProtectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHotlinkProtectionListResponse">ZoneSettingHotlinkProtectionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### HTTP2s

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP2UpdateResponse">ZoneSettingHTTP2UpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP2ListResponse">ZoneSettingHTTP2ListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/http2">client.Zones.Settings.HTTP2s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP2Service.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP2UpdateParams">ZoneSettingHTTP2UpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP2UpdateResponse">ZoneSettingHTTP2UpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/http2">client.Zones.Settings.HTTP2s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP2Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP2ListResponse">ZoneSettingHTTP2ListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### HTTP3s

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP3UpdateResponse">ZoneSettingHTTP3UpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP3ListResponse">ZoneSettingHTTP3ListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/http3">client.Zones.Settings.HTTP3s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP3Service.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP3UpdateParams">ZoneSettingHTTP3UpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP3UpdateResponse">ZoneSettingHTTP3UpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/http3">client.Zones.Settings.HTTP3s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP3Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingHTTP3ListResponse">ZoneSettingHTTP3ListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ImageResizings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingImageResizingUpdateResponse">ZoneSettingImageResizingUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingImageResizingListResponse">ZoneSettingImageResizingListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/image_resizing">client.Zones.Settings.ImageResizings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingImageResizingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingImageResizingUpdateParams">ZoneSettingImageResizingUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingImageResizingUpdateResponse">ZoneSettingImageResizingUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/image_resizing">client.Zones.Settings.ImageResizings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingImageResizingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingImageResizingListResponse">ZoneSettingImageResizingListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### IPGeolocations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPGeolocationUpdateResponse">ZoneSettingIPGeolocationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPGeolocationListResponse">ZoneSettingIPGeolocationListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/ip_geolocation">client.Zones.Settings.IPGeolocations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPGeolocationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPGeolocationUpdateParams">ZoneSettingIPGeolocationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPGeolocationUpdateResponse">ZoneSettingIPGeolocationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/ip_geolocation">client.Zones.Settings.IPGeolocations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPGeolocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPGeolocationListResponse">ZoneSettingIPGeolocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### IPV6s

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPV6UpdateResponse">ZoneSettingIPV6UpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPV6ListResponse">ZoneSettingIPV6ListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/ipv6">client.Zones.Settings.IPV6s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPV6Service.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPV6UpdateParams">ZoneSettingIPV6UpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPV6UpdateResponse">ZoneSettingIPV6UpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/ipv6">client.Zones.Settings.IPV6s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPV6Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingIPV6ListResponse">ZoneSettingIPV6ListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### MinTLSVersions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinTLSVersionUpdateResponse">ZoneSettingMinTLSVersionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinTLSVersionListResponse">ZoneSettingMinTLSVersionListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/min_tls_version">client.Zones.Settings.MinTLSVersions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinTLSVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinTLSVersionUpdateParams">ZoneSettingMinTLSVersionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinTLSVersionUpdateResponse">ZoneSettingMinTLSVersionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/min_tls_version">client.Zones.Settings.MinTLSVersions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinTLSVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinTLSVersionListResponse">ZoneSettingMinTLSVersionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Minifies

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinifyUpdateResponse">ZoneSettingMinifyUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinifyListResponse">ZoneSettingMinifyListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/minify">client.Zones.Settings.Minifies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinifyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinifyUpdateParams">ZoneSettingMinifyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinifyUpdateResponse">ZoneSettingMinifyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/minify">client.Zones.Settings.Minifies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinifyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMinifyListResponse">ZoneSettingMinifyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Mirages

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMirageUpdateResponse">ZoneSettingMirageUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMirageListResponse">ZoneSettingMirageListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/mirage">client.Zones.Settings.Mirages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMirageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMirageUpdateParams">ZoneSettingMirageUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMirageUpdateResponse">ZoneSettingMirageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/mirage">client.Zones.Settings.Mirages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMirageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMirageListResponse">ZoneSettingMirageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### MobileRedirects

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMobileRedirectUpdateResponse">ZoneSettingMobileRedirectUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMobileRedirectListResponse">ZoneSettingMobileRedirectListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/mobile_redirect">client.Zones.Settings.MobileRedirects.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMobileRedirectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMobileRedirectUpdateParams">ZoneSettingMobileRedirectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMobileRedirectUpdateResponse">ZoneSettingMobileRedirectUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/mobile_redirect">client.Zones.Settings.MobileRedirects.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMobileRedirectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingMobileRedirectListResponse">ZoneSettingMobileRedirectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### NELs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingNELUpdateResponse">ZoneSettingNELUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingNELListResponse">ZoneSettingNELListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/nel">client.Zones.Settings.NELs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingNELService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingNELUpdateParams">ZoneSettingNELUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingNELUpdateResponse">ZoneSettingNELUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/nel">client.Zones.Settings.NELs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingNELService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingNELListResponse">ZoneSettingNELListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OpportunisticEncryptions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticEncryptionUpdateResponse">ZoneSettingOpportunisticEncryptionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticEncryptionListResponse">ZoneSettingOpportunisticEncryptionListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/opportunistic_encryption">client.Zones.Settings.OpportunisticEncryptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticEncryptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticEncryptionUpdateParams">ZoneSettingOpportunisticEncryptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticEncryptionUpdateResponse">ZoneSettingOpportunisticEncryptionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/opportunistic_encryption">client.Zones.Settings.OpportunisticEncryptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticEncryptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticEncryptionListResponse">ZoneSettingOpportunisticEncryptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OpportunisticOnions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticOnionUpdateResponse">ZoneSettingOpportunisticOnionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticOnionListResponse">ZoneSettingOpportunisticOnionListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/opportunistic_onion">client.Zones.Settings.OpportunisticOnions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticOnionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticOnionUpdateParams">ZoneSettingOpportunisticOnionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticOnionUpdateResponse">ZoneSettingOpportunisticOnionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/opportunistic_onion">client.Zones.Settings.OpportunisticOnions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticOnionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOpportunisticOnionListResponse">ZoneSettingOpportunisticOnionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OrangeToOranges

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOrangeToOrangeUpdateResponse">ZoneSettingOrangeToOrangeUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOrangeToOrangeListResponse">ZoneSettingOrangeToOrangeListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/orange_to_orange">client.Zones.Settings.OrangeToOranges.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOrangeToOrangeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOrangeToOrangeUpdateParams">ZoneSettingOrangeToOrangeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOrangeToOrangeUpdateResponse">ZoneSettingOrangeToOrangeUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/orange_to_orange">client.Zones.Settings.OrangeToOranges.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOrangeToOrangeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOrangeToOrangeListResponse">ZoneSettingOrangeToOrangeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OriginErrorPagePassThrus

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginErrorPagePassThrusUpdateResponse">ZoneSettingOriginErrorPagePassThrusUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginErrorPagePassThrusListResponse">ZoneSettingOriginErrorPagePassThrusListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/origin_error_page_pass_thru">client.Zones.Settings.OriginErrorPagePassThrus.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginErrorPagePassThrusService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginErrorPagePassThrusUpdateParams">ZoneSettingOriginErrorPagePassThrusUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginErrorPagePassThrusUpdateResponse">ZoneSettingOriginErrorPagePassThrusUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/origin_error_page_pass_thru">client.Zones.Settings.OriginErrorPagePassThrus.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginErrorPagePassThrusService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginErrorPagePassThrusListResponse">ZoneSettingOriginErrorPagePassThrusListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OriginMaxHTTPVersions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginMaxHTTPVersionUpdateResponse">ZoneSettingOriginMaxHTTPVersionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginMaxHTTPVersionListResponse">ZoneSettingOriginMaxHTTPVersionListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/origin_max_http_version">client.Zones.Settings.OriginMaxHTTPVersions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginMaxHTTPVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginMaxHTTPVersionUpdateParams">ZoneSettingOriginMaxHTTPVersionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginMaxHTTPVersionUpdateResponse">ZoneSettingOriginMaxHTTPVersionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/origin_max_http_version">client.Zones.Settings.OriginMaxHTTPVersions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginMaxHTTPVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingOriginMaxHTTPVersionListResponse">ZoneSettingOriginMaxHTTPVersionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Polishes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPolishUpdateResponse">ZoneSettingPolishUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPolishListResponse">ZoneSettingPolishListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/polish">client.Zones.Settings.Polishes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPolishService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPolishUpdateParams">ZoneSettingPolishUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPolishUpdateResponse">ZoneSettingPolishUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/polish">client.Zones.Settings.Polishes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPolishService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPolishListResponse">ZoneSettingPolishListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### PrefetchPreloads

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPrefetchPreloadUpdateResponse">ZoneSettingPrefetchPreloadUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPrefetchPreloadListResponse">ZoneSettingPrefetchPreloadListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/prefetch_preload">client.Zones.Settings.PrefetchPreloads.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPrefetchPreloadService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPrefetchPreloadUpdateParams">ZoneSettingPrefetchPreloadUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPrefetchPreloadUpdateResponse">ZoneSettingPrefetchPreloadUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/prefetch_preload">client.Zones.Settings.PrefetchPreloads.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPrefetchPreloadService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPrefetchPreloadListResponse">ZoneSettingPrefetchPreloadListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### PrivacyPasses

### ProxyReadTimeouts

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingProxyReadTimeoutUpdateResponse">ZoneSettingProxyReadTimeoutUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingProxyReadTimeoutListResponse">ZoneSettingProxyReadTimeoutListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/proxy_read_timeout">client.Zones.Settings.ProxyReadTimeouts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingProxyReadTimeoutService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingProxyReadTimeoutUpdateParams">ZoneSettingProxyReadTimeoutUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingProxyReadTimeoutUpdateResponse">ZoneSettingProxyReadTimeoutUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/proxy_read_timeout">client.Zones.Settings.ProxyReadTimeouts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingProxyReadTimeoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingProxyReadTimeoutListResponse">ZoneSettingProxyReadTimeoutListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### PseudoIpv4s

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPseudoIpv4UpdateResponse">ZoneSettingPseudoIpv4UpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPseudoIpv4ListResponse">ZoneSettingPseudoIpv4ListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/pseudo_ipv4">client.Zones.Settings.PseudoIpv4s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPseudoIpv4Service.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPseudoIpv4UpdateParams">ZoneSettingPseudoIpv4UpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPseudoIpv4UpdateResponse">ZoneSettingPseudoIpv4UpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/pseudo_ipv4">client.Zones.Settings.PseudoIpv4s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPseudoIpv4Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingPseudoIpv4ListResponse">ZoneSettingPseudoIpv4ListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ResponseBufferings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingResponseBufferingUpdateResponse">ZoneSettingResponseBufferingUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingResponseBufferingListResponse">ZoneSettingResponseBufferingListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/response_buffering">client.Zones.Settings.ResponseBufferings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingResponseBufferingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingResponseBufferingUpdateParams">ZoneSettingResponseBufferingUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingResponseBufferingUpdateResponse">ZoneSettingResponseBufferingUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/response_buffering">client.Zones.Settings.ResponseBufferings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingResponseBufferingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingResponseBufferingListResponse">ZoneSettingResponseBufferingListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### RocketLoaders

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingRocketLoaderUpdateResponse">ZoneSettingRocketLoaderUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingRocketLoaderListResponse">ZoneSettingRocketLoaderListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/rocket_loader">client.Zones.Settings.RocketLoaders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingRocketLoaderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingRocketLoaderUpdateParams">ZoneSettingRocketLoaderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingRocketLoaderUpdateResponse">ZoneSettingRocketLoaderUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/rocket_loader">client.Zones.Settings.RocketLoaders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingRocketLoaderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingRocketLoaderListResponse">ZoneSettingRocketLoaderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### SecurityHeaders

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityHeaderUpdateResponse">ZoneSettingSecurityHeaderUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityHeaderListResponse">ZoneSettingSecurityHeaderListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/security_header">client.Zones.Settings.SecurityHeaders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityHeaderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityHeaderUpdateParams">ZoneSettingSecurityHeaderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityHeaderUpdateResponse">ZoneSettingSecurityHeaderUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/security_header">client.Zones.Settings.SecurityHeaders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityHeaderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityHeaderListResponse">ZoneSettingSecurityHeaderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### SecurityLevels

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityLevelUpdateResponse">ZoneSettingSecurityLevelUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityLevelListResponse">ZoneSettingSecurityLevelListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/security_level">client.Zones.Settings.SecurityLevels.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityLevelService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityLevelUpdateParams">ZoneSettingSecurityLevelUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityLevelUpdateResponse">ZoneSettingSecurityLevelUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/security_level">client.Zones.Settings.SecurityLevels.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityLevelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSecurityLevelListResponse">ZoneSettingSecurityLevelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ServerSideExcludes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingServerSideExcludeUpdateResponse">ZoneSettingServerSideExcludeUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingServerSideExcludeListResponse">ZoneSettingServerSideExcludeListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/server_side_exclude">client.Zones.Settings.ServerSideExcludes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingServerSideExcludeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingServerSideExcludeUpdateParams">ZoneSettingServerSideExcludeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingServerSideExcludeUpdateResponse">ZoneSettingServerSideExcludeUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/server_side_exclude">client.Zones.Settings.ServerSideExcludes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingServerSideExcludeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingServerSideExcludeListResponse">ZoneSettingServerSideExcludeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### SortQueryStringForCaches

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSortQueryStringForCachUpdateResponse">ZoneSettingSortQueryStringForCachUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSortQueryStringForCachListResponse">ZoneSettingSortQueryStringForCachListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/sort_query_string_for_cache">client.Zones.Settings.SortQueryStringForCaches.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSortQueryStringForCachService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSortQueryStringForCachUpdateParams">ZoneSettingSortQueryStringForCachUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSortQueryStringForCachUpdateResponse">ZoneSettingSortQueryStringForCachUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/sort_query_string_for_cache">client.Zones.Settings.SortQueryStringForCaches.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSortQueryStringForCachService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSortQueryStringForCachListResponse">ZoneSettingSortQueryStringForCachListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### SSLs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLUpdateResponse">ZoneSettingSSLUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLListResponse">ZoneSettingSSLListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/ssl">client.Zones.Settings.SSLs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLUpdateParams">ZoneSettingSSLUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLUpdateResponse">ZoneSettingSSLUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/ssl">client.Zones.Settings.SSLs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLListResponse">ZoneSettingSSLListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### SSLRecommenders

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLRecommenderUpdateResponse">ZoneSettingSSLRecommenderUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLRecommenderListResponse">ZoneSettingSSLRecommenderListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/ssl_recommender">client.Zones.Settings.SSLRecommenders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLRecommenderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLRecommenderUpdateParams">ZoneSettingSSLRecommenderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLRecommenderUpdateResponse">ZoneSettingSSLRecommenderUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/ssl_recommender">client.Zones.Settings.SSLRecommenders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLRecommenderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingSSLRecommenderListResponse">ZoneSettingSSLRecommenderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### TLS1_3s

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTls1_3UpdateResponse">ZoneSettingTls1_3UpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTls1_3ListResponse">ZoneSettingTls1_3ListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/tls_1_3">client.Zones.Settings.TLS1_3s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLS1_3Service.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLS1_3UpdateParams">ZoneSettingTLS1_3UpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTls1_3UpdateResponse">ZoneSettingTls1_3UpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/tls_1_3">client.Zones.Settings.TLS1_3s.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLS1_3Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTls1_3ListResponse">ZoneSettingTls1_3ListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### TLSClientAuths

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLSClientAuthUpdateResponse">ZoneSettingTLSClientAuthUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLSClientAuthListResponse">ZoneSettingTLSClientAuthListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/tls_client_auth">client.Zones.Settings.TLSClientAuths.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLSClientAuthService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLSClientAuthUpdateParams">ZoneSettingTLSClientAuthUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLSClientAuthUpdateResponse">ZoneSettingTLSClientAuthUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/tls_client_auth">client.Zones.Settings.TLSClientAuths.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLSClientAuthService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTLSClientAuthListResponse">ZoneSettingTLSClientAuthListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### TrueClientIPHeaders

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTrueClientIPHeaderUpdateResponse">ZoneSettingTrueClientIPHeaderUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTrueClientIPHeaderListResponse">ZoneSettingTrueClientIPHeaderListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/true_client_ip_header">client.Zones.Settings.TrueClientIPHeaders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTrueClientIPHeaderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTrueClientIPHeaderUpdateParams">ZoneSettingTrueClientIPHeaderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTrueClientIPHeaderUpdateResponse">ZoneSettingTrueClientIPHeaderUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/true_client_ip_header">client.Zones.Settings.TrueClientIPHeaders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTrueClientIPHeaderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingTrueClientIPHeaderListResponse">ZoneSettingTrueClientIPHeaderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### WAFs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWAFUpdateResponse">ZoneSettingWAFUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWAFListResponse">ZoneSettingWAFListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/waf">client.Zones.Settings.WAFs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWAFService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWAFUpdateParams">ZoneSettingWAFUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWAFUpdateResponse">ZoneSettingWAFUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/waf">client.Zones.Settings.WAFs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWAFService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWAFListResponse">ZoneSettingWAFListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Webps

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebpUpdateResponse">ZoneSettingWebpUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebpListResponse">ZoneSettingWebpListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/webp">client.Zones.Settings.Webps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebpService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebpUpdateParams">ZoneSettingWebpUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebpUpdateResponse">ZoneSettingWebpUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/webp">client.Zones.Settings.Webps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebpService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebpListResponse">ZoneSettingWebpListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Websockets

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebsocketUpdateResponse">ZoneSettingWebsocketUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebsocketListResponse">ZoneSettingWebsocketListResponse</a>

Methods:

- <code title="patch /zones/{zone_identifier}/settings/websockets">client.Zones.Settings.Websockets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebsocketService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebsocketUpdateParams">ZoneSettingWebsocketUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebsocketUpdateResponse">ZoneSettingWebsocketUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/settings/websockets">client.Zones.Settings.Websockets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebsocketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneSettingWebsocketListResponse">ZoneSettingWebsocketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AI

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIRunResponse">AIRunResponse</a>

Methods:

- <code title="post /apiv4/accounts/{account_identifier}/ai/run/{model_name}">client.AI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIRunParams">AIRunParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIRunResponse">AIRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Huggingface

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIHuggingfaceDistilbertSst2Int8Response">AIHuggingfaceDistilbertSst2Int8Response</a>

Methods:

- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/huggingface/distilbert-sst-2-int8">client.AI.Huggingface.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIHuggingfaceService.DistilbertSst2Int8">DistilbertSst2Int8</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIHuggingfaceDistilbertSst2Int8Params">AIHuggingfaceDistilbertSst2Int8Params</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIHuggingfaceDistilbertSst2Int8Response">AIHuggingfaceDistilbertSst2Int8Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Baai

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeBaseEnV1_5Response">AIBaaiBgeBaseEnV1_5Response</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeLargeEnV1_5Response">AIBaaiBgeLargeEnV1_5Response</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeSmallEnV1_5Response">AIBaaiBgeSmallEnV1_5Response</a>

Methods:

- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/baai/bge-base-en-v1.5">client.AI.Baai.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiService.BgeBaseEnV1_5">BgeBaseEnV1_5</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeBaseEnV1_5Params">AIBaaiBgeBaseEnV1_5Params</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeBaseEnV1_5Response">AIBaaiBgeBaseEnV1_5Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/baai/bge-large-en-v1.5">client.AI.Baai.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiService.BgeLargeEnV1_5">BgeLargeEnV1_5</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeLargeEnV1_5Params">AIBaaiBgeLargeEnV1_5Params</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeLargeEnV1_5Response">AIBaaiBgeLargeEnV1_5Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/baai/bge-small-en-v1.5">client.AI.Baai.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiService.BgeSmallEnV1_5">BgeSmallEnV1_5</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeSmallEnV1_5Params">AIBaaiBgeSmallEnV1_5Params</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIBaaiBgeSmallEnV1_5Response">AIBaaiBgeSmallEnV1_5Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OpenAI

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIOpenAIWhisperResponse">AIOpenAIWhisperResponse</a>

Methods:

- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/openai/whisper">client.AI.OpenAI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIOpenAIService.Whisper">Whisper</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIOpenAIWhisperResponse">AIOpenAIWhisperResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Microsoft

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMicrosoftResnet50Response">AIMicrosoftResnet50Response</a>

Methods:

- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/microsoft/resnet-50">client.AI.Microsoft.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMicrosoftService.Resnet50">Resnet50</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMicrosoftResnet50Response">AIMicrosoftResnet50Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Meta

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaLlama2_7bChatFp16Response">AIMetaLlama2_7bChatFp16Response</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaLlama2_7bChatInt8Response">AIMetaLlama2_7bChatInt8Response</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaM2m100_1_2bResponse">AIMetaM2m100_1_2bResponse</a>

Methods:

- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/meta/llama-2-7b-chat-fp16">client.AI.Meta.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaService.Llama2_7bChatFp16">Llama2_7bChatFp16</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaLlama2_7bChatFp16Params">AIMetaLlama2_7bChatFp16Params</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaLlama2_7bChatFp16Response">AIMetaLlama2_7bChatFp16Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/meta/llama-2-7b-chat-int8">client.AI.Meta.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaService.Llama2_7bChatInt8">Llama2_7bChatInt8</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaLlama2_7bChatInt8Params">AIMetaLlama2_7bChatInt8Params</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaLlama2_7bChatInt8Response">AIMetaLlama2_7bChatInt8Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/meta/m2m100-1.2b">client.AI.Meta.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaService.M2m100_1_2b">M2m100_1_2b</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaM2m100_1_2bParams">AIMetaM2m100_1_2bParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMetaM2m100_1_2bResponse">AIMetaM2m100_1_2bResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Mistral

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMistralMistral7bInstructV0_1Response">AIMistralMistral7bInstructV0_1Response</a>

Methods:

- <code title="post /apiv4/accounts/{account_identifier}/ai/run/@cf/mistral/mistral-7b-instruct-v0.1">client.AI.Mistral.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMistralService.Mistral7bInstructV0_1">Mistral7bInstructV0_1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMistralMistral7bInstructV0_1Params">AIMistralMistral7bInstructV0_1Params</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIMistralMistral7bInstructV0_1Response">AIMistralMistral7bInstructV0_1Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
