// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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

// Creates and returns a new Hyperdrive configuration.
func (r *ConfigService) New(ctx context.Context, params ConfigNewParams, opts ...option.RequestOption) (res *ConfigNewResponse, err error) {
	var env ConfigNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates and returns the specified Hyperdrive configuration.
func (r *ConfigService) Update(ctx context.Context, hyperdriveID string, params ConfigUpdateParams, opts ...option.RequestOption) (res *ConfigUpdateResponse, err error) {
	var env ConfigUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", params.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of Hyperdrives
func (r *ConfigService) List(ctx context.Context, query ConfigListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ConfigListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Returns a list of Hyperdrives
func (r *ConfigService) ListAutoPaging(ctx context.Context, query ConfigListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ConfigListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the specified Hyperdrive.
func (r *ConfigService) Delete(ctx context.Context, hyperdriveID string, body ConfigDeleteParams, opts ...option.RequestOption) (res *ConfigDeleteResponse, err error) {
	var env ConfigDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", body.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches and returns the specified Hyperdrive configuration. Custom caching
// settings are not kept if caching is disabled.
func (r *ConfigService) Edit(ctx context.Context, hyperdriveID string, params ConfigEditParams, opts ...option.RequestOption) (res *ConfigEditResponse, err error) {
	var env ConfigEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", params.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified Hyperdrive configuration.
func (r *ConfigService) Get(ctx context.Context, hyperdriveID string, query ConfigGetParams, opts ...option.RequestOption) (res *ConfigGetResponse, err error) {
	var env ConfigGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", query.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConfigNewResponse struct {
	Caching ConfigNewResponseCaching `json:"caching,required"`
	// Identifier
	ID     string                  `json:"id"`
	Name   string                  `json:"name"`
	Origin ConfigNewResponseOrigin `json:"origin"`
	JSON   configNewResponseJSON   `json:"-"`
}

// configNewResponseJSON contains the JSON metadata for the struct
// [ConfigNewResponse]
type configNewResponseJSON struct {
	Caching     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigNewResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                        `json:"stale_while_revalidate"`
	JSON                 configNewResponseCachingJSON `json:"-"`
	union                ConfigNewResponseCachingUnion
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

func (r configNewResponseCachingJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigNewResponseCaching) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigNewResponseCaching{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigNewResponseCachingUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [hyperdrive.ConfigNewResponseCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigNewResponseCachingHyperdriveHyperdriveCachingEnabled].
func (r ConfigNewResponseCaching) AsUnion() ConfigNewResponseCachingUnion {
	return r.union
}

// Union satisfied by
// [hyperdrive.ConfigNewResponseCachingHyperdriveHyperdriveCachingCommon] or
// [hyperdrive.ConfigNewResponseCachingHyperdriveHyperdriveCachingEnabled].
type ConfigNewResponseCachingUnion interface {
	implementsHyperdriveConfigNewResponseCaching()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigNewResponseCachingUnion)(nil)).Elem(),
		"disabled",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigNewResponseCachingHyperdriveHyperdriveCachingCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigNewResponseCachingHyperdriveHyperdriveCachingEnabled{}),
		},
	)
}

type ConfigNewResponseCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool                                                          `json:"disabled"`
	JSON     configNewResponseCachingHyperdriveHyperdriveCachingCommonJSON `json:"-"`
}

// configNewResponseCachingHyperdriveHyperdriveCachingCommonJSON contains the JSON
// metadata for the struct
// [ConfigNewResponseCachingHyperdriveHyperdriveCachingCommon]
type configNewResponseCachingHyperdriveHyperdriveCachingCommonJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseCachingHyperdriveHyperdriveCachingCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseCachingHyperdriveHyperdriveCachingCommonJSON) RawJSON() string {
	return r.raw
}

func (r ConfigNewResponseCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveConfigNewResponseCaching() {
}

type ConfigNewResponseCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                                                          `json:"stale_while_revalidate"`
	JSON                 configNewResponseCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
}

// configNewResponseCachingHyperdriveHyperdriveCachingEnabledJSON contains the JSON
// metadata for the struct
// [ConfigNewResponseCachingHyperdriveHyperdriveCachingEnabled]
type configNewResponseCachingHyperdriveHyperdriveCachingEnabledJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigNewResponseCachingHyperdriveHyperdriveCachingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseCachingHyperdriveHyperdriveCachingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigNewResponseCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveConfigNewResponseCaching() {
}

type ConfigNewResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigNewResponseOriginScheme `json:"scheme,required"`
	// The user of your origin database.
	User string `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port  int64                       `json:"port"`
	JSON  configNewResponseOriginJSON `json:"-"`
	union ConfigNewResponseOriginUnion
}

// configNewResponseOriginJSON contains the JSON metadata for the struct
// [ConfigNewResponseOrigin]
type configNewResponseOriginJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
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
// Possible runtime types of the union are
// [hyperdrive.ConfigNewResponseOriginPublicDatabase],
// [hyperdrive.ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
func (r ConfigNewResponseOrigin) AsUnion() ConfigNewResponseOriginUnion {
	return r.union
}

// Union satisfied by [hyperdrive.ConfigNewResponseOriginPublicDatabase] or
// [hyperdrive.ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
type ConfigNewResponseOriginUnion interface {
	implementsHyperdriveConfigNewResponseOrigin()
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
	)
}

type ConfigNewResponseOriginPublicDatabase struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigNewResponseOriginPublicDatabaseScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                    `json:"user,required"`
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

func (r ConfigNewResponseOriginPublicDatabase) implementsHyperdriveConfigNewResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewResponseOriginPublicDatabaseScheme string

const (
	ConfigNewResponseOriginPublicDatabaseSchemePostgres   ConfigNewResponseOriginPublicDatabaseScheme = "postgres"
	ConfigNewResponseOriginPublicDatabaseSchemePostgresql ConfigNewResponseOriginPublicDatabaseScheme = "postgresql"
)

func (r ConfigNewResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigNewResponseOriginPublicDatabaseSchemePostgres, ConfigNewResponseOriginPublicDatabaseSchemePostgresql:
		return true
	}
	return false
}

type ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id,required"`
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                                                   `json:"user,required"`
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

func (r ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveConfigNewResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
)

func (r ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigNewResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewResponseOriginScheme string

const (
	ConfigNewResponseOriginSchemePostgres   ConfigNewResponseOriginScheme = "postgres"
	ConfigNewResponseOriginSchemePostgresql ConfigNewResponseOriginScheme = "postgresql"
)

func (r ConfigNewResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigNewResponseOriginSchemePostgres, ConfigNewResponseOriginSchemePostgresql:
		return true
	}
	return false
}

type ConfigUpdateResponse struct {
	Caching ConfigUpdateResponseCaching `json:"caching,required"`
	// Identifier
	ID     string                     `json:"id"`
	Name   string                     `json:"name"`
	Origin ConfigUpdateResponseOrigin `json:"origin"`
	JSON   configUpdateResponseJSON   `json:"-"`
}

// configUpdateResponseJSON contains the JSON metadata for the struct
// [ConfigUpdateResponse]
type configUpdateResponseJSON struct {
	Caching     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                           `json:"stale_while_revalidate"`
	JSON                 configUpdateResponseCachingJSON `json:"-"`
	union                ConfigUpdateResponseCachingUnion
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

func (r configUpdateResponseCachingJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigUpdateResponseCaching) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigUpdateResponseCaching{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigUpdateResponseCachingUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [hyperdrive.ConfigUpdateResponseCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigUpdateResponseCachingHyperdriveHyperdriveCachingEnabled].
func (r ConfigUpdateResponseCaching) AsUnion() ConfigUpdateResponseCachingUnion {
	return r.union
}

// Union satisfied by
// [hyperdrive.ConfigUpdateResponseCachingHyperdriveHyperdriveCachingCommon] or
// [hyperdrive.ConfigUpdateResponseCachingHyperdriveHyperdriveCachingEnabled].
type ConfigUpdateResponseCachingUnion interface {
	implementsHyperdriveConfigUpdateResponseCaching()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigUpdateResponseCachingUnion)(nil)).Elem(),
		"disabled",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigUpdateResponseCachingHyperdriveHyperdriveCachingCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigUpdateResponseCachingHyperdriveHyperdriveCachingEnabled{}),
		},
	)
}

type ConfigUpdateResponseCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool                                                             `json:"disabled"`
	JSON     configUpdateResponseCachingHyperdriveHyperdriveCachingCommonJSON `json:"-"`
}

// configUpdateResponseCachingHyperdriveHyperdriveCachingCommonJSON contains the
// JSON metadata for the struct
// [ConfigUpdateResponseCachingHyperdriveHyperdriveCachingCommon]
type configUpdateResponseCachingHyperdriveHyperdriveCachingCommonJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseCachingHyperdriveHyperdriveCachingCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseCachingHyperdriveHyperdriveCachingCommonJSON) RawJSON() string {
	return r.raw
}

func (r ConfigUpdateResponseCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveConfigUpdateResponseCaching() {
}

type ConfigUpdateResponseCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                                                             `json:"stale_while_revalidate"`
	JSON                 configUpdateResponseCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
}

// configUpdateResponseCachingHyperdriveHyperdriveCachingEnabledJSON contains the
// JSON metadata for the struct
// [ConfigUpdateResponseCachingHyperdriveHyperdriveCachingEnabled]
type configUpdateResponseCachingHyperdriveHyperdriveCachingEnabledJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigUpdateResponseCachingHyperdriveHyperdriveCachingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseCachingHyperdriveHyperdriveCachingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigUpdateResponseCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveConfigUpdateResponseCaching() {
}

type ConfigUpdateResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigUpdateResponseOriginScheme `json:"scheme,required"`
	// The user of your origin database.
	User string `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port  int64                          `json:"port"`
	JSON  configUpdateResponseOriginJSON `json:"-"`
	union ConfigUpdateResponseOriginUnion
}

// configUpdateResponseOriginJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseOrigin]
type configUpdateResponseOriginJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
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
// [hyperdrive.ConfigUpdateResponseOriginPublicDatabase],
// [hyperdrive.ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
func (r ConfigUpdateResponseOrigin) AsUnion() ConfigUpdateResponseOriginUnion {
	return r.union
}

// Union satisfied by [hyperdrive.ConfigUpdateResponseOriginPublicDatabase] or
// [hyperdrive.ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
type ConfigUpdateResponseOriginUnion interface {
	implementsHyperdriveConfigUpdateResponseOrigin()
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
	)
}

type ConfigUpdateResponseOriginPublicDatabase struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigUpdateResponseOriginPublicDatabaseScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                       `json:"user,required"`
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

func (r ConfigUpdateResponseOriginPublicDatabase) implementsHyperdriveConfigUpdateResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateResponseOriginPublicDatabaseScheme string

const (
	ConfigUpdateResponseOriginPublicDatabaseSchemePostgres   ConfigUpdateResponseOriginPublicDatabaseScheme = "postgres"
	ConfigUpdateResponseOriginPublicDatabaseSchemePostgresql ConfigUpdateResponseOriginPublicDatabaseScheme = "postgresql"
)

func (r ConfigUpdateResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseOriginPublicDatabaseSchemePostgres, ConfigUpdateResponseOriginPublicDatabaseSchemePostgresql:
		return true
	}
	return false
}

type ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id,required"`
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                                                      `json:"user,required"`
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

func (r ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveConfigUpdateResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
)

func (r ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigUpdateResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateResponseOriginScheme string

const (
	ConfigUpdateResponseOriginSchemePostgres   ConfigUpdateResponseOriginScheme = "postgres"
	ConfigUpdateResponseOriginSchemePostgresql ConfigUpdateResponseOriginScheme = "postgresql"
)

func (r ConfigUpdateResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseOriginSchemePostgres, ConfigUpdateResponseOriginSchemePostgresql:
		return true
	}
	return false
}

type ConfigListResponse struct {
	Caching ConfigListResponseCaching `json:"caching,required"`
	// Identifier
	ID     string                   `json:"id"`
	Name   string                   `json:"name"`
	Origin ConfigListResponseOrigin `json:"origin"`
	JSON   configListResponseJSON   `json:"-"`
}

// configListResponseJSON contains the JSON metadata for the struct
// [ConfigListResponse]
type configListResponseJSON struct {
	Caching     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigListResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                         `json:"stale_while_revalidate"`
	JSON                 configListResponseCachingJSON `json:"-"`
	union                ConfigListResponseCachingUnion
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

func (r configListResponseCachingJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigListResponseCaching) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigListResponseCaching{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigListResponseCachingUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [hyperdrive.ConfigListResponseCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigListResponseCachingHyperdriveHyperdriveCachingEnabled].
func (r ConfigListResponseCaching) AsUnion() ConfigListResponseCachingUnion {
	return r.union
}

// Union satisfied by
// [hyperdrive.ConfigListResponseCachingHyperdriveHyperdriveCachingCommon] or
// [hyperdrive.ConfigListResponseCachingHyperdriveHyperdriveCachingEnabled].
type ConfigListResponseCachingUnion interface {
	implementsHyperdriveConfigListResponseCaching()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigListResponseCachingUnion)(nil)).Elem(),
		"disabled",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigListResponseCachingHyperdriveHyperdriveCachingCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigListResponseCachingHyperdriveHyperdriveCachingEnabled{}),
		},
	)
}

type ConfigListResponseCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool                                                           `json:"disabled"`
	JSON     configListResponseCachingHyperdriveHyperdriveCachingCommonJSON `json:"-"`
}

// configListResponseCachingHyperdriveHyperdriveCachingCommonJSON contains the JSON
// metadata for the struct
// [ConfigListResponseCachingHyperdriveHyperdriveCachingCommon]
type configListResponseCachingHyperdriveHyperdriveCachingCommonJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigListResponseCachingHyperdriveHyperdriveCachingCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseCachingHyperdriveHyperdriveCachingCommonJSON) RawJSON() string {
	return r.raw
}

func (r ConfigListResponseCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveConfigListResponseCaching() {
}

type ConfigListResponseCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                                                           `json:"stale_while_revalidate"`
	JSON                 configListResponseCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
}

