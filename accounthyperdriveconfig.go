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

// AccountHyperdriveConfigService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountHyperdriveConfigService] method instead.
type AccountHyperdriveConfigService struct {
	Options []option.RequestOption
}

// NewAccountHyperdriveConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountHyperdriveConfigService(opts ...option.RequestOption) (r *AccountHyperdriveConfigService) {
	r = &AccountHyperdriveConfigService{}
	r.Options = opts
	return
}

// Creates and returns a new Hyperdrive configuration.
func (r *AccountHyperdriveConfigService) New(ctx context.Context, accountIdentifier string, body AccountHyperdriveConfigNewParams, opts ...option.RequestOption) (res *AccountHyperdriveConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of Hyperdrives
func (r *AccountHyperdriveConfigService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountHyperdriveConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountHyperdriveConfigNewResponse struct {
	Errors   []AccountHyperdriveConfigNewResponseError   `json:"errors"`
	Messages []AccountHyperdriveConfigNewResponseMessage `json:"messages"`
	Result   AccountHyperdriveConfigNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountHyperdriveConfigNewResponseSuccess `json:"success"`
	JSON    accountHyperdriveConfigNewResponseJSON    `json:"-"`
}

// accountHyperdriveConfigNewResponseJSON contains the JSON metadata for the struct
// [AccountHyperdriveConfigNewResponse]
type accountHyperdriveConfigNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHyperdriveConfigNewResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountHyperdriveConfigNewResponseErrorJSON `json:"-"`
}

// accountHyperdriveConfigNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigNewResponseError]
type accountHyperdriveConfigNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHyperdriveConfigNewResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountHyperdriveConfigNewResponseMessageJSON `json:"-"`
}

// accountHyperdriveConfigNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigNewResponseMessage]
type accountHyperdriveConfigNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHyperdriveConfigNewResponseResult struct {
	Name    string                                          `json:"name,required"`
	Origin  AccountHyperdriveConfigNewResponseResultOrigin  `json:"origin,required"`
	Caching AccountHyperdriveConfigNewResponseResultCaching `json:"caching"`
	JSON    accountHyperdriveConfigNewResponseResultJSON    `json:"-"`
}

// accountHyperdriveConfigNewResponseResultJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigNewResponseResult]
type accountHyperdriveConfigNewResponseResultJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHyperdriveConfigNewResponseResultOrigin struct {
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme AccountHyperdriveConfigNewResponseResultOriginScheme `json:"scheme"`
	// The user of your origin database.
	User string                                             `json:"user"`
	JSON accountHyperdriveConfigNewResponseResultOriginJSON `json:"-"`
}

// accountHyperdriveConfigNewResponseResultOriginJSON contains the JSON metadata
// for the struct [AccountHyperdriveConfigNewResponseResultOrigin]
type accountHyperdriveConfigNewResponseResultOriginJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigNewResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the URL scheme used to connect to your origin database.
type AccountHyperdriveConfigNewResponseResultOriginScheme string

const (
	AccountHyperdriveConfigNewResponseResultOriginSchemePostgres   AccountHyperdriveConfigNewResponseResultOriginScheme = "postgres"
	AccountHyperdriveConfigNewResponseResultOriginSchemePostgresql AccountHyperdriveConfigNewResponseResultOriginScheme = "postgresql"
)

type AccountHyperdriveConfigNewResponseResultCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                                               `json:"stale_while_revalidate"`
	JSON                 accountHyperdriveConfigNewResponseResultCachingJSON `json:"-"`
}

// accountHyperdriveConfigNewResponseResultCachingJSON contains the JSON metadata
// for the struct [AccountHyperdriveConfigNewResponseResultCaching]
type accountHyperdriveConfigNewResponseResultCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountHyperdriveConfigNewResponseResultCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountHyperdriveConfigNewResponseSuccess bool

const (
	AccountHyperdriveConfigNewResponseSuccessTrue AccountHyperdriveConfigNewResponseSuccess = true
)

