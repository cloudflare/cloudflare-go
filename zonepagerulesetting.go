// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZonePageruleSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZonePageruleSettingService]
// method instead.
type ZonePageruleSettingService struct {
	Options []option.RequestOption
}

// NewZonePageruleSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePageruleSettingService(opts ...option.RequestOption) (r *ZonePageruleSettingService) {
	r = &ZonePageruleSettingService{}
	r.Options = opts
	return
}

// Returns a list of settings (and their details) that Page Rules can apply to
// matching requests.
func (r *ZonePageruleSettingService) AvailablePageRulesSettingsListAvailablePageRulesSettings(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *PageruleSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/pagerules/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PageruleSettingsResponseCollection struct {
	Errors   []PageruleSettingsResponseCollectionError   `json:"errors"`
	Messages []PageruleSettingsResponseCollectionMessage `json:"messages"`
	// Settings available for the zone.
	Result []interface{} `json:"result"`
	// Whether the API call was successful
	Success PageruleSettingsResponseCollectionSuccess `json:"success"`
	JSON    pageruleSettingsResponseCollectionJSON    `json:"-"`
}

// pageruleSettingsResponseCollectionJSON contains the JSON metadata for the struct
// [PageruleSettingsResponseCollection]
type pageruleSettingsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleSettingsResponseCollectionError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    pageruleSettingsResponseCollectionErrorJSON `json:"-"`
}

// pageruleSettingsResponseCollectionErrorJSON contains the JSON metadata for the
// struct [PageruleSettingsResponseCollectionError]
type pageruleSettingsResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingsResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleSettingsResponseCollectionMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    pageruleSettingsResponseCollectionMessageJSON `json:"-"`
}

// pageruleSettingsResponseCollectionMessageJSON contains the JSON metadata for the
// struct [PageruleSettingsResponseCollectionMessage]
type pageruleSettingsResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingsResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageruleSettingsResponseCollectionSuccess bool

const (
	PageruleSettingsResponseCollectionSuccessTrue PageruleSettingsResponseCollectionSuccess = true
)
