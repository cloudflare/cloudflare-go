// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// ConfigService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// Creates and returns a new Hyperdrive configuration. For a PlanetScale
// integration, the Cloudflare account must already be linked to PlanetScale in the
// Hyperdrive dashboard.
func (r *ConfigService) New(ctx context.Context, params ConfigNewParams, opts ...option.RequestOption) (res *ConfigNewResponse, err error) {
	var env ConfigNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces and returns the specified Hyperdrive configuration. The request must
// include the name and complete origin connection details. Omitted caching
// settings are reset to their defaults, while omitted mTLS settings and origin
// connection limits are preserved. Use the update operation to modify only
// selected fields.
func (r *ConfigService) Update(ctx context.Context, hyperdriveID string, params ConfigUpdateParams, opts ...option.RequestOption) (res *ConfigUpdateResponse, err error) {
	var env ConfigUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", params.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a list of Hyperdrives.
func (r *ConfigService) List(ctx context.Context, params ConfigListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ConfigListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of Hyperdrives.
func (r *ConfigService) ListAutoPaging(ctx context.Context, params ConfigListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ConfigListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes the specified Hyperdrive.
func (r *ConfigService) Delete(ctx context.Context, hyperdriveID string, body ConfigDeleteParams, opts ...option.RequestOption) (res *ConfigDeleteResponse, err error) {
	var env ConfigDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", body.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates and returns the specified fields of the Hyperdrive configuration. Custom
// caching settings are not kept if caching is disabled.
func (r *ConfigService) Edit(ctx context.Context, hyperdriveID string, params ConfigEditParams, opts ...option.RequestOption) (res *ConfigEditResponse, err error) {
	var env ConfigEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", params.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the specified Hyperdrive configuration.
func (r *ConfigService) Get(ctx context.Context, hyperdriveID string, query ConfigGetParams, opts ...option.RequestOption) (res *ConfigGetResponse, err error) {
	var env ConfigGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", query.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ConfigNewResponse struct {
	// Define configurations using a unique string identifier.
	ID      string                   `json:"id" api:"required"`
	Caching ConfigNewResponseCaching `json:"caching" api:"required"`
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name string `json:"name" api:"required"`
	// Combines database connection fields with exactly one supported network location.
	Origin ConfigNewResponseOrigin `json:"origin" api:"required"`
	// Defines the creation time of the Hyperdrive configuration.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Connects to a PlanetScale database using credentials managed by Cloudflare. The
	// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
	// dashboard.
	Integration ConfigNewResponseIntegration `json:"integration"`
	// Defines the last modified time of the Hyperdrive configuration.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS ConfigNewResponseMTLS `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit int64 `json:"origin_connection_limit"`
	// Defines the last time the Hyperdrive connection pool was explicitly restarted
	// via the restart endpoint. Omitted if the pool has never been explicitly
	// restarted.
	RestartedOn time.Time             `json:"restarted_on" api:"nullable" format:"date-time"`
	JSON        configNewResponseJSON `json:"-"`
}

// configNewResponseJSON contains the JSON metadata for the struct
// [ConfigNewResponse]
type configNewResponseJSON struct {
	ID                    apijson.Field
	Caching               apijson.Field
	Name                  apijson.Field
	Origin                apijson.Field
	CreatedOn             apijson.Field
	Integration           apijson.Field
	ModifiedOn            apijson.Field
	MTLS                  apijson.Field
	OriginConnectionLimit apijson.Field
	RestartedOn           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigNewResponseCaching struct {
	// Defines whether caching is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// Defines the maximum duration (in seconds) items persist in the cache.
	MaxAge int64 `json:"max_age"`
	// Defines the number of seconds the cache may serve a stale response.
	StaleWhileRevalidate int64                        `json:"stale_while_revalidate"`
	JSON                 configNewResponseCachingJSON `json:"-"`
}

// configNewResponseCachingJSON contains the JSON metadata for the struct
// [ConfigNewResponseCaching]
type configNewResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigNewResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseCachingJSON) RawJSON() string {
	return r.raw
}

// Combines database connection fields with exactly one supported network location.
type ConfigNewResponseOrigin struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigNewResponseOriginScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string `json:"user" api:"required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string                      `json:"service_id"`
	JSON      configNewResponseOriginJSON `json:"-"`
	union     ConfigNewResponseOriginUnion
}

// configNewResponseOriginJSON contains the JSON metadata for the struct
// [ConfigNewResponseOrigin]
type configNewResponseOriginJSON struct {
	Database       apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Host           apijson.Field
	Port           apijson.Field
	ServiceID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r configNewResponseOriginJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigNewResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigNewResponseOrigin{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigNewResponseOriginUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ConfigNewResponseOriginPublicDatabase],
// [ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel],
// [ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPC].
func (r ConfigNewResponseOrigin) AsUnion() ConfigNewResponseOriginUnion {
	return r.union
}

// Combines database connection fields with exactly one supported network location.
//
// Union satisfied by [ConfigNewResponseOriginPublicDatabase],
// [ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel] or
// [ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPC].
type ConfigNewResponseOriginUnion interface {
	implementsConfigNewResponseOrigin()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigNewResponseOriginUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigNewResponseOriginPublicDatabase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPC{}),
		},
	)
}

type ConfigNewResponseOriginPublicDatabase struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host" api:"required"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigNewResponseOriginPublicDatabaseScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                    `json:"user" api:"required"`
	JSON configNewResponseOriginPublicDatabaseJSON `json:"-"`
}

// configNewResponseOriginPublicDatabaseJSON contains the JSON metadata for the
// struct [ConfigNewResponseOriginPublicDatabase]
type configNewResponseOriginPublicDatabaseJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseOriginPublicDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseOriginPublicDatabaseJSON) RawJSON() string {
	return r.raw
}

func (r ConfigNewResponseOriginPublicDatabase) implementsConfigNewResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewResponseOriginPublicDatabaseScheme string

const (
	ConfigNewResponseOriginPublicDatabaseSchemePostgres   ConfigNewResponseOriginPublicDatabaseScheme = "postgres"
	ConfigNewResponseOriginPublicDatabaseSchemePostgresql ConfigNewResponseOriginPublicDatabaseScheme = "postgresql"
	ConfigNewResponseOriginPublicDatabaseSchemeMysql      ConfigNewResponseOriginPublicDatabaseScheme = "mysql"
)

func (r ConfigNewResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigNewResponseOriginPublicDatabaseSchemePostgres, ConfigNewResponseOriginPublicDatabaseSchemePostgresql, ConfigNewResponseOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id" api:"required"`
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                                                   `json:"user" api:"required"`
	JSON configNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
}

// configNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON
// contains the JSON metadata for the struct
// [ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type configNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsConfigNewResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

type ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPC struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCScheme `json:"scheme" api:"required"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string `json:"service_id" api:"required"`
	// Set the user of your origin database.
	User string                                                         `json:"user" api:"required"`
	JSON configNewResponseOriginDatabaseReachableThroughAWorkersVPCJSON `json:"-"`
}

// configNewResponseOriginDatabaseReachableThroughAWorkersVPCJSON contains the JSON
// metadata for the struct
// [ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPC]
type configNewResponseOriginDatabaseReachableThroughAWorkersVPCJSON struct {
	Database    apijson.Field
	Scheme      apijson.Field
	ServiceID   apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseOriginDatabaseReachableThroughAWorkersVPCJSON) RawJSON() string {
	return r.raw
}

func (r ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPC) implementsConfigNewResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCScheme string

const (
	ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres   ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgres"
	ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgresql"
	ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql      ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "mysql"
)

func (r ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCScheme) IsKnown() bool {
	switch r {
	case ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres, ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql, ConfigNewResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewResponseOriginScheme string

const (
	ConfigNewResponseOriginSchemePostgres   ConfigNewResponseOriginScheme = "postgres"
	ConfigNewResponseOriginSchemePostgresql ConfigNewResponseOriginScheme = "postgresql"
	ConfigNewResponseOriginSchemeMysql      ConfigNewResponseOriginScheme = "mysql"
)

func (r ConfigNewResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigNewResponseOriginSchemePostgres, ConfigNewResponseOriginSchemePostgresql, ConfigNewResponseOriginSchemeMysql:
		return true
	}
	return false
}

// Connects to a PlanetScale database using credentials managed by Cloudflare. The
// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
// dashboard.
type ConfigNewResponseIntegration struct {
	// The name of the PlanetScale database branch.
	DatabaseBranchName string `json:"database_branch_name" api:"required"`
	// The name of the PlanetScale database.
	DatabaseName string `json:"database_name" api:"required"`
	// The database integration used by this operation.
	Integration ConfigNewResponseIntegrationIntegration `json:"integration" api:"required"`
	// The name of the PlanetScale organization.
	OrganizationName string `json:"organization_name" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigNewResponseIntegrationScheme `json:"scheme" api:"required"`
	// The database name to use when connecting. Defaults to `postgres` for PostgreSQL
	// and `mysql` for MySQL.
	CustomDatabaseName string                           `json:"custom_database_name"`
	JSON               configNewResponseIntegrationJSON `json:"-"`
}

// configNewResponseIntegrationJSON contains the JSON metadata for the struct
// [ConfigNewResponseIntegration]
type configNewResponseIntegrationJSON struct {
	DatabaseBranchName apijson.Field
	DatabaseName       apijson.Field
	Integration        apijson.Field
	OrganizationName   apijson.Field
	Scheme             apijson.Field
	CustomDatabaseName apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConfigNewResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// The database integration used by this operation.
type ConfigNewResponseIntegrationIntegration string

const (
	ConfigNewResponseIntegrationIntegrationPlanetscale ConfigNewResponseIntegrationIntegration = "planetscale"
)

func (r ConfigNewResponseIntegrationIntegration) IsKnown() bool {
	switch r {
	case ConfigNewResponseIntegrationIntegrationPlanetscale:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewResponseIntegrationScheme string

const (
	ConfigNewResponseIntegrationSchemePostgres   ConfigNewResponseIntegrationScheme = "postgres"
	ConfigNewResponseIntegrationSchemePostgresql ConfigNewResponseIntegrationScheme = "postgresql"
	ConfigNewResponseIntegrationSchemeMysql      ConfigNewResponseIntegrationScheme = "mysql"
)

func (r ConfigNewResponseIntegrationScheme) IsKnown() bool {
	switch r {
	case ConfigNewResponseIntegrationSchemePostgres, ConfigNewResponseIntegrationSchemePostgresql, ConfigNewResponseIntegrationSchemeMysql:
		return true
	}
	return false
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type ConfigNewResponseMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID string `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID string `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode string                    `json:"sslmode"`
	JSON    configNewResponseMTLSJSON `json:"-"`
}

// configNewResponseMTLSJSON contains the JSON metadata for the struct
// [ConfigNewResponseMTLS]
type configNewResponseMTLSJSON struct {
	CACertificateID   apijson.Field
	MTLSCertificateID apijson.Field
	Sslmode           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigNewResponseMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseMTLSJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponse struct {
	// Define configurations using a unique string identifier.
	ID      string                      `json:"id" api:"required"`
	Caching ConfigUpdateResponseCaching `json:"caching" api:"required"`
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name string `json:"name" api:"required"`
	// Combines database connection fields with exactly one supported network location.
	Origin ConfigUpdateResponseOrigin `json:"origin" api:"required"`
	// Defines the creation time of the Hyperdrive configuration.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Connects to a PlanetScale database using credentials managed by Cloudflare. The
	// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
	// dashboard.
	Integration ConfigUpdateResponseIntegration `json:"integration"`
	// Defines the last modified time of the Hyperdrive configuration.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS ConfigUpdateResponseMTLS `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit int64 `json:"origin_connection_limit"`
	// Defines the last time the Hyperdrive connection pool was explicitly restarted
	// via the restart endpoint. Omitted if the pool has never been explicitly
	// restarted.
	RestartedOn time.Time                `json:"restarted_on" api:"nullable" format:"date-time"`
	JSON        configUpdateResponseJSON `json:"-"`
}

// configUpdateResponseJSON contains the JSON metadata for the struct
// [ConfigUpdateResponse]
type configUpdateResponseJSON struct {
	ID                    apijson.Field
	Caching               apijson.Field
	Name                  apijson.Field
	Origin                apijson.Field
	CreatedOn             apijson.Field
	Integration           apijson.Field
	ModifiedOn            apijson.Field
	MTLS                  apijson.Field
	OriginConnectionLimit apijson.Field
	RestartedOn           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseCaching struct {
	// Defines whether caching is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// Defines the maximum duration (in seconds) items persist in the cache.
	MaxAge int64 `json:"max_age"`
	// Defines the number of seconds the cache may serve a stale response.
	StaleWhileRevalidate int64                           `json:"stale_while_revalidate"`
	JSON                 configUpdateResponseCachingJSON `json:"-"`
}

// configUpdateResponseCachingJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseCaching]
type configUpdateResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigUpdateResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseCachingJSON) RawJSON() string {
	return r.raw
}

// Combines database connection fields with exactly one supported network location.
type ConfigUpdateResponseOrigin struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigUpdateResponseOriginScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string `json:"user" api:"required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string                         `json:"service_id"`
	JSON      configUpdateResponseOriginJSON `json:"-"`
	union     ConfigUpdateResponseOriginUnion
}

// configUpdateResponseOriginJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseOrigin]
type configUpdateResponseOriginJSON struct {
	Database       apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Host           apijson.Field
	Port           apijson.Field
	ServiceID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r configUpdateResponseOriginJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigUpdateResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigUpdateResponseOrigin{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigUpdateResponseOriginUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigUpdateResponseOriginPublicDatabase],
// [ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel],
// [ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPC].
func (r ConfigUpdateResponseOrigin) AsUnion() ConfigUpdateResponseOriginUnion {
	return r.union
}

// Combines database connection fields with exactly one supported network location.
//
// Union satisfied by [ConfigUpdateResponseOriginPublicDatabase],
// [ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel] or
// [ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPC].
type ConfigUpdateResponseOriginUnion interface {
	implementsConfigUpdateResponseOrigin()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigUpdateResponseOriginUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigUpdateResponseOriginPublicDatabase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPC{}),
		},
	)
}

type ConfigUpdateResponseOriginPublicDatabase struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host" api:"required"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigUpdateResponseOriginPublicDatabaseScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                       `json:"user" api:"required"`
	JSON configUpdateResponseOriginPublicDatabaseJSON `json:"-"`
}

// configUpdateResponseOriginPublicDatabaseJSON contains the JSON metadata for the
// struct [ConfigUpdateResponseOriginPublicDatabase]
type configUpdateResponseOriginPublicDatabaseJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseOriginPublicDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseOriginPublicDatabaseJSON) RawJSON() string {
	return r.raw
}

func (r ConfigUpdateResponseOriginPublicDatabase) implementsConfigUpdateResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateResponseOriginPublicDatabaseScheme string

const (
	ConfigUpdateResponseOriginPublicDatabaseSchemePostgres   ConfigUpdateResponseOriginPublicDatabaseScheme = "postgres"
	ConfigUpdateResponseOriginPublicDatabaseSchemePostgresql ConfigUpdateResponseOriginPublicDatabaseScheme = "postgresql"
	ConfigUpdateResponseOriginPublicDatabaseSchemeMysql      ConfigUpdateResponseOriginPublicDatabaseScheme = "mysql"
)

func (r ConfigUpdateResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseOriginPublicDatabaseSchemePostgres, ConfigUpdateResponseOriginPublicDatabaseSchemePostgresql, ConfigUpdateResponseOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id" api:"required"`
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                                                      `json:"user" api:"required"`
	JSON configUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
}

// configUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON
// contains the JSON metadata for the struct
// [ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type configUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsConfigUpdateResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

type ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPC struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCScheme `json:"scheme" api:"required"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string `json:"service_id" api:"required"`
	// Set the user of your origin database.
	User string                                                            `json:"user" api:"required"`
	JSON configUpdateResponseOriginDatabaseReachableThroughAWorkersVPCJSON `json:"-"`
}

// configUpdateResponseOriginDatabaseReachableThroughAWorkersVPCJSON contains the
// JSON metadata for the struct
// [ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPC]
type configUpdateResponseOriginDatabaseReachableThroughAWorkersVPCJSON struct {
	Database    apijson.Field
	Scheme      apijson.Field
	ServiceID   apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseOriginDatabaseReachableThroughAWorkersVPCJSON) RawJSON() string {
	return r.raw
}

func (r ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPC) implementsConfigUpdateResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCScheme string

const (
	ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres   ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgres"
	ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgresql"
	ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql      ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "mysql"
)

func (r ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres, ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql, ConfigUpdateResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateResponseOriginScheme string

const (
	ConfigUpdateResponseOriginSchemePostgres   ConfigUpdateResponseOriginScheme = "postgres"
	ConfigUpdateResponseOriginSchemePostgresql ConfigUpdateResponseOriginScheme = "postgresql"
	ConfigUpdateResponseOriginSchemeMysql      ConfigUpdateResponseOriginScheme = "mysql"
)

func (r ConfigUpdateResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseOriginSchemePostgres, ConfigUpdateResponseOriginSchemePostgresql, ConfigUpdateResponseOriginSchemeMysql:
		return true
	}
	return false
}

// Connects to a PlanetScale database using credentials managed by Cloudflare. The
// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
// dashboard.
type ConfigUpdateResponseIntegration struct {
	// The name of the PlanetScale database branch.
	DatabaseBranchName string `json:"database_branch_name" api:"required"`
	// The name of the PlanetScale database.
	DatabaseName string `json:"database_name" api:"required"`
	// The database integration used by this operation.
	Integration ConfigUpdateResponseIntegrationIntegration `json:"integration" api:"required"`
	// The name of the PlanetScale organization.
	OrganizationName string `json:"organization_name" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigUpdateResponseIntegrationScheme `json:"scheme" api:"required"`
	// The database name to use when connecting. Defaults to `postgres` for PostgreSQL
	// and `mysql` for MySQL.
	CustomDatabaseName string                              `json:"custom_database_name"`
	JSON               configUpdateResponseIntegrationJSON `json:"-"`
}

// configUpdateResponseIntegrationJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseIntegration]
type configUpdateResponseIntegrationJSON struct {
	DatabaseBranchName apijson.Field
	DatabaseName       apijson.Field
	Integration        apijson.Field
	OrganizationName   apijson.Field
	Scheme             apijson.Field
	CustomDatabaseName apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConfigUpdateResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// The database integration used by this operation.
type ConfigUpdateResponseIntegrationIntegration string

const (
	ConfigUpdateResponseIntegrationIntegrationPlanetscale ConfigUpdateResponseIntegrationIntegration = "planetscale"
)

func (r ConfigUpdateResponseIntegrationIntegration) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseIntegrationIntegrationPlanetscale:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateResponseIntegrationScheme string

const (
	ConfigUpdateResponseIntegrationSchemePostgres   ConfigUpdateResponseIntegrationScheme = "postgres"
	ConfigUpdateResponseIntegrationSchemePostgresql ConfigUpdateResponseIntegrationScheme = "postgresql"
	ConfigUpdateResponseIntegrationSchemeMysql      ConfigUpdateResponseIntegrationScheme = "mysql"
)

func (r ConfigUpdateResponseIntegrationScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseIntegrationSchemePostgres, ConfigUpdateResponseIntegrationSchemePostgresql, ConfigUpdateResponseIntegrationSchemeMysql:
		return true
	}
	return false
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type ConfigUpdateResponseMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID string `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID string `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode string                       `json:"sslmode"`
	JSON    configUpdateResponseMTLSJSON `json:"-"`
}

// configUpdateResponseMTLSJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseMTLS]
type configUpdateResponseMTLSJSON struct {
	CACertificateID   apijson.Field
	MTLSCertificateID apijson.Field
	Sslmode           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigUpdateResponseMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseMTLSJSON) RawJSON() string {
	return r.raw
}

type ConfigListResponse struct {
	// Define configurations using a unique string identifier.
	ID      string                    `json:"id" api:"required"`
	Caching ConfigListResponseCaching `json:"caching" api:"required"`
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name string `json:"name" api:"required"`
	// Combines database connection fields with exactly one supported network location.
	Origin ConfigListResponseOrigin `json:"origin" api:"required"`
	// Defines the creation time of the Hyperdrive configuration.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Connects to a PlanetScale database using credentials managed by Cloudflare. The
	// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
	// dashboard.
	Integration ConfigListResponseIntegration `json:"integration"`
	// Defines the last modified time of the Hyperdrive configuration.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS ConfigListResponseMTLS `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit int64 `json:"origin_connection_limit"`
	// Defines the last time the Hyperdrive connection pool was explicitly restarted
	// via the restart endpoint. Omitted if the pool has never been explicitly
	// restarted.
	RestartedOn time.Time              `json:"restarted_on" api:"nullable" format:"date-time"`
	JSON        configListResponseJSON `json:"-"`
}

// configListResponseJSON contains the JSON metadata for the struct
// [ConfigListResponse]
type configListResponseJSON struct {
	ID                    apijson.Field
	Caching               apijson.Field
	Name                  apijson.Field
	Origin                apijson.Field
	CreatedOn             apijson.Field
	Integration           apijson.Field
	ModifiedOn            apijson.Field
	MTLS                  apijson.Field
	OriginConnectionLimit apijson.Field
	RestartedOn           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigListResponseCaching struct {
	// Defines whether caching is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// Defines the maximum duration (in seconds) items persist in the cache.
	MaxAge int64 `json:"max_age"`
	// Defines the number of seconds the cache may serve a stale response.
	StaleWhileRevalidate int64                         `json:"stale_while_revalidate"`
	JSON                 configListResponseCachingJSON `json:"-"`
}

// configListResponseCachingJSON contains the JSON metadata for the struct
// [ConfigListResponseCaching]
type configListResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigListResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseCachingJSON) RawJSON() string {
	return r.raw
}

// Combines database connection fields with exactly one supported network location.
type ConfigListResponseOrigin struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigListResponseOriginScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string `json:"user" api:"required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string                       `json:"service_id"`
	JSON      configListResponseOriginJSON `json:"-"`
	union     ConfigListResponseOriginUnion
}

// configListResponseOriginJSON contains the JSON metadata for the struct
// [ConfigListResponseOrigin]
type configListResponseOriginJSON struct {
	Database       apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Host           apijson.Field
	Port           apijson.Field
	ServiceID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r configListResponseOriginJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigListResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigListResponseOrigin{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigListResponseOriginUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigListResponseOriginPublicDatabase],
// [ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel],
// [ConfigListResponseOriginDatabaseReachableThroughAWorkersVPC].
func (r ConfigListResponseOrigin) AsUnion() ConfigListResponseOriginUnion {
	return r.union
}

// Combines database connection fields with exactly one supported network location.
//
// Union satisfied by [ConfigListResponseOriginPublicDatabase],
// [ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel] or
// [ConfigListResponseOriginDatabaseReachableThroughAWorkersVPC].
type ConfigListResponseOriginUnion interface {
	implementsConfigListResponseOrigin()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigListResponseOriginUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigListResponseOriginPublicDatabase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigListResponseOriginDatabaseReachableThroughAWorkersVPC{}),
		},
	)
}

type ConfigListResponseOriginPublicDatabase struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host" api:"required"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigListResponseOriginPublicDatabaseScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                     `json:"user" api:"required"`
	JSON configListResponseOriginPublicDatabaseJSON `json:"-"`
}

// configListResponseOriginPublicDatabaseJSON contains the JSON metadata for the
// struct [ConfigListResponseOriginPublicDatabase]
type configListResponseOriginPublicDatabaseJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigListResponseOriginPublicDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseOriginPublicDatabaseJSON) RawJSON() string {
	return r.raw
}

func (r ConfigListResponseOriginPublicDatabase) implementsConfigListResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigListResponseOriginPublicDatabaseScheme string

const (
	ConfigListResponseOriginPublicDatabaseSchemePostgres   ConfigListResponseOriginPublicDatabaseScheme = "postgres"
	ConfigListResponseOriginPublicDatabaseSchemePostgresql ConfigListResponseOriginPublicDatabaseScheme = "postgresql"
	ConfigListResponseOriginPublicDatabaseSchemeMysql      ConfigListResponseOriginPublicDatabaseScheme = "mysql"
)

func (r ConfigListResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigListResponseOriginPublicDatabaseSchemePostgres, ConfigListResponseOriginPublicDatabaseSchemePostgresql, ConfigListResponseOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id" api:"required"`
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                                                    `json:"user" api:"required"`
	JSON configListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
}

// configListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON
// contains the JSON metadata for the struct
// [ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type configListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsConfigListResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

type ConfigListResponseOriginDatabaseReachableThroughAWorkersVPC struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCScheme `json:"scheme" api:"required"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string `json:"service_id" api:"required"`
	// Set the user of your origin database.
	User string                                                          `json:"user" api:"required"`
	JSON configListResponseOriginDatabaseReachableThroughAWorkersVPCJSON `json:"-"`
}

// configListResponseOriginDatabaseReachableThroughAWorkersVPCJSON contains the
// JSON metadata for the struct
// [ConfigListResponseOriginDatabaseReachableThroughAWorkersVPC]
type configListResponseOriginDatabaseReachableThroughAWorkersVPCJSON struct {
	Database    apijson.Field
	Scheme      apijson.Field
	ServiceID   apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigListResponseOriginDatabaseReachableThroughAWorkersVPC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseOriginDatabaseReachableThroughAWorkersVPCJSON) RawJSON() string {
	return r.raw
}

func (r ConfigListResponseOriginDatabaseReachableThroughAWorkersVPC) implementsConfigListResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCScheme string

const (
	ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres   ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgres"
	ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgresql"
	ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql      ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "mysql"
)

func (r ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCScheme) IsKnown() bool {
	switch r {
	case ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres, ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql, ConfigListResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigListResponseOriginScheme string

const (
	ConfigListResponseOriginSchemePostgres   ConfigListResponseOriginScheme = "postgres"
	ConfigListResponseOriginSchemePostgresql ConfigListResponseOriginScheme = "postgresql"
	ConfigListResponseOriginSchemeMysql      ConfigListResponseOriginScheme = "mysql"
)

func (r ConfigListResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigListResponseOriginSchemePostgres, ConfigListResponseOriginSchemePostgresql, ConfigListResponseOriginSchemeMysql:
		return true
	}
	return false
}

// Connects to a PlanetScale database using credentials managed by Cloudflare. The
// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
// dashboard.
type ConfigListResponseIntegration struct {
	// The name of the PlanetScale database branch.
	DatabaseBranchName string `json:"database_branch_name" api:"required"`
	// The name of the PlanetScale database.
	DatabaseName string `json:"database_name" api:"required"`
	// The database integration used by this operation.
	Integration ConfigListResponseIntegrationIntegration `json:"integration" api:"required"`
	// The name of the PlanetScale organization.
	OrganizationName string `json:"organization_name" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigListResponseIntegrationScheme `json:"scheme" api:"required"`
	// The database name to use when connecting. Defaults to `postgres` for PostgreSQL
	// and `mysql` for MySQL.
	CustomDatabaseName string                            `json:"custom_database_name"`
	JSON               configListResponseIntegrationJSON `json:"-"`
}

// configListResponseIntegrationJSON contains the JSON metadata for the struct
// [ConfigListResponseIntegration]
type configListResponseIntegrationJSON struct {
	DatabaseBranchName apijson.Field
	DatabaseName       apijson.Field
	Integration        apijson.Field
	OrganizationName   apijson.Field
	Scheme             apijson.Field
	CustomDatabaseName apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConfigListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// The database integration used by this operation.
type ConfigListResponseIntegrationIntegration string

const (
	ConfigListResponseIntegrationIntegrationPlanetscale ConfigListResponseIntegrationIntegration = "planetscale"
)

func (r ConfigListResponseIntegrationIntegration) IsKnown() bool {
	switch r {
	case ConfigListResponseIntegrationIntegrationPlanetscale:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigListResponseIntegrationScheme string

const (
	ConfigListResponseIntegrationSchemePostgres   ConfigListResponseIntegrationScheme = "postgres"
	ConfigListResponseIntegrationSchemePostgresql ConfigListResponseIntegrationScheme = "postgresql"
	ConfigListResponseIntegrationSchemeMysql      ConfigListResponseIntegrationScheme = "mysql"
)

func (r ConfigListResponseIntegrationScheme) IsKnown() bool {
	switch r {
	case ConfigListResponseIntegrationSchemePostgres, ConfigListResponseIntegrationSchemePostgresql, ConfigListResponseIntegrationSchemeMysql:
		return true
	}
	return false
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type ConfigListResponseMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID string `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID string `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode string                     `json:"sslmode"`
	JSON    configListResponseMTLSJSON `json:"-"`
}

// configListResponseMTLSJSON contains the JSON metadata for the struct
// [ConfigListResponseMTLS]
type configListResponseMTLSJSON struct {
	CACertificateID   apijson.Field
	MTLSCertificateID apijson.Field
	Sslmode           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigListResponseMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseMTLSJSON) RawJSON() string {
	return r.raw
}

type ConfigDeleteResponse = interface{}

type ConfigEditResponse struct {
	// Define configurations using a unique string identifier.
	ID      string                    `json:"id" api:"required"`
	Caching ConfigEditResponseCaching `json:"caching" api:"required"`
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name string `json:"name" api:"required"`
	// Combines database connection fields with exactly one supported network location.
	Origin ConfigEditResponseOrigin `json:"origin" api:"required"`
	// Defines the creation time of the Hyperdrive configuration.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Connects to a PlanetScale database using credentials managed by Cloudflare. The
	// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
	// dashboard.
	Integration ConfigEditResponseIntegration `json:"integration"`
	// Defines the last modified time of the Hyperdrive configuration.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS ConfigEditResponseMTLS `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit int64 `json:"origin_connection_limit"`
	// Defines the last time the Hyperdrive connection pool was explicitly restarted
	// via the restart endpoint. Omitted if the pool has never been explicitly
	// restarted.
	RestartedOn time.Time              `json:"restarted_on" api:"nullable" format:"date-time"`
	JSON        configEditResponseJSON `json:"-"`
}

// configEditResponseJSON contains the JSON metadata for the struct
// [ConfigEditResponse]
type configEditResponseJSON struct {
	ID                    apijson.Field
	Caching               apijson.Field
	Name                  apijson.Field
	Origin                apijson.Field
	CreatedOn             apijson.Field
	Integration           apijson.Field
	ModifiedOn            apijson.Field
	MTLS                  apijson.Field
	OriginConnectionLimit apijson.Field
	RestartedOn           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ConfigEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigEditResponseCaching struct {
	// Defines whether caching is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// Defines the maximum duration (in seconds) items persist in the cache.
	MaxAge int64 `json:"max_age"`
	// Defines the number of seconds the cache may serve a stale response.
	StaleWhileRevalidate int64                         `json:"stale_while_revalidate"`
	JSON                 configEditResponseCachingJSON `json:"-"`
}

// configEditResponseCachingJSON contains the JSON metadata for the struct
// [ConfigEditResponseCaching]
type configEditResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigEditResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseCachingJSON) RawJSON() string {
	return r.raw
}

// Combines database connection fields with exactly one supported network location.
type ConfigEditResponseOrigin struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigEditResponseOriginScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string `json:"user" api:"required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string                       `json:"service_id"`
	JSON      configEditResponseOriginJSON `json:"-"`
	union     ConfigEditResponseOriginUnion
}

// configEditResponseOriginJSON contains the JSON metadata for the struct
// [ConfigEditResponseOrigin]
type configEditResponseOriginJSON struct {
	Database       apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Host           apijson.Field
	Port           apijson.Field
	ServiceID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r configEditResponseOriginJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigEditResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigEditResponseOrigin{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigEditResponseOriginUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigEditResponseOriginPublicDatabase],
// [ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel],
// [ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPC].
func (r ConfigEditResponseOrigin) AsUnion() ConfigEditResponseOriginUnion {
	return r.union
}

// Combines database connection fields with exactly one supported network location.
//
// Union satisfied by [ConfigEditResponseOriginPublicDatabase],
// [ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel] or
// [ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPC].
type ConfigEditResponseOriginUnion interface {
	implementsConfigEditResponseOrigin()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigEditResponseOriginUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigEditResponseOriginPublicDatabase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPC{}),
		},
	)
}

type ConfigEditResponseOriginPublicDatabase struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host" api:"required"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigEditResponseOriginPublicDatabaseScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                     `json:"user" api:"required"`
	JSON configEditResponseOriginPublicDatabaseJSON `json:"-"`
}

// configEditResponseOriginPublicDatabaseJSON contains the JSON metadata for the
// struct [ConfigEditResponseOriginPublicDatabase]
type configEditResponseOriginPublicDatabaseJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseOriginPublicDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseOriginPublicDatabaseJSON) RawJSON() string {
	return r.raw
}

func (r ConfigEditResponseOriginPublicDatabase) implementsConfigEditResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditResponseOriginPublicDatabaseScheme string

const (
	ConfigEditResponseOriginPublicDatabaseSchemePostgres   ConfigEditResponseOriginPublicDatabaseScheme = "postgres"
	ConfigEditResponseOriginPublicDatabaseSchemePostgresql ConfigEditResponseOriginPublicDatabaseScheme = "postgresql"
	ConfigEditResponseOriginPublicDatabaseSchemeMysql      ConfigEditResponseOriginPublicDatabaseScheme = "mysql"
)

func (r ConfigEditResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigEditResponseOriginPublicDatabaseSchemePostgres, ConfigEditResponseOriginPublicDatabaseSchemePostgresql, ConfigEditResponseOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id" api:"required"`
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                                                    `json:"user" api:"required"`
	JSON configEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
}

// configEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON
// contains the JSON metadata for the struct
// [ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type configEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsConfigEditResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

type ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPC struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCScheme `json:"scheme" api:"required"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string `json:"service_id" api:"required"`
	// Set the user of your origin database.
	User string                                                          `json:"user" api:"required"`
	JSON configEditResponseOriginDatabaseReachableThroughAWorkersVPCJSON `json:"-"`
}

// configEditResponseOriginDatabaseReachableThroughAWorkersVPCJSON contains the
// JSON metadata for the struct
// [ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPC]
type configEditResponseOriginDatabaseReachableThroughAWorkersVPCJSON struct {
	Database    apijson.Field
	Scheme      apijson.Field
	ServiceID   apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseOriginDatabaseReachableThroughAWorkersVPCJSON) RawJSON() string {
	return r.raw
}

func (r ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPC) implementsConfigEditResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCScheme string

const (
	ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres   ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgres"
	ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgresql"
	ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql      ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "mysql"
)

func (r ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCScheme) IsKnown() bool {
	switch r {
	case ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres, ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql, ConfigEditResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditResponseOriginScheme string

const (
	ConfigEditResponseOriginSchemePostgres   ConfigEditResponseOriginScheme = "postgres"
	ConfigEditResponseOriginSchemePostgresql ConfigEditResponseOriginScheme = "postgresql"
	ConfigEditResponseOriginSchemeMysql      ConfigEditResponseOriginScheme = "mysql"
)

func (r ConfigEditResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigEditResponseOriginSchemePostgres, ConfigEditResponseOriginSchemePostgresql, ConfigEditResponseOriginSchemeMysql:
		return true
	}
	return false
}

// Connects to a PlanetScale database using credentials managed by Cloudflare. The
// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
// dashboard.
type ConfigEditResponseIntegration struct {
	// The name of the PlanetScale database branch.
	DatabaseBranchName string `json:"database_branch_name" api:"required"`
	// The name of the PlanetScale database.
	DatabaseName string `json:"database_name" api:"required"`
	// The database integration used by this operation.
	Integration ConfigEditResponseIntegrationIntegration `json:"integration" api:"required"`
	// The name of the PlanetScale organization.
	OrganizationName string `json:"organization_name" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigEditResponseIntegrationScheme `json:"scheme" api:"required"`
	// The database name to use when connecting. Defaults to `postgres` for PostgreSQL
	// and `mysql` for MySQL.
	CustomDatabaseName string                            `json:"custom_database_name"`
	JSON               configEditResponseIntegrationJSON `json:"-"`
}

// configEditResponseIntegrationJSON contains the JSON metadata for the struct
// [ConfigEditResponseIntegration]
type configEditResponseIntegrationJSON struct {
	DatabaseBranchName apijson.Field
	DatabaseName       apijson.Field
	Integration        apijson.Field
	OrganizationName   apijson.Field
	Scheme             apijson.Field
	CustomDatabaseName apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConfigEditResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// The database integration used by this operation.
type ConfigEditResponseIntegrationIntegration string

const (
	ConfigEditResponseIntegrationIntegrationPlanetscale ConfigEditResponseIntegrationIntegration = "planetscale"
)

func (r ConfigEditResponseIntegrationIntegration) IsKnown() bool {
	switch r {
	case ConfigEditResponseIntegrationIntegrationPlanetscale:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditResponseIntegrationScheme string

const (
	ConfigEditResponseIntegrationSchemePostgres   ConfigEditResponseIntegrationScheme = "postgres"
	ConfigEditResponseIntegrationSchemePostgresql ConfigEditResponseIntegrationScheme = "postgresql"
	ConfigEditResponseIntegrationSchemeMysql      ConfigEditResponseIntegrationScheme = "mysql"
)

func (r ConfigEditResponseIntegrationScheme) IsKnown() bool {
	switch r {
	case ConfigEditResponseIntegrationSchemePostgres, ConfigEditResponseIntegrationSchemePostgresql, ConfigEditResponseIntegrationSchemeMysql:
		return true
	}
	return false
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type ConfigEditResponseMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID string `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID string `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode string                     `json:"sslmode"`
	JSON    configEditResponseMTLSJSON `json:"-"`
}

// configEditResponseMTLSJSON contains the JSON metadata for the struct
// [ConfigEditResponseMTLS]
type configEditResponseMTLSJSON struct {
	CACertificateID   apijson.Field
	MTLSCertificateID apijson.Field
	Sslmode           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigEditResponseMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseMTLSJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponse struct {
	// Define configurations using a unique string identifier.
	ID      string                   `json:"id" api:"required"`
	Caching ConfigGetResponseCaching `json:"caching" api:"required"`
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name string `json:"name" api:"required"`
	// Combines database connection fields with exactly one supported network location.
	Origin ConfigGetResponseOrigin `json:"origin" api:"required"`
	// Defines the creation time of the Hyperdrive configuration.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Connects to a PlanetScale database using credentials managed by Cloudflare. The
	// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
	// dashboard.
	Integration ConfigGetResponseIntegration `json:"integration"`
	// Defines the last modified time of the Hyperdrive configuration.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS ConfigGetResponseMTLS `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit int64 `json:"origin_connection_limit"`
	// Defines the last time the Hyperdrive connection pool was explicitly restarted
	// via the restart endpoint. Omitted if the pool has never been explicitly
	// restarted.
	RestartedOn time.Time             `json:"restarted_on" api:"nullable" format:"date-time"`
	JSON        configGetResponseJSON `json:"-"`
}

// configGetResponseJSON contains the JSON metadata for the struct
// [ConfigGetResponse]
type configGetResponseJSON struct {
	ID                    apijson.Field
	Caching               apijson.Field
	Name                  apijson.Field
	Origin                apijson.Field
	CreatedOn             apijson.Field
	Integration           apijson.Field
	ModifiedOn            apijson.Field
	MTLS                  apijson.Field
	OriginConnectionLimit apijson.Field
	RestartedOn           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseCaching struct {
	// Defines whether caching is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// Defines the maximum duration (in seconds) items persist in the cache.
	MaxAge int64 `json:"max_age"`
	// Defines the number of seconds the cache may serve a stale response.
	StaleWhileRevalidate int64                        `json:"stale_while_revalidate"`
	JSON                 configGetResponseCachingJSON `json:"-"`
}

// configGetResponseCachingJSON contains the JSON metadata for the struct
// [ConfigGetResponseCaching]
type configGetResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigGetResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseCachingJSON) RawJSON() string {
	return r.raw
}

// Combines database connection fields with exactly one supported network location.
type ConfigGetResponseOrigin struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigGetResponseOriginScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string `json:"user" api:"required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string                      `json:"service_id"`
	JSON      configGetResponseOriginJSON `json:"-"`
	union     ConfigGetResponseOriginUnion
}

// configGetResponseOriginJSON contains the JSON metadata for the struct
// [ConfigGetResponseOrigin]
type configGetResponseOriginJSON struct {
	Database       apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Host           apijson.Field
	Port           apijson.Field
	ServiceID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r configGetResponseOriginJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigGetResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigGetResponseOrigin{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigGetResponseOriginUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ConfigGetResponseOriginPublicDatabase],
// [ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel],
// [ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPC].
func (r ConfigGetResponseOrigin) AsUnion() ConfigGetResponseOriginUnion {
	return r.union
}

// Combines database connection fields with exactly one supported network location.
//
// Union satisfied by [ConfigGetResponseOriginPublicDatabase],
// [ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel] or
// [ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPC].
type ConfigGetResponseOriginUnion interface {
	implementsConfigGetResponseOrigin()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigGetResponseOriginUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigGetResponseOriginPublicDatabase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPC{}),
		},
	)
}

type ConfigGetResponseOriginPublicDatabase struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host string `json:"host" api:"required"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port int64 `json:"port" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigGetResponseOriginPublicDatabaseScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                    `json:"user" api:"required"`
	JSON configGetResponseOriginPublicDatabaseJSON `json:"-"`
}

// configGetResponseOriginPublicDatabaseJSON contains the JSON metadata for the
// struct [ConfigGetResponseOriginPublicDatabase]
type configGetResponseOriginPublicDatabaseJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseOriginPublicDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseOriginPublicDatabaseJSON) RawJSON() string {
	return r.raw
}

func (r ConfigGetResponseOriginPublicDatabase) implementsConfigGetResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigGetResponseOriginPublicDatabaseScheme string

const (
	ConfigGetResponseOriginPublicDatabaseSchemePostgres   ConfigGetResponseOriginPublicDatabaseScheme = "postgres"
	ConfigGetResponseOriginPublicDatabaseSchemePostgresql ConfigGetResponseOriginPublicDatabaseScheme = "postgresql"
	ConfigGetResponseOriginPublicDatabaseSchemeMysql      ConfigGetResponseOriginPublicDatabaseScheme = "mysql"
)

func (r ConfigGetResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigGetResponseOriginPublicDatabaseSchemePostgres, ConfigGetResponseOriginPublicDatabaseSchemePostgresql, ConfigGetResponseOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id" api:"required"`
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User string                                                                   `json:"user" api:"required"`
	JSON configGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
}

// configGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON
// contains the JSON metadata for the struct
// [ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type configGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsConfigGetResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

type ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPC struct {
	// Set the name of your origin database.
	Database string `json:"database" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCScheme `json:"scheme" api:"required"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID string `json:"service_id" api:"required"`
	// Set the user of your origin database.
	User string                                                         `json:"user" api:"required"`
	JSON configGetResponseOriginDatabaseReachableThroughAWorkersVPCJSON `json:"-"`
}

// configGetResponseOriginDatabaseReachableThroughAWorkersVPCJSON contains the JSON
// metadata for the struct
// [ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPC]
type configGetResponseOriginDatabaseReachableThroughAWorkersVPCJSON struct {
	Database    apijson.Field
	Scheme      apijson.Field
	ServiceID   apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseOriginDatabaseReachableThroughAWorkersVPCJSON) RawJSON() string {
	return r.raw
}

func (r ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPC) implementsConfigGetResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCScheme string

const (
	ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres   ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgres"
	ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "postgresql"
	ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql      ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCScheme = "mysql"
)

func (r ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCScheme) IsKnown() bool {
	switch r {
	case ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgres, ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql, ConfigGetResponseOriginDatabaseReachableThroughAWorkersVPCSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigGetResponseOriginScheme string

const (
	ConfigGetResponseOriginSchemePostgres   ConfigGetResponseOriginScheme = "postgres"
	ConfigGetResponseOriginSchemePostgresql ConfigGetResponseOriginScheme = "postgresql"
	ConfigGetResponseOriginSchemeMysql      ConfigGetResponseOriginScheme = "mysql"
)

func (r ConfigGetResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigGetResponseOriginSchemePostgres, ConfigGetResponseOriginSchemePostgresql, ConfigGetResponseOriginSchemeMysql:
		return true
	}
	return false
}

// Connects to a PlanetScale database using credentials managed by Cloudflare. The
// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
// dashboard.
type ConfigGetResponseIntegration struct {
	// The name of the PlanetScale database branch.
	DatabaseBranchName string `json:"database_branch_name" api:"required"`
	// The name of the PlanetScale database.
	DatabaseName string `json:"database_name" api:"required"`
	// The database integration used by this operation.
	Integration ConfigGetResponseIntegrationIntegration `json:"integration" api:"required"`
	// The name of the PlanetScale organization.
	OrganizationName string `json:"organization_name" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigGetResponseIntegrationScheme `json:"scheme" api:"required"`
	// The database name to use when connecting. Defaults to `postgres` for PostgreSQL
	// and `mysql` for MySQL.
	CustomDatabaseName string                           `json:"custom_database_name"`
	JSON               configGetResponseIntegrationJSON `json:"-"`
}

// configGetResponseIntegrationJSON contains the JSON metadata for the struct
// [ConfigGetResponseIntegration]
type configGetResponseIntegrationJSON struct {
	DatabaseBranchName apijson.Field
	DatabaseName       apijson.Field
	Integration        apijson.Field
	OrganizationName   apijson.Field
	Scheme             apijson.Field
	CustomDatabaseName apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConfigGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// The database integration used by this operation.
type ConfigGetResponseIntegrationIntegration string

const (
	ConfigGetResponseIntegrationIntegrationPlanetscale ConfigGetResponseIntegrationIntegration = "planetscale"
)

func (r ConfigGetResponseIntegrationIntegration) IsKnown() bool {
	switch r {
	case ConfigGetResponseIntegrationIntegrationPlanetscale:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigGetResponseIntegrationScheme string

const (
	ConfigGetResponseIntegrationSchemePostgres   ConfigGetResponseIntegrationScheme = "postgres"
	ConfigGetResponseIntegrationSchemePostgresql ConfigGetResponseIntegrationScheme = "postgresql"
	ConfigGetResponseIntegrationSchemeMysql      ConfigGetResponseIntegrationScheme = "mysql"
)

func (r ConfigGetResponseIntegrationScheme) IsKnown() bool {
	switch r {
	case ConfigGetResponseIntegrationSchemePostgres, ConfigGetResponseIntegrationSchemePostgresql, ConfigGetResponseIntegrationSchemeMysql:
		return true
	}
	return false
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type ConfigGetResponseMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID string `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID string `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode string                    `json:"sslmode"`
	JSON    configGetResponseMTLSJSON `json:"-"`
}

// configGetResponseMTLSJSON contains the JSON metadata for the struct
// [ConfigGetResponseMTLS]
type configGetResponseMTLSJSON struct {
	CACertificateID   apijson.Field
	MTLSCertificateID apijson.Field
	Sslmode           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigGetResponseMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseMTLSJSON) RawJSON() string {
	return r.raw
}

type ConfigNewParams struct {
	// Define configurations using a unique string identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A request to create a Hyperdrive configuration using exactly one of
	// caller-supplied origin credentials or a managed integration.
	Body ConfigNewParamsBodyUnion `json:"body" api:"required"`
}

func (r ConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// A request to create a Hyperdrive configuration using exactly one of
// caller-supplied origin credentials or a managed integration.
type ConfigNewParamsBody struct {
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name        param.Field[string]      `json:"name" api:"required"`
	Caching     param.Field[interface{}] `json:"caching"`
	Integration param.Field[interface{}] `json:"integration"`
	MTLS        param.Field[interface{}] `json:"mtls"`
	Origin      param.Field[interface{}] `json:"origin"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit param.Field[int64] `json:"origin_connection_limit"`
}

func (r ConfigNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBody) implementsConfigNewParamsBodyUnion() {}

// A request to create a Hyperdrive configuration using exactly one of
// caller-supplied origin credentials or a managed integration.
//
// Satisfied by
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOrigin],
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegration],
// [ConfigNewParamsBody].
type ConfigNewParamsBodyUnion interface {
	implementsConfigNewParamsBodyUnion()
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOrigin struct {
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name param.Field[string] `json:"name" api:"required"`
	// Combines database connection fields with exactly one supported network location.
	Origin      param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginUnion]  `json:"origin" api:"required"`
	Caching     param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingUnion] `json:"caching"`
	Integration param.Field[interface{}]                                                               `json:"integration"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginMTLS] `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit param.Field[int64] `json:"origin_connection_limit"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOrigin) implementsConfigNewParamsBodyUnion() {
}

// Combines database connection fields with exactly one supported network location.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOrigin struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database" api:"required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginScheme] `json:"scheme" api:"required"`
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

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOrigin) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginUnion() {
}

// Combines database connection fields with exactly one supported network location.
//
// Satisfied by
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabase],
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnel],
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPC],
// [ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOrigin].
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginUnion interface {
	implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginUnion()
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabase struct {
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
	Scheme param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseScheme] `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user" api:"required"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabase) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginUnion() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseScheme string

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseSchemePostgres   ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseScheme = "postgres"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseSchemePostgresql ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseScheme = "postgresql"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseSchemeMysql      ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseScheme = "mysql"
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseSchemePostgres, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseSchemePostgresql, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
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
	Scheme param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme] `json:"scheme" api:"required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user" api:"required"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginUnion() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPC struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database" api:"required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCScheme] `json:"scheme" api:"required"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID param.Field[string] `json:"service_id" api:"required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user" api:"required"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPC) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPC) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginUnion() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCScheme string

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCSchemePostgres   ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCScheme = "postgres"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCScheme = "postgresql"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCSchemeMysql      ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCScheme = "mysql"
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCScheme) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCSchemePostgres, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCSchemePostgresql, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginDatabaseReachableThroughAWorkersVPCSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginScheme string

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginSchemePostgres   ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginScheme = "postgres"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginSchemePostgresql ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginScheme = "postgresql"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginSchemeMysql      ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginScheme = "mysql"
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginScheme) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginSchemePostgres, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginSchemePostgresql, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginOriginSchemeMysql:
		return true
	}
	return false
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCaching struct {
	Disabled             param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabled] `json:"disabled"`
	MaxAge               param.Field[int64]                                                                        `json:"max_age"`
	StaleWhileRevalidate param.Field[int64]                                                                        `json:"stale_while_revalidate"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCaching) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingUnion() {
}

// Satisfied by
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabled],
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabled],
// [ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCaching].
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingUnion interface {
	implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingUnion()
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabled struct {
	Disabled             param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabledDisabled] `json:"disabled" api:"required"`
	MaxAge               param.Field[int64]                                                                                                                 `json:"max_age"`
	StaleWhileRevalidate param.Field[int64]                                                                                                                 `json:"stale_while_revalidate"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabled) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingUnion() {
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabledDisabled bool

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabledDisabledTrue ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabledDisabled = true
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabledDisabled) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateDisabledDisabledTrue:
		return true
	}
	return false
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabled struct {
	Disabled param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabledDisabled] `json:"disabled"`
	// Specify the maximum duration (in seconds) items should persist in the cache.
	// Defaults to 60 seconds if not specified.
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Defaults to
	// 15 seconds if not specified.
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabled) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingUnion() {
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabledDisabled bool

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabledDisabledFalse ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabledDisabled = false
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabledDisabled) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingHyperdriveHyperdriveCachingCreateEnabledDisabledFalse:
		return true
	}
	return false
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabled bool

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabledTrue  ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabled = true
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabledFalse ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabled = false
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabled) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabledTrue, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginCachingDisabledFalse:
		return true
	}
	return false
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID param.Field[string] `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID param.Field[string] `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode param.Field[string] `json:"sslmode"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithOriginMTLS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegration struct {
	// Connects to a PlanetScale database using credentials managed by Cloudflare. The
	// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
	// dashboard.
	Integration param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegration] `json:"integration" api:"required"`
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API.
	Name    param.Field[string]                                                                         `json:"name" api:"required"`
	Caching param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingUnion] `json:"caching"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS   param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationMTLS] `json:"mtls"`
	Origin param.Field[interface{}]                                                            `json:"origin"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit param.Field[int64] `json:"origin_connection_limit"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegration) implementsConfigNewParamsBodyUnion() {
}

// Connects to a PlanetScale database using credentials managed by Cloudflare. The
// Cloudflare account must already be linked to PlanetScale in the Hyperdrive
// dashboard.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegration struct {
	// The name of the PlanetScale database branch.
	DatabaseBranchName param.Field[string] `json:"database_branch_name" api:"required"`
	// The name of the PlanetScale database.
	DatabaseName param.Field[string] `json:"database_name" api:"required"`
	// The database integration used by this operation.
	Integration param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationIntegration] `json:"integration" api:"required"`
	// The name of the PlanetScale organization.
	OrganizationName param.Field[string] `json:"organization_name" api:"required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationScheme] `json:"scheme" api:"required"`
	// The database name to use when connecting. Defaults to `postgres` for PostgreSQL
	// and `mysql` for MySQL.
	CustomDatabaseName param.Field[string] `json:"custom_database_name"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The database integration used by this operation.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationIntegration string

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationIntegrationPlanetscale ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationIntegration = "planetscale"
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationIntegration) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationIntegrationPlanetscale:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationScheme string

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationSchemePostgres   ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationScheme = "postgres"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationSchemePostgresql ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationScheme = "postgresql"
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationSchemeMysql      ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationScheme = "mysql"
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationScheme) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationSchemePostgres, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationSchemePostgresql, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationIntegrationSchemeMysql:
		return true
	}
	return false
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCaching struct {
	Disabled             param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabled] `json:"disabled"`
	MaxAge               param.Field[int64]                                                                             `json:"max_age"`
	StaleWhileRevalidate param.Field[int64]                                                                             `json:"stale_while_revalidate"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCaching) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingUnion() {
}

// Satisfied by
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabled],
// [hyperdrive.ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabled],
// [ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCaching].
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingUnion interface {
	implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingUnion()
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabled struct {
	Disabled             param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabledDisabled] `json:"disabled" api:"required"`
	MaxAge               param.Field[int64]                                                                                                                      `json:"max_age"`
	StaleWhileRevalidate param.Field[int64]                                                                                                                      `json:"stale_while_revalidate"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabled) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingUnion() {
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabledDisabled bool

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabledDisabledTrue ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabledDisabled = true
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabledDisabled) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateDisabledDisabledTrue:
		return true
	}
	return false
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabled struct {
	Disabled param.Field[ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabledDisabled] `json:"disabled"`
	// Specify the maximum duration (in seconds) items should persist in the cache.
	// Defaults to 60 seconds if not specified.
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Defaults to
	// 15 seconds if not specified.
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabled) implementsConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingUnion() {
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabledDisabled bool

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabledDisabledFalse ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabledDisabled = false
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabledDisabled) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingHyperdriveHyperdriveCachingCreateEnabledDisabledFalse:
		return true
	}
	return false
}

type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabled bool

const (
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabledTrue  ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabled = true
	ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabledFalse ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabled = false
)

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabled) IsKnown() bool {
	switch r {
	case ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabledTrue, ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationCachingDisabledFalse:
		return true
	}
	return false
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID param.Field[string] `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID param.Field[string] `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode param.Field[string] `json:"sslmode"`
}

func (r ConfigNewParamsBodyHyperdriveHyperdriveConfigCreateWithIntegrationMTLS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   ConfigNewResponse     `json:"result" api:"required"`
	// Return the status of the API call success.
	Success ConfigNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    configNewResponseEnvelopeJSON    `json:"-"`
}

// configNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigNewResponseEnvelope]
type configNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type ConfigNewResponseEnvelopeSuccess bool

const (
	ConfigNewResponseEnvelopeSuccessTrue ConfigNewResponseEnvelopeSuccess = true
)

func (r ConfigNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigUpdateParams struct {
	// Define configurations using a unique string identifier.
	AccountID  param.Field[string] `path:"account_id" api:"required"`
	Hyperdrive HyperdriveParam     `json:"hyperdrive" api:"required"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Hyperdrive)
}

type ConfigUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   ConfigUpdateResponse  `json:"result" api:"required"`
	// Return the status of the API call success.
	Success ConfigUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    configUpdateResponseEnvelopeJSON    `json:"-"`
}

// configUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseEnvelope]
type configUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type ConfigUpdateResponseEnvelopeSuccess bool

const (
	ConfigUpdateResponseEnvelopeSuccessTrue ConfigUpdateResponseEnvelopeSuccess = true
)

func (r ConfigUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigListParams struct {
	// Define configurations using a unique string identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ConfigListParams]'s query parameters as `url.Values`.
func (r ConfigListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConfigDeleteParams struct {
	// Define configurations using a unique string identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConfigDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   ConfigDeleteResponse  `json:"result" api:"required,nullable"`
	// Return the status of the API call success.
	Success ConfigDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    configDeleteResponseEnvelopeJSON    `json:"-"`
}

// configDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigDeleteResponseEnvelope]
type configDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type ConfigDeleteResponseEnvelopeSuccess bool

const (
	ConfigDeleteResponseEnvelopeSuccessTrue ConfigDeleteResponseEnvelopeSuccess = true
)

func (r ConfigDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigEditParams struct {
	// Define configurations using a unique string identifier.
	AccountID param.Field[string]                       `path:"account_id" api:"required"`
	Caching   param.Field[ConfigEditParamsCachingUnion] `json:"caching"`
	// mTLS configuration for the origin connection. Cannot be used with VPC Service
	// origins; TLS must be managed on the VPC Service.
	MTLS param.Field[ConfigEditParamsMTLS] `json:"mtls"`
	// The name of the Hyperdrive configuration. Used to identify the configuration in
	// the Cloudflare dashboard and API. An empty value leaves the name unchanged.
	Name param.Field[string] `json:"name"`
	// Connect to a database through a Workers VPC Service. TLS settings (mTLS,
	// sslmode) cannot be configured on the Hyperdrive when using a VPC Service origin;
	// TLS must be managed on the VPC Service itself.
	Origin param.Field[ConfigEditParamsOriginUnion] `json:"origin"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts. If not
	// specified, defaults to 20 for free tier and 60 for paid tier. Certain
	// Cloudflare-managed origins may be permitted a higher limit. Contact Cloudflare
	// if you need a higher limit.
	OriginConnectionLimit param.Field[int64] `json:"origin_connection_limit"`
}

func (r ConfigEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigEditParamsCaching struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
	// Specify the maximum duration (in seconds) items should persist in the cache.
	// Defaults to 60 seconds if not specified.
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Defaults to
	// 15 seconds if not specified.
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigEditParamsCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsCaching) implementsConfigEditParamsCachingUnion() {}

// Satisfied by
// [hyperdrive.ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigEditParamsCachingHyperdriveHyperdriveCachingEnabled],
// [ConfigEditParamsCaching].
type ConfigEditParamsCachingUnion interface {
	implementsConfigEditParamsCachingUnion()
}

type ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
}

func (r ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon) implementsConfigEditParamsCachingUnion() {
}

type ConfigEditParamsCachingHyperdriveHyperdriveCachingEnabled struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
	// Specify the maximum duration (in seconds) items should persist in the cache.
	// Defaults to 60 seconds if not specified.
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Defaults to
	// 15 seconds if not specified.
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigEditParamsCachingHyperdriveHyperdriveCachingEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsCachingHyperdriveHyperdriveCachingEnabled) implementsConfigEditParamsCachingUnion() {
}

// mTLS configuration for the origin connection. Cannot be used with VPC Service
// origins; TLS must be managed on the VPC Service.
type ConfigEditParamsMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID param.Field[string] `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID param.Field[string] `json:"mtls_certificate_id"`
	// PostgreSQL accepts `require`, `verify-ca`, and `verify-full`. MySQL accepts
	// `REQUIRED`, `VERIFY_CA`, and `VERIFY_IDENTITY`. The verify modes require a CA
	// certificate; the require modes cannot be used with a CA certificate.
	Sslmode param.Field[string] `json:"sslmode"`
}

func (r ConfigEditParamsMTLS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Connect to a database through a Workers VPC Service. TLS settings (mTLS,
// sslmode) cannot be configured on the Hyperdrive when using a VPC Service origin;
// TLS must be managed on the VPC Service itself.
type ConfigEditParamsOrigin struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// Set the name of your origin database.
	Database param.Field[string] `json:"database"`
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host param.Field[string] `json:"host"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port param.Field[int64] `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigEditParamsOriginScheme] `json:"scheme"`
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID param.Field[string] `json:"service_id"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r ConfigEditParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOrigin) implementsConfigEditParamsOriginUnion() {}

// Connect to a database through a Workers VPC Service. TLS settings (mTLS,
// sslmode) cannot be configured on the Hyperdrive when using a VPC Service origin;
// TLS must be managed on the VPC Service itself.
//
// Satisfied by [hyperdrive.ConfigEditParamsOriginHyperdriveHyperdriveDatabase],
// [hyperdrive.ConfigEditParamsOriginHyperdriveInternetOrigin],
// [hyperdrive.ConfigEditParamsOriginHyperdriveOverAccessOrigin],
// [hyperdrive.ConfigEditParamsOriginHyperdriveVPCServiceOrigin],
// [ConfigEditParamsOrigin].
type ConfigEditParamsOriginUnion interface {
	implementsConfigEditParamsOriginUnion()
}

type ConfigEditParamsOriginHyperdriveHyperdriveDatabase struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme] `json:"scheme"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r ConfigEditParamsOriginHyperdriveHyperdriveDatabase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOriginHyperdriveHyperdriveDatabase) implementsConfigEditParamsOriginUnion() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme string

const (
	ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgres   ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme = "postgres"
	ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgresql ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme = "postgresql"
	ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemeMysql      ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme = "mysql"
)

func (r ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgres, ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgresql, ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemeMysql:
		return true
	}
	return false
}

type ConfigEditParamsOriginHyperdriveInternetOrigin struct {
	// Defines the publicly reachable hostname or IP of your origin database. Private,
	// loopback, and link-local IP addresses are not allowed.
	Host param.Field[string] `json:"host" api:"required"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or
	// 3306 for MySQL if not specified.
	Port param.Field[int64] `json:"port" api:"required"`
}

func (r ConfigEditParamsOriginHyperdriveInternetOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOriginHyperdriveInternetOrigin) implementsConfigEditParamsOriginUnion() {}

type ConfigEditParamsOriginHyperdriveOverAccessOrigin struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id" api:"required"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret" api:"required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host" api:"required"`
}

func (r ConfigEditParamsOriginHyperdriveOverAccessOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOriginHyperdriveOverAccessOrigin) implementsConfigEditParamsOriginUnion() {}

// Connect to a database through a Workers VPC Service. TLS settings (mTLS,
// sslmode) cannot be configured on the Hyperdrive when using a VPC Service origin;
// TLS must be managed on the VPC Service itself.
type ConfigEditParamsOriginHyperdriveVPCServiceOrigin struct {
	// The identifier of the Workers VPC Service to connect through. Hyperdrive will
	// egress through the specified VPC Service to reach the origin database.
	ServiceID param.Field[string] `json:"service_id" api:"required"`
}

func (r ConfigEditParamsOriginHyperdriveVPCServiceOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOriginHyperdriveVPCServiceOrigin) implementsConfigEditParamsOriginUnion() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditParamsOriginScheme string

const (
	ConfigEditParamsOriginSchemePostgres   ConfigEditParamsOriginScheme = "postgres"
	ConfigEditParamsOriginSchemePostgresql ConfigEditParamsOriginScheme = "postgresql"
	ConfigEditParamsOriginSchemeMysql      ConfigEditParamsOriginScheme = "mysql"
)

func (r ConfigEditParamsOriginScheme) IsKnown() bool {
	switch r {
	case ConfigEditParamsOriginSchemePostgres, ConfigEditParamsOriginSchemePostgresql, ConfigEditParamsOriginSchemeMysql:
		return true
	}
	return false
}

type ConfigEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   ConfigEditResponse    `json:"result" api:"required"`
	// Return the status of the API call success.
	Success ConfigEditResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    configEditResponseEnvelopeJSON    `json:"-"`
}

// configEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigEditResponseEnvelope]
type configEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type ConfigEditResponseEnvelopeSuccess bool

const (
	ConfigEditResponseEnvelopeSuccessTrue ConfigEditResponseEnvelopeSuccess = true
)

func (r ConfigEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigGetParams struct {
	// Define configurations using a unique string identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConfigGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   ConfigGetResponse     `json:"result" api:"required"`
	// Return the status of the API call success.
	Success ConfigGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    configGetResponseEnvelopeJSON    `json:"-"`
}

// configGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelope]
type configGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type ConfigGetResponseEnvelopeSuccess bool

const (
	ConfigGetResponseEnvelopeSuccessTrue ConfigGetResponseEnvelopeSuccess = true
)

func (r ConfigGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
