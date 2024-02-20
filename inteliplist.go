// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IntelIPListService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelIPListService] method
// instead.
type IntelIPListService struct {
	Options []option.RequestOption
}

// NewIntelIPListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelIPListService(opts ...option.RequestOption) (r *IntelIPListService) {
	r = &IntelIPListService{}
	r.Options = opts
	return
}

// Get IP Lists
func (r *IntelIPListService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]IntelIPListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIPListGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/ip-list", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelIPListGetResponse struct {
	ID          int64                      `json:"id"`
	Description string                     `json:"description"`
	Name        string                     `json:"name"`
	JSON        intelIPListGetResponseJSON `json:"-"`
}

// intelIPListGetResponseJSON contains the JSON metadata for the struct
// [IntelIPListGetResponse]
type intelIPListGetResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIPListGetResponseEnvelope struct {
	Errors   []IntelIPListGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIPListGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelIPListGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IntelIPListGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IntelIPListGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       intelIPListGetResponseEnvelopeJSON       `json:"-"`
}

// intelIPListGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelIPListGetResponseEnvelope]
type intelIPListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIPListGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    intelIPListGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIPListGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IntelIPListGetResponseEnvelopeErrors]
type intelIPListGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIPListGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    intelIPListGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIPListGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IntelIPListGetResponseEnvelopeMessages]
type intelIPListGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIPListGetResponseEnvelopeSuccess bool

const (
	IntelIPListGetResponseEnvelopeSuccessTrue IntelIPListGetResponseEnvelopeSuccess = true
)

type IntelIPListGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       intelIPListGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// intelIPListGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [IntelIPListGetResponseEnvelopeResultInfo]
type intelIPListGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
