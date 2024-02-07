// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// HyperdriveConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHyperdriveConfigService] method
// instead.
type HyperdriveConfigService struct {
	Options []option.RequestOption
}

// NewHyperdriveConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHyperdriveConfigService(opts ...option.RequestOption) (r *HyperdriveConfigService) {
	r = &HyperdriveConfigService{}
	r.Options = opts
	return
}

// Creates and returns a new Hyperdrive configuration.
func (r *HyperdriveConfigService) New(ctx context.Context, accountIdentifier string, body HyperdriveConfigNewParams, opts ...option.RequestOption) (res *HyperdriveConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified Hyperdrive configuration.
func (r *HyperdriveConfigService) Get(ctx context.Context, accountIdentifier string, hyperdriveIdentifier string, opts ...option.RequestOption) (res *HyperdriveConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountIdentifier, hyperdriveIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates and returns the specified Hyperdrive configuration.
func (r *HyperdriveConfigService) Update(ctx context.Context, accountIdentifier string, hyperdriveIdentifier string, body HyperdriveConfigUpdateParams, opts ...option.RequestOption) (res *HyperdriveConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountIdentifier, hyperdriveIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of Hyperdrives
func (r *HyperdriveConfigService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]HyperdriveConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes the specified Hyperdrive.
func (r *HyperdriveConfigService) Delete(ctx context.Context, accountIdentifier string, hyperdriveIdentifier string, opts ...option.RequestOption) (res *HyperdriveConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountIdentifier, hyperdriveIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HyperdriveConfigNewResponse struct {
	Name    string                             `json:"name,required"`
	Origin  HyperdriveConfigNewResponseOrigin  `json:"origin,required"`
	Caching HyperdriveConfigNewResponseCaching `json:"caching"`
	JSON    hyperdriveConfigNewResponseJSON    `json:"-"`
}

// hyperdriveConfigNewResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigNewResponse]
type hyperdriveConfigNewResponseJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigNewResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveConfigNewResponseOriginScheme `json:"scheme"`
	// The user of your origin database.
	User string                                `json:"user"`
	JSON hyperdriveConfigNewResponseOriginJSON `json:"-"`
}

// hyperdriveConfigNewResponseOriginJSON contains the JSON metadata for the struct
// [HyperdriveConfigNewResponseOrigin]
type hyperdriveConfigNewResponseOriginJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveConfigNewResponseOriginScheme string

const (
	HyperdriveConfigNewResponseOriginSchemePostgres   HyperdriveConfigNewResponseOriginScheme = "postgres"
	HyperdriveConfigNewResponseOriginSchemePostgresql HyperdriveConfigNewResponseOriginScheme = "postgresql"
)

type HyperdriveConfigNewResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                                  `json:"stale_while_revalidate"`
	JSON                 hyperdriveConfigNewResponseCachingJSON `json:"-"`
}

// hyperdriveConfigNewResponseCachingJSON contains the JSON metadata for the struct
// [HyperdriveConfigNewResponseCaching]
type hyperdriveConfigNewResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponse struct {
	Name    string                             `json:"name,required"`
	Origin  HyperdriveConfigGetResponseOrigin  `json:"origin,required"`
	Caching HyperdriveConfigGetResponseCaching `json:"caching"`
	JSON    hyperdriveConfigGetResponseJSON    `json:"-"`
}

// hyperdriveConfigGetResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigGetResponse]
type hyperdriveConfigGetResponseJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveConfigGetResponseOriginScheme `json:"scheme"`
	// The user of your origin database.
	User string                                `json:"user"`
	JSON hyperdriveConfigGetResponseOriginJSON `json:"-"`
}

// hyperdriveConfigGetResponseOriginJSON contains the JSON metadata for the struct
// [HyperdriveConfigGetResponseOrigin]
type hyperdriveConfigGetResponseOriginJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveConfigGetResponseOriginScheme string

const (
	HyperdriveConfigGetResponseOriginSchemePostgres   HyperdriveConfigGetResponseOriginScheme = "postgres"
	HyperdriveConfigGetResponseOriginSchemePostgresql HyperdriveConfigGetResponseOriginScheme = "postgresql"
)

type HyperdriveConfigGetResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                                  `json:"stale_while_revalidate"`
	JSON                 hyperdriveConfigGetResponseCachingJSON `json:"-"`
}

// hyperdriveConfigGetResponseCachingJSON contains the JSON metadata for the struct
// [HyperdriveConfigGetResponseCaching]
type hyperdriveConfigGetResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponse struct {
	Name    string                                `json:"name,required"`
	Origin  HyperdriveConfigUpdateResponseOrigin  `json:"origin,required"`
	Caching HyperdriveConfigUpdateResponseCaching `json:"caching"`
	JSON    hyperdriveConfigUpdateResponseJSON    `json:"-"`
}

// hyperdriveConfigUpdateResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigUpdateResponse]
type hyperdriveConfigUpdateResponseJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveConfigUpdateResponseOriginScheme `json:"scheme"`
	// The user of your origin database.
	User string                                   `json:"user"`
	JSON hyperdriveConfigUpdateResponseOriginJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseOriginJSON contains the JSON metadata for the
// struct [HyperdriveConfigUpdateResponseOrigin]
type hyperdriveConfigUpdateResponseOriginJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveConfigUpdateResponseOriginScheme string

const (
	HyperdriveConfigUpdateResponseOriginSchemePostgres   HyperdriveConfigUpdateResponseOriginScheme = "postgres"
	HyperdriveConfigUpdateResponseOriginSchemePostgresql HyperdriveConfigUpdateResponseOriginScheme = "postgresql"
)

type HyperdriveConfigUpdateResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                                     `json:"stale_while_revalidate"`
	JSON                 hyperdriveConfigUpdateResponseCachingJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseCachingJSON contains the JSON metadata for the
// struct [HyperdriveConfigUpdateResponseCaching]
type hyperdriveConfigUpdateResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigListResponse struct {
	Name    string                              `json:"name,required"`
	Origin  HyperdriveConfigListResponseOrigin  `json:"origin,required"`
	Caching HyperdriveConfigListResponseCaching `json:"caching"`
	JSON    hyperdriveConfigListResponseJSON    `json:"-"`
}

// hyperdriveConfigListResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigListResponse]
type hyperdriveConfigListResponseJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigListResponseOrigin struct {
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveConfigListResponseOriginScheme `json:"scheme"`
	// The user of your origin database.
	User string                                 `json:"user"`
	JSON hyperdriveConfigListResponseOriginJSON `json:"-"`
}

// hyperdriveConfigListResponseOriginJSON contains the JSON metadata for the struct
// [HyperdriveConfigListResponseOrigin]
type hyperdriveConfigListResponseOriginJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveConfigListResponseOriginScheme string

const (
	HyperdriveConfigListResponseOriginSchemePostgres   HyperdriveConfigListResponseOriginScheme = "postgres"
	HyperdriveConfigListResponseOriginSchemePostgresql HyperdriveConfigListResponseOriginScheme = "postgresql"
)

type HyperdriveConfigListResponseCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                                   `json:"stale_while_revalidate"`
	JSON                 hyperdriveConfigListResponseCachingJSON `json:"-"`
}

// hyperdriveConfigListResponseCachingJSON contains the JSON metadata for the
// struct [HyperdriveConfigListResponseCaching]
type hyperdriveConfigListResponseCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveConfigListResponseCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigDeleteResponse = interface{}

type HyperdriveConfigNewParams struct {
	Name    param.Field[string]                           `json:"name,required"`
	Origin  param.Field[HyperdriveConfigNewParamsOrigin]  `json:"origin,required"`
	Caching param.Field[HyperdriveConfigNewParamsCaching] `json:"caching"`
}

func (r HyperdriveConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigNewParamsOrigin struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveConfigNewParamsOriginScheme] `json:"scheme"`
	// The user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r HyperdriveConfigNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveConfigNewParamsOriginScheme string

const (
	HyperdriveConfigNewParamsOriginSchemePostgres   HyperdriveConfigNewParamsOriginScheme = "postgres"
	HyperdriveConfigNewParamsOriginSchemePostgresql HyperdriveConfigNewParamsOriginScheme = "postgresql"
)

type HyperdriveConfigNewParamsCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveConfigNewParamsCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigNewResponseEnvelope struct {
	Errors   []HyperdriveConfigNewResponseEnvelopeErrors   `json:"errors"`
	Messages []HyperdriveConfigNewResponseEnvelopeMessages `json:"messages"`
	Result   HyperdriveConfigNewResponse                   `json:"result"`
	// Whether the API call was successful
	Success HyperdriveConfigNewResponseEnvelopeSuccess `json:"success"`
	JSON    hyperdriveConfigNewResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigNewResponseEnvelope]
type hyperdriveConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    hyperdriveConfigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HyperdriveConfigNewResponseEnvelopeErrors]
type hyperdriveConfigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    hyperdriveConfigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HyperdriveConfigNewResponseEnvelopeMessages]
type hyperdriveConfigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigNewResponseEnvelopeSuccess bool

const (
	HyperdriveConfigNewResponseEnvelopeSuccessTrue HyperdriveConfigNewResponseEnvelopeSuccess = true
)

