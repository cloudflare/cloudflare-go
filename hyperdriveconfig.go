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

// Returns the specified Hyperdrive configuration.
func (r *HyperdriveConfigService) Get(ctx context.Context, accountIdentifier string, hyperdriveIdentifier string, opts ...option.RequestOption) (res *HyperdriveConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountIdentifier, hyperdriveIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates and returns the specified Hyperdrive configuration.
func (r *HyperdriveConfigService) Update(ctx context.Context, accountIdentifier string, hyperdriveIdentifier string, body HyperdriveConfigUpdateParams, opts ...option.RequestOption) (res *HyperdriveConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountIdentifier, hyperdriveIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes the specified Hyperdrive.
func (r *HyperdriveConfigService) Delete(ctx context.Context, accountIdentifier string, hyperdriveIdentifier string, opts ...option.RequestOption) (res *HyperdriveConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountIdentifier, hyperdriveIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type HyperdriveConfigGetResponse struct {
	Errors   []HyperdriveConfigGetResponseError   `json:"errors"`
	Messages []HyperdriveConfigGetResponseMessage `json:"messages"`
	Result   HyperdriveConfigGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HyperdriveConfigGetResponseSuccess `json:"success"`
	JSON    hyperdriveConfigGetResponseJSON    `json:"-"`
}

// hyperdriveConfigGetResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigGetResponse]
type hyperdriveConfigGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    hyperdriveConfigGetResponseErrorJSON `json:"-"`
}

// hyperdriveConfigGetResponseErrorJSON contains the JSON metadata for the struct
// [HyperdriveConfigGetResponseError]
type hyperdriveConfigGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    hyperdriveConfigGetResponseMessageJSON `json:"-"`
}

// hyperdriveConfigGetResponseMessageJSON contains the JSON metadata for the struct
// [HyperdriveConfigGetResponseMessage]
type hyperdriveConfigGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseResult struct {
	Name    string                                   `json:"name,required"`
	Origin  HyperdriveConfigGetResponseResultOrigin  `json:"origin,required"`
	Caching HyperdriveConfigGetResponseResultCaching `json:"caching"`
	JSON    hyperdriveConfigGetResponseResultJSON    `json:"-"`
}

// hyperdriveConfigGetResponseResultJSON contains the JSON metadata for the struct
// [HyperdriveConfigGetResponseResult]
type hyperdriveConfigGetResponseResultJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseResultOrigin struct {
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveConfigGetResponseResultOriginScheme `json:"scheme"`
	// The user of your origin database.
	User string                                      `json:"user"`
	JSON hyperdriveConfigGetResponseResultOriginJSON `json:"-"`
}

// hyperdriveConfigGetResponseResultOriginJSON contains the JSON metadata for the
// struct [HyperdriveConfigGetResponseResultOrigin]
type hyperdriveConfigGetResponseResultOriginJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveConfigGetResponseResultOriginScheme string

const (
	HyperdriveConfigGetResponseResultOriginSchemePostgres   HyperdriveConfigGetResponseResultOriginScheme = "postgres"
	HyperdriveConfigGetResponseResultOriginSchemePostgresql HyperdriveConfigGetResponseResultOriginScheme = "postgresql"
)

type HyperdriveConfigGetResponseResultCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                                        `json:"stale_while_revalidate"`
	JSON                 hyperdriveConfigGetResponseResultCachingJSON `json:"-"`
}

// hyperdriveConfigGetResponseResultCachingJSON contains the JSON metadata for the
// struct [HyperdriveConfigGetResponseResultCaching]
type hyperdriveConfigGetResponseResultCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseResultCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigGetResponseSuccess bool

const (
	HyperdriveConfigGetResponseSuccessTrue HyperdriveConfigGetResponseSuccess = true
)

