// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWaitingRoomSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWaitingRoomSettingService]
// method instead.
type ZoneWaitingRoomSettingService struct {
	Options []option.RequestOption
}

// NewZoneWaitingRoomSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWaitingRoomSettingService(opts ...option.RequestOption) (r *ZoneWaitingRoomSettingService) {
	r = &ZoneWaitingRoomSettingService{}
	r.Options = opts
	return
}

// Update zone-level Waiting Room settings
func (r *ZoneWaitingRoomSettingService) Update(ctx context.Context, zoneIdentifier string, body ZoneWaitingRoomSettingUpdateParams, opts ...option.RequestOption) (res *shared.WaitingRoomSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get zone-level Waiting Room settings
func (r *ZoneWaitingRoomSettingService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *shared.WaitingRoomSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch zone-level Waiting Room settings
func (r *ZoneWaitingRoomSettingService) Patch(ctx context.Context, zoneIdentifier string, body ZoneWaitingRoomSettingPatchParams, opts ...option.RequestOption) (res *shared.WaitingRoomSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneWaitingRoomSettingUpdateParams struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass param.Field[bool] `json:"search_engine_crawler_bypass"`
}

func (r ZoneWaitingRoomSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWaitingRoomSettingPatchParams struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass param.Field[bool] `json:"search_engine_crawler_bypass"`
}

func (r ZoneWaitingRoomSettingPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
