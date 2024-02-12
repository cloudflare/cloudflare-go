# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountGetResponse">AccountGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountUpdateResponse">AccountUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountListResponse">AccountListResponse</a>

Methods:

- <code title="get /accounts/{identifier}">client.Accounts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier interface{}) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountGetResponse">AccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{identifier}">client.Accounts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier interface{}, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountUpdateParams">AccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountUpdateResponse">AccountUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountListParams">AccountListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# IPs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IPListResponse">IPListResponse</a>

Methods:

- <code title="get /ips">client.IPs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IPService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IPListParams">IPListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IPListResponse">IPListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Zones

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneNewResponse">ZoneNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneListResponse">ZoneListResponse</a>

Methods:

- <code title="post /zones">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneNewParams">ZoneNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneNewResponse">ZoneNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneListParams">ZoneListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZoneListResponse">ZoneListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AI

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIRunModelResponse">AIRunModelResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/ai/run/{model_name}">client.AI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIService.RunModel">RunModel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIRunModelParams">AIRunModelParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AIRunModelResponse">AIRunModelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Huggingface

## Baai

## OpenAI

## Microsoft

## Meta

## Mistral

# LoadBalancers

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerNewResponse">LoadBalancerNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerGetResponse">LoadBalancerGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerUpdateResponse">LoadBalancerUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerListResponse">LoadBalancerListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerDeleteResponse">LoadBalancerDeleteResponse</a>

Methods:

- <code title="post /zones/{identifier}/load_balancers">client.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerNewParams">LoadBalancerNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerNewResponse">LoadBalancerNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{identifier1}/load_balancers/{identifier}">client.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier1 <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerGetResponse">LoadBalancerGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{identifier1}/load_balancers/{identifier}">client.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier1 <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerUpdateParams">LoadBalancerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerUpdateResponse">LoadBalancerUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{identifier}/load_balancers">client.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerListResponse">LoadBalancerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{identifier1}/load_balancers/{identifier}">client.LoadBalancers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier1 <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerDeleteResponse">LoadBalancerDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Monitors

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorGetResponse">LoadBalancerMonitorGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorUpdateResponse">LoadBalancerMonitorUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorDeleteResponse">LoadBalancerMonitorDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse">LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse">LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/load_balancers/monitors/{identifier}">client.LoadBalancers.Monitors.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorGetResponse">LoadBalancerMonitorGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/load_balancers/monitors/{identifier}">client.LoadBalancers.Monitors.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorUpdateParams">LoadBalancerMonitorUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorUpdateResponse">LoadBalancerMonitorUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/load_balancers/monitors/{identifier}">client.LoadBalancers.Monitors.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorDeleteResponse">LoadBalancerMonitorDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/load_balancers/monitors">client.LoadBalancers.Monitors.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorService.AccountLoadBalancerMonitorsNewMonitor">AccountLoadBalancerMonitorsNewMonitor</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams">LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse">LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/load_balancers/monitors">client.LoadBalancers.Monitors.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorService.AccountLoadBalancerMonitorsListMonitors">AccountLoadBalancerMonitorsListMonitors</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse">LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Previews

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorResponse">LoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/load_balancers/monitors/{identifier}/preview">client.LoadBalancers.Monitors.Previews.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorPreviewService.AccountLoadBalancerMonitorsPreviewMonitor">AccountLoadBalancerMonitorsPreviewMonitor</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParams">LoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorResponse">LoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### References

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse">LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/load_balancers/monitors/{identifier}/references">client.LoadBalancers.Monitors.References.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorReferenceService.AccountLoadBalancerMonitorsListMonitorReferences">AccountLoadBalancerMonitorsListMonitorReferences</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse">LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Pools

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolGetResponse">LoadBalancerPoolGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolUpdateResponse">LoadBalancerPoolUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolDeleteResponse">LoadBalancerPoolDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse">LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse">LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse">LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/load_balancers/pools/{identifier}">client.LoadBalancers.Pools.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolGetResponse">LoadBalancerPoolGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/load_balancers/pools/{identifier}">client.LoadBalancers.Pools.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolUpdateParams">LoadBalancerPoolUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolUpdateResponse">LoadBalancerPoolUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/load_balancers/pools/{identifier}">client.LoadBalancers.Pools.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolDeleteResponse">LoadBalancerPoolDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/load_balancers/pools">client.LoadBalancers.Pools.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolService.AccountLoadBalancerPoolsNewPool">AccountLoadBalancerPoolsNewPool</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams">LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse">LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/load_balancers/pools">client.LoadBalancers.Pools.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolService.AccountLoadBalancerPoolsListPools">AccountLoadBalancerPoolsListPools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams">LoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse">LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_identifier}/load_balancers/pools">client.LoadBalancers.Pools.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolService.AccountLoadBalancerPoolsPatchPools">AccountLoadBalancerPoolsPatchPools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams">LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse">LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Health

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse">LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/load_balancers/pools/{identifier}/health">client.LoadBalancers.Pools.Health.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolHealthService.AccountLoadBalancerPoolsPoolHealthDetails">AccountLoadBalancerPoolsPoolHealthDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse">LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Previews

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolPreviewAccountLoadBalancerPoolsPreviewPoolResponse">LoadBalancerPoolPreviewAccountLoadBalancerPoolsPreviewPoolResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/load_balancers/pools/{identifier}/preview">client.LoadBalancers.Pools.Previews.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolPreviewService.AccountLoadBalancerPoolsPreviewPool">AccountLoadBalancerPoolsPreviewPool</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolPreviewAccountLoadBalancerPoolsPreviewPoolParams">LoadBalancerPoolPreviewAccountLoadBalancerPoolsPreviewPoolParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolPreviewAccountLoadBalancerPoolsPreviewPoolResponse">LoadBalancerPoolPreviewAccountLoadBalancerPoolsPreviewPoolResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### References

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse">LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/load_balancers/pools/{identifier}/references">client.LoadBalancers.Pools.References.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolReferenceService.AccountLoadBalancerPoolsListPoolReferences">AccountLoadBalancerPoolsListPoolReferences</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse">LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Previews

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPreviewGetResponse">LoadBalancerPreviewGetResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/load_balancers/preview/{preview_id}">client.LoadBalancers.Previews.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPreviewService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, previewID interface{}) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerPreviewGetResponse">LoadBalancerPreviewGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Regions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerRegionGetResponse">LoadBalancerRegionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerRegionLoadBalancerRegionsListRegionsResponse">LoadBalancerRegionLoadBalancerRegionsListRegionsResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/load_balancers/regions/{region_code}">client.LoadBalancers.Regions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerRegionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, regionCode <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerRegionGetParamsRegionCode">LoadBalancerRegionGetParamsRegionCode</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerRegionGetResponse">LoadBalancerRegionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/load_balancers/regions">client.LoadBalancers.Regions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerRegionService.LoadBalancerRegionsListRegions">LoadBalancerRegionsListRegions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerRegionLoadBalancerRegionsListRegionsParams">LoadBalancerRegionLoadBalancerRegionsListRegionsParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerRegionLoadBalancerRegionsListRegionsResponse">LoadBalancerRegionLoadBalancerRegionsListRegionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Searches

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerSearchListResponse">LoadBalancerSearchListResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/load_balancers/search">client.LoadBalancers.Searches.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerSearchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerSearchListParams">LoadBalancerSearchListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LoadBalancerSearchListResponse">LoadBalancerSearchListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Access

## Apps

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppGetResponse">AccessAppGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppUpdateResponse">AccessAppUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppDeleteResponse">AccessAppDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppAccessApplicationsAddAnApplicationResponse">AccessAppAccessApplicationsAddAnApplicationResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppAccessApplicationsListAccessApplicationsResponse">AccessAppAccessApplicationsListAccessApplicationsResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/access/apps/{app_id}">client.Access.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, appID <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppGetParamsAppID">AccessAppGetParamsAppID</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppGetResponse">AccessAppGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{account_or_zone}/{account_or_zone_id}/access/apps/{app_id}">client.Access.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, appID <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppUpdateParamsVariant0AppID">AccessAppUpdateParamsVariant0AppID</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppUpdateParams">AccessAppUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppUpdateResponse">AccessAppUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{account_or_zone}/{account_or_zone_id}/access/apps/{app_id}">client.Access.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, appID <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppDeleteParamsAppID">AccessAppDeleteParamsAppID</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppDeleteResponse">AccessAppDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{account_or_zone}/{account_or_zone_id}/access/apps">client.Access.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppService.AccessApplicationsAddAnApplication">AccessApplicationsAddAnApplication</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppAccessApplicationsAddAnApplicationParams">AccessAppAccessApplicationsAddAnApplicationParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppAccessApplicationsAddAnApplicationResponse">AccessAppAccessApplicationsAddAnApplicationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/apps">client.Access.Apps.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppService.AccessApplicationsListAccessApplications">AccessApplicationsListAccessApplications</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppAccessApplicationsListAccessApplicationsResponse">AccessAppAccessApplicationsListAccessApplicationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Cas

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaDeleteResponse">AccessAppCaDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse">AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse">AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse">AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse</a>

Methods:

- <code title="delete /{account_or_zone}/{account_or_zone_id}/access/apps/{uuid}/ca">client.Access.Apps.Cas.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaDeleteResponse">AccessAppCaDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{account_or_zone}/{account_or_zone_id}/access/apps/{uuid}/ca">client.Access.Apps.Cas.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaService.AccessShortLivedCertificateCAsNewAShortLivedCertificateCa">AccessShortLivedCertificateCAsNewAShortLivedCertificateCa</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse">AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/apps/{uuid}/ca">client.Access.Apps.Cas.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaService.AccessShortLivedCertificateCAsGetAShortLivedCertificateCa">AccessShortLivedCertificateCAsGetAShortLivedCertificateCa</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse">AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/apps/ca">client.Access.Apps.Cas.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaService.AccessShortLivedCertificateCAsListShortLivedCertificateCAs">AccessShortLivedCertificateCAsListShortLivedCertificateCAs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse">AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### RevokeTokens

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponse">AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponse</a>

Methods:

