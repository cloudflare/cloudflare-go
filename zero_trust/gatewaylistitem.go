// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *GatewayListItemService) List(ctx context.Context, listID string, query GatewayListItemListParams, opts ...option.RequestOption) (res *[][]GatewayListItemListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListItemListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s/items", query.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayListItemListResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string                          `json:"value"`
	JSON  gatewayListItemListResponseJSON `json:"-"`
}

// gatewayListItemListResponseJSON contains the JSON metadata for the struct
// [GatewayListItemListResponse]
type gatewayListItemListResponseJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListItemListResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayListItemListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayListItemListResponseEnvelope struct {
	Errors   []GatewayListItemListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListItemListResponseEnvelopeMessages `json:"messages,required"`
	Result   [][]GatewayListItemListResponse               `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayListItemListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayListItemListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayListItemListResponseEnvelopeJSON       `json:"-"`
}

// gatewayListItemListResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayListItemListResponseEnvelope]
type gatewayListItemListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListItemListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayListItemListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayListItemListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListItemListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayListItemListResponseEnvelopeErrors]
type gatewayListItemListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListItemListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayListItemListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayListItemListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListItemListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayListItemListResponseEnvelopeMessages]
type gatewayListItemListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListItemListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListItemListResponseEnvelopeSuccess bool

const (
	GatewayListItemListResponseEnvelopeSuccessTrue GatewayListItemListResponseEnvelopeSuccess = true
)

func (r GatewayListItemListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListItemListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListItemListResponseEnvelopeResultInfo struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                           `json:"total_count"`
	JSON       gatewayListItemListResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayListItemListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [GatewayListItemListResponseEnvelopeResultInfo]
type gatewayListItemListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListItemListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListItemListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
