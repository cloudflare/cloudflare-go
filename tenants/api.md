# Tenants

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants">tenants</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#Tenant">Tenant</a>

Methods:

- <code title="get /tenants/{tenant_id}">client.Tenants.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#TenantService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants">tenants</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#Tenant">Tenant</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountTypes

Methods:

- <code title="get /tenants/{tenant_id}/account_types">client.Tenants.AccountTypes.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#AccountTypeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/builtin#string">string</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants">tenants</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#TenantAccount">TenantAccount</a>

Methods:

- <code title="get /tenants/{tenant_id}/accounts">client.Tenants.Accounts.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants">tenants</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#TenantAccount">TenantAccount</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Entitlements

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants">tenants</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#TenantEntitlements">TenantEntitlements</a>

Methods:

- <code title="get /tenants/{tenant_id}/entitlements">client.Tenants.Entitlements.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#EntitlementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants">tenants</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#TenantEntitlements">TenantEntitlements</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Memberships

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants">tenants</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#TenantMembership">TenantMembership</a>

Methods:

- <code title="get /tenants/{tenant_id}/memberships">client.Tenants.Memberships.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#MembershipService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tenantID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants">tenants</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/tenants#TenantMembership">TenantMembership</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