- <code title="post /{account_or_zone}/{account_or_zone_id}/access/apps/{app_id}/revoke_tokens">client.Access.Apps.RevokeTokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppRevokeTokenService.AccessApplicationsRevokeServiceTokens">AccessApplicationsRevokeServiceTokens</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, appID <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensParamsAppID">AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensParamsAppID</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponse">AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### UserPolicyChecks

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse">AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/access/apps/{app_id}/user_policy_checks">client.Access.Apps.UserPolicyChecks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppUserPolicyCheckService.AccessApplicationsTestAccessPolicies">AccessApplicationsTestAccessPolicies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, appID <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesParamsAppID">AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesParamsAppID</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse">AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Policies

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyGetResponse">AccessAppPolicyGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyUpdateResponse">AccessAppPolicyUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyDeleteResponse">AccessAppPolicyDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse">AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyAccessPoliciesListAccessPoliciesResponse">AccessAppPolicyAccessPoliciesListAccessPoliciesResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/access/apps/{uuid1}/policies/{uuid}">client.Access.Apps.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid1 <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyGetResponse">AccessAppPolicyGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{account_or_zone}/{account_or_zone_id}/access/apps/{uuid1}/policies/{uuid}">client.Access.Apps.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid1 <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyUpdateParams">AccessAppPolicyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyUpdateResponse">AccessAppPolicyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{account_or_zone}/{account_or_zone_id}/access/apps/{uuid1}/policies/{uuid}">client.Access.Apps.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid1 <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyDeleteResponse">AccessAppPolicyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{account_or_zone}/{account_or_zone_id}/access/apps/{uuid}/policies">client.Access.Apps.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyService.AccessPoliciesNewAnAccessPolicy">AccessPoliciesNewAnAccessPolicy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyAccessPoliciesNewAnAccessPolicyParams">AccessAppPolicyAccessPoliciesNewAnAccessPolicyParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse">AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/apps/{uuid}/policies">client.Access.Apps.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyService.AccessPoliciesListAccessPolicies">AccessPoliciesListAccessPolicies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessAppPolicyAccessPoliciesListAccessPoliciesResponse">AccessAppPolicyAccessPoliciesListAccessPoliciesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Certificates

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateGetResponse">AccessCertificateGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateUpdateResponse">AccessCertificateUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateDeleteResponse">AccessCertificateDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse">AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse">AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/access/certificates/{uuid}">client.Access.Certificates.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateGetResponse">AccessCertificateGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{account_or_zone}/{account_or_zone_id}/access/certificates/{uuid}">client.Access.Certificates.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateUpdateParams">AccessCertificateUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateUpdateResponse">AccessCertificateUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{account_or_zone}/{account_or_zone_id}/access/certificates/{uuid}">client.Access.Certificates.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateDeleteResponse">AccessCertificateDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{account_or_zone}/{account_or_zone_id}/access/certificates">client.Access.Certificates.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateService.AccessMTLSAuthenticationAddAnMTLSCertificate">AccessMTLSAuthenticationAddAnMTLSCertificate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateAccessMTLSAuthenticationAddAnMTLSCertificateParams">AccessCertificateAccessMTLSAuthenticationAddAnMTLSCertificateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse">AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/certificates">client.Access.Certificates.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateService.AccessMTLSAuthenticationListMTLSCertificates">AccessMTLSAuthenticationListMTLSCertificates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse">AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Settings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateSettingUpdateResponse">AccessCertificateSettingUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateSettingListResponse">AccessCertificateSettingListResponse</a>

Methods:

- <code title="put /{account_or_zone}/{account_or_zone_id}/access/certificates/settings">client.Access.Certificates.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateSettingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateSettingUpdateParams">AccessCertificateSettingUpdateParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateSettingUpdateResponse">AccessCertificateSettingUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/certificates/settings">client.Access.Certificates.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateSettingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCertificateSettingListResponse">AccessCertificateSettingListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Groups

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupGetResponse">AccessGroupGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupUpdateResponse">AccessGroupUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupDeleteResponse">AccessGroupDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupAccessGroupsNewAnAccessGroupResponse">AccessGroupAccessGroupsNewAnAccessGroupResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupAccessGroupsListAccessGroupsResponse">AccessGroupAccessGroupsListAccessGroupsResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/access/groups/{uuid}">client.Access.Groups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupGetResponse">AccessGroupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{account_or_zone}/{account_or_zone_id}/access/groups/{uuid}">client.Access.Groups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupUpdateParams">AccessGroupUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupUpdateResponse">AccessGroupUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{account_or_zone}/{account_or_zone_id}/access/groups/{uuid}">client.Access.Groups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupDeleteResponse">AccessGroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{account_or_zone}/{account_or_zone_id}/access/groups">client.Access.Groups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupService.AccessGroupsNewAnAccessGroup">AccessGroupsNewAnAccessGroup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupAccessGroupsNewAnAccessGroupParams">AccessGroupAccessGroupsNewAnAccessGroupParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupAccessGroupsNewAnAccessGroupResponse">AccessGroupAccessGroupsNewAnAccessGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/groups">client.Access.Groups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupService.AccessGroupsListAccessGroups">AccessGroupsListAccessGroups</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessGroupAccessGroupsListAccessGroupsResponse">AccessGroupAccessGroupsListAccessGroupsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## IdentityProviders

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderGetResponse">AccessIdentityProviderGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderUpdateResponse">AccessIdentityProviderUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderDeleteResponse">AccessIdentityProviderDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse">AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse">AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/access/identity_providers/{uuid}">client.Access.IdentityProviders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderGetResponse">AccessIdentityProviderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{account_or_zone}/{account_or_zone_id}/access/identity_providers/{uuid}">client.Access.IdentityProviders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderUpdateParams">AccessIdentityProviderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderUpdateResponse">AccessIdentityProviderUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{account_or_zone}/{account_or_zone_id}/access/identity_providers/{uuid}">client.Access.IdentityProviders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderDeleteResponse">AccessIdentityProviderDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{account_or_zone}/{account_or_zone_id}/access/identity_providers">client.Access.IdentityProviders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderService.AccessIdentityProvidersAddAnAccessIdentityProvider">AccessIdentityProvidersAddAnAccessIdentityProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams">AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse">AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/identity_providers">client.Access.IdentityProviders.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderService.AccessIdentityProvidersListAccessIdentityProviders">AccessIdentityProvidersListAccessIdentityProviders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse">AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Organizations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse">AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse">AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse">AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse</a>

Methods:

- <code title="post /{account_or_zone}/{account_or_zone_id}/access/organizations">client.Access.Organizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationService.ZeroTrustOrganizationNewYourZeroTrustOrganization">ZeroTrustOrganizationNewYourZeroTrustOrganization</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams">AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse">AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/organizations">client.Access.Organizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationService.ZeroTrustOrganizationGetYourZeroTrustOrganization">ZeroTrustOrganizationGetYourZeroTrustOrganization</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse">AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{account_or_zone}/{account_or_zone_id}/access/organizations">client.Access.Organizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationService.ZeroTrustOrganizationUpdateYourZeroTrustOrganization">ZeroTrustOrganizationUpdateYourZeroTrustOrganization</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams">AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse">AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ServiceTokens

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenUpdateResponse">AccessServiceTokenUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenDeleteResponse">AccessServiceTokenDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse">AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenAccessServiceTokensListServiceTokensResponse">AccessServiceTokenAccessServiceTokensListServiceTokensResponse</a>

Methods:

- <code title="put /{account_or_zone}/{account_or_zone_id}/access/service_tokens/{uuid}">client.Access.ServiceTokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenUpdateParams">AccessServiceTokenUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenUpdateResponse">AccessServiceTokenUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{account_or_zone}/{account_or_zone_id}/access/service_tokens/{uuid}">client.Access.ServiceTokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenDeleteResponse">AccessServiceTokenDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{account_or_zone}/{account_or_zone_id}/access/service_tokens">client.Access.ServiceTokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenService.AccessServiceTokensNewAServiceToken">AccessServiceTokensNewAServiceToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenAccessServiceTokensNewAServiceTokenParams">AccessServiceTokenAccessServiceTokensNewAServiceTokenParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse">AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{account_or_zone}/{account_or_zone_id}/access/service_tokens">client.Access.ServiceTokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenService.AccessServiceTokensListServiceTokens">AccessServiceTokensListServiceTokens</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenAccessServiceTokensListServiceTokensResponse">AccessServiceTokenAccessServiceTokensListServiceTokensResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Refreshes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse">AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse</a>

Methods:

- <code title="post /accounts/{identifier}/access/service_tokens/{uuid}/refresh">client.Access.ServiceTokens.Refreshes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenRefreshService.AccessServiceTokensRefreshAServiceToken">AccessServiceTokensRefreshAServiceToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse">AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Rotates

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse">AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse</a>

Methods:

- <code title="post /accounts/{identifier}/access/service_tokens/{uuid}/rotate">client.Access.ServiceTokens.Rotates.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenRotateService.AccessServiceTokensRotateAServiceToken">AccessServiceTokensRotateAServiceToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse">AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Bookmarks

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkGetResponse">AccessBookmarkGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkUpdateResponse">AccessBookmarkUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkDeleteResponse">AccessBookmarkDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse">AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse</a>

Methods:

- <code title="get /accounts/{identifier}/access/bookmarks/{uuid}">client.Access.Bookmarks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier interface{}, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkGetResponse">AccessBookmarkGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{identifier}/access/bookmarks/{uuid}">client.Access.Bookmarks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier interface{}, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkUpdateResponse">AccessBookmarkUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{identifier}/access/bookmarks/{uuid}">client.Access.Bookmarks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier interface{}, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkDeleteResponse">AccessBookmarkDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{identifier}/access/bookmarks">client.Access.Bookmarks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkService.AccessBookmarkApplicationsDeprecatedListBookmarkApplications">AccessBookmarkApplicationsDeprecatedListBookmarkApplications</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier interface{}) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse">AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Keys

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse">AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse">AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse</a>

Methods:

- <code title="get /accounts/{identifier}/access/keys">client.Access.Keys.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyService.AccessKeyConfigurationGetTheAccessKeyConfiguration">AccessKeyConfigurationGetTheAccessKeyConfiguration</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse">AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{identifier}/access/keys">client.Access.Keys.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyService.AccessKeyConfigurationUpdateTheAccessKeyConfiguration">AccessKeyConfigurationUpdateTheAccessKeyConfiguration</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams">AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse">AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Rotates

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse">AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse</a>

Methods:

- <code title="post /accounts/{identifier}/access/keys/rotate">client.Access.Keys.Rotates.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyRotateService.AccessKeyConfigurationRotateAccessKeys">AccessKeyConfigurationRotateAccessKeys</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse">AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

### AccessRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse">AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse</a>

Methods:

- <code title="get /accounts/{identifier}/access/logs/access_requests">client.Access.Logs.AccessRequests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessLogAccessRequestService.AccessAuthenticationLogsGetAccessAuthenticationLogs">AccessAuthenticationLogsGetAccessAuthenticationLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse">AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Seats

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessSeatZeroTrustSeatsUpdateAUserSeatResponse">AccessSeatZeroTrustSeatsUpdateAUserSeatResponse</a>

Methods:

- <code title="patch /accounts/{identifier}/access/seats">client.Access.Seats.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessSeatService.ZeroTrustSeatsUpdateAUserSeat">ZeroTrustSeatsUpdateAUserSeat</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessSeatZeroTrustSeatsUpdateAUserSeatParams">AccessSeatZeroTrustSeatsUpdateAUserSeatParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessSeatZeroTrustSeatsUpdateAUserSeatResponse">AccessSeatZeroTrustSeatsUpdateAUserSeatResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserListResponse">AccessUserListResponse</a>

Methods:

