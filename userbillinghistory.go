// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserBillingHistoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserBillingHistoryService] method
// instead.
type UserBillingHistoryService struct {
	Options []option.RequestOption
}

// NewUserBillingHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserBillingHistoryService(opts ...option.RequestOption) (r *UserBillingHistoryService) {
	r = &UserBillingHistoryService{}
	r.Options = opts
	return
}

// Accesses your billing history object.
func (r *UserBillingHistoryService) List(ctx context.Context, query UserBillingHistoryListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[UserBillingHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/billing/history"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Accesses your billing history object.
func (r *UserBillingHistoryService) ListAutoPaging(ctx context.Context, query UserBillingHistoryListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[UserBillingHistoryListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

type UserBillingHistoryListResponse struct {
	// Billing item identifier tag.
	ID string `json:"id,required"`
	// The billing item action.
	Action string `json:"action,required"`
	// The amount associated with this billing item.
	Amount float64 `json:"amount,required"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency,required"`
	// The billing item description.
	Description string `json:"description,required"`
	// When the billing item was created.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The billing item type.
	Type string                             `json:"type,required"`
	Zone UserBillingHistoryListResponseZone `json:"zone,required"`
	JSON userBillingHistoryListResponseJSON `json:"-"`
}

// userBillingHistoryListResponseJSON contains the JSON metadata for the struct
// [UserBillingHistoryListResponse]
type userBillingHistoryListResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	OccurredAt  apijson.Field
	Type        apijson.Field
	Zone        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryListResponseZone struct {
	Name interface{}                            `json:"name"`
	JSON userBillingHistoryListResponseZoneJSON `json:"-"`
}

// userBillingHistoryListResponseZoneJSON contains the JSON metadata for the struct
// [UserBillingHistoryListResponseZone]
type userBillingHistoryListResponseZoneJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryListResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryListParams struct {
	// Field to order billing history by.
	Order param.Field[UserBillingHistoryListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserBillingHistoryListParams]'s query parameters as
// `url.Values`.
func (r UserBillingHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to order billing history by.
type UserBillingHistoryListParamsOrder string

const (
	UserBillingHistoryListParamsOrderType      UserBillingHistoryListParamsOrder = "type"
	UserBillingHistoryListParamsOrderOccuredAt UserBillingHistoryListParamsOrder = "occured_at"
	UserBillingHistoryListParamsOrderAction    UserBillingHistoryListParamsOrder = "action"
)
