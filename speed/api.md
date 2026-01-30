# Speed

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#LabeledRegion">LabeledRegion</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#LighthouseReport">LighthouseReport</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Trend">Trend</a>

## Schedule

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Schedule">Schedule</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleNewResponse">ScheduleNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleDeleteResponse">ScheduleDeleteResponse</a>

Methods:

- <code title="post /zones/{zone_id}/speed_api/schedule/{url}">client.Speed.Schedule.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleNewParams">ScheduleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleNewResponse">ScheduleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/speed_api/schedule/{url}">client.Speed.Schedule.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleDeleteParams">ScheduleDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleDeleteResponse">ScheduleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/schedule/{url}">client.Speed.Schedule.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#ScheduleGetParams">ScheduleGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Schedule">Schedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Availabilities

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Availability">Availability</a>

Methods:

- <code title="get /zones/{zone_id}/speed_api/availabilities">client.Speed.Availabilities.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#AvailabilityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#AvailabilityListParams">AvailabilityListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Availability">Availability</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Pages

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageListResponse">PageListResponse</a>

Methods:

- <code title="get /zones/{zone_id}/speed_api/pages">client.Speed.Pages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageListParams">PageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageListResponse">PageListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/pages/{url}/trend">client.Speed.Pages.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageService.Trend">Trend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTrendParams">PageTrendParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Trend">Trend</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Tests

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Test">Test</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestDeleteResponse">PageTestDeleteResponse</a>

Methods:

- <code title="post /zones/{zone_id}/speed_api/pages/{url}/tests">client.Speed.Pages.Tests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestNewParams">PageTestNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Test">Test</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/pages/{url}/tests">client.Speed.Pages.Tests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestListParams">PageTestListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Test">Test</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/speed_api/pages/{url}/tests">client.Speed.Pages.Tests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestDeleteParams">PageTestDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestDeleteResponse">PageTestDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/speed_api/pages/{url}/tests/{test_id}">client.Speed.Pages.Tests.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#PageTestGetParams">PageTestGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed">speed</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/speed#Test">Test</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
