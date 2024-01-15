// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingZarazHistoryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingZarazHistoryService] method instead.
type ZoneSettingZarazHistoryService struct {
	Options []option.RequestOption
}

// NewZoneSettingZarazHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZarazHistoryService(opts ...option.RequestOption) (r *ZoneSettingZarazHistoryService) {
	r = &ZoneSettingZarazHistoryService{}
	r.Options = opts
	return
}

// Restores a historical published Zaraz configuration by ID for a zone.
func (r *ZoneSettingZarazHistoryService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingZarazHistoryUpdateParams, opts ...option.RequestOption) (res *ZarazConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/v2/history", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists a history of published Zaraz configuration records for a zone.
func (r *ZoneSettingZarazHistoryService) List(ctx context.Context, zoneIdentifier string, query ZoneSettingZarazHistoryListParams, opts ...option.RequestOption) (res *ZoneSettingZarazHistoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/v2/history", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSettingZarazHistoryListResponse struct {
	Errors   []ZoneSettingZarazHistoryListResponseError   `json:"errors"`
	Messages []ZoneSettingZarazHistoryListResponseMessage `json:"messages"`
	Result   []ZoneSettingZarazHistoryListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success bool                                    `json:"success"`
	JSON    zoneSettingZarazHistoryListResponseJSON `json:"-"`
}

// zoneSettingZarazHistoryListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingZarazHistoryListResponse]
type zoneSettingZarazHistoryListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazHistoryListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingZarazHistoryListResponseErrorJSON `json:"-"`
}

// zoneSettingZarazHistoryListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingZarazHistoryListResponseError]
type zoneSettingZarazHistoryListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazHistoryListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazHistoryListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingZarazHistoryListResponseMessageJSON `json:"-"`
}

// zoneSettingZarazHistoryListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingZarazHistoryListResponseMessage]
type zoneSettingZarazHistoryListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazHistoryListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazHistoryListResponseResult struct {
	// ID of the configuration
	ID int64 `json:"id"`
	// Date and time the configuration was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Configuration description provided by the user who published this configuration
	Description string `json:"description"`
	// Date and time the configuration was last updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Alpha-numeric ID of the account user who published the configuration
	UserID string                                        `json:"userId"`
	JSON   zoneSettingZarazHistoryListResponseResultJSON `json:"-"`
}

// zoneSettingZarazHistoryListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingZarazHistoryListResponseResult]
type zoneSettingZarazHistoryListResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	UpdatedAt   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazHistoryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazHistoryUpdateParams struct {
	// ID of the Zaraz configuration to restore.
	Body param.Field[int64] `json:"body,required"`
}

func (r ZoneSettingZarazHistoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneSettingZarazHistoryListParams struct {
	// Maximum amount of results to list. Default value is 10.
	Limit param.Field[int64] `query:"limit"`
	// Ordinal number to start listing the results with. Default value is 0.
	Offset param.Field[int64] `query:"offset"`
	// The field to sort by. Default is updated_at.
	SortField param.Field[ZoneSettingZarazHistoryListParamsSortField] `query:"sortField"`
	// Sorting order. Default is DESC.
	SortOrder param.Field[ZoneSettingZarazHistoryListParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [ZoneSettingZarazHistoryListParams]'s query parameters as
// `url.Values`.
func (r ZoneSettingZarazHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort by. Default is updated_at.
type ZoneSettingZarazHistoryListParamsSortField string

const (
	ZoneSettingZarazHistoryListParamsSortFieldID          ZoneSettingZarazHistoryListParamsSortField = "id"
	ZoneSettingZarazHistoryListParamsSortFieldUserID      ZoneSettingZarazHistoryListParamsSortField = "user_id"
	ZoneSettingZarazHistoryListParamsSortFieldDescription ZoneSettingZarazHistoryListParamsSortField = "description"
	ZoneSettingZarazHistoryListParamsSortFieldCreatedAt   ZoneSettingZarazHistoryListParamsSortField = "created_at"
	ZoneSettingZarazHistoryListParamsSortFieldUpdatedAt   ZoneSettingZarazHistoryListParamsSortField = "updated_at"
)

// Sorting order. Default is DESC.
type ZoneSettingZarazHistoryListParamsSortOrder string

const (
	ZoneSettingZarazHistoryListParamsSortOrderDesc ZoneSettingZarazHistoryListParamsSortOrder = "DESC"
	ZoneSettingZarazHistoryListParamsSortOrderAsc  ZoneSettingZarazHistoryListParamsSortOrder = "ASC"
)
