// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WaitingRoomSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWaitingRoomSettingService] method
// instead.
type WaitingRoomSettingService struct {
	Options []option.RequestOption
}

// NewWaitingRoomSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaitingRoomSettingService(opts ...option.RequestOption) (r *WaitingRoomSettingService) {
	r = &WaitingRoomSettingService{}
	r.Options = opts
	return
}

// Patch zone-level Waiting Room settings
func (r *WaitingRoomSettingService) Edit(ctx context.Context, zoneIdentifier string, body WaitingRoomSettingEditParams, opts ...option.RequestOption) (res *WaitingRoomSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingroomZoneSettingsResponse
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get zone-level Waiting Room settings
func (r *WaitingRoomSettingService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *WaitingRoomSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingroomZoneSettingsResponse
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WaitingroomZoneSettingsResponse struct {
	Result WaitingroomZoneSettingsResponseResult `json:"result,required"`
	JSON   waitingroomZoneSettingsResponseJSON   `json:"-"`
}

// waitingroomZoneSettingsResponseJSON contains the JSON metadata for the struct
// [WaitingroomZoneSettingsResponse]
type waitingroomZoneSettingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingroomZoneSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingroomZoneSettingsResponseResult struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass bool                                      `json:"search_engine_crawler_bypass,required"`
	JSON                      waitingroomZoneSettingsResponseResultJSON `json:"-"`
}

// waitingroomZoneSettingsResponseResultJSON contains the JSON metadata for the
// struct [WaitingroomZoneSettingsResponseResult]
type waitingroomZoneSettingsResponseResultJSON struct {
	SearchEngineCrawlerBypass apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WaitingroomZoneSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomSettingEditResponse struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass bool                               `json:"search_engine_crawler_bypass,required"`
	JSON                      waitingRoomSettingEditResponseJSON `json:"-"`
}

// waitingRoomSettingEditResponseJSON contains the JSON metadata for the struct
// [WaitingRoomSettingEditResponse]
type waitingRoomSettingEditResponseJSON struct {
	SearchEngineCrawlerBypass apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WaitingRoomSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomSettingGetResponse struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass bool                              `json:"search_engine_crawler_bypass,required"`
	JSON                      waitingRoomSettingGetResponseJSON `json:"-"`
}

// waitingRoomSettingGetResponseJSON contains the JSON metadata for the struct
// [WaitingRoomSettingGetResponse]
type waitingRoomSettingGetResponseJSON struct {
	SearchEngineCrawlerBypass apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WaitingRoomSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomSettingEditParams struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass param.Field[bool] `json:"search_engine_crawler_bypass"`
}

func (r WaitingRoomSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
