## 0.64.0 (Unreleased)

## 0.63.0 (15th March, 2023)

BREAKING CHANGES:

* tunnel: renamed `Tunnel` to `GetTunnel` ([#1227](https://github.com/cloudflare/cloudflare-go/issues/1227))
* tunnel: renamed `Tunnels` to `ListTunnels` ([#1227](https://github.com/cloudflare/cloudflare-go/issues/1227))

ENHANCEMENTS:

* access_organization: add ui_read_only_toggle_reason field ([#1181](https://github.com/cloudflare/cloudflare-go/issues/1181))
* added audit_ssh to gateway actions, updated gateway rule settings ([#1226](https://github.com/cloudflare/cloudflare-go/issues/1226))
* addressing: Add `Address Map` support ([#1232](https://github.com/cloudflare/cloudflare-go/issues/1232))
* teams_account: add support for `check_disks` ([#1197](https://github.com/cloudflare/cloudflare-go/issues/1197))
* tunnel: updated parameters to latest API docs ([#1227](https://github.com/cloudflare/cloudflare-go/issues/1227))

DEPENDENCIES:

* deps: bumps github.com/urfave/cli/v2 from 2.24.4 to 2.25.0 ([#1229](https://github.com/cloudflare/cloudflare-go/issues/1229))
* deps: bumps golang.org/x/net from 0.7.0 to 0.8.0 ([#1228](https://github.com/cloudflare/cloudflare-go/issues/1228))

## 0.62.0 (1st March, 2023)

ENHANCEMENTS:

* dex_test: add CRUD functionality for DEX test configurations ([#1209](https://github.com/cloudflare/cloudflare-go/issues/1209))
* dlp: Adds support for partial payload logging ([#1212](https://github.com/cloudflare/cloudflare-go/issues/1212))
* teams_accounts: Add new root_certificate_installation_enabled field ([#1208](https://github.com/cloudflare/cloudflare-go/issues/1208))
* teams_rules: Add `untrusted_cert` rule setting ([#1214](https://github.com/cloudflare/cloudflare-go/issues/1214))
* tunnels: automatically paginate `ListTunnels` ([#1206](https://github.com/cloudflare/cloudflare-go/issues/1206))

BUG FIXES:

* dex_test: use dex test types and json struct mappings instead of managed networks ([#1213](https://github.com/cloudflare/cloudflare-go/issues/1213))
* dns: dont reuse DNSListResponse when using pagination to avoid Proxied pointer overwrite ([#1222](https://github.com/cloudflare/cloudflare-go/issues/1222))

DEPENDENCIES:

* deps: bumps github.com/stretchr/testify from 1.8.1 to 1.8.2 ([#1220](https://github.com/cloudflare/cloudflare-go/issues/1220))
* deps: bumps github.com/urfave/cli/v2 from 2.24.3 to 2.24.4 ([#1210](https://github.com/cloudflare/cloudflare-go/issues/1210))
* deps: bumps golang.org/x/net from 0.0.0-20220722155237-a158d28d115b to 0.7.0 ([#1218](https://github.com/cloudflare/cloudflare-go/issues/1218))
* deps: bumps golang.org/x/net from 0.0.0-20220722155237-a158d28d115b to 0.7.0 ([#1219](https://github.com/cloudflare/cloudflare-go/issues/1219))
* deps: bumps golang.org/x/text from 0.3.7 to 0.3.8 ([#1215](https://github.com/cloudflare/cloudflare-go/issues/1215))
* deps: bumps golang.org/x/text from 0.3.7 to 0.3.8 ([#1216](https://github.com/cloudflare/cloudflare-go/issues/1216))
* deps: bumps golang.org/x/time from 0.0.0-20220224211638-0e9765cccd65 to 0.3.0 ([#1217](https://github.com/cloudflare/cloudflare-go/issues/1217))

## 0.61.0 (15th February, 2023)

ENHANCEMENTS:

* cloudflare: make it clearer when we hit a server error and to retry later ([#1207](https://github.com/cloudflare/cloudflare-go/issues/1207))
* devices_policy: Add new exclude_office_ips field to policy ([#1205](https://github.com/cloudflare/cloudflare-go/issues/1205))
* dlp_profile: Use int rather than uint for allowed_match_count field ([#1200](https://github.com/cloudflare/cloudflare-go/issues/1200))

BUG FIXES:

* dns: always send `tags` to allow clearing ([#1196](https://github.com/cloudflare/cloudflare-go/issues/1196))
* stream: renamed `RequiredSignedURLs` to `RequireSignedURLs` ([#1202](https://github.com/cloudflare/cloudflare-go/issues/1202))

DEPENDENCIES:

* deps: bumps github.com/urfave/cli/v2 from 2.24.2 to 2.24.3 ([#1199](https://github.com/cloudflare/cloudflare-go/issues/1199))

## 0.60.0 (1st February, 2023)

BREAKING CHANGES:

* queues: UpdateQueue has been updated to match the API and now correctly updates a Queue's name ([#1188](https://github.com/cloudflare/cloudflare-go/issues/1188))

ENHANCEMENTS:

* dlp_profile: Add new allowed_match_count field to profiles ([#1193](https://github.com/cloudflare/cloudflare-go/issues/1193))
* dns: allow sending empty strings to remove comments ([#1195](https://github.com/cloudflare/cloudflare-go/issues/1195))
* magic_transit_ipsec_tunnel: makes customer endpoint an optional field for ipsec tunnel creation ([#1185](https://github.com/cloudflare/cloudflare-go/issues/1185))
* rulesets: add support for `score_per_period` and `score_response_header_name` ([#1183](https://github.com/cloudflare/cloudflare-go/issues/1183))

DEPENDENCIES:

* deps: bumps dependabot/fetch-metadata from 1.3.5 to 1.3.6 ([#1184](https://github.com/cloudflare/cloudflare-go/issues/1184))
* deps: bumps github.com/urfave/cli/v2 from 2.23.7 to 2.24.1 ([#1180](https://github.com/cloudflare/cloudflare-go/issues/1180))
* deps: bumps github.com/urfave/cli/v2 from 2.24.1 to 2.24.2 ([#1191](https://github.com/cloudflare/cloudflare-go/issues/1191))
* deps: bumps goreleaser/goreleaser-action from 4.1.0 to 4.2.0 ([#1192](https://github.com/cloudflare/cloudflare-go/issues/1192))

## 0.59.0 (January 18th, 2023)

BREAKING CHANGES:

* dns: remove these read-only fields from `UpdateDNSRecordParams`: `CreatedOn`, `ModifiedOn`, `Meta`, `ZoneID`, `ZoneName`, `Proxiable`, and `Locked` ([#1170](https://github.com/cloudflare/cloudflare-go/issues/1170))
* dns: the fields `CreatedOn` and `ModifiedOn` are removed from `ListDNSRecordsParams` ([#1173](https://github.com/cloudflare/cloudflare-go/issues/1173))

NOTES:

* dns: remove additional lookup from `Update` operations when `Name` or `Type` was omitted ([#1170](https://github.com/cloudflare/cloudflare-go/issues/1170))

ENHANCEMENTS:

* access_organization: add user_seat_expiration_inactive_time field ([#1159](https://github.com/cloudflare/cloudflare-go/issues/1159))
* dns: `GetDNSRecord`, `UpdateDNSRecord` and `DeleteDNSRecord` now return the new, dedicated error `ErrMissingDNSRecordID` when an empty DNS record ID is given. ([#1174](https://github.com/cloudflare/cloudflare-go/issues/1174))
* dns: the URL parameter `tag-match` for listing DNS records is now supported as the field `TagMatch` in `ListDNSRecordsParams` ([#1173](https://github.com/cloudflare/cloudflare-go/issues/1173))
* dns: update default `per_page` attribute to 100 records ([#1171](https://github.com/cloudflare/cloudflare-go/issues/1171))
* teams_rules: adds support for Egress Policies ([#1142](https://github.com/cloudflare/cloudflare-go/issues/1142))
* workers: Add support for compatibility_date and compatibility_flags when upoading a worker script ([#1177](https://github.com/cloudflare/cloudflare-go/issues/1177))
* workers: script upload now supports Queues bindings ([#1176](https://github.com/cloudflare/cloudflare-go/issues/1176))

BUG FIXES:

* dns: don't send "priority" for list operations as it isn't supported and is only used for internal filtering ([#1167](https://github.com/cloudflare/cloudflare-go/issues/1167))
* dns: the field `Tags` in `ListDNSRecordsParams` was not correctly serialized into URL queries ([#1173](https://github.com/cloudflare/cloudflare-go/issues/1173))
* managednetworks: Update should be PUT ([#1172](https://github.com/cloudflare/cloudflare-go/issues/1172))

## 0.58.1 (January 5th, 2023)

ENHANCEMENTS:

* cloudflare: automatically redact sensitive values from HTTP interactions ([#1164](https://github.com/cloudflare/cloudflare-go/issues/1164))

## 0.58.0 (January 4th, 2023)

BREAKING CHANGES:

* dns: `DNSRecord` has been renamed to `GetDNSRecord` ([#1151](https://github.com/cloudflare/cloudflare-go/issues/1151))
* dns: `DNSRecords` has been renamed to `ListDNSRecords` ([#1151](https://github.com/cloudflare/cloudflare-go/issues/1151))
* dns: method signatures have been updated to align with the upcoming client conventions ([#1151](https://github.com/cloudflare/cloudflare-go/issues/1151))
* origin_ca: renamed to `CreateOriginCertificate` to `CreateOriginCACertificate` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
* origin_ca: renamed to `OriginCARootCertificate` to `GetOriginCARootCertificate` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
* origin_ca: renamed to `OriginCertificate` to `GetOriginCACertificate` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
* origin_ca: renamed to `OriginCertificates` to `ListOriginCACertificates` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
* origin_ca: renamed to `RevokeOriginCertificate` to `RevokeOriginCACertificate` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))

ENHANCEMENTS:

* dns: add support for tags and comments ([#1151](https://github.com/cloudflare/cloudflare-go/issues/1151))
* mtls_certificate: add support for managing mTLS certificates and assocations ([#1150](https://github.com/cloudflare/cloudflare-go/issues/1150))
* origin_ca: add support for using API keys, API tokens or API User service keys for interacting with Origin CA endpoints ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
* workers: Add support for workers logpush enablement on script upload ([#1160](https://github.com/cloudflare/cloudflare-go/issues/1160))

BUG FIXES:

* email_routing_destination: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
* email_routing_rules: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
* filter: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
* firewall_rules: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
* lockdown: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
* queue: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
* teams_list: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
* workers_kv: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))

DEPENDENCIES:

* deps: bumps github.com/hashicorp/go-retryablehttp from 0.7.1 to 0.7.2 ([#1162](https://github.com/cloudflare/cloudflare-go/issues/1162))

## 0.57.1 (December 23rd, 2022)

ENHANCEMENTS:

* tiered_cache: Add support for Tiered Caching interactions for setting Smart and Generic topologies ([#1149](https://github.com/cloudflare/cloudflare-go/issues/1149))

BUG FIXES:

* workers: correctly set `body` value for non-ES module uploads ([#1155](https://github.com/cloudflare/cloudflare-go/issues/1155))

## 0.57.0 (December 22nd, 2022)

BREAKING CHANGES:

* workers: API operations now target account level resources instead of older zone level resources (these are a 1:1 now) ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
* workers: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
* workers_bindings: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
* workers_cron_triggers: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
* workers_kv: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
* workers_routes: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
* workers_secrets: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
* workers_tails: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))

NOTES:

* workers: all worker methods have been split into product ownership(-ish) files ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
* workers: all worker methods now require an explicit `ResourceContainer` for endpoints instead of relying on the globally defined `api.AccountID` ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))

ENHANCEMENTS:

* managed_networks: add CRUD functionality for managednetworks ([#1148](https://github.com/cloudflare/cloudflare-go/issues/1148))

DEPENDENCIES:

* deps: bumps goreleaser/goreleaser-action from 3.2.0 to 4.1.0 ([#1146](https://github.com/cloudflare/cloudflare-go/issues/1146))

## 0.56.0 (December 5th, 2022)

BREAKING CHANGES:

* pages: Changed the type of EnvVars in PagesProjectDeploymentConfigEnvironment & PagesProjectDeployment in order to properly support secrets. ([#1136](https://github.com/cloudflare/cloudflare-go/issues/1136))

NOTES:

* pages: removed the v1 logs endpoint for Pages deployments. Please switch to v2: https://developers.cloudflare.com/api/operations/pages-deployment-get-deployment-logs ([#1135](https://github.com/cloudflare/cloudflare-go/issues/1135))

ENHANCEMENTS:

* cache_rules: add ignore option to query string struct ([#1140](https://github.com/cloudflare/cloudflare-go/issues/1140))
* pages: Updates bindings and other Functions related propreties. Service bindings, secrets, fail open/close and usage model are all now supported. ([#1136](https://github.com/cloudflare/cloudflare-go/issues/1136))
* workers: Support for Workers Analytics Engine bindings ([#1133](https://github.com/cloudflare/cloudflare-go/issues/1133))

DEPENDENCIES:

* deps: bumps github.com/urfave/cli/v2 from 2.23.5 to 2.23.6 ([#1139](https://github.com/cloudflare/cloudflare-go/issues/1139))

## 0.55.0 (November 23th, 2022)

BREAKING CHANGES:

* workers_kv: `CreateWorkersKVNamespace` has been updated to match the experimental client method signatures (https://github.com/cloudflare/cloudflare-go/blob/master/docs/experimental.md). ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `DeleteWorkersKVBulk` has been renamed to `DeleteWorkersKVEntries`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `DeleteWorkersKVNamespace` has been updated to match the experimental client method signatures (https://github.com/cloudflare/cloudflare-go/blob/master/docs/experimental.md). ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `DeleteWorkersKV` has been renamed to `DeleteWorkersKVEntry`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `ListWorkersKVNamespaces` has been updated to match the experimental client method signatures (https://github.com/cloudflare/cloudflare-go/blob/master/docs/experimental.md). ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `ListWorkersKVsWithOptions` has been removed. Use `ListWorkersKVKeys` instead and pass in the options. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `ListWorkersKVs` has been renamed to `ListWorkersKVKeys`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `ReadWorkersKV` has been renamed to `GetWorkersKV`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `UpdateWorkersKVNamespace` has been updated to match the experimental client method signatures (https://github.com/cloudflare/cloudflare-go/blob/master/docs/experimental.md). ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `WriteWorkersKVBulk` has been renamed to `WriteWorkersKVEntries`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
* workers_kv: `WriteWorkersKV` has been renamed to `WriteWorkersKVEntry`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))

ENHANCEMENTS:

* device_posture_rule: add input fields crowdstrike ([#1126](https://github.com/cloudflare/cloudflare-go/issues/1126))
* queue: add support queue API ([#1131](https://github.com/cloudflare/cloudflare-go/issues/1131))
* r2: Add support for listing R2 buckets ([#1063](https://github.com/cloudflare/cloudflare-go/issues/1063))
* workers_domain: add support for workers domain API ([#1130](https://github.com/cloudflare/cloudflare-go/issues/1130))
* workers_kv: `ListWorkersKVNamespaces` automatically paginates all results unless `PerPage` is defined. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))

DEPENDENCIES:

* deps: bumps github.com/urfave/cli/v2 from 2.23.4 to 2.23.5 ([#1127](https://github.com/cloudflare/cloudflare-go/issues/1127))

## 0.54.0 (November 9th, 2022)

ENHANCEMENTS:

* access: add support for service token rotation ([#1120](https://github.com/cloudflare/cloudflare-go/issues/1120))
* deps: fix import grouping, code formatting and enable goimports linter ([#1121](https://github.com/cloudflare/cloudflare-go/issues/1121))

DEPENDENCIES:

* deps: bumps dependabot/fetch-metadata from 1.3.4 to 1.3.5 ([#1123](https://github.com/cloudflare/cloudflare-go/issues/1123))
* deps: bumps github.com/urfave/cli/v2 from 2.20.3 to 2.23.0 ([#1122](https://github.com/cloudflare/cloudflare-go/issues/1122))
* deps: bumps github.com/urfave/cli/v2 from 2.23.0 to 2.23.2 ([#1124](https://github.com/cloudflare/cloudflare-go/issues/1124))
* deps: bumps github.com/urfave/cli/v2 from 2.23.2 to 2.23.4 ([#1125](https://github.com/cloudflare/cloudflare-go/issues/1125))

## 0.53.0 (October 26th, 2022)

BREAKING CHANGES:

* account_member: `CreateAccountMember` has been updated to accept a `CreateAccountMemberParams` struct instead of multiple parameters ([#1095](https://github.com/cloudflare/cloudflare-go/issues/1095))
* teams_list: updated methods to match the experimental client format ([#1114](https://github.com/cloudflare/cloudflare-go/issues/1114))

ENHANCEMENTS:

* account_member: add support for domain scoped roles ([#1095](https://github.com/cloudflare/cloudflare-go/issues/1095))
* cloudflare: expose `Messages` from the `Response` object ([#1106](https://github.com/cloudflare/cloudflare-go/issues/1106))
* dlp: Adds support for DLP resources ([#1111](https://github.com/cloudflare/cloudflare-go/issues/1111))
* teams_list: `List` operations now automatically paginate ([#1114](https://github.com/cloudflare/cloudflare-go/issues/1114))
* total_tls: adds support for TotalTLS ([#1105](https://github.com/cloudflare/cloudflare-go/issues/1105))
* waiting_room: add support for waiting room rules ([#1102](https://github.com/cloudflare/cloudflare-go/issues/1102))

DEPENDENCIES:

* deps: `ioutil` package is being deprecated in favor of `io` ([#1116](https://github.com/cloudflare/cloudflare-go/issues/1116))
* deps: bumps github.com/stretchr/testify from 1.8.0 to 1.8.1 ([#1119](https://github.com/cloudflare/cloudflare-go/issues/1119))
* deps: bumps github.com/urfave/cli/v2 from 2.19.2 to 2.20.2 ([#1108](https://github.com/cloudflare/cloudflare-go/issues/1108))
* deps: bumps github.com/urfave/cli/v2 from 2.20.2 to 2.20.3 ([#1118](https://github.com/cloudflare/cloudflare-go/issues/1118))
* deps: bumps goreleaser/goreleaser-action from 3.1.0 to 3.2.0 ([#1112](https://github.com/cloudflare/cloudflare-go/issues/1112))
* deps: remove `github.com/pkg/errors` in favor of `errors` ([#1117](https://github.com/cloudflare/cloudflare-go/issues/1117))

## 0.52.0 (October 12th, 2022)

ENHANCEMENTS:

* access: add UI read-only field to organizations ([#1104](https://github.com/cloudflare/cloudflare-go/issues/1104))
* devices_policy: Add support for additional device settings policies ([#1090](https://github.com/cloudflare/cloudflare-go/issues/1090))
* rulesets: add support for `sensitivity_level` to override all rule sensitivity ([#1093](https://github.com/cloudflare/cloudflare-go/issues/1093))

DEPENDENCIES:

* deps: bumps dependabot/fetch-metadata from 1.3.3 to 1.3.4 ([#1097](https://github.com/cloudflare/cloudflare-go/issues/1097))
* deps: bumps github.com/urfave/cli/v2 from 2.16.3 to 2.17.1 ([#1094](https://github.com/cloudflare/cloudflare-go/issues/1094))
* deps: bumps github.com/urfave/cli/v2 from 2.17.1 to 2.19.2 ([#1103](https://github.com/cloudflare/cloudflare-go/issues/1103))

## 0.51.0 (September 28th, 2022)

BREAKING CHANGES:

* load_balancing: update method signatures to match experimental conventions ([#1084](https://github.com/cloudflare/cloudflare-go/issues/1084))

ENHANCEMENTS:

* device_posture_rule: add input fields for linux OS ([#1087](https://github.com/cloudflare/cloudflare-go/issues/1087))
* load_balancing: support adaptive_routing and location_strategy ([#1091](https://github.com/cloudflare/cloudflare-go/issues/1091))

BUG FIXES:

* user-agent-blocking-rules: add missing managed_challenge validation and removed the deprecated whitelist one ([#1089](https://github.com/cloudflare/cloudflare-go/issues/1089))

## 0.50.0 (September 14, 2022)

ENHANCEMENTS:

* auditlogs: add support for hide_user_logs filter parameter ([#1075](https://github.com/cloudflare/cloudflare-go/issues/1075))

BUG FIXES:

* cloudflare: exiting closer to the source on context timeouts to improve error messaging and better defend from potential edge cases ([#1080](https://github.com/cloudflare/cloudflare-go/issues/1080))
* origin certificate: Fix API auth type used ([#1082](https://github.com/cloudflare/cloudflare-go/issues/1082))

DEPENDENCIES:

* deps: bumps github.com/urfave/cli/v2 from 2.11.2 to 2.14.0 ([#1077](https://github.com/cloudflare/cloudflare-go/issues/1077))
* deps: bumps github.com/urfave/cli/v2 from 2.14.0 to 2.14.1 ([#1081](https://github.com/cloudflare/cloudflare-go/issues/1081))
* deps: bumps github.com/urfave/cli/v2 from 2.14.1 to 2.15.0 ([#1085](https://github.com/cloudflare/cloudflare-go/issues/1085))
* deps: bumps github.com/urfave/cli/v2 from 2.15.0 to 2.16.3 ([#1086](https://github.com/cloudflare/cloudflare-go/issues/1086))

## 0.49.0 (August 31st, 2022)

ENHANCEMENTS:

* access_service_token: add support for refreshing an existing token in place ([#1074](https://github.com/cloudflare/cloudflare-go/issues/1074))
* api: addded context and headers to Raw method ([#1068](https://github.com/cloudflare/cloudflare-go/issues/1068))
* api_shield: add GET/PUT for API Shield Configuration ([#1059](https://github.com/cloudflare/cloudflare-go/issues/1059))
* pages_project: Add `kv_namespaces`, `durable_object_namespaces`, `r2_buckets`, and `d1_databases` bindings to deployment config ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
* pages_project: Add `preview_deployment_setting`, `preview_branch_includes`, and `preview_branch_excludes` to source config ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
* pages_project: Add `production_branch` field ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
* teams_account: add support for `os_distro_name` and `os_distro_revision` ([#1073](https://github.com/cloudflare/cloudflare-go/issues/1073))
* url_normalization_settings: Add APIs to get and update URL normalization settings ([#1071](https://github.com/cloudflare/cloudflare-go/issues/1071))
* workers: Support for multipart encoding for DownloadWorker on a module-format Worker script ([#1040](https://github.com/cloudflare/cloudflare-go/issues/1040))

BUG FIXES:

* cloudflare: fix nil dereference error in makeRequestWithAuthTypeAndHeaders ([#1072](https://github.com/cloudflare/cloudflare-go/issues/1072))
* email_routing_rules: Fix response for email routing catch all rule. ([#1070](https://github.com/cloudflare/cloudflare-go/issues/1070))
* email_routing_settings: change enable endpoint from `enabled` to `enable` ([#1060](https://github.com/cloudflare/cloudflare-go/issues/1060))
* stream: Update pctComplete to string from int ([#1066](https://github.com/cloudflare/cloudflare-go/issues/1066))

DEPENDENCIES:

* deps: bumps goreleaser/goreleaser-action from 3.0.0 to 3.1.0 ([#1067](https://github.com/cloudflare/cloudflare-go/issues/1067))

## 0.48.0 (August 22nd, 2022)

ENHANCEMENTS:

* errors: add some error type convenience functions for mocking and inspection ([#1047](https://github.com/cloudflare/cloudflare-go/issues/1047))
* pages_project: Add compatibility date and compatibility_flags to pages deployment configs ([#1051](https://github.com/cloudflare/cloudflare-go/issues/1051))
* teams_account: add support for `suppress_footer` ([#1053](https://github.com/cloudflare/cloudflare-go/issues/1053))

BUG FIXES:

* r2: fix create bucket endpoint ([#1035](https://github.com/cloudflare/cloudflare-go/issues/1035))
* tunnel_configuration: Remove unnecessary double-unmarshalling due to changes in the API ([#1046](https://github.com/cloudflare/cloudflare-go/issues/1046))

## 0.47.1 (August 18th, 2022)

BUG FIXES:

* zonelockdown: add `Priority` to `ZoneLockdownCreateParams` and `ZoneLockdownUpdateParams` ([#1052](https://github.com/cloudflare/cloudflare-go/issues/1052))

## 0.47.0 (August 17th, 2022)

BREAKING CHANGES:

* certificate_packs: deprecate "custom" configuration for ACM everywhere ([#1032](https://github.com/cloudflare/cloudflare-go/issues/1032))

ENHANCEMENTS:

* cloudflare: make it clear when the rate limit retries have been exhausted ([#1043](https://github.com/cloudflare/cloudflare-go/issues/1043))
* email_routing_destination: Adds support for the email routing destination API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
* email_routing_rules: Adds support for the email routing rules API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
* email_routing_settings: Adds support for the email routing settings API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
* filter: fix double endpoint calls & moving towards common method signature ([#1016](https://github.com/cloudflare/cloudflare-go/issues/1016))
* firewall_rule: fix double endpoint calls & moving towards common method signature ([#1016](https://github.com/cloudflare/cloudflare-go/issues/1016))
* lockdown: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1017](https://github.com/cloudflare/cloudflare-go/issues/1017))
* r2: Add in support for creating and deleting R2 buckets ([#1028](https://github.com/cloudflare/cloudflare-go/issues/1028))
* rulesets: add support for `http_config_settings` phase and supporting actions ([#1036](https://github.com/cloudflare/cloudflare-go/issues/1036))
* workers-account-settings: Add in support for Workers account settings API ([#1027](https://github.com/cloudflare/cloudflare-go/issues/1027))
* workers-subdomain: Add in support Workers Subdomain API ([#1031](https://github.com/cloudflare/cloudflare-go/issues/1031))
* workers-tail: Add in support for Workers tail API ([#1026](https://github.com/cloudflare/cloudflare-go/issues/1026))
* workers: Add support for attaching a worker to a domain ([#1014](https://github.com/cloudflare/cloudflare-go/issues/1014))
* workers: Add support to upload module workers ([#1010](https://github.com/cloudflare/cloudflare-go/issues/1010))

BUG FIXES:

* email_routing_destination: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
* email_routing_rules: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
* email_routing_settings: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
* tunnel_routes: Fix not removing route when it contains virtual network ([#1030](https://github.com/cloudflare/cloudflare-go/issues/1030))
* workers_test: Fix incorrect test from PR #1014 ([#1048](https://github.com/cloudflare/cloudflare-go/issues/1048))
* workers_test: Use application/json mime-type in headers ([#1049](https://github.com/cloudflare/cloudflare-go/issues/1049))

DEPENDENCIES:

* deps: bumps golang.org/x/tools/gopls from 0.9.3 to 0.9.4 ([#1044](https://github.com/cloudflare/cloudflare-go/issues/1044))
* deps: bumps github.com/golangci/golangci-lint from 1.47.3 to 1.48.0 ([#1020](https://github.com/cloudflare/cloudflare-go/issues/1020))
* deps: bumps github.com/urfave/cli/v2 from 2.11.1 to 2.11.2 ([#1042](https://github.com/cloudflare/cloudflare-go/issues/1042))
* deps: bumps golang.org/x/tools/gopls from 0.9.1 to 0.9.2 ([#1037](https://github.com/cloudflare/cloudflare-go/issues/1037))
* deps: bumps golang.org/x/tools/gopls from 0.9.2 to 0.9.3 ([#1039](https://github.com/cloudflare/cloudflare-go/issues/1039))

## 0.46.0 (3rd August, 2022)

NOTES:

* docs: add release notes ([#1001](https://github.com/cloudflare/cloudflare-go/issues/1001))

ENHANCEMENTS:

* filter: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1004](https://github.com/cloudflare/cloudflare-go/issues/1004))
* firewall_rule: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1004](https://github.com/cloudflare/cloudflare-go/issues/1004))
* rulesets: add support for `http_custom_errors` phase ([#998](https://github.com/cloudflare/cloudflare-go/issues/998))
* rulesets: add support for `serve_error` action ([#998](https://github.com/cloudflare/cloudflare-go/issues/998))

BUG FIXES:

* access_application: fix inability to set bool values to false ([#1006](https://github.com/cloudflare/cloudflare-go/issues/1006))
* rulesets: fix sni action parameter ([#1002](https://github.com/cloudflare/cloudflare-go/issues/1002))

DEPENDENCIES:

* provider: bumps github.com/golangci/golangci-lint from 1.47.1 to 1.47.2 ([#1005](https://github.com/cloudflare/cloudflare-go/issues/1005))
* provider: bumps github.com/golangci/golangci-lint from 1.47.2 to 1.47.3 ([#1008](https://github.com/cloudflare/cloudflare-go/issues/1008))
* provider: bumps github.com/urfave/cli/v2 from 2.11.0 to 2.11.1 ([#1003](https://github.com/cloudflare/cloudflare-go/issues/1003))

## 0.45.0 (July 20th, 2022)
