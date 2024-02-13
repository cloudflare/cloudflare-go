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
func (r *IntelIPListService) IPListGetIPLists(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]IntelIPListIPListGetIPListsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIPListIPListGetIPListsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/ip-list", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelIPListIPListGetIPListsResponse struct {
	ID          int64                                   `json:"id"`
	Description string                                  `json:"description"`
	Name        string                                  `json:"name"`
	JSON        intelIPListIPListGetIPListsResponseJSON `json:"-"`
}

// intelIPListIPListGetIPListsResponseJSON contains the JSON metadata for the
// struct [IntelIPListIPListGetIPListsResponse]
type intelIPListIPListGetIPListsResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListIPListGetIPListsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIPListIPListGetIPListsResponseEnvelope struct {
	Errors   []IntelIPListIPListGetIPListsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIPListIPListGetIPListsResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelIPListIPListGetIPListsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IntelIPListIPListGetIPListsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IntelIPListIPListGetIPListsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       intelIPListIPListGetIPListsResponseEnvelopeJSON       `json:"-"`
}

// intelIPListIPListGetIPListsResponseEnvelopeJSON contains the JSON metadata for
// the struct [IntelIPListIPListGetIPListsResponseEnvelope]
type intelIPListIPListGetIPListsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListIPListGetIPListsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIPListIPListGetIPListsResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    intelIPListIPListGetIPListsResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIPListIPListGetIPListsResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IntelIPListIPListGetIPListsResponseEnvelopeErrors]
type intelIPListIPListGetIPListsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListIPListGetIPListsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIPListIPListGetIPListsResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    intelIPListIPListGetIPListsResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIPListIPListGetIPListsResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IntelIPListIPListGetIPListsResponseEnvelopeMessages]
type intelIPListIPListGetIPListsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListIPListGetIPListsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIPListIPListGetIPListsResponseEnvelopeSuccess bool

const (
	IntelIPListIPListGetIPListsResponseEnvelopeSuccessTrue IntelIPListIPListGetIPListsResponseEnvelopeSuccess = true
)

type IntelIPListIPListGetIPListsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       intelIPListIPListGetIPListsResponseEnvelopeResultInfoJSON `json:"-"`
}

// intelIPListIPListGetIPListsResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [IntelIPListIPListGetIPListsResponseEnvelopeResultInfo]
type intelIPListIPListGetIPListsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPListIPListGetIPListsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
