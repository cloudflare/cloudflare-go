# Logs

## LogExplorer

### Query

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerQuerySqlResponse">LogExplorerQuerySqlResponse</a>

Methods:

- <code title="post /{accounts_or_zones}/{account_or_zone_id}/logs/explorer/query/sql">client.Logs.LogExplorer.Query.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerQueryService.Sql">Sql</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerQuerySqlParams">LogExplorerQuerySqlParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerQuerySqlResponse">LogExplorerQuerySqlResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Datasets

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#CreateRequestParam">CreateRequestParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#UpdateRequestParam">UpdateRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#Dataset">Dataset</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#DatasetSummary">DatasetSummary</a>

Methods:

- <code title="post /{accounts_or_zones}/{account_or_zone_id}/logs/explorer/datasets">client.Logs.LogExplorer.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetNewParams">LogExplorerDatasetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /{accounts_or_zones}/{account_or_zone_id}/logs/explorer/datasets/{dataset_id}">client.Logs.LogExplorer.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetUpdateParams">LogExplorerDatasetUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{accounts_or_zones}/{account_or_zone_id}/logs/explorer/datasets">client.Logs.LogExplorer.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetListParams">LogExplorerDatasetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#DatasetSummary">DatasetSummary</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{accounts_or_zones}/{account_or_zone_id}/logs/explorer/datasets/{dataset_id}">client.Logs.LogExplorer.Datasets.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetGetParams">LogExplorerDatasetGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Available

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#AvailableDataset">AvailableDataset</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#AvailableList">AvailableList</a>

Methods:

- <code title="get /{accounts_or_zones}/{account_or_zone_id}/logs/explorer/datasets/available">client.Logs.LogExplorer.Datasets.Available.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetAvailableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#LogExplorerDatasetAvailableListParams">LogExplorerDatasetAvailableListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#AvailableDataset">AvailableDataset</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Control

### Retention

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlRetentionNewResponse">ControlRetentionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlRetentionGetResponse">ControlRetentionGetResponse</a>

Methods:

- <code title="post /zones/{zone_id}/logs/control/retention/flag">client.Logs.Control.Retention.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlRetentionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlRetentionNewParams">ControlRetentionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlRetentionNewResponse">ControlRetentionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/logs/control/retention/flag">client.Logs.Control.Retention.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlRetentionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlRetentionGetParams">ControlRetentionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlRetentionGetResponse">ControlRetentionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Cmb

#### Config

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#CmbConfigParam">CmbConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#CmbConfig">CmbConfig</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlCmbConfigDeleteResponse">ControlCmbConfigDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_id}/logs/control/cmb/config">client.Logs.Control.Cmb.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlCmbConfigService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlCmbConfigNewParams">ControlCmbConfigNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#CmbConfig">CmbConfig</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/logs/control/cmb/config">client.Logs.Control.Cmb.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlCmbConfigService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlCmbConfigDeleteParams">ControlCmbConfigDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlCmbConfigDeleteResponse">ControlCmbConfigDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/logs/control/cmb/config">client.Logs.Control.Cmb.Config.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlCmbConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ControlCmbConfigGetParams">ControlCmbConfigGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#CmbConfig">CmbConfig</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RayID

Methods:

- <code title="get /zones/{zone_id}/logs/rayids/{ray_id}">client.Logs.RayID.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#RayIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, RayID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#RayIDGetParams">RayIDGetParams</a>) (\*interface{}, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Received

Methods:

- <code title="get /zones/{zone_id}/logs/received">client.Logs.Received.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ReceivedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ReceivedGetParams">ReceivedGetParams</a>) (\*interface{}, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Fields

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ReceivedFieldGetResponse">ReceivedFieldGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/logs/received/fields">client.Logs.Received.Fields.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ReceivedFieldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ReceivedFieldGetParams">ReceivedFieldGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs">logs</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/logs#ReceivedFieldGetResponse">ReceivedFieldGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