// configListResponseCachingHyperdriveHyperdriveCachingEnabledJSON contains the
// JSON metadata for the struct
// [ConfigListResponseCachingHyperdriveHyperdriveCachingEnabled]
type configListResponseCachingHyperdriveHyperdriveCachingEnabledJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigListResponseCachingHyperdriveHyperdriveCachingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseCachingHyperdriveHyperdriveCachingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigListResponseCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveConfigListResponseCaching() {
}

type ConfigListResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigListResponseOriginScheme `json:"scheme,required"`
	// The user of your origin database.
	User string `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port  int64                        `json:"port"`
	JSON  configListResponseOriginJSON `json:"-"`
	union ConfigListResponseOriginUnion
}

// configListResponseOriginJSON contains the JSON metadata for the struct
// [ConfigListResponseOrigin]
type configListResponseOriginJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
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
// [hyperdrive.ConfigListResponseOriginPublicDatabase],
// [hyperdrive.ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
func (r ConfigListResponseOrigin) AsUnion() ConfigListResponseOriginUnion {
	return r.union
}

// Union satisfied by [hyperdrive.ConfigListResponseOriginPublicDatabase] or
// [hyperdrive.ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
type ConfigListResponseOriginUnion interface {
	implementsHyperdriveConfigListResponseOrigin()
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
	)
}

type ConfigListResponseOriginPublicDatabase struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigListResponseOriginPublicDatabaseScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                     `json:"user,required"`
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

func (r ConfigListResponseOriginPublicDatabase) implementsHyperdriveConfigListResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigListResponseOriginPublicDatabaseScheme string

const (
	ConfigListResponseOriginPublicDatabaseSchemePostgres   ConfigListResponseOriginPublicDatabaseScheme = "postgres"
	ConfigListResponseOriginPublicDatabaseSchemePostgresql ConfigListResponseOriginPublicDatabaseScheme = "postgresql"
)

func (r ConfigListResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigListResponseOriginPublicDatabaseSchemePostgres, ConfigListResponseOriginPublicDatabaseSchemePostgresql:
		return true
	}
	return false
}

type ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id,required"`
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                                                    `json:"user,required"`
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

func (r ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveConfigListResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
)

func (r ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigListResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigListResponseOriginScheme string

const (
	ConfigListResponseOriginSchemePostgres   ConfigListResponseOriginScheme = "postgres"
	ConfigListResponseOriginSchemePostgresql ConfigListResponseOriginScheme = "postgresql"
)

func (r ConfigListResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigListResponseOriginSchemePostgres, ConfigListResponseOriginSchemePostgresql:
		return true
	}
	return false
}

type ConfigDeleteResponse = interface{}

type ConfigEditResponse struct {
	Caching ConfigEditResponseCaching `json:"caching,required"`
	// Identifier
	ID     string                   `json:"id"`
	Name   string                   `json:"name"`
	Origin ConfigEditResponseOrigin `json:"origin"`
	JSON   configEditResponseJSON   `json:"-"`
}

// configEditResponseJSON contains the JSON metadata for the struct
// [ConfigEditResponse]
type configEditResponseJSON struct {
	Caching     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigEditResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                         `json:"stale_while_revalidate"`
	JSON                 configEditResponseCachingJSON `json:"-"`
	union                ConfigEditResponseCachingUnion
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

func (r configEditResponseCachingJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigEditResponseCaching) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigEditResponseCaching{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigEditResponseCachingUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [hyperdrive.ConfigEditResponseCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigEditResponseCachingHyperdriveHyperdriveCachingEnabled].
func (r ConfigEditResponseCaching) AsUnion() ConfigEditResponseCachingUnion {
	return r.union
}

// Union satisfied by
// [hyperdrive.ConfigEditResponseCachingHyperdriveHyperdriveCachingCommon] or
// [hyperdrive.ConfigEditResponseCachingHyperdriveHyperdriveCachingEnabled].
type ConfigEditResponseCachingUnion interface {
	implementsHyperdriveConfigEditResponseCaching()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigEditResponseCachingUnion)(nil)).Elem(),
		"disabled",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigEditResponseCachingHyperdriveHyperdriveCachingCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigEditResponseCachingHyperdriveHyperdriveCachingEnabled{}),
		},
	)
}

type ConfigEditResponseCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool                                                           `json:"disabled"`
	JSON     configEditResponseCachingHyperdriveHyperdriveCachingCommonJSON `json:"-"`
}

// configEditResponseCachingHyperdriveHyperdriveCachingCommonJSON contains the JSON
// metadata for the struct
// [ConfigEditResponseCachingHyperdriveHyperdriveCachingCommon]
type configEditResponseCachingHyperdriveHyperdriveCachingCommonJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseCachingHyperdriveHyperdriveCachingCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseCachingHyperdriveHyperdriveCachingCommonJSON) RawJSON() string {
	return r.raw
}

