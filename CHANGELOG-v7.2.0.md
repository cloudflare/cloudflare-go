## 7.2.0 (2026-05-06)

Full Changelog: [v7.1.0...v7.2.0](https://github.com/cloudflare/cloudflare-go/compare/v7.1.0...v7.2.0)

This release adds new sub-resources for AI Gateway billing, Cache Origin Cloud Regions v2, Radar BGP analytics, and Load Balancer monitor group references.

---

## Features

### Cache - Origin Cloud Regions v2 (`client.Cache.OriginCloudRegions`)

- **NEW SUB-RESOURCE**: Origin Cloud Regions service using the v2 `/origin/cloud_regions` endpoints
  - `Update()` -- create or update a cloud region mapping for an origin IP
  - `List()` -- list all cloud region mappings (paginated)
  - `Delete()` -- remove a cloud region mapping
  - `BulkDelete()` -- remove multiple cloud region mappings
  - `BulkUpdate()` -- create or update multiple cloud region mappings
  - `Get()` -- get a specific cloud region mapping
  - `SupportedRegions()` -- list supported cloud regions and vendors

### AI Gateway - Billing (`client.AIGateway.Billing`)

- **NEW SUB-RESOURCE**: Billing service for AI Gateway accounts
  - `CreditBalance()` -- get current credit balance
  - `UsageHistory()` -- get historical usage data
  - `InvoiceHistory()` -- get past invoices
  - `InvoicePreview()` -- preview upcoming invoice

### AI Gateway - Billing Spending Limit (`client.AIGateway.Billing.SpendingLimit`)

- **NEW SUB-RESOURCE**: Manage spending limits for AI Gateway billing
  - `New()` -- create a spending limit
  - `Delete()` -- remove a spending limit
  - `Get()` -- get the current spending limit

### AI Gateway - Billing Topup (`client.AIGateway.Billing.Topup`)

- **NEW SUB-RESOURCE**: Manage billing top-ups for AI Gateway
  - `New()` -- create a top-up
  - `Status()` -- check top-up payment status

### AI Gateway - Billing Topup Config (`client.AIGateway.Billing.Topup.Config`)

- **NEW SUB-RESOURCE**: Configure automatic top-ups
  - `New()` -- create a top-up config
  - `Delete()` -- remove a top-up config
  - `Get()` -- get the current top-up config

### Cache - Smart Tiered Cache (`client.Cache.SmartTieredCache`)

- Added `New()` method (POST) to enable smart tiered cache topology

### Radar - BGP IP Top (`client.Radar.BGP.IPs.Top`)

- **NEW SUB-RESOURCE**: Top ASes by BGP IP announcements
  - `Ases()` -- get top ASes for BGP IP data

### Radar - BGP RPKI ROAs (`client.Radar.BGP.RPKI.Roas`)

- **NEW SUB-RESOURCE**: RPKI Route Origin Authorization timeseries
  - `Timeseries()` -- get ROA timeseries data

### Load Balancers - Monitor Group References (`client.LoadBalancers.MonitorGroups.References`)

- **NEW SUB-RESOURCE**: Retrieve resources referencing a monitor group
  - `Get()` -- list references for a monitor group

---

## Bug Fixes

- **Codegen**: Updated generated types and methods for `ai_search`, `cloudforce_one`, `resource_sharing`, `url_scanner`, and `user` packages