type HyperdriveConfigGetResponseEnvelope struct {
	Errors   []HyperdriveConfigGetResponseEnvelopeErrors   `json:"errors"`
	Messages []HyperdriveConfigGetResponseEnvelopeMessages `json:"messages"`
	Result   HyperdriveConfigGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success HyperdriveConfigGetResponseEnvelopeSuccess `json:"success"`
	JSON    hyperdriveConfigGetResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigGetResponseEnvelope]
type hyperdriveConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    hyperdriveConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HyperdriveConfigGetResponseEnvelopeErrors]
type hyperdriveConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    hyperdriveConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HyperdriveConfigGetResponseEnvelopeMessages]
type hyperdriveConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigGetResponseEnvelopeSuccess bool

const (
	HyperdriveConfigGetResponseEnvelopeSuccessTrue HyperdriveConfigGetResponseEnvelopeSuccess = true
)

type HyperdriveConfigUpdateParams struct {
	Name    param.Field[string]                              `json:"name,required"`
	Origin  param.Field[HyperdriveConfigUpdateParamsOrigin]  `json:"origin,required"`
	Caching param.Field[HyperdriveConfigUpdateParamsCaching] `json:"caching"`
}

func (r HyperdriveConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigUpdateParamsOrigin struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveConfigUpdateParamsOriginScheme] `json:"scheme"`
	// The user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r HyperdriveConfigUpdateParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveConfigUpdateParamsOriginScheme string

const (
	HyperdriveConfigUpdateParamsOriginSchemePostgres   HyperdriveConfigUpdateParamsOriginScheme = "postgres"
	HyperdriveConfigUpdateParamsOriginSchemePostgresql HyperdriveConfigUpdateParamsOriginScheme = "postgresql"
)

type HyperdriveConfigUpdateParamsCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveConfigUpdateParamsCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigUpdateResponseEnvelope struct {
	Errors   []HyperdriveConfigUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []HyperdriveConfigUpdateResponseEnvelopeMessages `json:"messages"`
	Result   HyperdriveConfigUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success HyperdriveConfigUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    hyperdriveConfigUpdateResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigUpdateResponseEnvelope]
type hyperdriveConfigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hyperdriveConfigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HyperdriveConfigUpdateResponseEnvelopeErrors]
type hyperdriveConfigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    hyperdriveConfigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HyperdriveConfigUpdateResponseEnvelopeMessages]
type hyperdriveConfigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigUpdateResponseEnvelopeSuccess bool

const (
	HyperdriveConfigUpdateResponseEnvelopeSuccessTrue HyperdriveConfigUpdateResponseEnvelopeSuccess = true
)

type HyperdriveConfigListResponseEnvelope struct {
	Errors   []HyperdriveConfigListResponseEnvelopeErrors   `json:"errors"`
	Messages []HyperdriveConfigListResponseEnvelopeMessages `json:"messages"`
	Result   []HyperdriveConfigListResponse                 `json:"result"`
	// Whether the API call was successful
	Success HyperdriveConfigListResponseEnvelopeSuccess `json:"success"`
	JSON    hyperdriveConfigListResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigListResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigListResponseEnvelope]
type hyperdriveConfigListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    hyperdriveConfigListResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HyperdriveConfigListResponseEnvelopeErrors]
type hyperdriveConfigListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hyperdriveConfigListResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HyperdriveConfigListResponseEnvelopeMessages]
type hyperdriveConfigListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigListResponseEnvelopeSuccess bool

const (
	HyperdriveConfigListResponseEnvelopeSuccessTrue HyperdriveConfigListResponseEnvelopeSuccess = true
)

type HyperdriveConfigDeleteResponseEnvelope struct {
	Errors   []HyperdriveConfigDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []HyperdriveConfigDeleteResponseEnvelopeMessages `json:"messages"`
	Result   HyperdriveConfigDeleteResponse                   `json:"result,nullable"`
	// Whether the API call was successful
	Success HyperdriveConfigDeleteResponseEnvelopeSuccess `json:"success"`
	JSON    hyperdriveConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigDeleteResponseEnvelope]
type hyperdriveConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hyperdriveConfigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HyperdriveConfigDeleteResponseEnvelopeErrors]
type hyperdriveConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    hyperdriveConfigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HyperdriveConfigDeleteResponseEnvelopeMessages]
type hyperdriveConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigDeleteResponseEnvelopeSuccess bool

const (
	HyperdriveConfigDeleteResponseEnvelopeSuccessTrue HyperdriveConfigDeleteResponseEnvelopeSuccess = true
)