func (r ConfigEditResponseCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveConfigEditResponseCaching() {
}

type ConfigEditResponseCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                                                           `json:"stale_while_revalidate"`
	JSON                 configEditResponseCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
}

// configEditResponseCachingHyperdriveHyperdriveCachingEnabledJSON contains the
// JSON metadata for the struct
// [ConfigEditResponseCachingHyperdriveHyperdriveCachingEnabled]
type configEditResponseCachingHyperdriveHyperdriveCachingEnabledJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigEditResponseCachingHyperdriveHyperdriveCachingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseCachingHyperdriveHyperdriveCachingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigEditResponseCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveConfigEditResponseCaching() {
}

type ConfigEditResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigEditResponseOriginScheme `json:"scheme,required"`
	// The user of your origin database.
	User string `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port  int64                        `json:"port"`
	JSON  configEditResponseOriginJSON `json:"-"`
	union ConfigEditResponseOriginUnion
}

// configEditResponseOriginJSON contains the JSON metadata for the struct
// [ConfigEditResponseOrigin]
type configEditResponseOriginJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
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
// [hyperdrive.ConfigEditResponseOriginPublicDatabase],
// [hyperdrive.ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
func (r ConfigEditResponseOrigin) AsUnion() ConfigEditResponseOriginUnion {
	return r.union
}

// Union satisfied by [hyperdrive.ConfigEditResponseOriginPublicDatabase] or
// [hyperdrive.ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
type ConfigEditResponseOriginUnion interface {
	implementsHyperdriveConfigEditResponseOrigin()
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
	)
}

type ConfigEditResponseOriginPublicDatabase struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigEditResponseOriginPublicDatabaseScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                     `json:"user,required"`
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

func (r ConfigEditResponseOriginPublicDatabase) implementsHyperdriveConfigEditResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditResponseOriginPublicDatabaseScheme string

const (
	ConfigEditResponseOriginPublicDatabaseSchemePostgres   ConfigEditResponseOriginPublicDatabaseScheme = "postgres"
	ConfigEditResponseOriginPublicDatabaseSchemePostgresql ConfigEditResponseOriginPublicDatabaseScheme = "postgresql"
)

func (r ConfigEditResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigEditResponseOriginPublicDatabaseSchemePostgres, ConfigEditResponseOriginPublicDatabaseSchemePostgresql:
		return true
	}
	return false
}

type ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id,required"`
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                                                    `json:"user,required"`
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

func (r ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveConfigEditResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
)

func (r ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigEditResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditResponseOriginScheme string

const (
	ConfigEditResponseOriginSchemePostgres   ConfigEditResponseOriginScheme = "postgres"
	ConfigEditResponseOriginSchemePostgresql ConfigEditResponseOriginScheme = "postgresql"
)

func (r ConfigEditResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigEditResponseOriginSchemePostgres, ConfigEditResponseOriginSchemePostgresql:
		return true
	}
	return false
}

type ConfigGetResponse struct {
	Caching ConfigGetResponseCaching `json:"caching,required"`
	// Identifier
	ID     string                  `json:"id"`
	Name   string                  `json:"name"`
	Origin ConfigGetResponseOrigin `json:"origin"`
	JSON   configGetResponseJSON   `json:"-"`
}

// configGetResponseJSON contains the JSON metadata for the struct
// [ConfigGetResponse]
type configGetResponseJSON struct {
	Caching     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                        `json:"stale_while_revalidate"`
	JSON                 configGetResponseCachingJSON `json:"-"`
	union                ConfigGetResponseCachingUnion
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

func (r configGetResponseCachingJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigGetResponseCaching) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigGetResponseCaching{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigGetResponseCachingUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [hyperdrive.ConfigGetResponseCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigGetResponseCachingHyperdriveHyperdriveCachingEnabled].
func (r ConfigGetResponseCaching) AsUnion() ConfigGetResponseCachingUnion {
	return r.union
}

// Union satisfied by
// [hyperdrive.ConfigGetResponseCachingHyperdriveHyperdriveCachingCommon] or
// [hyperdrive.ConfigGetResponseCachingHyperdriveHyperdriveCachingEnabled].
type ConfigGetResponseCachingUnion interface {
	implementsHyperdriveConfigGetResponseCaching()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigGetResponseCachingUnion)(nil)).Elem(),
		"disabled",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigGetResponseCachingHyperdriveHyperdriveCachingCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigGetResponseCachingHyperdriveHyperdriveCachingEnabled{}),
		},
	)
}

type ConfigGetResponseCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool                                                          `json:"disabled"`
	JSON     configGetResponseCachingHyperdriveHyperdriveCachingCommonJSON `json:"-"`
}

