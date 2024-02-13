// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// GatewayListItemService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayListItemService] method
// instead.
type GatewayListItemService struct {
	Options []option.RequestOption
}

// NewGatewayListItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayListItemService(opts ...option.RequestOption) (r *GatewayListItemService) {
	r = &GatewayListItemService{}
	r.Options = opts
	return
}

// Fetches all items in a single Zero Trust list.
func (r *GatewayListItemService) ZeroTrustListsZeroTrustListItems(ctx context.Context, accountID interface{}, listID string, opts ...option.RequestOption) (res *[][]GatewayListItemZeroTrustListsZeroTrustListItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayListItemZeroTrustListsZeroTrustListItemsResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string                                                      `json:"value"`
	JSON  gatewayListItemZeroTrustListsZeroTrustListItemsResponseJSON `json:"-"`
}

// gatewayListItemZeroTrustListsZeroTrustListItemsResponseJSON contains the JSON
// metadata for the struct
// [GatewayListItemZeroTrustListsZeroTrustListItemsResponse]
type gatewayListItemZeroTrustListsZeroTrustListItemsResponseJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemZeroTrustListsZeroTrustListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelope struct {
	Errors   []GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeMessages `json:"messages,required"`
	Result   [][]GatewayListItemZeroTrustListsZeroTrustListItemsResponse               `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeJSON       `json:"-"`
}

// gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelope]
type gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeErrors]
type gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeMessages]
type gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeSuccess bool

const (
	GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeSuccessTrue GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeSuccess = true
)

type GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeResultInfo struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                                       `json:"total_count"`
	JSON       gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeResultInfo]
type gatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemZeroTrustListsZeroTrustListItemsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
