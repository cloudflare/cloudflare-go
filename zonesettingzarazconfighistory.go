// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingZarazConfigHistoryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingZarazConfigHistoryService] method instead.
type ZoneSettingZarazConfigHistoryService struct {
	Options []option.RequestOption
}

// NewZoneSettingZarazConfigHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingZarazConfigHistoryService(opts ...option.RequestOption) (r *ZoneSettingZarazConfigHistoryService) {
	r = &ZoneSettingZarazConfigHistoryService{}
	r.Options = opts
	return
}

// Gets a history of published Zaraz configurations by ID(s) for a zone.
func (r *ZoneSettingZarazConfigHistoryService) Get(ctx context.Context, zoneIdentifier string, query ZoneSettingZarazConfigHistoryGetParams, opts ...option.RequestOption) (res *ZoneSettingZarazConfigHistoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/v2/history/configs", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSettingZarazConfigHistoryGetResponse struct {
	Errors   []ZoneSettingZarazConfigHistoryGetResponseError   `json:"errors"`
	Messages []ZoneSettingZarazConfigHistoryGetResponseMessage `json:"messages"`
	// Object where keys are numericc onfiguration IDs
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success bool                                         `json:"success"`
	JSON    zoneSettingZarazConfigHistoryGetResponseJSON `json:"-"`
}

// zoneSettingZarazConfigHistoryGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingZarazConfigHistoryGetResponse]
type zoneSettingZarazConfigHistoryGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazConfigHistoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazConfigHistoryGetResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingZarazConfigHistoryGetResponseErrorJSON `json:"-"`
}

// zoneSettingZarazConfigHistoryGetResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingZarazConfigHistoryGetResponseError]
type zoneSettingZarazConfigHistoryGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazConfigHistoryGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazConfigHistoryGetResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingZarazConfigHistoryGetResponseMessageJSON `json:"-"`
}

// zoneSettingZarazConfigHistoryGetResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingZarazConfigHistoryGetResponseMessage]
type zoneSettingZarazConfigHistoryGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazConfigHistoryGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazConfigHistoryGetParams struct {
	// Comma separated list of Zaraz configuration IDs
	IDs param.Field[[]int64] `query:"ids,required"`
}

// URLQuery serializes [ZoneSettingZarazConfigHistoryGetParams]'s query parameters
// as `url.Values`.
func (r ZoneSettingZarazConfigHistoryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