- <code title="get /accounts/{identifier}/access/users">client.Access.Users.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserListResponse">AccessUserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ActiveSessions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserActiveSessionGetResponse">AccessUserActiveSessionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserActiveSessionListResponse">AccessUserActiveSessionListResponse</a>

Methods:

- <code title="get /accounts/{identifier}/access/users/{id}/active_sessions/{nonce}">client.Access.Users.ActiveSessions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserActiveSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, nonce <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserActiveSessionGetResponse">AccessUserActiveSessionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{identifier}/access/users/{id}/active_sessions">client.Access.Users.ActiveSessions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserActiveSessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserActiveSessionListResponse">AccessUserActiveSessionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### LastSeenIdentity

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserLastSeenIdentityGetResponse">AccessUserLastSeenIdentityGetResponse</a>

Methods:

- <code title="get /accounts/{identifier}/access/users/{id}/last_seen_identity">client.Access.Users.LastSeenIdentity.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserLastSeenIdentityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserLastSeenIdentityGetResponse">AccessUserLastSeenIdentityGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### FailedLogins

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse">AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse</a>

Methods:

- <code title="get /accounts/{identifier}/access/users/{id}/failed_logins">client.Access.Users.FailedLogins.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserFailedLoginService.ZeroTrustUsersGetFailedLogins">ZeroTrustUsersGetFailedLogins</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse">AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CustomPages

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageNewResponse">AccessCustomPageNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageGetResponse">AccessCustomPageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageUpdateResponse">AccessCustomPageUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageListResponse">AccessCustomPageListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageDeleteResponse">AccessCustomPageDeleteResponse</a>

Methods:

- <code title="post /accounts/{identifier}/access/custom_pages">client.Access.CustomPages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageNewParams">AccessCustomPageNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageNewResponse">AccessCustomPageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{identifier}/access/custom_pages/{uuid}">client.Access.CustomPages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageGetResponse">AccessCustomPageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{identifier}/access/custom_pages/{uuid}">client.Access.CustomPages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageUpdateParams">AccessCustomPageUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageUpdateResponse">AccessCustomPageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{identifier}/access/custom_pages">client.Access.CustomPages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageListResponse">AccessCustomPageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{identifier}/access/custom_pages/{uuid}">client.Access.CustomPages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, uuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessCustomPageDeleteResponse">AccessCustomPageDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tags

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagListResponse">AccessTagListResponse</a>

Methods:

- <code title="get /accounts/{identifier}/access/tags">client.Access.Tags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagListResponse">AccessTagListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DNSRecords

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordGetResponse">DNSRecordGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordUpdateResponse">DNSRecordUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordDeleteResponse">DNSRecordDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordDNSRecordsForAZoneNewDNSRecordResponse">DNSRecordDNSRecordsForAZoneNewDNSRecordResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordDNSRecordsForAZoneListDNSRecordsResponse">DNSRecordDNSRecordsForAZoneListDNSRecordsResponse</a>

Methods:

- <code title="get /zones/{zone_id}/dns_records/{dns_record_id}">client.DNSRecords.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, dnsRecordID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordGetResponse">DNSRecordGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_id}/dns_records/{dns_record_id}">client.DNSRecords.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, dnsRecordID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordUpdateParams">DNSRecordUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordUpdateResponse">DNSRecordUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/dns_records/{dns_record_id}">client.DNSRecords.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, dnsRecordID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordDeleteResponse">DNSRecordDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /zones/{zone_id}/dns_records">client.DNSRecords.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordService.DNSRecordsForAZoneNewDNSRecord">DNSRecordsForAZoneNewDNSRecord</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordDNSRecordsForAZoneNewDNSRecordParams">DNSRecordDNSRecordsForAZoneNewDNSRecordParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordDNSRecordsForAZoneNewDNSRecordResponse">DNSRecordDNSRecordsForAZoneNewDNSRecordResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/dns_records">client.DNSRecords.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordService.DNSRecordsForAZoneListDNSRecords">DNSRecordsForAZoneListDNSRecords</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordDNSRecordsForAZoneListDNSRecordsParams">DNSRecordDNSRecordsForAZoneListDNSRecordsParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordDNSRecordsForAZoneListDNSRecordsResponse">DNSRecordDNSRecordsForAZoneListDNSRecordsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Exports

Methods:

- <code title="get /zones/{zone_id}/dns_records/export">client.DNSRecords.Exports.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordExportService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Imports

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse">DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse</a>

Methods:

- <code title="post /zones/{zone_id}/dns_records/import">client.DNSRecords.Imports.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordImportService.DNSRecordsForAZoneImportDNSRecords">DNSRecordsForAZoneImportDNSRecords</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordImportDNSRecordsForAZoneImportDNSRecordsParams">DNSRecordImportDNSRecordsForAZoneImportDNSRecordsParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse">DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Scans

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse">DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse</a>

Methods:

- <code title="post /zones/{zone_id}/dns_records/scan">client.DNSRecords.Scans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordScanService.DNSRecordsForAZoneScanDNSRecords">DNSRecordsForAZoneScanDNSRecords</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse">DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Emails

## Routings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse">EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/email/routing">client.Emails.Routings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingService.EmailRoutingSettingsGetEmailRoutingSettings">EmailRoutingSettingsGetEmailRoutingSettings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse">EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Disables

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse">EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse</a>

Methods:

- <code title="post /zones/{zone_identifier}/email/routing/disable">client.Emails.Routings.Disables.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingDisableService.EmailRoutingSettingsDisableEmailRouting">EmailRoutingSettingsDisableEmailRouting</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse">EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### DNS

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse">EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/email/routing/dns">client.Emails.Routings.DNS.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingDNSService.EmailRoutingSettingsEmailRoutingDNSSettings">EmailRoutingSettingsEmailRoutingDNSSettings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse">EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Enables

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse">EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse</a>

Methods:

- <code title="post /zones/{zone_identifier}/email/routing/enable">client.Emails.Routings.Enables.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingEnableService.EmailRoutingSettingsEnableEmailRouting">EmailRoutingSettingsEnableEmailRouting</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse">EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Rules

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleGetResponse">EmailRoutingRuleGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleUpdateResponse">EmailRoutingRuleUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleDeleteResponse">EmailRoutingRuleDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse">EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse">EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/email/routing/rules/{rule_identifier}">client.Emails.Routings.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, ruleIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleGetResponse">EmailRoutingRuleGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_identifier}/email/routing/rules/{rule_identifier}">client.Emails.Routings.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, ruleIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleUpdateParams">EmailRoutingRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleUpdateResponse">EmailRoutingRuleUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_identifier}/email/routing/rules/{rule_identifier}">client.Emails.Routings.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, ruleIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleDeleteResponse">EmailRoutingRuleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /zones/{zone_identifier}/email/routing/rules">client.Emails.Routings.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleService.EmailRoutingRoutingRulesNewRoutingRule">EmailRoutingRoutingRulesNewRoutingRule</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams">EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse">EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/email/routing/rules">client.Emails.Routings.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleService.EmailRoutingRoutingRulesListRoutingRules">EmailRoutingRoutingRulesListRoutingRules</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams">EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse">EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### CatchAlls

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse">EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse">EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/email/routing/rules/catch_all">client.Emails.Routings.Rules.CatchAlls.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleCatchAllService.EmailRoutingRoutingRulesGetCatchAllRule">EmailRoutingRoutingRulesGetCatchAllRule</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse">EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_identifier}/email/routing/rules/catch_all">client.Emails.Routings.Rules.CatchAlls.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleCatchAllService.EmailRoutingRoutingRulesUpdateCatchAllRule">EmailRoutingRoutingRulesUpdateCatchAllRule</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams">EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse">EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Addresses

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressGetResponse">EmailRoutingAddressGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressDeleteResponse">EmailRoutingAddressDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse">EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse">EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/email/routing/addresses/{destination_address_identifier}">client.Emails.Routings.Addresses.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, destinationAddressIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressGetResponse">EmailRoutingAddressGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/email/routing/addresses/{destination_address_identifier}">client.Emails.Routings.Addresses.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, destinationAddressIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressDeleteResponse">EmailRoutingAddressDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/email/routing/addresses">client.Emails.Routings.Addresses.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressService.EmailRoutingDestinationAddressesNewADestinationAddress">EmailRoutingDestinationAddressesNewADestinationAddress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams">EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse">EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/email/routing/addresses">client.Emails.Routings.Addresses.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressService.EmailRoutingDestinationAddressesListDestinationAddresses">EmailRoutingDestinationAddressesListDestinationAddresses</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams">EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse">EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccountMembers

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberGetResponse">AccountMemberGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberUpdateResponse">AccountMemberUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberDeleteResponse">AccountMemberDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberAccountMembersAddMemberResponse">AccountMemberAccountMembersAddMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberAccountMembersListMembersResponse">AccountMemberAccountMembersListMembersResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/members/{identifier}">client.AccountMembers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier interface{}, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberGetResponse">AccountMemberGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/members/{identifier}">client.AccountMembers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier interface{}, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberUpdateParams">AccountMemberUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberUpdateResponse">AccountMemberUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/members/{identifier}">client.AccountMembers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier interface{}, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberDeleteResponse">AccountMemberDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/members">client.AccountMembers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberService.AccountMembersAddMember">AccountMembersAddMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier interface{}, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberAccountMembersAddMemberParams">AccountMemberAccountMembersAddMemberParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberAccountMembersAddMemberResponse">AccountMemberAccountMembersAddMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/members">client.AccountMembers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberService.AccountMembersListMembers">AccountMembersListMembers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier interface{}, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberAccountMembersListMembersParams">AccountMemberAccountMembersListMembersParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccountMemberAccountMembersListMembersResponse">AccountMemberAccountMembersListMembersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tunnels

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelGetResponse">TunnelGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelDeleteResponse">TunnelDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelArgoTunnelNewAnArgoTunnelResponse">TunnelArgoTunnelNewAnArgoTunnelResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelArgoTunnelListArgoTunnelsResponse">TunnelArgoTunnelListArgoTunnelsResponse</a>

Methods:

- <code title="get /accounts/{account_id}/tunnels/{tunnel_id}">client.Tunnels.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, tunnelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelGetResponse">TunnelGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/tunnels/{tunnel_id}">client.Tunnels.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, tunnelID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelDeleteParams">TunnelDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelDeleteResponse">TunnelDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_id}/tunnels">client.Tunnels.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelService.ArgoTunnelNewAnArgoTunnel">ArgoTunnelNewAnArgoTunnel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelArgoTunnelNewAnArgoTunnelParams">TunnelArgoTunnelNewAnArgoTunnelParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelArgoTunnelNewAnArgoTunnelResponse">TunnelArgoTunnelNewAnArgoTunnelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/tunnels">client.Tunnels.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelService.ArgoTunnelListArgoTunnels">ArgoTunnelListArgoTunnels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelArgoTunnelListArgoTunnelsParams">TunnelArgoTunnelListArgoTunnelsParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelArgoTunnelListArgoTunnelsResponse">TunnelArgoTunnelListArgoTunnelsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelConnectionDeleteResponse">TunnelConnectionDeleteResponse</a>

