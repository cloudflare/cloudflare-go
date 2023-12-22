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

// ZoneCustomPageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneCustomPageService] method
// instead.
type ZoneCustomPageService struct {
	Options []option.RequestOption
}

// NewZoneCustomPageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCustomPageService(opts ...option.RequestOption) (r *ZoneCustomPageService) {
	r = &ZoneCustomPageService{}
	r.Options = opts
	return
}

// Fetches the details of a custom page.
func (r *ZoneCustomPageService) Get(ctx context.Context, zoneIdentifier string, identifier ZoneCustomPageGetParamsIdentifier, opts ...option.RequestOption) (res *CustomPagesResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_pages/%v", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration of an existing custom page.
func (r *ZoneCustomPageService) Update(ctx context.Context, zoneIdentifier string, identifier ZoneCustomPageUpdateParamsIdentifier, body ZoneCustomPageUpdateParams, opts ...option.RequestOption) (res *CustomPagesResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_pages/%v", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all the custom pages at the zone level.
func (r *ZoneCustomPageService) CustomPagesForAZoneListCustomPages(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *CustomPagesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_pages", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The name of the custom page type.
type ZoneCustomPageGetParamsIdentifier string

const (
	ZoneCustomPageGetParamsIdentifierBasicChallenge   ZoneCustomPageGetParamsIdentifier = "basic_challenge"
	ZoneCustomPageGetParamsIdentifierManagedChallenge ZoneCustomPageGetParamsIdentifier = "managed_challenge"
	ZoneCustomPageGetParamsIdentifierWafBlock         ZoneCustomPageGetParamsIdentifier = "waf_block"
	ZoneCustomPageGetParamsIdentifierRatelimitBlock   ZoneCustomPageGetParamsIdentifier = "ratelimit_block"
	ZoneCustomPageGetParamsIdentifierCountryChallenge ZoneCustomPageGetParamsIdentifier = "country_challenge"
	ZoneCustomPageGetParamsIdentifierIPBlock          ZoneCustomPageGetParamsIdentifier = "ip_block"
	ZoneCustomPageGetParamsIdentifierUnderAttack      ZoneCustomPageGetParamsIdentifier = "under_attack"
	ZoneCustomPageGetParamsIdentifier500Errors        ZoneCustomPageGetParamsIdentifier = "500_errors"
	ZoneCustomPageGetParamsIdentifier1000Errors       ZoneCustomPageGetParamsIdentifier = "1000_errors"
)

type ZoneCustomPageUpdateParams struct {
	// The custom page state.
	State param.Field[ZoneCustomPageUpdateParamsState] `json:"state,required"`
	// The URL associated with the custom page.
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r ZoneCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of the custom page type.
type ZoneCustomPageUpdateParamsIdentifier string

const (
	ZoneCustomPageUpdateParamsIdentifierBasicChallenge   ZoneCustomPageUpdateParamsIdentifier = "basic_challenge"
	ZoneCustomPageUpdateParamsIdentifierManagedChallenge ZoneCustomPageUpdateParamsIdentifier = "managed_challenge"
	ZoneCustomPageUpdateParamsIdentifierWafBlock         ZoneCustomPageUpdateParamsIdentifier = "waf_block"
	ZoneCustomPageUpdateParamsIdentifierRatelimitBlock   ZoneCustomPageUpdateParamsIdentifier = "ratelimit_block"
	ZoneCustomPageUpdateParamsIdentifierCountryChallenge ZoneCustomPageUpdateParamsIdentifier = "country_challenge"
	ZoneCustomPageUpdateParamsIdentifierIPBlock          ZoneCustomPageUpdateParamsIdentifier = "ip_block"
	ZoneCustomPageUpdateParamsIdentifierUnderAttack      ZoneCustomPageUpdateParamsIdentifier = "under_attack"
	ZoneCustomPageUpdateParamsIdentifier500Errors        ZoneCustomPageUpdateParamsIdentifier = "500_errors"
	ZoneCustomPageUpdateParamsIdentifier1000Errors       ZoneCustomPageUpdateParamsIdentifier = "1000_errors"
)

// The custom page state.
type ZoneCustomPageUpdateParamsState string

const (
	ZoneCustomPageUpdateParamsStateDefault    ZoneCustomPageUpdateParamsState = "default"
	ZoneCustomPageUpdateParamsStateCustomized ZoneCustomPageUpdateParamsState = "customized"
)