// configGetResponseCachingHyperdriveHyperdriveCachingCommonJSON contains the JSON
// metadata for the struct
// [ConfigGetResponseCachingHyperdriveHyperdriveCachingCommon]
type configGetResponseCachingHyperdriveHyperdriveCachingCommonJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseCachingHyperdriveHyperdriveCachingCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseCachingHyperdriveHyperdriveCachingCommonJSON) RawJSON() string {
	return r.raw
}

func (r ConfigGetResponseCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveConfigGetResponseCaching() {
}

type ConfigGetResponseCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                                                          `json:"stale_while_revalidate"`
	JSON                 configGetResponseCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
}

// configGetResponseCachingHyperdriveHyperdriveCachingEnabledJSON contains the JSON
// metadata for the struct
// [ConfigGetResponseCachingHyperdriveHyperdriveCachingEnabled]
type configGetResponseCachingHyperdriveHyperdriveCachingEnabledJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigGetResponseCachingHyperdriveHyperdriveCachingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseCachingHyperdriveHyperdriveCachingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigGetResponseCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveConfigGetResponseCaching() {
}

type ConfigGetResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigGetResponseOriginScheme `json:"scheme,required"`
	// The user of your origin database.
	User string `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port  int64                       `json:"port"`
	JSON  configGetResponseOriginJSON `json:"-"`
	union ConfigGetResponseOriginUnion
}

// configGetResponseOriginJSON contains the JSON metadata for the struct
// [ConfigGetResponseOrigin]
type configGetResponseOriginJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
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
// Possible runtime types of the union are
// [hyperdrive.ConfigGetResponseOriginPublicDatabase],
// [hyperdrive.ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
func (r ConfigGetResponseOrigin) AsUnion() ConfigGetResponseOriginUnion {
	return r.union
}

// Union satisfied by [hyperdrive.ConfigGetResponseOriginPublicDatabase] or
// [hyperdrive.ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
type ConfigGetResponseOriginUnion interface {
	implementsHyperdriveConfigGetResponseOrigin()
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
	)
}

type ConfigGetResponseOriginPublicDatabase struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigGetResponseOriginPublicDatabaseScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                    `json:"user,required"`
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

func (r ConfigGetResponseOriginPublicDatabase) implementsHyperdriveConfigGetResponseOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigGetResponseOriginPublicDatabaseScheme string

const (
	ConfigGetResponseOriginPublicDatabaseSchemePostgres   ConfigGetResponseOriginPublicDatabaseScheme = "postgres"
	ConfigGetResponseOriginPublicDatabaseSchemePostgresql ConfigGetResponseOriginPublicDatabaseScheme = "postgresql"
)

func (r ConfigGetResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigGetResponseOriginPublicDatabaseSchemePostgres, ConfigGetResponseOriginPublicDatabaseSchemePostgresql:
		return true
	}
	return false
}

type ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id,required"`
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                                                   `json:"user,required"`
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

func (r ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveConfigGetResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
)

func (r ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigGetResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigGetResponseOriginScheme string

const (
	ConfigGetResponseOriginSchemePostgres   ConfigGetResponseOriginScheme = "postgres"
	ConfigGetResponseOriginSchemePostgresql ConfigGetResponseOriginScheme = "postgresql"
)

func (r ConfigGetResponseOriginScheme) IsKnown() bool {
	switch r {
	case ConfigGetResponseOriginSchemePostgres, ConfigGetResponseOriginSchemePostgresql:
		return true
	}
	return false
}

type ConfigNewParams struct {
	// Identifier
	AccountID param.Field[string]                      `path:"account_id,required"`
	Name      param.Field[string]                      `json:"name,required"`
	Origin    param.Field[ConfigNewParamsOriginUnion]  `json:"origin,required"`
	Caching   param.Field[ConfigNewParamsCachingUnion] `json:"caching"`
}

func (r ConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigNewParamsOrigin struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigNewParamsOriginScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID param.Field[string] `json:"access_client_id"`
	// The Client Secret of the Access token to use when connecting to the origin
	// database. This value is write-only and never returned by the API.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
}

func (r ConfigNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsOrigin) implementsHyperdriveConfigNewParamsOriginUnion() {}

// Satisfied by [hyperdrive.ConfigNewParamsOriginPublicDatabase],
// [hyperdrive.ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnel],
// [ConfigNewParamsOrigin].
type ConfigNewParamsOriginUnion interface {
	implementsHyperdriveConfigNewParamsOriginUnion()
}

type ConfigNewParamsOriginPublicDatabase struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigNewParamsOriginPublicDatabaseScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r ConfigNewParamsOriginPublicDatabase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsOriginPublicDatabase) implementsHyperdriveConfigNewParamsOriginUnion() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewParamsOriginPublicDatabaseScheme string

const (
	ConfigNewParamsOriginPublicDatabaseSchemePostgres   ConfigNewParamsOriginPublicDatabaseScheme = "postgres"
	ConfigNewParamsOriginPublicDatabaseSchemePostgresql ConfigNewParamsOriginPublicDatabaseScheme = "postgresql"
)

func (r ConfigNewParamsOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigNewParamsOriginPublicDatabaseSchemePostgres, ConfigNewParamsOriginPublicDatabaseSchemePostgresql:
		return true
	}
	return false
}

type ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// The Client Secret of the Access token to use when connecting to the origin
	// database. This value is write-only and never returned by the API.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveConfigNewParamsOriginUnion() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
)

func (r ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigNewParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigNewParamsOriginScheme string

const (
	ConfigNewParamsOriginSchemePostgres   ConfigNewParamsOriginScheme = "postgres"
	ConfigNewParamsOriginSchemePostgresql ConfigNewParamsOriginScheme = "postgresql"
)

func (r ConfigNewParamsOriginScheme) IsKnown() bool {
	switch r {
	case ConfigNewParamsOriginSchemePostgres, ConfigNewParamsOriginSchemePostgresql:
		return true
	}
	return false
}

type ConfigNewParamsCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigNewParamsCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsCaching) implementsHyperdriveConfigNewParamsCachingUnion() {}

// Satisfied by
// [hyperdrive.ConfigNewParamsCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigNewParamsCachingHyperdriveHyperdriveCachingEnabled],
// [ConfigNewParamsCaching].
type ConfigNewParamsCachingUnion interface {
	implementsHyperdriveConfigNewParamsCachingUnion()
}

type ConfigNewParamsCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
}

func (r ConfigNewParamsCachingHyperdriveHyperdriveCachingCommon) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveConfigNewParamsCachingUnion() {
}

type ConfigNewParamsCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigNewParamsCachingHyperdriveHyperdriveCachingEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigNewParamsCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveConfigNewParamsCachingUnion() {
}

type ConfigNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigNewResponse     `json:"result,required"`
	// Whether the API call was successful
	Success ConfigNewResponseEnvelopeSuccess `json:"success,required"`
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

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string]                         `path:"account_id,required"`
	Name      param.Field[string]                         `json:"name,required"`
	Origin    param.Field[ConfigUpdateParamsOriginUnion]  `json:"origin,required"`
	Caching   param.Field[ConfigUpdateParamsCachingUnion] `json:"caching"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsOrigin struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigUpdateParamsOriginScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID param.Field[string] `json:"access_client_id"`
	// The Client Secret of the Access token to use when connecting to the origin
	// database. This value is write-only and never returned by the API.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
}

func (r ConfigUpdateParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsOrigin) implementsHyperdriveConfigUpdateParamsOriginUnion() {}

// Satisfied by [hyperdrive.ConfigUpdateParamsOriginPublicDatabase],
// [hyperdrive.ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnel],
// [ConfigUpdateParamsOrigin].
type ConfigUpdateParamsOriginUnion interface {
	implementsHyperdriveConfigUpdateParamsOriginUnion()
}

type ConfigUpdateParamsOriginPublicDatabase struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigUpdateParamsOriginPublicDatabaseScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r ConfigUpdateParamsOriginPublicDatabase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsOriginPublicDatabase) implementsHyperdriveConfigUpdateParamsOriginUnion() {}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateParamsOriginPublicDatabaseScheme string

const (
	ConfigUpdateParamsOriginPublicDatabaseSchemePostgres   ConfigUpdateParamsOriginPublicDatabaseScheme = "postgres"
	ConfigUpdateParamsOriginPublicDatabaseSchemePostgresql ConfigUpdateParamsOriginPublicDatabaseScheme = "postgresql"
)

func (r ConfigUpdateParamsOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsOriginPublicDatabaseSchemePostgres, ConfigUpdateParamsOriginPublicDatabaseSchemePostgresql:
		return true
	}
	return false
}

type ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// The Client Secret of the Access token to use when connecting to the origin
	// database. This value is write-only and never returned by the API.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveConfigUpdateParamsOriginUnion() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
)

func (r ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, ConfigUpdateParamsOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigUpdateParamsOriginScheme string

const (
	ConfigUpdateParamsOriginSchemePostgres   ConfigUpdateParamsOriginScheme = "postgres"
	ConfigUpdateParamsOriginSchemePostgresql ConfigUpdateParamsOriginScheme = "postgresql"
)

func (r ConfigUpdateParamsOriginScheme) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsOriginSchemePostgres, ConfigUpdateParamsOriginSchemePostgresql:
		return true
	}
	return false
}

type ConfigUpdateParamsCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigUpdateParamsCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsCaching) implementsHyperdriveConfigUpdateParamsCachingUnion() {}

// Satisfied by
// [hyperdrive.ConfigUpdateParamsCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigUpdateParamsCachingHyperdriveHyperdriveCachingEnabled],
// [ConfigUpdateParamsCaching].
type ConfigUpdateParamsCachingUnion interface {
	implementsHyperdriveConfigUpdateParamsCachingUnion()
}

type ConfigUpdateParamsCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
}

func (r ConfigUpdateParamsCachingHyperdriveHyperdriveCachingCommon) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveConfigUpdateParamsCachingUnion() {
}

type ConfigUpdateParamsCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigUpdateParamsCachingHyperdriveHyperdriveCachingEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveConfigUpdateParamsCachingUnion() {
}

type ConfigUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigUpdateResponse  `json:"result,required"`
	// Whether the API call was successful
	Success ConfigUpdateResponseEnvelopeSuccess `json:"success,required"`
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

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigDeleteResponse  `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
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

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string]                       `path:"account_id,required"`
	Caching   param.Field[ConfigEditParamsCachingUnion] `json:"caching"`
	Name      param.Field[string]                       `json:"name"`
	Origin    param.Field[ConfigEditParamsOriginUnion]  `json:"origin"`
}