Methods:

- <code title="delete /accounts/{account_id}/tunnels/{tunnel_id}/connections">client.Tunnels.Connections.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelConnectionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, tunnelID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelConnectionDeleteParams">TunnelConnectionDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TunnelConnectionDeleteResponse">TunnelConnectionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# D1

## Databases

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseNewResponse">D1DatabaseNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseListResponse">D1DatabaseListResponse</a>

Methods:

- <code title="post /accounts/{account_id}/d1/database">client.D1.Databases.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseNewParams">D1DatabaseNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseNewResponse">D1DatabaseNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/d1/database">client.D1.Databases.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseListParams">D1DatabaseListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseListResponse">D1DatabaseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Database

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseGetResponse">D1DatabaseGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseDeleteResponse">D1DatabaseDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseQueryResponse">D1DatabaseQueryResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/d1/database/{database_identifier}">client.D1.Database.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, databaseIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseGetResponse">D1DatabaseGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/d1/database/{database_identifier}">client.D1.Database.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, databaseIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseDeleteResponse">D1DatabaseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/d1/database/{database_identifier}/query">client.D1.Database.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, databaseIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseQueryParams">D1DatabaseQueryParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#D1DatabaseQueryResponse">D1DatabaseQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Dex

## Colos

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexColoListResponse">DexColoListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/colos">client.Dex.Colos.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexColoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexColoListParams">DexColoListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexColoListResponse">DexColoListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FleetStatus

### Devices

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusDeviceListResponse">DexFleetStatusDeviceListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/fleet-status/devices">client.Dex.FleetStatus.Devices.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusDeviceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusDeviceListParams">DexFleetStatusDeviceListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusDeviceListResponse">DexFleetStatusDeviceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Live

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusLiveListResponse">DexFleetStatusLiveListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/fleet-status/live">client.Dex.FleetStatus.Live.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusLiveService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusLiveListParams">DexFleetStatusLiveListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusLiveListResponse">DexFleetStatusLiveListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OverTime

Methods:

- <code title="get /accounts/{account_id}/dex/fleet-status/over-time">client.Dex.FleetStatus.OverTime.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusOverTimeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexFleetStatusOverTimeListParams">DexFleetStatusOverTimeListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## HTTPTests

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexHTTPTestGetResponse">DexHTTPTestGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/http-tests/{test_id}">client.Dex.HTTPTests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexHTTPTestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexHTTPTestGetParams">DexHTTPTestGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexHTTPTestGetResponse">DexHTTPTestGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Percentiles

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexHTTPTestPercentileListResponse">DexHTTPTestPercentileListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/http-tests/{test_id}/percentiles">client.Dex.HTTPTests.Percentiles.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexHTTPTestPercentileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexHTTPTestPercentileListParams">DexHTTPTestPercentileListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexHTTPTestPercentileListResponse">DexHTTPTestPercentileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tests

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTestListResponse">DexTestListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/tests">client.Dex.Tests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTestListParams">DexTestListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTestListResponse">DexTestListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### UniqueDevices

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTestUniqueDeviceListResponse">DexTestUniqueDeviceListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/tests/unique-devices">client.Dex.Tests.UniqueDevices.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTestUniqueDeviceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTestUniqueDeviceListParams">DexTestUniqueDeviceListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTestUniqueDeviceListResponse">DexTestUniqueDeviceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TracerouteTestResults

### NetworkPath

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestResultNetworkPathListResponse">DexTracerouteTestResultNetworkPathListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/traceroute-test-results/{test_result_id}/network-path">client.Dex.TracerouteTestResults.NetworkPath.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestResultNetworkPathService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, testResultID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestResultNetworkPathListResponse">DexTracerouteTestResultNetworkPathListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TracerouteTests

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestGetResponse">DexTracerouteTestGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestNetworkPathResponse">DexTracerouteTestNetworkPathResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestPercentilesResponse">DexTracerouteTestPercentilesResponse</a>

Methods:

- <code title="get /accounts/{account_id}/dex/traceroute-tests/{test_id}">client.Dex.TracerouteTests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestGetParams">DexTracerouteTestGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestGetResponse">DexTracerouteTestGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/dex/traceroute-tests/{test_id}/network-path">client.Dex.TracerouteTests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestService.NetworkPath">NetworkPath</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestNetworkPathParams">DexTracerouteTestNetworkPathParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestNetworkPathResponse">DexTracerouteTestNetworkPathResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/dex/traceroute-tests/{test_id}/percentiles">client.Dex.TracerouteTests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestService.Percentiles">Percentiles</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestPercentilesParams">DexTracerouteTestPercentilesParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DexTracerouteTestPercentilesResponse">DexTracerouteTestPercentilesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# R2

## Buckets

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketNewResponse">R2BucketNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketGetResponse">R2BucketGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketListResponse">R2BucketListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketDeleteResponse">R2BucketDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_id}/r2/buckets">client.R2.Buckets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketNewParams">R2BucketNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketNewResponse">R2BucketNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/r2/buckets/{bucket_name}">client.R2.Buckets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, bucketName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketGetResponse">R2BucketGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/r2/buckets">client.R2.Buckets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketListParams">R2BucketListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketListResponse">R2BucketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/r2/buckets/{bucket_name}">client.R2.Buckets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, bucketName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#R2BucketDeleteResponse">R2BucketDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Stream

## AudioTracks

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackUpdateResponse">StreamAudioTrackUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackListResponse">StreamAudioTrackListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackDeleteResponse">StreamAudioTrackDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackCopyResponse">StreamAudioTrackCopyResponse</a>

Methods:

- <code title="patch /accounts/{account_id}/stream/{identifier}/audio/{audio_identifier}">client.Stream.AudioTracks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, audioIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackUpdateParams">StreamAudioTrackUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackUpdateResponse">StreamAudioTrackUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/stream/{identifier}/audio">client.Stream.AudioTracks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackListResponse">StreamAudioTrackListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/stream/{identifier}/audio/{audio_identifier}">client.Stream.AudioTracks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, audioIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackDeleteResponse">StreamAudioTrackDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_id}/stream/{identifier}/audio/copy">client.Stream.AudioTracks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackService.Copy">Copy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackCopyParams">StreamAudioTrackCopyParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamAudioTrackCopyResponse">StreamAudioTrackCopyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Videos

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamVideoStorageUsageResponse">StreamVideoStorageUsageResponse</a>

Methods:

- <code title="get /accounts/{account_id}/stream/storage-usage">client.Stream.Videos.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamVideoService.StorageUsage">StorageUsage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamVideoStorageUsageParams">StreamVideoStorageUsageParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#StreamVideoStorageUsageResponse">StreamVideoStorageUsageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Teamnet

## Routes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteNewResponse">TeamnetRouteNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteUpdateResponse">TeamnetRouteUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteDeleteResponse">TeamnetRouteDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_id}/teamnet/routes">client.Teamnet.Routes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteNewParams">TeamnetRouteNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteNewResponse">TeamnetRouteNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/teamnet/routes/{route_id}">client.Teamnet.Routes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, routeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteUpdateParams">TeamnetRouteUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteUpdateResponse">TeamnetRouteUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/teamnet/routes/{route_id}">client.Teamnet.Routes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, routeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#TeamnetRouteDeleteResponse">TeamnetRouteDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WarpConnector

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorNewResponse">WarpConnectorNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorGetResponse">WarpConnectorGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorUpdateResponse">WarpConnectorUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorListResponse">WarpConnectorListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorDeleteResponse">WarpConnectorDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_id}/warp_connector">client.WarpConnector.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorNewParams">WarpConnectorNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorNewResponse">WarpConnectorNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/warp_connector/{tunnel_id}">client.WarpConnector.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, tunnelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorGetResponse">WarpConnectorGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/warp_connector/{tunnel_id}">client.WarpConnector.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, tunnelID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorUpdateParams">WarpConnectorUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorUpdateResponse">WarpConnectorUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/warp_connector">client.WarpConnector.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorListParams">WarpConnectorListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorListResponse">WarpConnectorListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/warp_connector/{tunnel_id}">client.WarpConnector.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, tunnelID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorDeleteParams">WarpConnectorDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WarpConnectorDeleteResponse">WarpConnectorDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Dispatchers

## Scripts

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptGetResponse">DispatcherScriptGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptUpdateResponse">DispatcherScriptUpdateResponse</a>

Methods:

- <code title="get /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}">client.Dispatchers.Scripts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, dispatchNamespace <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptGetResponse">DispatcherScriptGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}">client.Dispatchers.Scripts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, dispatchNamespace <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptUpdateParams">DispatcherScriptUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptUpdateResponse">DispatcherScriptUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}">client.Dispatchers.Scripts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, dispatchNamespace <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DispatcherScriptDeleteParams">DispatcherScriptDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# WorkersForPlatforms

## Dispatch

### Namespaces

#### Scripts

##### Content

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptContentUpdateResponse">WorkersForPlatformDispatchNamespaceScriptContentUpdateResponse</a>

Methods:

- <code title="get /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/content">client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Content.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptContentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, dispatchNamespace <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/content">client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Content.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptContentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, dispatchNamespace <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptContentUpdateParams">WorkersForPlatformDispatchNamespaceScriptContentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptContentUpdateResponse">WorkersForPlatformDispatchNamespaceScriptContentUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Settings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptSettingGetResponse">WorkersForPlatformDispatchNamespaceScriptSettingGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse">WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse</a>

Methods:

- <code title="get /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/settings">client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptSettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, dispatchNamespace <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptSettingGetResponse">WorkersForPlatformDispatchNamespaceScriptSettingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/settings">client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptSettingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, dispatchNamespace <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptSettingUpdateParams">WorkersForPlatformDispatchNamespaceScriptSettingUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse">WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WorkerDomains

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerDomainGetResponse">WorkerDomainGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/workers/domains/{domain_id}">client.WorkerDomains.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerDomainService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}, domainID interface{}) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerDomainGetResponse">WorkerDomainGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/workers/domains/{domain_id}">client.WorkerDomains.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerDomainService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}, domainID interface{}) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# WorkerScripts

## Content

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptContentUpdateResponse">WorkerScriptContentUpdateResponse</a>

Methods:

- <code title="put /accounts/{account_id}/workers/scripts/{script_name}/content">client.WorkerScripts.Content.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptContentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptContentUpdateParams">WorkerScriptContentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptContentUpdateResponse">WorkerScriptContentUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ContentV2

Methods:

- <code title="get /accounts/{account_id}/workers/scripts/{script_name}/content/v2">client.WorkerScripts.ContentV2.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptContentV2Service.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Settings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptSettingGetResponse">WorkerScriptSettingGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptSettingUpdateResponse">WorkerScriptSettingUpdateResponse</a>

Methods:

