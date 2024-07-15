// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

type Configuration struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme ConfigurationScheme `json:"scheme,required"`
	// The user of your origin database.
	User string `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID string `json:"access_client_id"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64             `json:"port"`
	JSON configurationJSON `json:"-"`
}

// configurationJSON contains the JSON metadata for the struct [Configuration]
type configurationJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationJSON) RawJSON() string {
	return r.raw
}

// Specifies the URL scheme used to connect to your origin database.
type ConfigurationScheme string

const (
	ConfigurationSchemePostgres   ConfigurationScheme = "postgres"
	ConfigurationSchemePostgresql ConfigurationScheme = "postgresql"
	ConfigurationSchemeMysql      ConfigurationScheme = "mysql"
)

func (r ConfigurationScheme) IsKnown() bool {
	switch r {
	case ConfigurationSchemePostgres, ConfigurationSchemePostgresql, ConfigurationSchemeMysql:
		return true
	}
	return false
}

type ConfigurationParam struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[ConfigurationScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
	// The Client ID of the Access token to use when connecting to the origin database
	AccessClientID param.Field[string] `json:"access_client_id"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
}

func (r ConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Hyperdrive struct {
	Caching HyperdriveCaching `json:"caching"`
	Name    string            `json:"name"`
	Origin  Configuration     `json:"origin"`
	JSON    hyperdriveJSON    `json:"-"`
}

// hyperdriveJSON contains the JSON metadata for the struct [Hyperdrive]
type hyperdriveJSON struct {
	Caching     apijson.Field
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Hyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveJSON) RawJSON() string {
	return r.raw
}

type HyperdriveCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                 `json:"stale_while_revalidate"`
	JSON                 hyperdriveCachingJSON `json:"-"`
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

func (r *HyperdriveCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveCachingJSON) RawJSON() string {
	return r.raw
}

type HyperdriveParam struct {
	Caching param.Field[HyperdriveCachingParam] `json:"caching"`
	Name    param.Field[string]                 `json:"name"`
	Origin  param.Field[ConfigurationParam]     `json:"origin"`
}

func (r HyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveCachingParam struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveCachingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
