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

// SecondaryDNSTsigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSTsigService] method
// instead.
type SecondaryDNSTsigService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSTsigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSTsigService(opts ...option.RequestOption) (r *SecondaryDNSTsigService) {
	r = &SecondaryDNSTsigService{}
	r.Options = opts
	return
}

// Modify TSIG.
func (r *SecondaryDNSTsigService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body SecondaryDNSTsigUpdateParams, opts ...option.RequestOption) (res *SecondaryDNSTsigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete TSIG.
func (r *SecondaryDNSTsigService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SecondaryDNSTsigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get TSIG.
func (r *SecondaryDNSTsigService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SecondaryDNSTsigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create TSIG.
func (r *SecondaryDNSTsigService) SecondaryDNSTsigNewTsig(ctx context.Context, accountIdentifier interface{}, body SecondaryDNSTsigSecondaryDNSTsigNewTsigParams, opts ...option.RequestOption) (res *SecondaryDNSTsigSecondaryDNSTsigNewTsigResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List TSIGs.
func (r *SecondaryDNSTsigService) SecondaryDNSTsigListTsiGs(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *[]SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSTsigUpdateResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                             `json:"secret,required"`
	JSON   secondaryDNSTsigUpdateResponseJSON `json:"-"`
}

// secondaryDNSTsigUpdateResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSTsigUpdateResponse]
type secondaryDNSTsigUpdateResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigDeleteResponse struct {
	ID   interface{}                        `json:"id"`
	JSON secondaryDNSTsigDeleteResponseJSON `json:"-"`
}

// secondaryDNSTsigDeleteResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSTsigDeleteResponse]
type secondaryDNSTsigDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigGetResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                          `json:"secret,required"`
	JSON   secondaryDNSTsigGetResponseJSON `json:"-"`
}

// secondaryDNSTsigGetResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSTsigGetResponse]
type secondaryDNSTsigGetResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigSecondaryDNSTsigNewTsigResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                                              `json:"secret,required"`
	JSON   secondaryDNSTsigSecondaryDNSTsigNewTsigResponseJSON `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigNewTsigResponseJSON contains the JSON metadata
// for the struct [SecondaryDNSTsigSecondaryDNSTsigNewTsigResponse]
type secondaryDNSTsigSecondaryDNSTsigNewTsigResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigNewTsigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                                                `json:"secret,required"`
	JSON   secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseJSON `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseJSON contains the JSON metadata
// for the struct [SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse]
type secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigUpdateParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r SecondaryDNSTsigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSTsigUpdateResponseEnvelope struct {
	Errors   []SecondaryDNSTsigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSTsigUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTsigUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSTsigUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSTsigUpdateResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSTsigUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSTsigUpdateResponseEnvelope]
type secondaryDNSTsigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDNSTsigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSTsigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSTsigUpdateResponseEnvelopeErrors]
type secondaryDNSTsigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    secondaryDNSTsigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSTsigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSTsigUpdateResponseEnvelopeMessages]
type secondaryDNSTsigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSTsigUpdateResponseEnvelopeSuccess bool

const (
	SecondaryDNSTsigUpdateResponseEnvelopeSuccessTrue SecondaryDNSTsigUpdateResponseEnvelopeSuccess = true
)

type SecondaryDNSTsigDeleteResponseEnvelope struct {
	Errors   []SecondaryDNSTsigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSTsigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTsigDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSTsigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSTsigDeleteResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSTsigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSTsigDeleteResponseEnvelope]
type secondaryDNSTsigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDNSTsigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSTsigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSTsigDeleteResponseEnvelopeErrors]
type secondaryDNSTsigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    secondaryDNSTsigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSTsigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSTsigDeleteResponseEnvelopeMessages]
type secondaryDNSTsigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSTsigDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDNSTsigDeleteResponseEnvelopeSuccessTrue SecondaryDNSTsigDeleteResponseEnvelopeSuccess = true
)

type SecondaryDNSTsigGetResponseEnvelope struct {
	Errors   []SecondaryDNSTsigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSTsigGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTsigGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSTsigGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSTsigGetResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSTsigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSTsigGetResponseEnvelope]
type secondaryDNSTsigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    secondaryDNSTsigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSTsigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDNSTsigGetResponseEnvelopeErrors]
type secondaryDNSTsigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDNSTsigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSTsigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDNSTsigGetResponseEnvelopeMessages]
type secondaryDNSTsigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSTsigGetResponseEnvelopeSuccess bool

const (
	SecondaryDNSTsigGetResponseEnvelopeSuccessTrue SecondaryDNSTsigGetResponseEnvelopeSuccess = true
)

type SecondaryDNSTsigSecondaryDNSTsigNewTsigParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r SecondaryDNSTsigSecondaryDNSTsigNewTsigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelope struct {
	Errors   []SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTsigSecondaryDNSTsigNewTsigResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelope]
type secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeErrors]
type secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeMessages]
type secondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeSuccess bool

const (
	SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeSuccessTrue SecondaryDNSTsigSecondaryDNSTsigNewTsigResponseEnvelopeSuccess = true
)

type SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelope struct {
	Errors   []SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeMessages `json:"messages,required"`
	Result   []SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeJSON       `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelope]
type secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeErrors]
type secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeMessages]
type secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeSuccess bool

const (
	SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeSuccessTrue SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeSuccess = true
)

type SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeResultInfoJSON `json:"-"`
}

// secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeResultInfo]
type secondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
