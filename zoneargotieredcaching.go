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

// ZoneArgoTieredCachingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneArgoTieredCachingService]
// method instead.
type ZoneArgoTieredCachingService struct {
	Options []option.RequestOption
}

// NewZoneArgoTieredCachingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneArgoTieredCachingService(opts ...option.RequestOption) (r *ZoneArgoTieredCachingService) {
	r = &ZoneArgoTieredCachingService{}
	r.Options = opts
	return
}

// Get Tiered Caching setting
func (r *ZoneArgoTieredCachingService) TieredCachingGetTieredCachingSetting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates enablement of Tiered Caching
func (r *ZoneArgoTieredCachingService) TieredCachingPatchTieredCachingSetting(ctx context.Context, zoneIdentifier string, body ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParams, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParams struct {
	// Enables Tiered Caching.
	Value param.Field[ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue] `json:"value,required"`
}

func (r ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Caching.
type ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue string

const (
	ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValueOn  ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue = "on"
	ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValueOff ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue = "off"
)
