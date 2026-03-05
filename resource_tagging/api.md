# ResourceTagging

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ResourceTaggingListResponse">ResourceTaggingListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/tags/resources">client.ResourceTagging.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ResourceTaggingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ResourceTaggingListParams">ResourceTaggingListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#CursorPaginationAfter">CursorPaginationAfter</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ResourceTaggingListResponse">ResourceTaggingListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountTags

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagUpdateResponse">AccountTagUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagGetResponse">AccountTagGetResponse</a>

Methods:

- <code title="put /accounts/{account_id}/tags">client.ResourceTagging.AccountTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagUpdateParams">AccountTagUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagUpdateResponse">AccountTagUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/tags">client.ResourceTagging.AccountTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagDeleteParams">AccountTagDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /accounts/{account_id}/tags">client.ResourceTagging.AccountTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagGetParams">AccountTagGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#AccountTagGetResponse">AccountTagGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ZoneTags

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagUpdateResponse">ZoneTagUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagGetResponse">ZoneTagGetResponse</a>

Methods:

- <code title="put /zones/{zone_id}/tags">client.ResourceTagging.ZoneTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagUpdateParams">ZoneTagUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagUpdateResponse">ZoneTagUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/tags">client.ResourceTagging.ZoneTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagDeleteParams">ZoneTagDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /zones/{zone_id}/tags">client.ResourceTagging.ZoneTags.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagGetParams">ZoneTagGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ZoneTagGetResponse">ZoneTagGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Keys

Methods:

- <code title="get /accounts/{account_id}/tags/keys">client.ResourceTagging.Keys.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#KeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#KeyListParams">KeyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#CursorPaginationAfter">CursorPaginationAfter</a>[<a href="https://pkg.go.dev/builtin#string">string</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Values

Methods:

- <code title="get /accounts/{account_id}/tags/values/{tag_key}">client.ResourceTagging.Values.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ValueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tagKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging">resource_tagging</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/resource_tagging#ValueListParams">ValueListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#CursorPaginationAfter">CursorPaginationAfter</a>[<a href="https://pkg.go.dev/builtin#string">string</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
