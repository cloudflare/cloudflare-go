// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/tidwall/gjson"
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

type Hyperdrive struct {
	// Identifier
	ID      string            `json:"id,required"`
	Name    string            `json:"name,required"`
	Origin  HyperdriveOrigin  `json:"origin,required"`
	Caching HyperdriveCaching `json:"caching"`
	JSON    hyperdriveJSON    `json:"-"`
}

// hyperdriveJSON contains the JSON metadata for the struct [Hyperdrive]
type hyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Hyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveJSON) RawJSON() string {
	return r.raw
}

type HyperdriveOrigin struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveOriginScheme `json:"scheme,required"`
	// The user of your origin database.
	User string `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port  int64                `json:"port"`
	JSON  hyperdriveOriginJSON `json:"-"`
	union HyperdriveOriginUnion
}

// hyperdriveOriginJSON contains the JSON metadata for the struct
// [HyperdriveOrigin]
type hyperdriveOriginJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r hyperdriveOriginJSON) RawJSON() string {
	return r.raw
}

func (r *HyperdriveOrigin) UnmarshalJSON(data []byte) (err error) {
	*r = HyperdriveOrigin{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HyperdriveOriginUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [hyperdrive.HyperdriveOriginPublicDatabase],
// [hyperdrive.HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel].
func (r HyperdriveOrigin) AsUnion() HyperdriveOriginUnion {
	return r.union
}

// Union satisfied by [hyperdrive.HyperdriveOriginPublicDatabase] or
// [hyperdrive.HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel].
type HyperdriveOriginUnion interface {
	implementsHyperdriveHyperdriveOrigin()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HyperdriveOriginUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveOriginPublicDatabase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel{}),
		},
	)
}

type HyperdriveOriginPublicDatabase struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveOriginPublicDatabaseScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                             `json:"user,required"`
	JSON hyperdriveOriginPublicDatabaseJSON `json:"-"`
}

// hyperdriveOriginPublicDatabaseJSON contains the JSON metadata for the struct
// [HyperdriveOriginPublicDatabase]
type hyperdriveOriginPublicDatabaseJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveOriginPublicDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveOriginPublicDatabaseJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveOriginPublicDatabase) implementsHyperdriveHyperdriveOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginPublicDatabaseScheme string

const (
	HyperdriveOriginPublicDatabaseSchemePostgres   HyperdriveOriginPublicDatabaseScheme = "postgres"
	HyperdriveOriginPublicDatabaseSchemePostgresql HyperdriveOriginPublicDatabaseScheme = "postgresql"
)

func (r HyperdriveOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginPublicDatabaseSchemePostgres, HyperdriveOriginPublicDatabaseSchemePostgresql:
		return true
	}
	return false
}

type HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id,required"`
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                                                            `json:"user,required"`
	JSON hyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
}

// hyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON contains the
// JSON metadata for the struct
// [HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type hyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveHyperdriveOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
)

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginScheme string

const (
	HyperdriveOriginSchemePostgres   HyperdriveOriginScheme = "postgres"
	HyperdriveOriginSchemePostgresql HyperdriveOriginScheme = "postgresql"
)

func (r HyperdriveOriginScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginSchemePostgres, HyperdriveOriginSchemePostgresql:
		return true
	}
	return false
}

type HyperdriveCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                 `json:"stale_while_revalidate"`
	JSON                 hyperdriveCachingJSON `json:"-"`
	union                HyperdriveCachingUnion
}

// hyperdriveCachingJSON contains the JSON metadata for the struct
// [HyperdriveCaching]
type hyperdriveCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r hyperdriveCachingJSON) RawJSON() string {
	return r.raw
}

func (r *HyperdriveCaching) UnmarshalJSON(data []byte) (err error) {
	*r = HyperdriveCaching{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HyperdriveCachingUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingCommon],
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingEnabled].
func (r HyperdriveCaching) AsUnion() HyperdriveCachingUnion {
	return r.union
}

// Union satisfied by
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingCommon] or
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingEnabled].
type HyperdriveCachingUnion interface {
	implementsHyperdriveHyperdriveCaching()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HyperdriveCachingUnion)(nil)).Elem(),
		"disabled",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveCachingHyperdriveHyperdriveCachingCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveCachingHyperdriveHyperdriveCachingEnabled{}),
		},
	)
}

type HyperdriveCachingHyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool                                                   `json:"disabled"`
	JSON     hyperdriveCachingHyperdriveHyperdriveCachingCommonJSON `json:"-"`
}

// hyperdriveCachingHyperdriveHyperdriveCachingCommonJSON contains the JSON
// metadata for the struct [HyperdriveCachingHyperdriveHyperdriveCachingCommon]
type hyperdriveCachingHyperdriveHyperdriveCachingCommonJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveCachingHyperdriveHyperdriveCachingCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveCachingHyperdriveHyperdriveCachingCommonJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveHyperdriveCaching() {}

type HyperdriveCachingHyperdriveHyperdriveCachingEnabled struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                                                   `json:"stale_while_revalidate"`
	JSON                 hyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
}

// hyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON contains the JSON
// metadata for the struct [HyperdriveCachingHyperdriveHyperdriveCachingEnabled]
type hyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveCachingHyperdriveHyperdriveCachingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveHyperdriveCaching() {
}

type HyperdriveParam struct {
	Name    param.Field[string]                      `json:"name,required"`
	Origin  param.Field[HyperdriveOriginUnionParam]  `json:"origin,required"`
	Caching param.Field[HyperdriveCachingUnionParam] `json:"caching"`
}

func (r HyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveOriginParam struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveOriginScheme] `json:"scheme,required"`
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

func (r HyperdriveOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginParam) implementsHyperdriveHyperdriveOriginUnionParam() {}

// Satisfied by [hyperdrive.HyperdriveOriginPublicDatabaseParam],
// [hyperdrive.HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam],
// [HyperdriveOriginParam].
type HyperdriveOriginUnionParam interface {
	implementsHyperdriveHyperdriveOriginUnionParam()
}

type HyperdriveOriginPublicDatabaseParam struct {
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
	Scheme param.Field[HyperdriveOriginPublicDatabaseScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r HyperdriveOriginPublicDatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginPublicDatabaseParam) implementsHyperdriveHyperdriveOriginUnionParam() {}

type HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam struct {
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
	Scheme param.Field[HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) implementsHyperdriveHyperdriveOriginUnionParam() {
}

type HyperdriveCachingParam struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveCachingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingParam) implementsHyperdriveHyperdriveCachingUnionParam() {}

// Satisfied by
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingCommonParam],
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam],
// [HyperdriveCachingParam].
type HyperdriveCachingUnionParam interface {
	implementsHyperdriveHyperdriveCachingUnionParam()
}

type HyperdriveCachingHyperdriveHyperdriveCachingCommonParam struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingCommonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingCommonParam) implementsHyperdriveHyperdriveCachingUnionParam() {
}

type HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) implementsHyperdriveHyperdriveCachingUnionParam() {
}
