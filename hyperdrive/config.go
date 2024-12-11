// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
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
func (r *ConfigService) New(ctx context.Context, params ConfigNewParams, opts ...option.RequestOption) (res *Hyperdrive, err error) {
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
func (r *ConfigService) Update(ctx context.Context, hyperdriveID string, params ConfigUpdateParams, opts ...option.RequestOption) (res *Hyperdrive, err error) {
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
func (r *ConfigService) List(ctx context.Context, query ConfigListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Hyperdrive], err error) {
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
func (r *ConfigService) ListAutoPaging(ctx context.Context, query ConfigListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Hyperdrive] {
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
func (r *ConfigService) Edit(ctx context.Context, hyperdriveID string, params ConfigEditParams, opts ...option.RequestOption) (res *Hyperdrive, err error) {
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
func (r *ConfigService) Get(ctx context.Context, hyperdriveID string, query ConfigGetParams, opts ...option.RequestOption) (res *Hyperdrive, err error) {
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

type ConfigDeleteResponse = interface{}

type ConfigNewParams struct {
	// Identifier
	AccountID  param.Field[string] `path:"account_id,required"`
	Hyperdrive HyperdriveParam     `json:"hyperdrive,required"`
}

func (r ConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Hyperdrive)
}

type ConfigNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Hyperdrive            `json:"result,required"`
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
	AccountID  param.Field[string] `path:"account_id,required"`
	Hyperdrive HyperdriveParam     `json:"hyperdrive,required"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Hyperdrive)
}

type ConfigUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Hyperdrive            `json:"result,required"`
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
	Result   Hyperdrive            `json:"result,required"`
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
	Result   Hyperdrive            `json:"result,required"`
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
