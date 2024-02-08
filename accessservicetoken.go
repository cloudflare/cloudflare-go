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
	var env AccessServiceTokenUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a service token.
func (r *AccessServiceTokenService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessServiceTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *AccessServiceTokenService) AccessServiceTokensNewAServiceToken(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessServiceTokenAccessServiceTokensNewAServiceTokenParams, opts ...option.RequestOption) (res *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all service tokens.
func (r *AccessServiceTokenService) AccessServiceTokensListServiceTokens(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessServiceTokenAccessServiceTokensListServiceTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessServiceTokenUpdateResponse struct {
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
	Name      string                               `json:"name"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenUpdateResponseJSON `json:"-"`
}

// accessServiceTokenUpdateResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenUpdateResponse]
type accessServiceTokenUpdateResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponse struct {
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
	Name      string                               `json:"name"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenDeleteResponseJSON `json:"-"`
}

// accessServiceTokenDeleteResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenDeleteResponse]
type accessServiceTokenDeleteResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse struct {
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
	Name      string                                                            `json:"name"`
	UpdatedAt time.Time                                                         `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON contains the
// JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse]
type accessServiceTokenAccessServiceTokensNewAServiceTokenResponseJSON struct {
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

func (r *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensListServiceTokensResponse struct {
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
	Name      string                                                             `json:"name"`
	UpdatedAt time.Time                                                          `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenAccessServiceTokensListServiceTokensResponseJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseJSON contains the
// JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponse]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type AccessServiceTokenUpdateResponseEnvelope struct {
	Errors   []AccessServiceTokenUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenUpdateResponseEnvelope]
type accessServiceTokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessServiceTokenUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenUpdateResponseEnvelopeErrors]
type accessServiceTokenUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessServiceTokenUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenUpdateResponseEnvelopeMessages]
type accessServiceTokenUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenUpdateResponseEnvelopeSuccess bool

const (
	AccessServiceTokenUpdateResponseEnvelopeSuccessTrue AccessServiceTokenUpdateResponseEnvelopeSuccess = true
)

type AccessServiceTokenDeleteResponseEnvelope struct {
	Errors   []AccessServiceTokenDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenDeleteResponseEnvelope]
type accessServiceTokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessServiceTokenDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenDeleteResponseEnvelopeErrors]
type accessServiceTokenDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessServiceTokenDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenDeleteResponseEnvelopeMessages]
type accessServiceTokenDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenDeleteResponseEnvelopeSuccess bool

const (
	AccessServiceTokenDeleteResponseEnvelopeSuccessTrue AccessServiceTokenDeleteResponseEnvelopeSuccess = true
)

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

type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelope struct {
	Errors   []AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenAccessServiceTokensNewAServiceTokenResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelope]
type accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeErrors]
type accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeMessages]
type accessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeSuccess bool

const (
	AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeSuccessTrue AccessServiceTokenAccessServiceTokensNewAServiceTokenResponseEnvelopeSuccess = true
)

type AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelope struct {
	Errors   []AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessServiceTokenAccessServiceTokensListServiceTokensResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeJSON       `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelope]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeErrors struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeErrors]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeMessages struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeMessages]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeSuccess bool

const (
	AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeSuccessTrue AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeSuccess = true
)

type AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                              `json:"total_count"`
	JSON       accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeResultInfo]
type accessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenAccessServiceTokensListServiceTokensResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
