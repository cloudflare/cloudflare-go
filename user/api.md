# User

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#UserEditResponse">UserEditResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#UserGetResponse">UserGetResponse</a>

Methods:

- <code title="patch /user">client.User.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#UserService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#UserEditParams">UserEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#UserEditResponse">UserEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user">client.User.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#UserGetResponse">UserGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AuditLogs

Methods:

- <code title="get /user/audit_logs">client.User.AuditLogs.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#AuditLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#AuditLogListParams">AuditLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#AuditLog">AuditLog</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Billing

### History

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#BillingHistory">BillingHistory</a>

Methods:

- <code title="get /user/billing/history">client.User.Billing.History.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#BillingHistoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#BillingHistoryListParams">BillingHistoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#BillingHistory">BillingHistory</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Profile

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#BillingProfileGetResponse">BillingProfileGetResponse</a>

Methods:

- <code title="get /user/billing/profile">client.User.Billing.Profile.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#BillingProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#BillingProfileGetResponse">BillingProfileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Invites

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#Invite">Invite</a>

Methods:

- <code title="get /user/invites">client.User.Invites.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#InviteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#Invite">Invite</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /user/invites/{invite_id}">client.User.Invites.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#InviteService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inviteID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#InviteEditParams">InviteEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#Invite">Invite</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/invites/{invite_id}">client.User.Invites.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#InviteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inviteID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#Invite">Invite</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Organizations

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#Organization">Organization</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#OrganizationDeleteResponse">OrganizationDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#OrganizationGetResponse">OrganizationGetResponse</a>

Methods:

- <code title="get /user/organizations">client.User.Organizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#OrganizationListParams">OrganizationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#Organization">Organization</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /user/organizations/{organization_id}">client.User.Organizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#OrganizationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#OrganizationDeleteResponse">OrganizationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/organizations/{organization_id}">client.User.Organizations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#OrganizationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#OrganizationGetResponse">OrganizationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>

Methods:

- <code title="put /user/subscriptions/{identifier}">client.User.Subscriptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#SubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#SubscriptionUpdateParams">SubscriptionUpdateParams</a>) (\*interface{}, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /user/subscriptions/{identifier}">client.User.Subscriptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#SubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/subscriptions">client.User.Subscriptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#SubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#Subscription">Subscription</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tokens

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenNewResponse">TokenNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenDeleteResponse">TokenDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenVerifyResponse">TokenVerifyResponse</a>

Methods:

- <code title="post /user/tokens">client.User.Tokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenNewParams">TokenNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenNewResponse">TokenNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/tokens/{token_id}">client.User.Tokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenUpdateParams">TokenUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#Token">Token</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/tokens">client.User.Tokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenListParams">TokenListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#Token">Token</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /user/tokens/{token_id}">client.User.Tokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenDeleteResponse">TokenDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/tokens/{token_id}">client.User.Tokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#Token">Token</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/tokens/verify">client.User.Tokens.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenVerifyResponse">TokenVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### PermissionGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenPermissionGroupListResponse">TokenPermissionGroupListResponse</a>

Methods:

- <code title="get /user/tokens/permission_groups">client.User.Tokens.PermissionGroups.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenPermissionGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenPermissionGroupListParams">TokenPermissionGroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenPermissionGroupListResponse">TokenPermissionGroupListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Value

Methods:

- <code title="put /user/tokens/{token_id}/value">client.User.Tokens.Value.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenValueService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user">user</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/user#TokenValueUpdateParams">TokenValueUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/shared#TokenValue">TokenValue</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
