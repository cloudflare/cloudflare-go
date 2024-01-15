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

// AccountAccessServiceTokenService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAccessServiceTokenService] method instead.
type AccountAccessServiceTokenService struct {
	Options   []option.RequestOption
	Refreshes *AccountAccessServiceTokenRefreshService
	Rotates   *AccountAccessServiceTokenRotateService
}

// NewAccountAccessServiceTokenService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessServiceTokenService(opts ...option.RequestOption) (r *AccountAccessServiceTokenService) {
	r = &AccountAccessServiceTokenService{}
	r.Options = opts
	r.Refreshes = NewAccountAccessServiceTokenRefreshService(opts...)
	r.Rotates = NewAccountAccessServiceTokenRotateService(opts...)
	return
}

// Updates a configured service token.
func (r *AccountAccessServiceTokenService) Update(ctx context.Context, identifier string, uuid string, body AccountAccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *AccountAccessServiceTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a service token.
func (r *AccountAccessServiceTokenService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessServiceTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *AccountAccessServiceTokenService) AccessServiceTokensNewAServiceToken(ctx context.Context, identifier string, body AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenParams, opts ...option.RequestOption) (res *AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all service tokens.
func (r *AccountAccessServiceTokenService) AccessServiceTokensListServiceTokens(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessServiceTokenUpdateResponse struct {
	Errors   []AccountAccessServiceTokenUpdateResponseError   `json:"errors"`
	Messages []AccountAccessServiceTokenUpdateResponseMessage `json:"messages"`
	Result   AccountAccessServiceTokenUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessServiceTokenUpdateResponseSuccess `json:"success"`
	JSON    accountAccessServiceTokenUpdateResponseJSON    `json:"-"`
}

// accountAccessServiceTokenUpdateResponseJSON contains the JSON metadata for the
// struct [AccountAccessServiceTokenUpdateResponse]
type accountAccessServiceTokenUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenUpdateResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountAccessServiceTokenUpdateResponseErrorJSON `json:"-"`
}

// accountAccessServiceTokenUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountAccessServiceTokenUpdateResponseError]
type accountAccessServiceTokenUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenUpdateResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountAccessServiceTokenUpdateResponseMessageJSON `json:"-"`
}

// accountAccessServiceTokenUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountAccessServiceTokenUpdateResponseMessage]
type accountAccessServiceTokenUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenUpdateResponseResult struct {
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
	Name      string                                            `json:"name"`
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	JSON      accountAccessServiceTokenUpdateResponseResultJSON `json:"-"`
}

// accountAccessServiceTokenUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessServiceTokenUpdateResponseResult]
type accountAccessServiceTokenUpdateResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessServiceTokenUpdateResponseSuccess bool

const (
	AccountAccessServiceTokenUpdateResponseSuccessTrue AccountAccessServiceTokenUpdateResponseSuccess = true
)

type AccountAccessServiceTokenDeleteResponse struct {
	Errors   []AccountAccessServiceTokenDeleteResponseError   `json:"errors"`
	Messages []AccountAccessServiceTokenDeleteResponseMessage `json:"messages"`
	Result   AccountAccessServiceTokenDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessServiceTokenDeleteResponseSuccess `json:"success"`
	JSON    accountAccessServiceTokenDeleteResponseJSON    `json:"-"`
}

// accountAccessServiceTokenDeleteResponseJSON contains the JSON metadata for the
// struct [AccountAccessServiceTokenDeleteResponse]
type accountAccessServiceTokenDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenDeleteResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountAccessServiceTokenDeleteResponseErrorJSON `json:"-"`
}

// accountAccessServiceTokenDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountAccessServiceTokenDeleteResponseError]
type accountAccessServiceTokenDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenDeleteResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountAccessServiceTokenDeleteResponseMessageJSON `json:"-"`
}

// accountAccessServiceTokenDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountAccessServiceTokenDeleteResponseMessage]
type accountAccessServiceTokenDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenDeleteResponseResult struct {
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
	Name      string                                            `json:"name"`
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	JSON      accountAccessServiceTokenDeleteResponseResultJSON `json:"-"`
}

// accountAccessServiceTokenDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessServiceTokenDeleteResponseResult]
type accountAccessServiceTokenDeleteResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessServiceTokenDeleteResponseSuccess bool

const (
	AccountAccessServiceTokenDeleteResponseSuccessTrue AccountAccessServiceTokenDeleteResponseSuccess = true
)

type AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponse struct {
	Errors   []AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseError   `json:"errors"`
	Messages []AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessage `json:"messages"`
	Result   AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseSuccess `json:"success"`
	JSON    accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON    `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponse]
type accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseErrorJSON `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseError]
type accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessageJSON `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessage]
type accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResult struct {
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
	Name      string                                                                         `json:"name"`
	UpdatedAt time.Time                                                                      `json:"updated_at" format:"date-time"`
	JSON      accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResultJSON `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResult]
type accountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResultJSON struct {
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

func (r *AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseSuccess bool

const (
	AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseSuccessTrue AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenResponseSuccess = true
)

type AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponse struct {
	Errors     []AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseError    `json:"errors"`
	Messages   []AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseMessage  `json:"messages"`
	Result     []AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResult   `json:"result"`
	ResultInfo AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseSuccess `json:"success"`
	JSON    accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseJSON    `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponse]
type accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseError struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseErrorJSON `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseError]
type accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseMessage struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseMessageJSON `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseMessage]
type accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResult struct {
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
	Name      string                                                                          `json:"name"`
	UpdatedAt time.Time                                                                       `json:"updated_at" format:"date-time"`
	JSON      accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultJSON `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResult]
type accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                             `json:"total_count"`
	JSON       accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfoJSON `json:"-"`
}

// accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfo]
type accountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseSuccess bool

const (
	AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseSuccessTrue AccountAccessServiceTokenAccessServiceTokensListServiceTokensResponseSuccess = true
)

type AccountAccessServiceTokenUpdateParams struct {
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r AccountAccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
}

func (r AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