type HyperdriveConfigUpdateResponse struct {
	Errors   []HyperdriveConfigUpdateResponseError   `json:"errors"`
	Messages []HyperdriveConfigUpdateResponseMessage `json:"messages"`
	Result   HyperdriveConfigUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HyperdriveConfigUpdateResponseSuccess `json:"success"`
	JSON    hyperdriveConfigUpdateResponseJSON    `json:"-"`
}

// hyperdriveConfigUpdateResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigUpdateResponse]
type hyperdriveConfigUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    hyperdriveConfigUpdateResponseErrorJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseErrorJSON contains the JSON metadata for the
// struct [HyperdriveConfigUpdateResponseError]
type hyperdriveConfigUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    hyperdriveConfigUpdateResponseMessageJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseMessageJSON contains the JSON metadata for the
// struct [HyperdriveConfigUpdateResponseMessage]
type hyperdriveConfigUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseResult struct {
	Name    string                                      `json:"name,required"`
	Origin  HyperdriveConfigUpdateResponseResultOrigin  `json:"origin,required"`
	Caching HyperdriveConfigUpdateResponseResultCaching `json:"caching"`
	JSON    hyperdriveConfigUpdateResponseResultJSON    `json:"-"`
}

// hyperdriveConfigUpdateResponseResultJSON contains the JSON metadata for the
// struct [HyperdriveConfigUpdateResponseResult]
type hyperdriveConfigUpdateResponseResultJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseResultOrigin struct {
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveConfigUpdateResponseResultOriginScheme `json:"scheme"`
	// The user of your origin database.
	User string                                         `json:"user"`
	JSON hyperdriveConfigUpdateResponseResultOriginJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseResultOriginJSON contains the JSON metadata for
// the struct [HyperdriveConfigUpdateResponseResultOrigin]
type hyperdriveConfigUpdateResponseResultOriginJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveConfigUpdateResponseResultOriginScheme string

const (
	HyperdriveConfigUpdateResponseResultOriginSchemePostgres   HyperdriveConfigUpdateResponseResultOriginScheme = "postgres"
	HyperdriveConfigUpdateResponseResultOriginSchemePostgresql HyperdriveConfigUpdateResponseResultOriginScheme = "postgresql"
)

type HyperdriveConfigUpdateResponseResultCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                                           `json:"stale_while_revalidate"`
	JSON                 hyperdriveConfigUpdateResponseResultCachingJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseResultCachingJSON contains the JSON metadata for
// the struct [HyperdriveConfigUpdateResponseResultCaching]
type hyperdriveConfigUpdateResponseResultCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseResultCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigUpdateResponseSuccess bool

const (
	HyperdriveConfigUpdateResponseSuccessTrue HyperdriveConfigUpdateResponseSuccess = true
)

type HyperdriveConfigDeleteResponse struct {
	Errors   []HyperdriveConfigDeleteResponseError   `json:"errors"`
	Messages []HyperdriveConfigDeleteResponseMessage `json:"messages"`
	Result   interface{}                             `json:"result,nullable"`
	// Whether the API call was successful
	Success HyperdriveConfigDeleteResponseSuccess `json:"success"`
	JSON    hyperdriveConfigDeleteResponseJSON    `json:"-"`
}

// hyperdriveConfigDeleteResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigDeleteResponse]
type hyperdriveConfigDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    hyperdriveConfigDeleteResponseErrorJSON `json:"-"`
}

// hyperdriveConfigDeleteResponseErrorJSON contains the JSON metadata for the
// struct [HyperdriveConfigDeleteResponseError]
type hyperdriveConfigDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    hyperdriveConfigDeleteResponseMessageJSON `json:"-"`
}

// hyperdriveConfigDeleteResponseMessageJSON contains the JSON metadata for the
// struct [HyperdriveConfigDeleteResponseMessage]
type hyperdriveConfigDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigDeleteResponseSuccess bool

const (
	HyperdriveConfigDeleteResponseSuccessTrue HyperdriveConfigDeleteResponseSuccess = true
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
