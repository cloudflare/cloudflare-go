# Billing

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#BillingAddressValidationResponse">BillingAddressValidationResponse</a>

Methods:

- <code title="post /billing/address-validation">client.Billing.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#BillingService.AddressValidation">AddressValidation</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#BillingAddressValidationParams">BillingAddressValidationParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#BillingAddressValidationResponse">BillingAddressValidationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileNewResponse">ProfileNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileUpdateResponse">ProfileUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileGetResponse">ProfileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileUpdateBillingEmailResponse">ProfileUpdateBillingEmailResponse</a>

Methods:

- <code title="post /accounts/{account_id}/billing/profile">client.Billing.Profiles.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileNewParams">ProfileNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileNewResponse">ProfileNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{account_id}/billing/profile">client.Billing.Profiles.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileUpdateParams">ProfileUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileUpdateResponse">ProfileUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/billing/profile">client.Billing.Profiles.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileDeleteParams">ProfileDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /accounts/{account_id}/billing/profile">client.Billing.Profiles.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileGetParams">ProfileGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileGetResponse">ProfileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/billing/profile">client.Billing.Profiles.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileService.UpdateBillingEmail">UpdateBillingEmail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileUpdateBillingEmailParams">ProfileUpdateBillingEmailParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfileUpdateBillingEmailResponse">ProfileUpdateBillingEmailResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### PaymentMethod

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfilePaymentMethodNewResponse">ProfilePaymentMethodNewResponse</a>

Methods:

- <code title="post /accounts/{account_id}/billing/profile/payment-method">client.Billing.Profiles.PaymentMethod.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfilePaymentMethodService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfilePaymentMethodNewParams">ProfilePaymentMethodNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#ProfilePaymentMethodNewResponse">ProfilePaymentMethodNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Usage

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetResponse">UsageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageInfoV1Response">UsageGetAccountUsageInfoV1Response</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageV1Response">UsageGetAccountUsageV1Response</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageV2Response">UsageGetAccountUsageV2Response</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoResponse">UsagePaygoResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoInfoResponse">UsagePaygoInfoResponse</a>

Methods:

- <code title="get /accounts/{account_id}/billable/usage">client.Billing.Usage.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetParams">UsageGetParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetResponse">UsageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/billable-usage/info">client.Billing.Usage.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageService.GetAccountUsageInfoV1">GetAccountUsageInfoV1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageInfoV1Params">UsageGetAccountUsageInfoV1Params</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageInfoV1Response">UsageGetAccountUsageInfoV1Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/billable-usage">client.Billing.Usage.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageService.GetAccountUsageV1">GetAccountUsageV1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageV1Params">UsageGetAccountUsageV1Params</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageV1Response">UsageGetAccountUsageV1Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/billable/usage">client.Billing.Usage.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageService.GetAccountUsageV2">GetAccountUsageV2</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageV2Params">UsageGetAccountUsageV2Params</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageGetAccountUsageV2Response">UsageGetAccountUsageV2Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/billable-usage">client.Billing.Usage.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageService.Paygo">Paygo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoParams">UsagePaygoParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoResponse">UsagePaygoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/billable-usage/info">client.Billing.Usage.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsageService.PaygoInfo">PaygoInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoInfoParams">UsagePaygoInfoParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UsagePaygoInfoResponse">UsagePaygoInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Credits

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#CreditGetResponse">CreditGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/billing/credits">client.Billing.Credits.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#CreditService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#CreditGetParams">CreditGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#CreditGetResponse">CreditGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## History

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#HistoryListResponse">HistoryListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/billing/history">client.Billing.History.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#HistoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#HistoryListParams">HistoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#HistoryListResponse">HistoryListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BadDebt

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#BadDebtGetResponse">BadDebtGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/billing/bad-debt">client.Billing.BadDebt.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#BadDebtService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#BadDebtGetParams">BadDebtGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#BadDebtGetResponse">BadDebtGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## UnpaidInvoice

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UnpaidInvoiceGetResponse">UnpaidInvoiceGetResponse</a>

Methods:

- <code title="get /accounts/{account_id}/billing/unpaid-invoice">client.Billing.UnpaidInvoice.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UnpaidInvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UnpaidInvoiceGetParams">UnpaidInvoiceGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#UnpaidInvoiceGetResponse">UnpaidInvoiceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RatePlans

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#RatePlanGetResponse">RatePlanGetResponse</a>

Methods:

- <code title="get /billing/rate_plans/{public_key}">client.Billing.RatePlans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#RatePlanService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, publicKey <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing">billing</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/billing#RatePlanGetResponse">RatePlanGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