func (r ConfigEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigEditParamsCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigEditParamsCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsCaching) implementsHyperdriveConfigEditParamsCachingUnion() {}

// Satisfied by
// [hyperdrive.ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.ConfigEditParamsCachingHyperdriveHyperdriveCachingEnabled],
// [ConfigEditParamsCaching].
type ConfigEditParamsCachingUnion interface {
	implementsHyperdriveConfigEditParamsCachingUnion()
}

type ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
}

func (r ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveConfigEditParamsCachingUnion() {
}

type ConfigEditParamsCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r ConfigEditParamsCachingHyperdriveHyperdriveCachingEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveConfigEditParamsCachingUnion() {
}

type ConfigEditParamsOrigin struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID param.Field[string] `json:"access_client_id"`
	// The Client Secret of the Access token to use when connecting to the origin
	// database. This value is write-only and never returned by the API.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// The name of your origin database.
	Database param.Field[string] `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigEditParamsOriginScheme] `json:"scheme"`
	// The user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r ConfigEditParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOrigin) implementsHyperdriveConfigEditParamsOriginUnion() {}

// Satisfied by [hyperdrive.ConfigEditParamsOriginHyperdriveHyperdriveDatabase],
// [hyperdrive.ConfigEditParamsOriginHyperdriveHyperdriveInternetOrigin],
// [hyperdrive.ConfigEditParamsOriginHyperdriveHyperdriveOverAccessOrigin],
// [ConfigEditParamsOrigin].
type ConfigEditParamsOriginUnion interface {
	implementsHyperdriveConfigEditParamsOriginUnion()
}

type ConfigEditParamsOriginHyperdriveHyperdriveDatabase struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme] `json:"scheme"`
	// The user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r ConfigEditParamsOriginHyperdriveHyperdriveDatabase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOriginHyperdriveHyperdriveDatabase) implementsHyperdriveConfigEditParamsOriginUnion() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme string

const (
	ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgres   ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme = "postgres"
	ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgresql ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme = "postgresql"
)

func (r ConfigEditParamsOriginHyperdriveHyperdriveDatabaseScheme) IsKnown() bool {
	switch r {
	case ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgres, ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgresql:
		return true
	}
	return false
}

type ConfigEditParamsOriginHyperdriveHyperdriveInternetOrigin struct {
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port,required"`
}

func (r ConfigEditParamsOriginHyperdriveHyperdriveInternetOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOriginHyperdriveHyperdriveInternetOrigin) implementsHyperdriveConfigEditParamsOriginUnion() {
}

type ConfigEditParamsOriginHyperdriveHyperdriveOverAccessOrigin struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// The Client Secret of the Access token to use when connecting to the origin
	// database. This value is write-only and never returned by the API.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
}

func (r ConfigEditParamsOriginHyperdriveHyperdriveOverAccessOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigEditParamsOriginHyperdriveHyperdriveOverAccessOrigin) implementsHyperdriveConfigEditParamsOriginUnion() {
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigEditParamsOriginScheme string

const (
	ConfigEditParamsOriginSchemePostgres   ConfigEditParamsOriginScheme = "postgres"
	ConfigEditParamsOriginSchemePostgresql ConfigEditParamsOriginScheme = "postgresql"
)

func (r ConfigEditParamsOriginScheme) IsKnown() bool {
	switch r {
	case ConfigEditParamsOriginSchemePostgres, ConfigEditParamsOriginSchemePostgresql:
		return true
	}
	return false
}

type ConfigEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigEditResponse    `json:"result,required"`
	// Whether the API call was successful
	Success ConfigEditResponseEnvelopeSuccess `json:"success,required"`
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

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigGetResponse     `json:"result,required"`
	// Whether the API call was successful
	Success ConfigGetResponseEnvelopeSuccess `json:"success,required"`
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

// Whether the API call was successful
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