- <code title="get /accounts/{account_id}/workers/scripts/{script_name}/settings">client.WorkerScripts.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptSettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptSettingGetResponse">WorkerScriptSettingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/workers/scripts/{script_name}/settings">client.WorkerScripts.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptSettingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, scriptName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptSettingUpdateParams">WorkerScriptSettingUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#WorkerScriptSettingUpdateResponse">WorkerScriptSettingUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Zerotrust

## ConnectivitySettings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZerotrustConnectivitySettingGetResponse">ZerotrustConnectivitySettingGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZerotrustConnectivitySettingUpdateResponse">ZerotrustConnectivitySettingUpdateResponse</a>

Methods:

- <code title="get /accounts/{account_id}/zerotrust/connectivity_settings">client.Zerotrust.ConnectivitySettings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZerotrustConnectivitySettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZerotrustConnectivitySettingGetResponse">ZerotrustConnectivitySettingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/zerotrust/connectivity_settings">client.Zerotrust.ConnectivitySettings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZerotrustConnectivitySettingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZerotrustConnectivitySettingUpdateParams">ZerotrustConnectivitySettingUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZerotrustConnectivitySettingUpdateResponse">ZerotrustConnectivitySettingUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Addressing

## Prefixes

### BgpPrefixes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixGetResponse">AddressingPrefixBgpPrefixGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixUpdateResponse">AddressingPrefixBgpPrefixUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixListResponse">AddressingPrefixBgpPrefixListResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/addressing/prefixes/{prefix_identifier}/bgp/prefixes/{bgp_prefix_identifier}">client.Addressing.Prefixes.BgpPrefixes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, prefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, bgpPrefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixGetResponse">AddressingPrefixBgpPrefixGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_identifier}/addressing/prefixes/{prefix_identifier}/bgp/prefixes/{bgp_prefix_identifier}">client.Addressing.Prefixes.BgpPrefixes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, prefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, bgpPrefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixUpdateParams">AddressingPrefixBgpPrefixUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixUpdateResponse">AddressingPrefixBgpPrefixUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/addressing/prefixes/{prefix_identifier}/bgp/prefixes">client.Addressing.Prefixes.BgpPrefixes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, prefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBgpPrefixListResponse">AddressingPrefixBgpPrefixListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Bindings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingNewResponse">AddressingPrefixBindingNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingGetResponse">AddressingPrefixBindingGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingListResponse">AddressingPrefixBindingListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingDeleteResponse">AddressingPrefixBindingDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/addressing/prefixes/{prefix_identifier}/bindings">client.Addressing.Prefixes.Bindings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, prefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingNewParams">AddressingPrefixBindingNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingNewResponse">AddressingPrefixBindingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/addressing/prefixes/{prefix_identifier}/bindings/{binding_identifier}">client.Addressing.Prefixes.Bindings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, prefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, bindingIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingGetResponse">AddressingPrefixBindingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/addressing/prefixes/{prefix_identifier}/bindings">client.Addressing.Prefixes.Bindings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, prefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingListResponse">AddressingPrefixBindingListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/addressing/prefixes/{prefix_identifier}/bindings/{binding_identifier}">client.Addressing.Prefixes.Bindings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, prefixIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, bindingIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingPrefixBindingDeleteResponse">AddressingPrefixBindingDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Services

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingServiceListResponse">AddressingServiceListResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/addressing/services">client.Addressing.Services.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingServiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AddressingServiceListResponse">AddressingServiceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Challenges

## Widgets

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetNewResponse">ChallengeWidgetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetGetResponse">ChallengeWidgetGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetUpdateResponse">ChallengeWidgetUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetListResponse">ChallengeWidgetListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetDeleteResponse">ChallengeWidgetDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetRotateSecretResponse">ChallengeWidgetRotateSecretResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/challenges/widgets">client.Challenges.Widgets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetNewParams">ChallengeWidgetNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetNewResponse">ChallengeWidgetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/challenges/widgets/{sitekey}">client.Challenges.Widgets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, sitekey <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetGetResponse">ChallengeWidgetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/challenges/widgets/{sitekey}">client.Challenges.Widgets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, sitekey <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetUpdateParams">ChallengeWidgetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetUpdateResponse">ChallengeWidgetUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/challenges/widgets">client.Challenges.Widgets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetListParams">ChallengeWidgetListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetListResponse">ChallengeWidgetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/challenges/widgets/{sitekey}">client.Challenges.Widgets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, sitekey <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetDeleteResponse">ChallengeWidgetDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/challenges/widgets/{sitekey}/rotate_secret">client.Challenges.Widgets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetService.RotateSecret">RotateSecret</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, sitekey <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetRotateSecretParams">ChallengeWidgetRotateSecretParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ChallengeWidgetRotateSecretResponse">ChallengeWidgetRotateSecretResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Hyperdrive

## Configs

# Intel

## IndicatorFeeds

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedNewResponse">IntelIndicatorFeedNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedGetResponse">IntelIndicatorFeedGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedListResponse">IntelIndicatorFeedListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedPermissionsAddResponse">IntelIndicatorFeedPermissionsAddResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedPermissionsRemoveResponse">IntelIndicatorFeedPermissionsRemoveResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedPermissionsViewResponse">IntelIndicatorFeedPermissionsViewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedSnapshotResponse">IntelIndicatorFeedSnapshotResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/intel/indicator-feeds">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedNewParams">IntelIndicatorFeedNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedNewResponse">IntelIndicatorFeedNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/intel/indicator-feeds/{feed_id}">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, feedID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedGetResponse">IntelIndicatorFeedGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/intel/indicator-feeds">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedListResponse">IntelIndicatorFeedListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/intel/indicator-feeds/{feed_id}/data">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedService.Data">Data</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, feedID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/intel/indicator-feeds/permissions/add">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedService.PermissionsAdd">PermissionsAdd</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedPermissionsAddParams">IntelIndicatorFeedPermissionsAddParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedPermissionsAddResponse">IntelIndicatorFeedPermissionsAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/intel/indicator-feeds/permissions/remove">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedService.PermissionsRemove">PermissionsRemove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedPermissionsRemoveParams">IntelIndicatorFeedPermissionsRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedPermissionsRemoveResponse">IntelIndicatorFeedPermissionsRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/intel/indicator-feeds/permissions/view">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedService.PermissionsView">PermissionsView</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedPermissionsViewResponse">IntelIndicatorFeedPermissionsViewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/intel/indicator-feeds/{feed_id}/snapshot">client.Intel.IndicatorFeeds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedService.Snapshot">Snapshot</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, feedID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedSnapshotParams">IntelIndicatorFeedSnapshotParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelIndicatorFeedSnapshotResponse">IntelIndicatorFeedSnapshotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sinkholes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelSinkholeListResponse">IntelSinkholeListResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/intel/sinkholes">client.Intel.Sinkholes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelSinkholeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#IntelSinkholeListResponse">IntelSinkholeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rum

## SiteInfos

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoNewResponse">RumSiteInfoNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoGetResponse">RumSiteInfoGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoUpdateResponse">RumSiteInfoUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoListResponse">RumSiteInfoListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoDeleteResponse">RumSiteInfoDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/rum/site_info">client.Rum.SiteInfos.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoNewParams">RumSiteInfoNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoNewResponse">RumSiteInfoNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/rum/site_info/{site_identifier}">client.Rum.SiteInfos.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, siteIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoGetResponse">RumSiteInfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/rum/site_info/{site_identifier}">client.Rum.SiteInfos.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, siteIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoUpdateParams">RumSiteInfoUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoUpdateResponse">RumSiteInfoUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/rum/site_info/list">client.Rum.SiteInfos.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoListParams">RumSiteInfoListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoListResponse">RumSiteInfoListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/rum/site_info/{site_identifier}">client.Rum.SiteInfos.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, siteIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumSiteInfoDeleteResponse">RumSiteInfoDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Rules

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleNewResponse">RumRuleNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleUpdateResponse">RumRuleUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleListResponse">RumRuleListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleDeleteResponse">RumRuleDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/rum/v2/{ruleset_identifier}/rule">client.Rum.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, rulesetIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleNewParams">RumRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleNewResponse">RumRuleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/rum/v2/{ruleset_identifier}/rule/{rule_identifier}">client.Rum.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, rulesetIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, ruleIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleUpdateParams">RumRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleUpdateResponse">RumRuleUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/rum/v2/{ruleset_identifier}/rules">client.Rum.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, rulesetIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleListResponse">RumRuleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/rum/v2/{ruleset_identifier}/rule/{rule_identifier}">client.Rum.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, rulesetIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, ruleIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RumRuleDeleteResponse">RumRuleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Vectorize

## Indexes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexListResponse">VectorizeIndexListResponse</a>

Methods:

- <code title="get /accounts/{account_identifier}/vectorize/indexes">client.Vectorize.Indexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexListResponse">VectorizeIndexListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VectorizeIndexes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexNewResponse">VectorizeIndexNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexGetResponse">VectorizeIndexGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexUpdateResponse">VectorizeIndexUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexDeleteResponse">VectorizeIndexDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexDeleteByIDsResponse">VectorizeIndexDeleteByIDsResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexGetByIDsResponse">VectorizeIndexGetByIDsResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexInsertResponse">VectorizeIndexInsertResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexQueryResponse">VectorizeIndexQueryResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexUpsertResponse">VectorizeIndexUpsertResponse</a>

Methods:

- <code title="post /accounts/{account_identifier}/vectorize/indexes">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexNewParams">VectorizeIndexNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexNewResponse">VectorizeIndexNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_identifier}/vectorize/indexes/{index_name}">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, indexName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexGetResponse">VectorizeIndexGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_identifier}/vectorize/indexes/{index_name}">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, indexName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexUpdateParams">VectorizeIndexUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexUpdateResponse">VectorizeIndexUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_identifier}/vectorize/indexes/{index_name}">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, indexName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexDeleteResponse">VectorizeIndexDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/vectorize/indexes/{index_name}/delete-by-ids">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.DeleteByIDs">DeleteByIDs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, indexName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexDeleteByIDsParams">VectorizeIndexDeleteByIDsParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexDeleteByIDsResponse">VectorizeIndexDeleteByIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/vectorize/indexes/{index_name}/get-by-ids">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.GetByIDs">GetByIDs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, indexName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexGetByIDsParams">VectorizeIndexGetByIDsParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexGetByIDsResponse">VectorizeIndexGetByIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/vectorize/indexes/{index_name}/insert">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.Insert">Insert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, indexName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexInsertResponse">VectorizeIndexInsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/vectorize/indexes/{index_name}/query">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, indexName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexQueryParams">VectorizeIndexQueryParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexQueryResponse">VectorizeIndexQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_identifier}/vectorize/indexes/{index_name}/upsert">client.VectorizeIndexes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, indexName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#VectorizeIndexUpsertResponse">VectorizeIndexUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# URLScanner

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanResponse">URLScannerScanResponse</a>

Methods:

