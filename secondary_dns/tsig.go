// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TSIGService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTSIGService] method instead.
type TSIGService struct {
	Options []option.RequestOption
}

// NewTSIGService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTSIGService(opts ...option.RequestOption) (r *TSIGService) {
	r = &TSIGService{}
	r.Options = opts
	return
}

// Create TSIG.
func (r *TSIGService) New(ctx context.Context, params TSIGNewParams, opts ...option.RequestOption) (res *SecondaryDNSTSIG, err error) {
	opts = append(r.Options[:], opts...)
	var env TSIGNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify TSIG.
func (r *TSIGService) Update(ctx context.Context, tsigID string, params TSIGUpdateParams, opts ...option.RequestOption) (res *SecondaryDNSTSIG, err error) {
	opts = append(r.Options[:], opts...)
	var env TSIGUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", params.AccountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List TSIGs.
func (r *TSIGService) List(ctx context.Context, query TSIGListParams, opts ...option.RequestOption) (res *[]SecondaryDNSTSIG, err error) {
	opts = append(r.Options[:], opts...)
	var env TSIGListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete TSIG.
func (r *TSIGService) Delete(ctx context.Context, tsigID string, body TSIGDeleteParams, opts ...option.RequestOption) (res *TSIGDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TSIGDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", body.AccountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get TSIG.
func (r *TSIGService) Get(ctx context.Context, tsigID string, query TSIGGetParams, opts ...option.RequestOption) (res *SecondaryDNSTSIG, err error) {
	opts = append(r.Options[:], opts...)
	var env TSIGGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", query.AccountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSTSIG struct {
	ID string `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string               `json:"secret,required"`
	JSON   secondaryDnstsigJSON `json:"-"`
}

// secondaryDnstsigJSON contains the JSON metadata for the struct
// [SecondaryDNSTSIG]
type secondaryDnstsigJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSTSIG) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnstsigJSON) RawJSON() string {
	return r.raw
}

type TSIGDeleteResponse struct {
	ID   string                 `json:"id"`
	JSON tsigDeleteResponseJSON `json:"-"`
}

// tsigDeleteResponseJSON contains the JSON metadata for the struct
// [TSIGDeleteResponse]
type tsigDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TSIGNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r TSIGNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TSIGNewResponseEnvelope struct {
	Errors   []TSIGNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TSIGNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTSIG                  `json:"result,required"`
	// Whether the API call was successful
	Success TSIGNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tsigNewResponseEnvelopeJSON    `json:"-"`
}

// tsigNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGNewResponseEnvelope]
type tsigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TSIGNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    tsigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// tsigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TSIGNewResponseEnvelopeErrors]
type tsigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TSIGNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    tsigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// tsigNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TSIGNewResponseEnvelopeMessages]
type tsigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGNewResponseEnvelopeSuccess bool

const (
	TSIGNewResponseEnvelopeSuccessTrue TSIGNewResponseEnvelopeSuccess = true
)

func (r TSIGNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TSIGUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r TSIGUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TSIGUpdateResponseEnvelope struct {
	Errors   []TSIGUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TSIGUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTSIG                     `json:"result,required"`
	// Whether the API call was successful
	Success TSIGUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    tsigUpdateResponseEnvelopeJSON    `json:"-"`
}

// tsigUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGUpdateResponseEnvelope]
type tsigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TSIGUpdateResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    tsigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// tsigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TSIGUpdateResponseEnvelopeErrors]
type tsigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TSIGUpdateResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    tsigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// tsigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TSIGUpdateResponseEnvelopeMessages]
type tsigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGUpdateResponseEnvelopeSuccess bool

const (
	TSIGUpdateResponseEnvelopeSuccessTrue TSIGUpdateResponseEnvelopeSuccess = true
)

func (r TSIGUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TSIGListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type TSIGListResponseEnvelope struct {
	Errors   []TSIGListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TSIGListResponseEnvelopeMessages `json:"messages,required"`
	Result   []SecondaryDNSTSIG                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    TSIGListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo TSIGListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       tsigListResponseEnvelopeJSON       `json:"-"`
}

// tsigListResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGListResponseEnvelope]
type tsigListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TSIGListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    tsigListResponseEnvelopeErrorsJSON `json:"-"`
}

// tsigListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TSIGListResponseEnvelopeErrors]
type tsigListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TSIGListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    tsigListResponseEnvelopeMessagesJSON `json:"-"`
}

// tsigListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TSIGListResponseEnvelopeMessages]
type tsigListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGListResponseEnvelopeSuccess bool

const (
	TSIGListResponseEnvelopeSuccessTrue TSIGListResponseEnvelopeSuccess = true
)

func (r TSIGListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TSIGListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       tsigListResponseEnvelopeResultInfoJSON `json:"-"`
}

// tsigListResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [TSIGListResponseEnvelopeResultInfo]
type tsigListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type TSIGDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type TSIGDeleteResponseEnvelope struct {
	Errors   []TSIGDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TSIGDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TSIGDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TSIGDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tsigDeleteResponseEnvelopeJSON    `json:"-"`
}

// tsigDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGDeleteResponseEnvelope]
type tsigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TSIGDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    tsigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// tsigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TSIGDeleteResponseEnvelopeErrors]
type tsigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TSIGDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    tsigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// tsigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TSIGDeleteResponseEnvelopeMessages]
type tsigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGDeleteResponseEnvelopeSuccess bool

const (
	TSIGDeleteResponseEnvelopeSuccessTrue TSIGDeleteResponseEnvelopeSuccess = true
)

func (r TSIGDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TSIGGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type TSIGGetResponseEnvelope struct {
	Errors   []TSIGGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TSIGGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSTSIG                  `json:"result,required"`
	// Whether the API call was successful
	Success TSIGGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tsigGetResponseEnvelopeJSON    `json:"-"`
}

// tsigGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGGetResponseEnvelope]
type tsigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TSIGGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    tsigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// tsigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TSIGGetResponseEnvelopeErrors]
type tsigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TSIGGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    tsigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// tsigGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TSIGGetResponseEnvelopeMessages]
type tsigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGGetResponseEnvelopeSuccess bool

const (
	TSIGGetResponseEnvelopeSuccessTrue TSIGGetResponseEnvelopeSuccess = true
)

func (r TSIGGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
