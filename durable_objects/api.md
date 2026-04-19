# DurableObjects

## Namespaces

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects">durable_objects</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects#Namespace">Namespace</a>

Methods:

- <code title="get /accounts/{account_id}/workers/durable_objects/namespaces">client.DurableObjects.Namespaces.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects#NamespaceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects">durable_objects</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects#NamespaceListParams">NamespaceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects">durable_objects</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects#Namespace">Namespace</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Objects

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects">durable_objects</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects#DurableObject">DurableObject</a>

Methods:

- <code title="get /accounts/{account_id}/workers/durable_objects/namespaces/{id}/objects">client.DurableObjects.Namespaces.Objects.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects#NamespaceObjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects">durable_objects</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects#NamespaceObjectListParams">NamespaceObjectListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#CursorPaginationAfter">CursorPaginationAfter</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects">durable_objects</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/durable_objects#DurableObject">DurableObject</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