- <code title="get /accounts/{accountId}/urlscanner/scan">client.URLScanner.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanParams">URLScannerScanParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanResponse">URLScannerScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Scans

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanNewResponse">URLScannerScanNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanGetResponse">URLScannerScanGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanHarResponse">URLScannerScanHarResponse</a>

Methods:

- <code title="post /accounts/{accountId}/urlscanner/scan">client.URLScanner.Scans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanNewParams">URLScannerScanNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanNewResponse">URLScannerScanNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/urlscanner/scan/{scanId}">client.URLScanner.Scans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, scanID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanGetResponse">URLScannerScanGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/urlscanner/scan/{scanId}/har">client.URLScanner.Scans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanService.Har">Har</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, scanID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanHarResponse">URLScannerScanHarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/urlscanner/scan/{scanId}/screenshot">client.URLScanner.Scans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanService.Screenshot">Screenshot</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, scanID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#URLScannerScanScreenshotParams">URLScannerScanScreenshotParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Radar

## Attacks

### Layer3

#### TimeseriesGroups

##### Industry

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupIndustryListResponse">RadarAttackLayer3TimeseriesGroupIndustryListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/timeseries_groups/industry">client.Radar.Attacks.Layer3.TimeseriesGroups.Industry.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupIndustryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupIndustryListParams">RadarAttackLayer3TimeseriesGroupIndustryListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupIndustryListResponse">RadarAttackLayer3TimeseriesGroupIndustryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### IPVersion

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupIPVersionListResponse">RadarAttackLayer3TimeseriesGroupIPVersionListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/timeseries_groups/ip_version">client.Radar.Attacks.Layer3.TimeseriesGroups.IPVersion.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupIPVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupIPVersionListParams">RadarAttackLayer3TimeseriesGroupIPVersionListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupIPVersionListResponse">RadarAttackLayer3TimeseriesGroupIPVersionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Protocol

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupProtocolListResponse">RadarAttackLayer3TimeseriesGroupProtocolListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/timeseries_groups/protocol">client.Radar.Attacks.Layer3.TimeseriesGroups.Protocol.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupProtocolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupProtocolListParams">RadarAttackLayer3TimeseriesGroupProtocolListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupProtocolListResponse">RadarAttackLayer3TimeseriesGroupProtocolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Vector

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupVectorListResponse">RadarAttackLayer3TimeseriesGroupVectorListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/timeseries_groups/vector">client.Radar.Attacks.Layer3.TimeseriesGroups.Vector.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupVectorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupVectorListParams">RadarAttackLayer3TimeseriesGroupVectorListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupVectorListResponse">RadarAttackLayer3TimeseriesGroupVectorListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Vertical

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupVerticalListResponse">RadarAttackLayer3TimeseriesGroupVerticalListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/timeseries_groups/vertical">client.Radar.Attacks.Layer3.TimeseriesGroups.Vertical.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupVerticalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupVerticalListParams">RadarAttackLayer3TimeseriesGroupVerticalListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TimeseriesGroupVerticalListResponse">RadarAttackLayer3TimeseriesGroupVerticalListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Top

##### Attacks

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopAttackListResponse">RadarAttackLayer3TopAttackListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/top/attacks">client.Radar.Attacks.Layer3.Top.Attacks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopAttackService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopAttackListParams">RadarAttackLayer3TopAttackListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopAttackListResponse">RadarAttackLayer3TopAttackListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Industry

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopIndustryListResponse">RadarAttackLayer3TopIndustryListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/top/industry">client.Radar.Attacks.Layer3.Top.Industry.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopIndustryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopIndustryListParams">RadarAttackLayer3TopIndustryListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopIndustryListResponse">RadarAttackLayer3TopIndustryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Locations

###### Origin

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopLocationOriginListResponse">RadarAttackLayer3TopLocationOriginListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/top/locations/origin">client.Radar.Attacks.Layer3.Top.Locations.Origin.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopLocationOriginService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopLocationOriginListParams">RadarAttackLayer3TopLocationOriginListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopLocationOriginListResponse">RadarAttackLayer3TopLocationOriginListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

###### Target

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopLocationTargetListResponse">RadarAttackLayer3TopLocationTargetListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/top/locations/target">client.Radar.Attacks.Layer3.Top.Locations.Target.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopLocationTargetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopLocationTargetListParams">RadarAttackLayer3TopLocationTargetListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopLocationTargetListResponse">RadarAttackLayer3TopLocationTargetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Vertical

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopVerticalListResponse">RadarAttackLayer3TopVerticalListResponse</a>

Methods:

- <code title="get /radar/attacks/layer3/top/vertical">client.Radar.Attacks.Layer3.Top.Vertical.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopVerticalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopVerticalListParams">RadarAttackLayer3TopVerticalListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarAttackLayer3TopVerticalListResponse">RadarAttackLayer3TopVerticalListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Emails

### Security

#### Dmarc

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityDmarcListResponse">RadarEmailSecurityDmarcListResponse</a>

Methods:

- <code title="get /radar/email/security/timeseries_groups/dmarc">client.Radar.Emails.Security.Dmarc.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityDmarcService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityDmarcListParams">RadarEmailSecurityDmarcListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityDmarcListResponse">RadarEmailSecurityDmarcListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Malicious

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityMaliciousListResponse">RadarEmailSecurityMaliciousListResponse</a>

Methods:

- <code title="get /radar/email/security/timeseries_groups/malicious">client.Radar.Emails.Security.Malicious.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityMaliciousService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityMaliciousListParams">RadarEmailSecurityMaliciousListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityMaliciousListResponse">RadarEmailSecurityMaliciousListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Spam

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecuritySpamListResponse">RadarEmailSecuritySpamListResponse</a>

Methods:

- <code title="get /radar/email/security/timeseries_groups/spam">client.Radar.Emails.Security.Spam.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecuritySpamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecuritySpamListParams">RadarEmailSecuritySpamListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecuritySpamListResponse">RadarEmailSecuritySpamListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Spf

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecuritySpfListResponse">RadarEmailSecuritySpfListResponse</a>

Methods:

- <code title="get /radar/email/security/timeseries_groups/spf">client.Radar.Emails.Security.Spf.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecuritySpfService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecuritySpfListParams">RadarEmailSecuritySpfListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecuritySpfListResponse">RadarEmailSecuritySpfListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### ThreatCategory

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityThreatCategoryListResponse">RadarEmailSecurityThreatCategoryListResponse</a>

Methods:

- <code title="get /radar/email/security/timeseries_groups/threat_category">client.Radar.Emails.Security.ThreatCategory.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityThreatCategoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityThreatCategoryListParams">RadarEmailSecurityThreatCategoryListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityThreatCategoryListResponse">RadarEmailSecurityThreatCategoryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Top

##### Ases

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseListResponse">RadarEmailSecurityTopAseListResponse</a>

Methods:

- <code title="get /radar/email/security/top/ases">client.Radar.Emails.Security.Top.Ases.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseListParams">RadarEmailSecurityTopAseListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseListResponse">RadarEmailSecurityTopAseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

###### Arc

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseArcGetResponse">RadarEmailSecurityTopAseArcGetResponse</a>

Methods:

- <code title="get /radar/email/security/top/ases/arc/{arc}">client.Radar.Emails.Security.Top.Ases.Arc.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseArcService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, arc <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseArcGetParamsArc">RadarEmailSecurityTopAseArcGetParamsArc</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseArcGetParams">RadarEmailSecurityTopAseArcGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseArcGetResponse">RadarEmailSecurityTopAseArcGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

###### Dkim

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDkimGetResponse">RadarEmailSecurityTopAseDkimGetResponse</a>

Methods:

- <code title="get /radar/email/security/top/ases/dkim/{dkim}">client.Radar.Emails.Security.Top.Ases.Dkim.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDkimService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, dkim <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDkimGetParamsDkim">RadarEmailSecurityTopAseDkimGetParamsDkim</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDkimGetParams">RadarEmailSecurityTopAseDkimGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDkimGetResponse">RadarEmailSecurityTopAseDkimGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

###### Dmarc

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDmarcGetResponse">RadarEmailSecurityTopAseDmarcGetResponse</a>

Methods:

- <code title="get /radar/email/security/top/ases/dmarc/{dmarc}">client.Radar.Emails.Security.Top.Ases.Dmarc.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDmarcService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, dmarc <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDmarcGetParamsDmarc">RadarEmailSecurityTopAseDmarcGetParamsDmarc</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDmarcGetParams">RadarEmailSecurityTopAseDmarcGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEmailSecurityTopAseDmarcGetResponse">RadarEmailSecurityTopAseDmarcGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Entities

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEntityIPsResponse">RadarEntityIPsResponse</a>

Methods:

- <code title="get /radar/entities/ip">client.Radar.Entities.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEntityService.IPs">IPs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEntityIPsParams">RadarEntityIPsParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEntityIPsResponse">RadarEntityIPsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Asns

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEntityAsnRelResponse">RadarEntityAsnRelResponse</a>

Methods:

- <code title="get /radar/entities/asns/{asn}/rel">client.Radar.Entities.Asns.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEntityAsnService.Rel">Rel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, asn <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEntityAsnRelParams">RadarEntityAsnRelParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarEntityAsnRelResponse">RadarEntityAsnRelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## HTTP

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBotClassesResponse">RadarHTTPBotClassesResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBrowserFamiliesResponse">RadarHTTPBrowserFamiliesResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBrowsersResponse">RadarHTTPBrowsersResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPDeviceTypesResponse">RadarHTTPDeviceTypesResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHttphttpProtocolsResponse">RadarHttphttpProtocolsResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHttphttpVersionsResponse">RadarHttphttpVersionsResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHttpipVersionsResponse">RadarHttpipVersionsResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPOssResponse">RadarHTTPOssResponse</a>

Methods:

- <code title="get /radar/http/timeseries_groups/bot_class">client.Radar.HTTP.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPService.BotClasses">BotClasses</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBotClassesParams">RadarHTTPBotClassesParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBotClassesResponse">RadarHTTPBotClassesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /radar/http/timeseries_groups/browser_family">client.Radar.HTTP.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPService.BrowserFamilies">BrowserFamilies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBrowserFamiliesParams">RadarHTTPBrowserFamiliesParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBrowserFamiliesResponse">RadarHTTPBrowserFamiliesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /radar/http/timeseries_groups/browser">client.Radar.HTTP.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPService.Browsers">Browsers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBrowsersParams">RadarHTTPBrowsersParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPBrowsersResponse">RadarHTTPBrowsersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /radar/http/timeseries_groups/device_type">client.Radar.HTTP.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPService.DeviceTypes">DeviceTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPDeviceTypesParams">RadarHTTPDeviceTypesParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPDeviceTypesResponse">RadarHTTPDeviceTypesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /radar/http/timeseries_groups/http_protocol">client.Radar.HTTP.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPService.HTTPProtocols">HTTPProtocols</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPHTTPProtocolsParams">RadarHTTPHTTPProtocolsParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHttphttpProtocolsResponse">RadarHttphttpProtocolsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /radar/http/timeseries_groups/http_version">client.Radar.HTTP.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPService.HTTPVersions">HTTPVersions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPHTTPVersionsParams">RadarHTTPHTTPVersionsParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHttphttpVersionsResponse">RadarHttphttpVersionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /radar/http/timeseries_groups/ip_version">client.Radar.HTTP.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPService.IPVersions">IPVersions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPIPVersionsParams">RadarHTTPIPVersionsParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHttpipVersionsResponse">RadarHttpipVersionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /radar/http/timeseries_groups/os">client.Radar.HTTP.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPService.Oss">Oss</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPOssParams">RadarHTTPOssParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPOssResponse">RadarHTTPOssResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### TLSVersion

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHttptlsVersionListResponse">RadarHttptlsVersionListResponse</a>

