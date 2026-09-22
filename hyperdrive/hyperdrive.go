// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// HyperdriveService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHyperdriveService] method instead.
type HyperdriveService struct {
	Options []option.RequestOption
	Configs *ConfigService
}

// NewHyperdriveService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHyperdriveService(opts ...option.RequestOption) (r *HyperdriveService) {
	r = &HyperdriveService{}
	r.Options = opts
	r.Configs = NewConfigService(opts...)
	return
}

type HyperdriveParam struct {
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name param.Field[string] `json:"name" api:"required"`
	// Combines database connection fields with exactly one supported network location.
	Origin  param.Field[HyperdriveOriginUnionParam]  `json:"origin" api:"required"`
	Caching param.Field[HyperdriveCachingUnionParam] `json:"caching"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS param.Field[HyperdriveMTLSParam] `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit param.Field[int64] `json:"origin_connection_limit"`
}

func (r HyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Combines database connection fields with exactly one supported network location.
type HyperdriveOriginParam struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database" api:"required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveOriginScheme] `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user" api:"required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host param.Field[string] `json:"host"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port param.Field[int64] `json:"port"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID param.Field[string] `json:"service_id"`
}

func (r HyperdriveOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginParam) implementsHyperdriveOriginUnionParam() {}

// Combines database connection fields with exactly one supported network location.
//
// Satisfied by [hyperdrive.HyperdriveOriginPublicDatabaseParam],
// [hyperdrive.HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam],
// [hyperdrive.HyperdriveOriginDatabaseReachableThroughAWorkersVPCParam],
// [HyperdriveOriginParam].
type HyperdriveOriginUnionParam interface {
	implementsHyperdriveOriginUnionParam()
}

type HyperdriveOriginPublicDatabaseParam struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database" api:"required"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host param.Field[string] `json:"host" api:"required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password" api:"required"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port param.Field[int64] `json:"port" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveOriginPublicDatabaseScheme] `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user" api:"required"`
}

func (r HyperdriveOriginPublicDatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginPublicDatabaseParam) implementsHyperdriveOriginUnionParam() {}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginPublicDatabaseScheme string

const (
	HyperdriveOriginPublicDatabaseSchemePostgres   HyperdriveOriginPublicDatabaseScheme = "postgres"
	HyperdriveOriginPublicDatabaseSchemePostgresql HyperdriveOriginPublicDatabaseScheme = "postgresql"
	HyperdriveOriginPublicDatabaseSchemeMysql      HyperdriveOriginPublicDatabaseScheme = "mysql"
)

func (r HyperdriveOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginPublicDatabaseSchemePostgres, HyperdriveOriginPublicDatabaseSchemePostgresql, HyperdriveOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id" api:"required"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret" api:"required"`
	// Set the name of your origin database.
	Database param.Field[string] `json:"database" api:"required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host" api:"required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme] `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user" api:"required"`
}

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) implementsHyperdriveOriginUnionParam() {
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

type HyperdriveOriginDatabaseReachableThroughAWorkersVPCParam struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database" api:"required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveOriginDatabaseReachableThroughAWorkersVPCScheme] `json:"scheme" api:"required"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID param.Field[string] `json:"service_id" api:"required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user" api:"required"`
}

func (r HyperdriveOriginDatabaseReachableThroughAWorkersVPCParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginDatabaseReachableThroughAWorkersVPCParam) implementsHyperdriveOriginUnionParam() {
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginDatabaseReachableThroughAWorkersVPCScheme string

const (
	HyperdriveOriginDatabaseReachableThroughAWorkersVPCSchemePostgres   HyperdriveOriginDatabaseReachableThroughAWorkersVPCScheme = "postgres"
	HyperdriveOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql HyperdriveOriginDatabaseReachableThroughAWorkersVPCScheme = "postgresql"
	HyperdriveOriginDatabaseReachableThroughAWorkersVPCSchemeMysql      HyperdriveOriginDatabaseReachableThroughAWorkersVPCScheme = "mysql"
)

func (r HyperdriveOriginDatabaseReachableThroughAWorkersVPCScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginDatabaseReachableThroughAWorkersVPCSchemePostgres, HyperdriveOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql, HyperdriveOriginDatabaseReachableThroughAWorkersVPCSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginScheme string

const (
	HyperdriveOriginSchemePostgres   HyperdriveOriginScheme = "postgres"
	HyperdriveOriginSchemePostgresql HyperdriveOriginScheme = "postgresql"
	HyperdriveOriginSchemeMysql      HyperdriveOriginScheme = "mysql"
)

func (r HyperdriveOriginScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginSchemePostgres, HyperdriveOriginSchemePostgresql, HyperdriveOriginSchemeMysql:
		return true
	}
	return false
}

type HyperdriveCachingParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
	// Specify the maximum duration (in seconds) items should persist in the cache.
	// Defaults to 60 seconds if not specified.
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Defaults to
	// 15 seconds if not specified.
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveCachingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingParam) implementsHyperdriveCachingUnionParam() {}

// Satisfied by
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingCommonParam],
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam],
// [HyperdriveCachingParam].
type HyperdriveCachingUnionParam interface {
	implementsHyperdriveCachingUnionParam()
}

type HyperdriveCachingHyperdriveHyperdriveCachingCommonParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingCommonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingCommonParam) implementsHyperdriveCachingUnionParam() {
}

type HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
	// Specify the maximum duration (in seconds) items should persist in the cache.
	// Defaults to 60 seconds if not specified.
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Defaults to
	// 15 seconds if not specified.
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) implementsHyperdriveCachingUnionParam() {
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type HyperdriveMTLSParam struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID param.Field[string] `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID param.Field[string] `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode param.Field[string] `json:"sslmode"`
}

func (r HyperdriveMTLSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
