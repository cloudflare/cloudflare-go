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

// Create TSIG.
func (r *SecondaryDNSTsigService) New(ctx context.Context, accountID interface{}, body SecondaryDNSTsigNewParams, opts ...option.RequestOption) (res *SecondaryDNSTsigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List TSIGs.
func (r *SecondaryDNSTsigService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]SecondaryDNSTsigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete TSIG.
func (r *SecondaryDNSTsigService) Delete(ctx context.Context, accountID interface{}, tsigID interface{}, opts ...option.RequestOption) (res *SecondaryDNSTsigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get TSIG.
func (r *SecondaryDNSTsigService) Get(ctx context.Context, accountID interface{}, tsigID interface{}, opts ...option.RequestOption) (res *SecondaryDNSTsigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify TSIG.
func (r *SecondaryDNSTsigService) Replace(ctx context.Context, accountID interface{}, tsigID interface{}, body SecondaryDNSTsigReplaceParams, opts ...option.RequestOption) (res *SecondaryDNSTsigReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSTsigReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSTsigNewResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                          `json:"secret,required"`
	JSON   secondaryDNSTsigNewResponseJSON `json:"-"`
}

// secondaryDNSTsigNewResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSTsigNewResponse]
type secondaryDNSTsigNewResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigListResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                           `json:"secret,required"`
	JSON   secondaryDNSTsigListResponseJSON `json:"-"`
}

// secondaryDNSTsigListResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSTsigListResponse]
type secondaryDNSTsigListResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigListResponse) UnmarshalJSON(data []byte) (err error) {
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

type SecondaryDNSTsigReplaceResponse struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                              `json:"secret,required"`
	JSON   secondaryDNSTsigReplaceResponseJSON `json:"-"`
}

// secondaryDNSTsigReplaceResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSTsigReplaceResponse]
type secondaryDNSTsigReplaceResponseJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigNewParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r SecondaryDNSTsigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSTsigNewResponseEnvelope struct {
	Errors   []SecondaryDNSTsigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSTsigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTsigNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSTsigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSTsigNewResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSTsigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSTsigNewResponseEnvelope]
type secondaryDNSTsigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    secondaryDNSTsigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSTsigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDNSTsigNewResponseEnvelopeErrors]
type secondaryDNSTsigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDNSTsigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSTsigNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDNSTsigNewResponseEnvelopeMessages]
type secondaryDNSTsigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSTsigNewResponseEnvelopeSuccess bool

const (
	SecondaryDNSTsigNewResponseEnvelopeSuccessTrue SecondaryDNSTsigNewResponseEnvelopeSuccess = true
)

type SecondaryDNSTsigListResponseEnvelope struct {
	Errors   []SecondaryDNSTsigListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSTsigListResponseEnvelopeMessages `json:"messages,required"`
	Result   []SecondaryDNSTsigListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SecondaryDNSTsigListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SecondaryDNSTsigListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       secondaryDNSTsigListResponseEnvelopeJSON       `json:"-"`
}

// secondaryDNSTsigListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSTsigListResponseEnvelope]
type secondaryDNSTsigListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    secondaryDNSTsigListResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSTsigListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSTsigListResponseEnvelopeErrors]
type secondaryDNSTsigListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDNSTsigListResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSTsigListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDNSTsigListResponseEnvelopeMessages]
type secondaryDNSTsigListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSTsigListResponseEnvelopeSuccess bool

const (
	SecondaryDNSTsigListResponseEnvelopeSuccessTrue SecondaryDNSTsigListResponseEnvelopeSuccess = true
)

type SecondaryDNSTsigListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       secondaryDNSTsigListResponseEnvelopeResultInfoJSON `json:"-"`
}

// secondaryDNSTsigListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [SecondaryDNSTsigListResponseEnvelopeResultInfo]
type secondaryDNSTsigListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type SecondaryDNSTsigReplaceParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r SecondaryDNSTsigReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSTsigReplaceResponseEnvelope struct {
	Errors   []SecondaryDNSTsigReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSTsigReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTsigReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSTsigReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSTsigReplaceResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSTsigReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSTsigReplaceResponseEnvelope]
type secondaryDNSTsigReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigReplaceResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    secondaryDNSTsigReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSTsigReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSTsigReplaceResponseEnvelopeErrors]
type secondaryDNSTsigReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSTsigReplaceResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    secondaryDNSTsigReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSTsigReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSTsigReplaceResponseEnvelopeMessages]
type secondaryDNSTsigReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTsigReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSTsigReplaceResponseEnvelopeSuccess bool

const (
	SecondaryDNSTsigReplaceResponseEnvelopeSuccessTrue SecondaryDNSTsigReplaceResponseEnvelopeSuccess = true
)