Methods:

- <code title="get /radar/http/timeseries_groups/tls_version">client.Radar.HTTP.TLSVersion.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPTLSVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHTTPTLSVersionListParams">RadarHTTPTLSVersionListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarHttptlsVersionListResponse">RadarHttptlsVersionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Quality

### Iqi

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualityIqiGetResponse">RadarQualityIqiGetResponse</a>

Methods:

- <code title="get /radar/quality/iqi/summary">client.Radar.Quality.Iqi.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualityIqiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualityIqiGetParams">RadarQualityIqiGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualityIqiGetResponse">RadarQualityIqiGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### TimeseriesGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualityIqiTimeseriesGroupListResponse">RadarQualityIqiTimeseriesGroupListResponse</a>

Methods:

- <code title="get /radar/quality/iqi/timeseries_groups">client.Radar.Quality.Iqi.TimeseriesGroups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualityIqiTimeseriesGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualityIqiTimeseriesGroupListParams">RadarQualityIqiTimeseriesGroupListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualityIqiTimeseriesGroupListResponse">RadarQualityIqiTimeseriesGroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Speed

#### Histogram

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedHistogramGetResponse">RadarQualitySpeedHistogramGetResponse</a>

Methods:

- <code title="get /radar/quality/speed/histogram">client.Radar.Quality.Speed.Histogram.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedHistogramService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedHistogramGetParams">RadarQualitySpeedHistogramGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedHistogramGetResponse">RadarQualitySpeedHistogramGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Summary

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedSummaryGetResponse">RadarQualitySpeedSummaryGetResponse</a>

Methods:

- <code title="get /radar/quality/speed/summary">client.Radar.Quality.Speed.Summary.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedSummaryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedSummaryGetParams">RadarQualitySpeedSummaryGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedSummaryGetResponse">RadarQualitySpeedSummaryGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Top

##### Ases

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedTopAseListResponse">RadarQualitySpeedTopAseListResponse</a>

Methods:

- <code title="get /radar/quality/speed/top/ases">client.Radar.Quality.Speed.Top.Ases.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedTopAseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedTopAseListParams">RadarQualitySpeedTopAseListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedTopAseListResponse">RadarQualitySpeedTopAseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Locations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedTopLocationListResponse">RadarQualitySpeedTopLocationListResponse</a>

Methods:

- <code title="get /radar/quality/speed/top/locations">client.Radar.Quality.Speed.Top.Locations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedTopLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedTopLocationListParams">RadarQualitySpeedTopLocationListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarQualitySpeedTopLocationListResponse">RadarQualitySpeedTopLocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ranking

### TimeseriesGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarRankingTimeseriesGroupListResponse">RadarRankingTimeseriesGroupListResponse</a>

Methods:

- <code title="get /radar/ranking/timeseries_groups">client.Radar.Ranking.TimeseriesGroups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarRankingTimeseriesGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarRankingTimeseriesGroupListParams">RadarRankingTimeseriesGroupListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarRankingTimeseriesGroupListResponse">RadarRankingTimeseriesGroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TrafficAnomalies

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarTrafficAnomalyListResponse">RadarTrafficAnomalyListResponse</a>

Methods:

- <code title="get /radar/traffic_anomalies">client.Radar.TrafficAnomalies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarTrafficAnomalyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarTrafficAnomalyListParams">RadarTrafficAnomalyListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarTrafficAnomalyListResponse">RadarTrafficAnomalyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Locations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarTrafficAnomalyLocationListResponse">RadarTrafficAnomalyLocationListResponse</a>

Methods:

- <code title="get /radar/traffic_anomalies/locations">client.Radar.TrafficAnomalies.Locations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarTrafficAnomalyLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarTrafficAnomalyLocationListParams">RadarTrafficAnomalyLocationListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#RadarTrafficAnomalyLocationListResponse">RadarTrafficAnomalyLocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BotManagements

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#BotManagementGetResponse">BotManagementGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#BotManagementUpdateResponse">BotManagementUpdateResponse</a>

Methods:

- <code title="get /zones/{zone_id}/bot_management">client.BotManagements.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#BotManagementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#BotManagementGetResponse">BotManagementGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_id}/bot_management">client.BotManagements.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#BotManagementService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#BotManagementUpdateParams">BotManagementUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#BotManagementUpdateResponse">BotManagementUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CacheReserves

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheReserveNewResponse">CacheReserveNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheReserveClearResponse">CacheReserveClearResponse</a>

Methods:

- <code title="post /zones/{zone_id}/cache/cache_reserve_clear">client.CacheReserves.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheReserveService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheReserveNewResponse">CacheReserveNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/cache/cache_reserve_clear">client.CacheReserves.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheReserveService.Clear">Clear</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheReserveClearResponse">CacheReserveClearResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OriginPostQuantumEncryptions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#OriginPostQuantumEncryptionGetResponse">OriginPostQuantumEncryptionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#OriginPostQuantumEncryptionUpdateResponse">OriginPostQuantumEncryptionUpdateResponse</a>

Methods:

- <code title="get /zones/{zone_id}/cache/origin_post_quantum_encryption">client.OriginPostQuantumEncryptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#OriginPostQuantumEncryptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#OriginPostQuantumEncryptionGetResponse">OriginPostQuantumEncryptionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_id}/cache/origin_post_quantum_encryption">client.OriginPostQuantumEncryptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#OriginPostQuantumEncryptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#OriginPostQuantumEncryptionUpdateParams">OriginPostQuantumEncryptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#OriginPostQuantumEncryptionUpdateResponse">OriginPostQuantumEncryptionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cache

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheRegionalTieredCachesResponse">CacheRegionalTieredCachesResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheUpdateRegionalTieredCacheResponse">CacheUpdateRegionalTieredCacheResponse</a>

Methods:

- <code title="get /zones/{zone_id}/cache/regional_tiered_cache">client.Cache.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheService.RegionalTieredCaches">RegionalTieredCaches</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheRegionalTieredCachesResponse">CacheRegionalTieredCachesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_id}/cache/regional_tiered_cache">client.Cache.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheService.UpdateRegionalTieredCache">UpdateRegionalTieredCache</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheUpdateRegionalTieredCacheParams">CacheUpdateRegionalTieredCacheParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#CacheUpdateRegionalTieredCacheResponse">CacheUpdateRegionalTieredCacheResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Firewall

## WAF

### Packages

#### Groups

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupGetResponse">FirewallWAFPackageGroupGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupUpdateResponse">FirewallWAFPackageGroupUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupListResponse">FirewallWAFPackageGroupListResponse</a>

Methods:

- <code title="get /zones/{zone_id}/firewall/waf/packages/{package_id}/groups/{group_id}">client.Firewall.WAF.Packages.Groups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, packageID <a href="https://pkg.go.dev/builtin#string">string</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupGetResponse">FirewallWAFPackageGroupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_id}/firewall/waf/packages/{package_id}/groups/{group_id}">client.Firewall.WAF.Packages.Groups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, packageID <a href="https://pkg.go.dev/builtin#string">string</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupUpdateParams">FirewallWAFPackageGroupUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupUpdateResponse">FirewallWAFPackageGroupUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/firewall/waf/packages/{package_id}/groups">client.Firewall.WAF.Packages.Groups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, packageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupListParams">FirewallWAFPackageGroupListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageGroupListResponse">FirewallWAFPackageGroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Rules

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageRuleGetResponse">FirewallWAFPackageRuleGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageRuleUpdateResponse">FirewallWAFPackageRuleUpdateResponse</a>

Methods:

- <code title="get /zones/{zone_id}/firewall/waf/packages/{package_id}/rules/{rule_id}">client.Firewall.WAF.Packages.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, packageID <a href="https://pkg.go.dev/builtin#string">string</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageRuleGetResponse">FirewallWAFPackageRuleGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_id}/firewall/waf/packages/{package_id}/rules/{rule_id}">client.Firewall.WAF.Packages.Rules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, packageID <a href="https://pkg.go.dev/builtin#string">string</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageRuleUpdateParams">FirewallWAFPackageRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#FirewallWAFPackageRuleUpdateResponse">FirewallWAFPackageRuleUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Zaraz

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazWorkflowUpdateResponse">ZarazWorkflowUpdateResponse</a>

Methods:

- <code title="put /zones/{zone_id}/settings/zaraz/workflow">client.Zaraz.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazService.WorkflowUpdate">WorkflowUpdate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazWorkflowUpdateParams">ZarazWorkflowUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazWorkflowUpdateResponse">ZarazWorkflowUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Config

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazConfigGetResponse">ZarazConfigGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazConfigUpdateResponse">ZarazConfigUpdateResponse</a>

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/config">client.Zaraz.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazConfigGetResponse">ZarazConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_id}/settings/zaraz/config">client.Zaraz.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazConfigService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazConfigUpdateParams">ZarazConfigUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazConfigUpdateResponse">ZarazConfigUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Default

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazDefaultGetResponse">ZarazDefaultGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/default">client.Zaraz.Default.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazDefaultService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazDefaultGetResponse">ZarazDefaultGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Export

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazExportGetResponse">ZarazExportGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/export">client.Zaraz.Export.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazExportService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazExportGetResponse">ZarazExportGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## History

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryUpdateResponse">ZarazHistoryUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryListResponse">ZarazHistoryListResponse</a>

Methods:

- <code title="put /zones/{zone_id}/settings/zaraz/history">client.Zaraz.History.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryUpdateParams">ZarazHistoryUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryUpdateResponse">ZarazHistoryUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/settings/zaraz/history">client.Zaraz.History.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryListParams">ZarazHistoryListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryListResponse">ZarazHistoryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Configs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryConfigGetResponse">ZarazHistoryConfigGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/history/configs">client.Zaraz.History.Configs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryConfigGetParams">ZarazHistoryConfigGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazHistoryConfigGetResponse">ZarazHistoryConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Publish

Methods:

