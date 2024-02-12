// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// List all accounts you have ownership or verified access to.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *[]AccountListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountListResponseEnvelope
	path := "accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountListResponse = interface{}

type AccountListParams struct {
	// Direction to order results.
	Direction param.Field[AccountListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type AccountListParamsDirection string

const (
	AccountListParamsDirectionAsc  AccountListParamsDirection = "asc"
	AccountListParamsDirectionDesc AccountListParamsDirection = "desc"
)

type AccountListResponseEnvelope struct {
	Errors   []AccountListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccountListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccountListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccountListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accountListResponseEnvelopeJSON       `json:"-"`
}

// accountListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountListResponseEnvelope]
type accountListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountListResponseEnvelopeErrorsJSON `json:"-"`
}

// accountListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccountListResponseEnvelopeErrors]
type accountListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountListResponseEnvelopeMessagesJSON `json:"-"`
}

// accountListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountListResponseEnvelopeMessages]
type accountListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountListResponseEnvelopeSuccess bool

const (
	AccountListResponseEnvelopeSuccessTrue AccountListResponseEnvelopeSuccess = true
)

type AccountListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       accountListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accountListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccountListResponseEnvelopeResultInfo]
type accountListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
