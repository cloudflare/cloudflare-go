// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// IPListService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewIPListService] method instead.
type IPListService struct {
	Options []option.RequestOption
}

// NewIPListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPListService(opts ...option.RequestOption) (r *IPListService) {
	r = &IPListService{}
	r.Options = opts
	return
}

// Get IP Lists
func (r *IPListService) Get(ctx context.Context, query IPListGetParams, opts ...option.RequestOption) (res *[]IntelIPList, err error) {
	opts = append(r.Options[:], opts...)
	var env IPListGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/ip-list", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelIPList struct {
	ID          int64           `json:"id"`
	Description string          `json:"description"`
	Name        string          `json:"name"`
	JSON        intelIPListJSON `json:"-"`
}

// intelIPListJSON contains the JSON metadata for the struct [IntelIPList]
type intelIPListJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIPListJSON) RawJSON() string {
	return r.raw
}

type IPListGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IPListGetResponseEnvelope struct {
	Errors   []IPListGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPListGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelIPList                       `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IPListGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IPListGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ipListGetResponseEnvelopeJSON       `json:"-"`
}

// ipListGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPListGetResponseEnvelope]
type ipListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPListGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    ipListGetResponseEnvelopeErrorsJSON `json:"-"`
}

// ipListGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IPListGetResponseEnvelopeErrors]
type ipListGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPListGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    ipListGetResponseEnvelopeMessagesJSON `json:"-"`
}

// ipListGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IPListGetResponseEnvelopeMessages]
type ipListGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPListGetResponseEnvelopeSuccess bool

const (
	IPListGetResponseEnvelopeSuccessTrue IPListGetResponseEnvelopeSuccess = true
)

type IPListGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       ipListGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// ipListGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [IPListGetResponseEnvelopeResultInfo]
type ipListGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
