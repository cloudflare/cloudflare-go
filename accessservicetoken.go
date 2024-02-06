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

// AccessServiceTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessServiceTokenService] method
// instead.
type AccessServiceTokenService struct {
	Options   []option.RequestOption
	Refreshes *AccessServiceTokenRefreshService
	Rotates   *AccessServiceTokenRotateService
}

// NewAccessServiceTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessServiceTokenService(opts ...option.RequestOption) (r *AccessServiceTokenService) {
	r = &AccessServiceTokenService{}
	r.Options = opts
	r.Refreshes = NewAccessServiceTokenRefreshService(opts...)
	r.Rotates = NewAccessServiceTokenRotateService(opts...)
	return
}

// Updates a configured service token.
func (r *AccessServiceTokenService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *AccessServiceTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a service token.
func (r *AccessServiceTokenService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessServiceTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *AccessServiceTokenService) AccessServiceTokensNewAServiceToken(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessServiceTokenAccessServiceTokensNewAServiceTokenParams, opts ...option.RequestOption) (res *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all service tokens.
func (r *AccessServiceTokenService) AccessServiceTokensListServiceTokens(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessServiceTokenAccessServiceTokensListServiceTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessServiceTokenUpdateResponse struct {
	Errors   []AccessServiceTokenUpdateResponseError   `json:"errors"`
	Messages []AccessServiceTokenUpdateResponseMessage `json:"messages"`
	Result   AccessServiceTokenUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessServiceTokenUpdateResponseSuccess `json:"success"`
	JSON    accessServiceTokenUpdateResponseJSON    `json:"-"`
}

// accessServiceTokenUpdateResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenUpdateResponse]
type accessServiceTokenUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessServiceTokenUpdateResponseErrorJSON `json:"-"`
}

// accessServiceTokenUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccessServiceTokenUpdateResponseError]
type accessServiceTokenUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessServiceTokenUpdateResponseMessageJSON `json:"-"`
}

// accessServiceTokenUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccessServiceTokenUpdateResponseMessage]
type accessServiceTokenUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenUpdateResponseResult struct {
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
	Name      string                                     `json:"name"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenUpdateResponseResultJSON `json:"-"`
}

// accessServiceTokenUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccessServiceTokenUpdateResponseResult]
type accessServiceTokenUpdateResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenUpdateResponseSuccess bool

const (
	AccessServiceTokenUpdateResponseSuccessTrue AccessServiceTokenUpdateResponseSuccess = true
)

type AccessServiceTokenDeleteResponse struct {
	Errors   []AccessServiceTokenDeleteResponseError   `json:"errors"`
	Messages []AccessServiceTokenDeleteResponseMessage `json:"messages"`
	Result   AccessServiceTokenDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessServiceTokenDeleteResponseSuccess `json:"success"`
	JSON    accessServiceTokenDeleteResponseJSON    `json:"-"`
}

// accessServiceTokenDeleteResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenDeleteResponse]
type accessServiceTokenDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessServiceTokenDeleteResponseErrorJSON `json:"-"`
}

// accessServiceTokenDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccessServiceTokenDeleteResponseError]
type accessServiceTokenDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessServiceTokenDeleteResponseMessageJSON `json:"-"`
}

// accessServiceTokenDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccessServiceTokenDeleteResponseMessage]
type accessServiceTokenDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponseResult struct {
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
	Name      string                                     `json:"name"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenDeleteResponseResultJSON `json:"-"`
}

// accessServiceTokenDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccessServiceTokenDeleteResponseResult]
type accessServiceTokenDeleteResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenDeleteResponseSuccess bool

const (
	AccessServiceTokenDeleteResponseSuccessTrue AccessServiceTokenDeleteResponseSuccess = true
)

type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse struct {
	Errors   []AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseError   `json:"errors"`
	Messages []AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessage `json:"messages"`
	Result   AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseSuccess `json:"success"`
	JSON    accessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON    `json:"-"`
}

// accessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON contains the
// JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse]
type accessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseError struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    accessServiceTokenAccessServiceTokensNewAServiceTokenResponseErrorJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensNewAServiceTokenResponseErrorJSON contains
// the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseError]
type accessServiceTokenAccessServiceTokensNewAServiceTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessage struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessageJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessage]
type accessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResult struct {
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
	Name      string                                                                  `json:"name"`
	UpdatedAt time.Time                                                               `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenAccessServiceTokensNewAServiceTokenResponseResultJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensNewAServiceTokenResponseResultJSON contains
// the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResult]
type accessServiceTokenAccessServiceTokensNewAServiceTokenResponseResultJSON struct {
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

func (r *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseSuccess bool

const (
	AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseSuccessTrue AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseSuccess = true
)

type AccessServiceTokenAccessServiceTokensListServiceTokensResponse struct {
	Errors     []AccessServiceTokenAccessServiceTokensListServiceTokensResponseError    `json:"errors"`
	Messages   []AccessServiceTokenAccessServiceTokensListServiceTokensResponseMessage  `json:"messages"`
	Result     []AccessServiceTokenAccessServiceTokensListServiceTokensResponseResult   `json:"result"`
	ResultInfo AccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessServiceTokenAccessServiceTokensListServiceTokensResponseSuccess `json:"success"`
	JSON    accessServiceTokenAccessServiceTokensListServiceTokensResponseJSON    `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseJSON contains the
// JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponse]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensListServiceTokensResponseError struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accessServiceTokenAccessServiceTokensListServiceTokensResponseErrorJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseErrorJSON contains
// the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponseError]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensListServiceTokensResponseMessage struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accessServiceTokenAccessServiceTokensListServiceTokensResponseMessageJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponseMessage]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensListServiceTokensResponseResult struct {
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
	Name      string                                                                   `json:"name"`
	UpdatedAt time.Time                                                                `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenAccessServiceTokensListServiceTokensResponseResultJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseResultJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponseResult]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                      `json:"total_count"`
	JSON       accessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfoJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfo]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenAccessServiceTokensListServiceTokensResponseSuccess bool

const (
	AccessServiceTokenAccessServiceTokensListServiceTokensResponseSuccessTrue AccessServiceTokenAccessServiceTokensListServiceTokensResponseSuccess = true
)

type AccessServiceTokenUpdateParams struct {
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r AccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessServiceTokenAccessServiceTokensNewAServiceTokenParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
}

func (r AccessServiceTokenAccessServiceTokensNewAServiceTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
