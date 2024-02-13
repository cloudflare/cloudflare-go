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

// CustomNAvailabilityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomNAvailabilityService]
// method instead.
type CustomNAvailabilityService struct {
	Options []option.RequestOption
}

// NewCustomNAvailabilityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomNAvailabilityService(opts ...option.RequestOption) (r *CustomNAvailabilityService) {
	r = &CustomNAvailabilityService{}
	r.Options = opts
	return
}

// Get Eligible Zones for Account Custom Nameservers
func (r *CustomNAvailabilityService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomNAvailabilityGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/custom_ns/availability", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomNAvailabilityGetResponseEnvelope struct {
	Errors   []CustomNAvailabilityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNAvailabilityGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []string                                         `json:"result,required,nullable" format:"hostname"`
	// Whether the API call was successful
	Success    CustomNAvailabilityGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomNAvailabilityGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customNAvailabilityGetResponseEnvelopeJSON       `json:"-"`
}

// customNAvailabilityGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNAvailabilityGetResponseEnvelope]
type customNAvailabilityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNAvailabilityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNAvailabilityGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    customNAvailabilityGetResponseEnvelopeErrorsJSON `json:"-"`
}

// customNAvailabilityGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomNAvailabilityGetResponseEnvelopeErrors]
type customNAvailabilityGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNAvailabilityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomNAvailabilityGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    customNAvailabilityGetResponseEnvelopeMessagesJSON `json:"-"`
}

// customNAvailabilityGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CustomNAvailabilityGetResponseEnvelopeMessages]
type customNAvailabilityGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNAvailabilityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomNAvailabilityGetResponseEnvelopeSuccess bool

const (
	CustomNAvailabilityGetResponseEnvelopeSuccessTrue CustomNAvailabilityGetResponseEnvelopeSuccess = true
)

type CustomNAvailabilityGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       customNAvailabilityGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNAvailabilityGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [CustomNAvailabilityGetResponseEnvelopeResultInfo]
type customNAvailabilityGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNAvailabilityGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
