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
func (r *AccountAccessServiceTokenService) Update(ctx context.Context, identifier string, uuid string, body AccountAccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *ServiceTokensSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a service token.
func (r *AccountAccessServiceTokenService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ServiceTokensSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *AccountAccessServiceTokenService) AccessServiceTokensNewAServiceToken(ctx context.Context, identifier string, body AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenParams, opts ...option.RequestOption) (res *CreateResponse, err error) {
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

type CreateResponse struct {
	Errors   []CreateResponseError   `json:"errors"`
	Messages []CreateResponseMessage `json:"messages"`
	Result   CreateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success CreateResponseSuccess `json:"success"`
	JSON    createResponseJSON    `json:"-"`
}

// createResponseJSON contains the JSON metadata for the struct [CreateResponse]
type createResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreateResponseError struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    createResponseErrorJSON `json:"-"`
}

// createResponseErrorJSON contains the JSON metadata for the struct
// [CreateResponseError]
type createResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreateResponseMessage struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    createResponseMessageJSON `json:"-"`
}

// createResponseMessageJSON contains the JSON metadata for the struct
// [CreateResponseMessage]
type createResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreateResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The name of the service token.
	Name      string                   `json:"name"`
	UpdatedAt time.Time                `json:"updated_at" format:"date-time"`
	JSON      createResponseResultJSON `json:"-"`
}

// createResponseResultJSON contains the JSON metadata for the struct
// [CreateResponseResult]
type createResponseResultJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CreateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CreateResponseSuccess bool

const (
	CreateResponseSuccessTrue CreateResponseSuccess = true
)

type ServiceTokensSingleResponse struct {
	Errors   []ServiceTokensSingleResponseError   `json:"errors"`
	Messages []ServiceTokensSingleResponseMessage `json:"messages"`
	Result   ServiceTokensSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ServiceTokensSingleResponseSuccess `json:"success"`
	JSON    serviceTokensSingleResponseJSON    `json:"-"`
}

// serviceTokensSingleResponseJSON contains the JSON metadata for the struct
// [ServiceTokensSingleResponse]
type serviceTokensSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceTokensSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceTokensSingleResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    serviceTokensSingleResponseErrorJSON `json:"-"`
}

// serviceTokensSingleResponseErrorJSON contains the JSON metadata for the struct
// [ServiceTokensSingleResponseError]
type serviceTokensSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceTokensSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceTokensSingleResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    serviceTokensSingleResponseMessageJSON `json:"-"`
}

// serviceTokensSingleResponseMessageJSON contains the JSON metadata for the struct
// [ServiceTokensSingleResponseMessage]
type serviceTokensSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceTokensSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceTokensSingleResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The name of the service token.
	Name      string                                `json:"name"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      serviceTokensSingleResponseResultJSON `json:"-"`
}

// serviceTokensSingleResponseResultJSON contains the JSON metadata for the struct
// [ServiceTokensSingleResponseResult]
type serviceTokensSingleResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceTokensSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ServiceTokensSingleResponseSuccess bool

const (
	ServiceTokensSingleResponseSuccessTrue ServiceTokensSingleResponseSuccess = true
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
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r AccountAccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessServiceTokenAccessServiceTokensNewAServiceTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
