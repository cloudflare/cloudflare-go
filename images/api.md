# Images

## V1

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Image">Image</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1ListResponse">V1ListResponse</a>

Methods:

- <code title="post /accounts/{account_id}/images/v1">client.Images.V1.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1Service.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1NewParams">V1NewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/images/v1">client.Images.V1.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1ListParams">V1ListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePagination">V4PagePagination</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1ListResponse">V1ListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/images/v1/{image_id}">client.Images.V1.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1Service.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1DeleteParams">V1DeleteParams</a>) (\*interface{}, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/images/v1/{image_id}">client.Images.V1.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1Service.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1EditParams">V1EditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/images/v1/{image_id}">client.Images.V1.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1Service.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1GetParams">V1GetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Keys

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Key">Key</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyUpdateResponse">V1KeyUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyListResponse">V1KeyListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyDeleteResponse">V1KeyDeleteResponse</a>

Methods:

- <code title="put /accounts/{account_id}/images/v1/keys/{signing_key_name}">client.Images.V1.Keys.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, signingKeyName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyUpdateParams">V1KeyUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyUpdateResponse">V1KeyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/images/v1/keys">client.Images.V1.Keys.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyListParams">V1KeyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyListResponse">V1KeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/images/v1/keys/{signing_key_name}">client.Images.V1.Keys.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, signingKeyName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyDeleteParams">V1KeyDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1KeyDeleteResponse">V1KeyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Stats

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Stat">Stat</a>

Methods:

- <code title="get /accounts/{account_id}/images/v1/stats">client.Images.V1.Stats.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1StatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1StatGetParams">V1StatGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Stat">Stat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Variants

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Variant">Variant</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantNewResponse">V1VariantNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantEditResponse">V1VariantEditResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantGetResponse">V1VariantGetResponse</a>

Methods:

- <code title="post /accounts/{account_id}/images/v1/variants">client.Images.V1.Variants.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantNewParams">V1VariantNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantNewResponse">V1VariantNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/images/v1/variants">client.Images.V1.Variants.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantListParams">V1VariantListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#Variant">Variant</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/images/v1/variants/{variant_id}">client.Images.V1.Variants.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, variantID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantDeleteParams">V1VariantDeleteParams</a>) (\*interface{}, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/images/v1/variants/{variant_id}">client.Images.V1.Variants.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, variantID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantEditParams">V1VariantEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantEditResponse">V1VariantEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/images/v1/variants/{variant_id}">client.Images.V1.Variants.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, variantID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantGetParams">V1VariantGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1VariantGetResponse">V1VariantGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Blobs

Methods:

- <code title="get /accounts/{account_id}/images/v1/{image_id}/blob">client.Images.V1.Blobs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1BlobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V1BlobGetParams">V1BlobGetParams</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## V2

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V2ListResponse">V2ListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/images/v2">client.Images.V2.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V2Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V2ListParams">V2ListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V2ListResponse">V2ListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### DirectUploads

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V2DirectUploadNewResponse">V2DirectUploadNewResponse</a>

Methods:

- <code title="post /accounts/{account_id}/images/v2/direct_upload">client.Images.V2.DirectUploads.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V2DirectUploadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V2DirectUploadNewParams">V2DirectUploadNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images">images</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/images#V2DirectUploadNewResponse">V2DirectUploadNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
