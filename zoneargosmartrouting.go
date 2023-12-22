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

// ZoneArgoSmartRoutingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneArgoSmartRoutingService]
// method instead.
type ZoneArgoSmartRoutingService struct {
	Options []option.RequestOption
}

// NewZoneArgoSmartRoutingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneArgoSmartRoutingService(opts ...option.RequestOption) (r *ZoneArgoSmartRoutingService) {
	r = &ZoneArgoSmartRoutingService{}
	r.Options = opts
	return
}

// Get Argo Smart Routing setting
func (r *ZoneArgoSmartRoutingService) ArgoSmartRoutingGetArgoSmartRoutingSetting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates enablement of Argo Smart Routing.
func (r *ZoneArgoSmartRoutingService) ArgoSmartRoutingPatchArgoSmartRoutingSetting(ctx context.Context, zoneIdentifier string, body ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParams, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParams struct {
	// Enables Argo Smart Routing.
	Value param.Field[ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValue] `json:"value,required"`
}

func (r ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Argo Smart Routing.
type ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValue string

const (
	ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValueOn  ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValue = "on"
	ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValueOff ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValue = "off"
)