type AccountHyperdriveConfigListResponse struct {
	Errors   []AccountHyperdriveConfigListResponseError   `json:"errors"`
	Messages []AccountHyperdriveConfigListResponseMessage `json:"messages"`
	Result   []AccountHyperdriveConfigListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountHyperdriveConfigListResponseSuccess `json:"success"`
	JSON    accountHyperdriveConfigListResponseJSON    `json:"-"`
}

// accountHyperdriveConfigListResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigListResponse]
type accountHyperdriveConfigListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHyperdriveConfigListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountHyperdriveConfigListResponseErrorJSON `json:"-"`
}

// accountHyperdriveConfigListResponseErrorJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigListResponseError]
type accountHyperdriveConfigListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHyperdriveConfigListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountHyperdriveConfigListResponseMessageJSON `json:"-"`
}

// accountHyperdriveConfigListResponseMessageJSON contains the JSON metadata for
// the struct [AccountHyperdriveConfigListResponseMessage]
type accountHyperdriveConfigListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHyperdriveConfigListResponseResult struct {
	Name    string                                           `json:"name,required"`
	Origin  AccountHyperdriveConfigListResponseResultOrigin  `json:"origin,required"`
	Caching AccountHyperdriveConfigListResponseResultCaching `json:"caching"`
	JSON    accountHyperdriveConfigListResponseResultJSON    `json:"-"`
}

// accountHyperdriveConfigListResponseResultJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigListResponseResult]
type accountHyperdriveConfigListResponseResultJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHyperdriveConfigListResponseResultOrigin struct {
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme AccountHyperdriveConfigListResponseResultOriginScheme `json:"scheme"`
	// The user of your origin database.
	User string                                              `json:"user"`
	JSON accountHyperdriveConfigListResponseResultOriginJSON `json:"-"`
}

// accountHyperdriveConfigListResponseResultOriginJSON contains the JSON metadata
// for the struct [AccountHyperdriveConfigListResponseResultOrigin]
type accountHyperdriveConfigListResponseResultOriginJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigListResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the URL scheme used to connect to your origin database.
type AccountHyperdriveConfigListResponseResultOriginScheme string

const (
	AccountHyperdriveConfigListResponseResultOriginSchemePostgres   AccountHyperdriveConfigListResponseResultOriginScheme = "postgres"
	AccountHyperdriveConfigListResponseResultOriginSchemePostgresql AccountHyperdriveConfigListResponseResultOriginScheme = "postgresql"
)

type AccountHyperdriveConfigListResponseResultCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate int64                                                `json:"stale_while_revalidate"`
	JSON                 accountHyperdriveConfigListResponseResultCachingJSON `json:"-"`
}

// accountHyperdriveConfigListResponseResultCachingJSON contains the JSON metadata
// for the struct [AccountHyperdriveConfigListResponseResultCaching]
type accountHyperdriveConfigListResponseResultCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountHyperdriveConfigListResponseResultCaching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountHyperdriveConfigListResponseSuccess bool

const (
	AccountHyperdriveConfigListResponseSuccessTrue AccountHyperdriveConfigListResponseSuccess = true
)

type AccountHyperdriveConfigNewParams struct {
	Name    param.Field[string]                                  `json:"name,required"`
	Origin  param.Field[AccountHyperdriveConfigNewParamsOrigin]  `json:"origin,required"`
	Caching param.Field[AccountHyperdriveConfigNewParamsCaching] `json:"caching"`
}

func (r AccountHyperdriveConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHyperdriveConfigNewParamsOrigin struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[AccountHyperdriveConfigNewParamsOriginScheme] `json:"scheme"`
	// The user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r AccountHyperdriveConfigNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the URL scheme used to connect to your origin database.
type AccountHyperdriveConfigNewParamsOriginScheme string

const (
	AccountHyperdriveConfigNewParamsOriginSchemePostgres   AccountHyperdriveConfigNewParamsOriginScheme = "postgres"
	AccountHyperdriveConfigNewParamsOriginSchemePostgresql AccountHyperdriveConfigNewParamsOriginScheme = "postgresql"
)

type AccountHyperdriveConfigNewParamsCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled param.Field[bool] `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r AccountHyperdriveConfigNewParamsCaching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
