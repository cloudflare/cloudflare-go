# AI

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#AIRunResponseUnion">AIRunResponseUnion</a>

Methods:

- <code title="post /accounts/{account_id}/ai/run/{model_name}">client.AI.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#AIService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#AIRunParams">AIRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#AIRunResponseUnion">AIRunResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Finetunes

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneNewResponse">FinetuneNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneListResponse">FinetuneListResponse</a>

Methods:

- <code title="post /accounts/{account_id}/ai/finetunes">client.AI.Finetunes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneNewParams">FinetuneNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneNewResponse">FinetuneNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/ai/finetunes">client.AI.Finetunes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneListParams">FinetuneListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneListResponse">FinetuneListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneAssetNewResponse">FinetuneAssetNewResponse</a>

Methods:

- <code title="post /accounts/{account_id}/ai/finetunes/{finetune_id}/finetune-assets">client.AI.Finetunes.Assets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneAssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, finetuneID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneAssetNewParams">FinetuneAssetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetuneAssetNewResponse">FinetuneAssetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Public

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetunePublicListResponse">FinetunePublicListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/ai/finetunes/public">client.AI.Finetunes.Public.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetunePublicService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetunePublicListParams">FinetunePublicListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#FinetunePublicListResponse">FinetunePublicListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Authors

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#AuthorListResponse">AuthorListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/ai/authors/search">client.AI.Authors.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#AuthorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#AuthorListParams">AuthorListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#AuthorListResponse">AuthorListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tasks

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#TaskListResponse">TaskListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/ai/tasks/search">client.AI.Tasks.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#TaskService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#TaskListParams">TaskListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#TaskListResponse">TaskListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Models

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/ai/models/search">client.AI.Models.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ModelListParams">ModelListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ModelListResponse">ModelListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Schema

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ModelSchemaGetResponse">ModelSchemaGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/ai/models/schema">client.AI.Models.Schema.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ModelSchemaService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ModelSchemaGetParams">ModelSchemaGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ModelSchemaGetResponse">ModelSchemaGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ToMarkdown

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ToMarkdownSupportedResponse">ToMarkdownSupportedResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ToMarkdownTransformResponse">ToMarkdownTransformResponse</a>

Methods:

- <code title="get /accounts/{account_id}/ai/tomarkdown/supported">client.AI.ToMarkdown.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ToMarkdownService.Supported">Supported</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ToMarkdownSupportedParams">ToMarkdownSupportedParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ToMarkdownSupportedResponse">ToMarkdownSupportedResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_id}/ai/tomarkdown">client.AI.ToMarkdown.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ToMarkdownService.Transform">Transform</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ToMarkdownTransformParams">ToMarkdownTransformParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai">ai</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/ai#ToMarkdownTransformResponse">ToMarkdownTransformResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