- <code title="post /zones/{zone_id}/settings/zaraz/publish">client.Zaraz.Publish.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazPublishService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazPublishNewParams">ZarazPublishNewParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Workflow

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazWorkflowGetResponse">ZarazWorkflowGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/settings/zaraz/workflow">client.Zaraz.Workflow.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazWorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#ZarazWorkflowGetResponse">ZarazWorkflowGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SpeedAPI

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIAvailabilitiesListResponse">SpeedAPIAvailabilitiesListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIPagesListResponse">SpeedAPIPagesListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleDeleteResponse">SpeedAPIScheduleDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleGetResponse">SpeedAPIScheduleGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsNewResponse">SpeedAPITestsNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsDeleteResponse">SpeedAPITestsDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsListResponse">SpeedAPITestsListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsGetResponse">SpeedAPITestsGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITrendsListResponse">SpeedAPITrendsListResponse</a>

Methods:

- <code title="get /zones/{zone_id}/speed_api/availabilities">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.AvailabilitiesList">AvailabilitiesList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIAvailabilitiesListResponse">SpeedAPIAvailabilitiesListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/pages">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.PagesList">PagesList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIPagesListResponse">SpeedAPIPagesListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/speed_api/schedule/{url}">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.ScheduleDelete">ScheduleDelete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleDeleteParams">SpeedAPIScheduleDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleDeleteResponse">SpeedAPIScheduleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/schedule/{url}">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.ScheduleGet">ScheduleGet</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleGetParams">SpeedAPIScheduleGetParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleGetResponse">SpeedAPIScheduleGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /zones/{zone_id}/speed_api/pages/{url}/tests">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.TestsNew">TestsNew</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsNewParams">SpeedAPITestsNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsNewResponse">SpeedAPITestsNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/speed_api/pages/{url}/tests">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.TestsDelete">TestsDelete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsDeleteParams">SpeedAPITestsDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsDeleteResponse">SpeedAPITestsDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/pages/{url}/tests">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.TestsList">TestsList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsListParams">SpeedAPITestsListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsListResponse">SpeedAPITestsListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/pages/{url}/tests/{test_id}">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.TestsGet">TestsGet</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITestsGetResponse">SpeedAPITestsGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/pages/{url}/trend">client.SpeedAPI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIService.TrendsList">TrendsList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITrendsListParams">SpeedAPITrendsListParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPITrendsListResponse">SpeedAPITrendsListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Schedule

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleNewResponse">SpeedAPIScheduleNewResponse</a>

Methods:

- <code title="post /zones/{zone_id}/speed_api/schedule/{url}">client.SpeedAPI.Schedule.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleNewParams">SpeedAPIScheduleNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SpeedAPIScheduleNewResponse">SpeedAPIScheduleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DcvDelegation

## Uuid

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DcvDelegationUuidGetResponse">DcvDelegationUuidGetResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/dcv_delegation/uuid">client.DcvDelegation.Uuid.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DcvDelegationUuidService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DcvDelegationUuidGetResponse">DcvDelegationUuidGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Hostnames

## Settings

### TLS

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSGetResponse">HostnameSettingTLSGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSUpdateResponse">HostnameSettingTLSUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSDeleteResponse">HostnameSettingTLSDeleteResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/hostnames/settings/{tls_setting}">client.Hostnames.Settings.TLS.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, tlsSetting <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSGetParamsTLSSetting">HostnameSettingTLSGetParamsTLSSetting</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSGetResponse">HostnameSettingTLSGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_identifier}/hostnames/settings/{tls_setting}/{hostname}">client.Hostnames.Settings.TLS.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, tlsSetting <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSUpdateParamsTLSSetting">HostnameSettingTLSUpdateParamsTLSSetting</a>, hostname <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSUpdateParams">HostnameSettingTLSUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSUpdateResponse">HostnameSettingTLSUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_identifier}/hostnames/settings/{tls_setting}/{hostname}">client.Hostnames.Settings.TLS.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, tlsSetting <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSDeleteParamsTLSSetting">HostnameSettingTLSDeleteParamsTLSSetting</a>, hostname <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HostnameSettingTLSDeleteResponse">HostnameSettingTLSDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Logpush

## Datasets

### Fields

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushDatasetFieldListResponse">LogpushDatasetFieldListResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/logpush/datasets/{dataset_id}/fields">client.Logpush.Datasets.Fields.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushDatasetFieldService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushDatasetFieldListResponse">LogpushDatasetFieldListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushDatasetJobListResponse">LogpushDatasetJobListResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/logpush/datasets/{dataset_id}/jobs">client.Logpush.Datasets.Jobs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushDatasetJobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushDatasetJobListResponse">LogpushDatasetJobListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobGetResponse">LogpushJobGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobUpdateResponse">LogpushJobUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobDeleteResponse">LogpushJobDeleteResponse</a>

Methods:

- <code title="get /{account_or_zone}/{account_or_zone_id}/logpush/jobs/{job_id}">client.Logpush.Jobs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, jobID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobGetResponse">LogpushJobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{account_or_zone}/{account_or_zone_id}/logpush/jobs/{job_id}">client.Logpush.Jobs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, jobID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobUpdateParams">LogpushJobUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobUpdateResponse">LogpushJobUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{account_or_zone}/{account_or_zone_id}/logpush/jobs/{job_id}">client.Logpush.Jobs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountOrZone <a href="https://pkg.go.dev/builtin#string">string</a>, accountOrZoneID <a href="https://pkg.go.dev/builtin#string">string</a>, jobID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#LogpushJobDeleteResponse">LogpushJobDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Hold

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldGetResponse">HoldGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldEnforceResponse">HoldEnforceResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldRemoveResponse">HoldRemoveResponse</a>

Methods:

- <code title="get /zones/{zone_id}/hold">client.Hold.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldGetResponse">HoldGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /zones/{zone_id}/hold">client.Hold.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldService.Enforce">Enforce</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldEnforceParams">HoldEnforceParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldEnforceResponse">HoldEnforceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/hold">client.Hold.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldRemoveParams">HoldRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#HoldRemoveResponse">HoldRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PageShield

## Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldConnectionGetResponse">PageShieldConnectionGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/page_shield/connections/{connection_id}">client.PageShield.Connections.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldConnectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldConnectionGetResponse">PageShieldConnectionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Policies

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyNewResponse">PageShieldPolicyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyGetResponse">PageShieldPolicyGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyUpdateResponse">PageShieldPolicyUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyListResponse">PageShieldPolicyListResponse</a>

Methods:

- <code title="post /zones/{zone_id}/page_shield/policies">client.PageShield.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyNewParams">PageShieldPolicyNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyNewResponse">PageShieldPolicyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/page_shield/policies/{policy_id}">client.PageShield.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, policyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyGetResponse">PageShieldPolicyGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_id}/page_shield/policies/{policy_id}">client.PageShield.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, policyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyUpdateParams">PageShieldPolicyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyUpdateResponse">PageShieldPolicyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/page_shield/policies">client.PageShield.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyListResponse">PageShieldPolicyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/page_shield/policies/{policy_id}">client.PageShield.Policies.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#PageShieldPolicyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneID <a href="https://pkg.go.dev/builtin#string">string</a>, policyID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# FontSettings

# Snippets

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetGetResponse">SnippetGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetUpdateResponse">SnippetUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetListResponse">SnippetListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetDeleteResponse">SnippetDeleteResponse</a>

Methods:

- <code title="get /zones/{zone_identifier}/snippets/{snippet_name}">client.Snippets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, snippetName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetGetResponse">SnippetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_identifier}/snippets/{snippet_name}">client.Snippets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, snippetName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetUpdateParams">SnippetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetUpdateResponse">SnippetUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/snippets">client.Snippets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetListResponse">SnippetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_identifier}/snippets/{snippet_name}">client.Snippets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, snippetName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetDeleteResponse">SnippetDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Content

Methods:

- <code title="get /zones/{zone_identifier}/snippets/{snippet_name}/content">client.Snippets.Content.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetContentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, snippetName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SnippetRules

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetSnippetRuleUpdateResponse">SnippetSnippetRuleUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetSnippetRuleListResponse">SnippetSnippetRuleListResponse</a>

Methods:

- <code title="put /zones/{zone_identifier}/snippets/snippet_rules">client.Snippets.SnippetRules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetSnippetRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetSnippetRuleUpdateParams">SnippetSnippetRuleUpdateParams</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetSnippetRuleUpdateResponse">SnippetSnippetRuleUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_identifier}/snippets/snippet_rules">client.Snippets.SnippetRules.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetSnippetRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, zoneIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#SnippetSnippetRuleListResponse">SnippetSnippetRuleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Dlp

## Datasets

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetNewResponse">DlpDatasetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetGetResponse">DlpDatasetGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetUpdateResponse">DlpDatasetUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetListResponse">DlpDatasetListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetUploadResponse">DlpDatasetUploadResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetUploadPrepareResponse">DlpDatasetUploadPrepareResponse</a>

Methods:

- <code title="post /accounts/{account_id}/dlp/datasets">client.Dlp.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetNewParams">DlpDatasetNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetNewResponse">DlpDatasetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/dlp/datasets/{dataset_id}">client.Dlp.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetGetResponse">DlpDatasetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_id}/dlp/datasets/{dataset_id}">client.Dlp.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetUpdateParams">DlpDatasetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetUpdateResponse">DlpDatasetUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/dlp/datasets">client.Dlp.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetListResponse">DlpDatasetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/dlp/datasets/{dataset_id}">client.Dlp.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /accounts/{account_id}/dlp/datasets/{dataset_id}/upload/{version}">client.Dlp.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, version <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetUploadResponse">DlpDatasetUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_id}/dlp/datasets/{dataset_id}/upload">client.Dlp.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetService.UploadPrepare">UploadPrepare</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#DlpDatasetUploadPrepareResponse">DlpDatasetUploadPrepareResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Gateway

## AuditSSHSettings

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#GatewayAuditSSHSettingGetResponse">GatewayAuditSSHSettingGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#GatewayAuditSSHSettingUpdateResponse">GatewayAuditSSHSettingUpdateResponse</a>

Methods:

- <code title="get /accounts/{account_id}/gateway/audit_ssh_settings">client.Gateway.AuditSSHSettings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#GatewayAuditSSHSettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#GatewayAuditSSHSettingGetResponse">GatewayAuditSSHSettingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_id}/gateway/audit_ssh_settings">client.Gateway.AuditSSHSettings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#GatewayAuditSSHSettingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#GatewayAuditSSHSettingUpdateParams">GatewayAuditSSHSettingUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#GatewayAuditSSHSettingUpdateResponse">GatewayAuditSSHSettingUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccessTags

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagNewResponse">AccessTagNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagGetResponse">AccessTagGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagUpdateResponse">AccessTagUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagDeleteResponse">AccessTagDeleteResponse</a>

Methods:

- <code title="post /accounts/{identifier}/access/tags">client.AccessTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagNewParams">AccessTagNewParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagNewResponse">AccessTagNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{identifier}/access/tags/{name}">client.AccessTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagGetResponse">AccessTagGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{identifier}/access/tags/{name}">client.AccessTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, tagName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagUpdateParams">AccessTagUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagUpdateResponse">AccessTagUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{identifier}/access/tags/{name}">client.AccessTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go">cloudflare</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go#AccessTagDeleteResponse">AccessTagDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
