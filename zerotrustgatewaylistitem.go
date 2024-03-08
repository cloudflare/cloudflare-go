// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustGatewayListItemService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustGatewayListItemService] method instead.
type ZeroTrustGatewayListItemService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayListItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayListItemService(opts ...option.RequestOption) (r *ZeroTrustGatewayListItemService) {
	r = &ZeroTrustGatewayListItemService{}
	r.Options = opts
	return
}

// Fetches all items in a single Zero Trust list.
func (r *ZeroTrustGatewayListItemService) List(ctx context.Context, listID string, query ZeroTrustGatewayListItemListParams, opts ...option.RequestOption) (res *[][]ZeroTrustGatewayListItemListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayListItemListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s/items", query.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayListItemListResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string                                   `json:"value"`
	JSON  zeroTrustGatewayListItemListResponseJSON `json:"-"`
}

// zeroTrustGatewayListItemListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayListItemListResponse]
type zeroTrustGatewayListItemListResponseJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListItemListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayListItemListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayListItemListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayListItemListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayListItemListResponseEnvelopeMessages `json:"messages,required"`
	Result   [][]ZeroTrustGatewayListItemListResponse               `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustGatewayListItemListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustGatewayListItemListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustGatewayListItemListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustGatewayListItemListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayListItemListResponseEnvelope]
type zeroTrustGatewayListItemListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListItemListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListItemListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayListItemListResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustGatewayListItemListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayListItemListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayListItemListResponseEnvelopeErrors]
type zeroTrustGatewayListItemListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListItemListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListItemListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayListItemListResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustGatewayListItemListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayListItemListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayListItemListResponseEnvelopeMessages]
type zeroTrustGatewayListItemListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListItemListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListItemListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustGatewayListItemListResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayListItemListResponseEnvelopeSuccessTrue ZeroTrustGatewayListItemListResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayListItemListResponseEnvelopeResultInfo struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                    `json:"total_count"`
	JSON       zeroTrustGatewayListItemListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustGatewayListItemListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayListItemListResponseEnvelopeResultInfo]
type zeroTrustGatewayListItemListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListItemListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListItemListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
