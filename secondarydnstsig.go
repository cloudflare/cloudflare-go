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

// SecondaryDNSTSIGService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSTSIGService] method
// instead.
type SecondaryDNSTSIGService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSTSIGService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSTSIGService(opts ...option.RequestOption) (r *SecondaryDNSTSIGService) {
	r = &SecondaryDNSTSIGService{}
	r.Options = opts
	return
}

// Create TSIG.
func (r *SecondaryDNSTSIGService) New(ctx context.Context, accountID interface{}, body SecondaryDNSTSIGNewParams, opts ...option.RequestOption) (res *SecondaryDnstsigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnstsigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify TSIG.
func (r *SecondaryDNSTSIGService) Update(ctx context.Context, accountID interface{}, tsigID interface{}, body SecondaryDNSTSIGUpdateParams, opts ...option.RequestOption) (res *SecondaryDnstsigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnstsigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List TSIGs.
func (r *SecondaryDNSTSIGService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]SecondaryDnstsigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnstsigListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete TSIG.
func (r *SecondaryDNSTSIGService) Delete(ctx context.Context, accountID interface{}, tsigID interface{}, opts ...option.RequestOption) (res *SecondaryDnstsigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnstsigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get TSIG.
func (r *SecondaryDNSTSIGService) Get(ctx context.Context, accountID interface{}, tsigID interface{}, opts ...option.RequestOption) (res *SecondaryDnstsigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDnstsigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDnstsigNewResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                          `json:"secret,required"`
	JSON   secondaryDnstsigNewResponseJSON `json:"-"`
}

// secondaryDnstsigNewResponseJSON contains the JSON metadata for the struct
// [SecondaryDnstsigNewResponse]
type secondaryDnstsigNewResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigUpdateResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                             `json:"secret,required"`
	JSON   secondaryDnstsigUpdateResponseJSON `json:"-"`
}

// secondaryDnstsigUpdateResponseJSON contains the JSON metadata for the struct
// [SecondaryDnstsigUpdateResponse]
type secondaryDnstsigUpdateResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigListResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                           `json:"secret,required"`
	JSON   secondaryDnstsigListResponseJSON `json:"-"`
}

// secondaryDnstsigListResponseJSON contains the JSON metadata for the struct
// [SecondaryDnstsigListResponse]
type secondaryDnstsigListResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigDeleteResponse struct {
	ID   interface{}                        `json:"id"`
	JSON secondaryDnstsigDeleteResponseJSON `json:"-"`
}

// secondaryDnstsigDeleteResponseJSON contains the JSON metadata for the struct
// [SecondaryDnstsigDeleteResponse]
type secondaryDnstsigDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigGetResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                          `json:"secret,required"`
	JSON   secondaryDnstsigGetResponseJSON `json:"-"`
}

// secondaryDnstsigGetResponseJSON contains the JSON metadata for the struct
// [SecondaryDnstsigGetResponse]
type secondaryDnstsigGetResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTSIGNewParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r SecondaryDNSTSIGNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDnstsigNewResponseEnvelope struct {
	Errors   []SecondaryDnstsigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnstsigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnstsigNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnstsigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnstsigNewResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnstsigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDnstsigNewResponseEnvelope]
type secondaryDnstsigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    secondaryDnstsigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnstsigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDnstsigNewResponseEnvelopeErrors]
type secondaryDnstsigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDnstsigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnstsigNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDnstsigNewResponseEnvelopeMessages]
type secondaryDnstsigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDnstsigNewResponseEnvelopeSuccess bool

const (
	SecondaryDnstsigNewResponseEnvelopeSuccessTrue SecondaryDnstsigNewResponseEnvelopeSuccess = true
)

type SecondaryDNSTSIGUpdateParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r SecondaryDNSTSIGUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDnstsigUpdateResponseEnvelope struct {
	Errors   []SecondaryDnstsigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnstsigUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnstsigUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnstsigUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnstsigUpdateResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnstsigUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDnstsigUpdateResponseEnvelope]
type secondaryDnstsigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDnstsigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnstsigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDnstsigUpdateResponseEnvelopeErrors]
type secondaryDnstsigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    secondaryDnstsigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnstsigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDnstsigUpdateResponseEnvelopeMessages]
type secondaryDnstsigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDnstsigUpdateResponseEnvelopeSuccess bool

const (
	SecondaryDnstsigUpdateResponseEnvelopeSuccessTrue SecondaryDnstsigUpdateResponseEnvelopeSuccess = true
)

type SecondaryDnstsigListResponseEnvelope struct {
	Errors   []SecondaryDnstsigListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnstsigListResponseEnvelopeMessages `json:"messages,required"`
	Result   []SecondaryDnstsigListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SecondaryDnstsigListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SecondaryDnstsigListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       secondaryDnstsigListResponseEnvelopeJSON       `json:"-"`
}

// secondaryDnstsigListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDnstsigListResponseEnvelope]
type secondaryDnstsigListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    secondaryDnstsigListResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnstsigListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDnstsigListResponseEnvelopeErrors]
type secondaryDnstsigListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDnstsigListResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnstsigListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDnstsigListResponseEnvelopeMessages]
type secondaryDnstsigListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDnstsigListResponseEnvelopeSuccess bool

const (
	SecondaryDnstsigListResponseEnvelopeSuccessTrue SecondaryDnstsigListResponseEnvelopeSuccess = true
)

type SecondaryDnstsigListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       secondaryDnstsigListResponseEnvelopeResultInfoJSON `json:"-"`
}

// secondaryDnstsigListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [SecondaryDnstsigListResponseEnvelopeResultInfo]
type secondaryDnstsigListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigDeleteResponseEnvelope struct {
	Errors   []SecondaryDnstsigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnstsigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnstsigDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnstsigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnstsigDeleteResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnstsigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDnstsigDeleteResponseEnvelope]
type secondaryDnstsigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDnstsigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnstsigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDnstsigDeleteResponseEnvelopeErrors]
type secondaryDnstsigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    secondaryDnstsigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnstsigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDnstsigDeleteResponseEnvelopeMessages]
type secondaryDnstsigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDnstsigDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDnstsigDeleteResponseEnvelopeSuccessTrue SecondaryDnstsigDeleteResponseEnvelopeSuccess = true
)

type SecondaryDnstsigGetResponseEnvelope struct {
	Errors   []SecondaryDnstsigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDnstsigGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDnstsigGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDnstsigGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDnstsigGetResponseEnvelopeJSON    `json:"-"`
}

// secondaryDnstsigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDnstsigGetResponseEnvelope]
type secondaryDnstsigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    secondaryDnstsigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDnstsigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDnstsigGetResponseEnvelopeErrors]
type secondaryDnstsigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDnstsigGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDnstsigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDnstsigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDnstsigGetResponseEnvelopeMessages]
type secondaryDnstsigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDnstsigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDnstsigGetResponseEnvelopeSuccess bool

const (
	SecondaryDnstsigGetResponseEnvelopeSuccessTrue SecondaryDnstsigGetResponseEnvelopeSuccess = true
)
