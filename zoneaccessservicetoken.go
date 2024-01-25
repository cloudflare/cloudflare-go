// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAccessServiceTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAccessServiceTokenService]
// method instead.
type ZoneAccessServiceTokenService struct {
	Options []option.RequestOption
}

// NewZoneAccessServiceTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAccessServiceTokenService(opts ...option.RequestOption) (r *ZoneAccessServiceTokenService) {
	r = &ZoneAccessServiceTokenService{}
	r.Options = opts
	return
}

// Updates a configured service token.
func (r *ZoneAccessServiceTokenService) Update(ctx context.Context, identifier string, uuid string, body ZoneAccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *ZoneAccessServiceTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/service_tokens/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a service token.
func (r *ZoneAccessServiceTokenService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZoneAccessServiceTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/service_tokens/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to create a new
// service token.
func (r *ZoneAccessServiceTokenService) ZoneLevelAccessServiceTokensNewAServiceToken(ctx context.Context, identifier string, body ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenParams, opts ...option.RequestOption) (res *ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/service_tokens", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all service tokens.
func (r *ZoneAccessServiceTokenService) ZoneLevelAccessServiceTokensListServiceTokens(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/service_tokens", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAccessServiceTokenUpdateResponse struct {
	Errors   []ZoneAccessServiceTokenUpdateResponseError   `json:"errors"`
	Messages []ZoneAccessServiceTokenUpdateResponseMessage `json:"messages"`
	Result   ZoneAccessServiceTokenUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessServiceTokenUpdateResponseSuccess `json:"success"`
	JSON    zoneAccessServiceTokenUpdateResponseJSON    `json:"-"`
}

// zoneAccessServiceTokenUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneAccessServiceTokenUpdateResponse]
type zoneAccessServiceTokenUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneAccessServiceTokenUpdateResponseErrorJSON `json:"-"`
}

// zoneAccessServiceTokenUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAccessServiceTokenUpdateResponseError]
type zoneAccessServiceTokenUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneAccessServiceTokenUpdateResponseMessageJSON `json:"-"`
}

// zoneAccessServiceTokenUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAccessServiceTokenUpdateResponseMessage]
type zoneAccessServiceTokenUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenUpdateResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                         `json:"name"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      zoneAccessServiceTokenUpdateResponseResultJSON `json:"-"`
}

// zoneAccessServiceTokenUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneAccessServiceTokenUpdateResponseResult]
type zoneAccessServiceTokenUpdateResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessServiceTokenUpdateResponseSuccess bool

const (
	ZoneAccessServiceTokenUpdateResponseSuccessTrue ZoneAccessServiceTokenUpdateResponseSuccess = true
)

type ZoneAccessServiceTokenDeleteResponse struct {
	Errors   []ZoneAccessServiceTokenDeleteResponseError   `json:"errors"`
	Messages []ZoneAccessServiceTokenDeleteResponseMessage `json:"messages"`
	Result   ZoneAccessServiceTokenDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessServiceTokenDeleteResponseSuccess `json:"success"`
	JSON    zoneAccessServiceTokenDeleteResponseJSON    `json:"-"`
}

// zoneAccessServiceTokenDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneAccessServiceTokenDeleteResponse]
type zoneAccessServiceTokenDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneAccessServiceTokenDeleteResponseErrorJSON `json:"-"`
}

// zoneAccessServiceTokenDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAccessServiceTokenDeleteResponseError]
type zoneAccessServiceTokenDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneAccessServiceTokenDeleteResponseMessageJSON `json:"-"`
}

// zoneAccessServiceTokenDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAccessServiceTokenDeleteResponseMessage]
type zoneAccessServiceTokenDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenDeleteResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                         `json:"name"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      zoneAccessServiceTokenDeleteResponseResultJSON `json:"-"`
}

// zoneAccessServiceTokenDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneAccessServiceTokenDeleteResponseResult]
type zoneAccessServiceTokenDeleteResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessServiceTokenDeleteResponseSuccess bool

const (
	ZoneAccessServiceTokenDeleteResponseSuccessTrue ZoneAccessServiceTokenDeleteResponseSuccess = true
)

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponse struct {
	Errors   []ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseError   `json:"errors"`
	Messages []ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseMessage `json:"messages"`
	Result   ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseSuccess `json:"success"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseJSON    `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponse]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseErrorJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseError]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseMessageJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseMessage]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                                                               `json:"name"`
	UpdatedAt time.Time                                                                            `json:"updated_at" format:"date-time"`
	JSON      zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseResultJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseResult]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseResultJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseSuccess bool

const (
	ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseSuccessTrue ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenResponseSuccess = true
)

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponse struct {
	Errors     []ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseError    `json:"errors"`
	Messages   []ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessage  `json:"messages"`
	Result     []ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResult   `json:"result"`
	ResultInfo ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseSuccess `json:"success"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseJSON    `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponse]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseErrorJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseError]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessageJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessage]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                                                                `json:"name"`
	UpdatedAt time.Time                                                                             `json:"updated_at" format:"date-time"`
	JSON      zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResult]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                   `json:"total_count"`
	JSON       zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfoJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfo]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseSuccess bool

const (
	ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseSuccessTrue ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseSuccess = true
)

type ZoneAccessServiceTokenUpdateParams struct {
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r ZoneAccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
}

func (r ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
