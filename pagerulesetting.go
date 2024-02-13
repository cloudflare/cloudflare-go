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

// PageruleSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageruleSettingService] method
// instead.
type PageruleSettingService struct {
	Options []option.RequestOption
}

// NewPageruleSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageruleSettingService(opts ...option.RequestOption) (r *PageruleSettingService) {
	r = &PageruleSettingService{}
	r.Options = opts
	return
}

// Returns a list of settings (and their details) that Page Rules can apply to
// matching requests.
func (r *PageruleSettingService) AvailablePageRulesSettingsListAvailablePageRulesSettings(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponse = interface{}

type PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelope struct {
	Errors   []PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeMessages `json:"messages,required"`
	// Settings available for the zone.
	Result []PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponse `json:"result,required"`
	// Whether the API call was successful
	Success PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeJSON    `json:"-"`
}

// pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelope]
type pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeErrors struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeErrorsJSON `json:"-"`
}

// pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeErrors]
type pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeMessages struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeMessagesJSON `json:"-"`
}

// pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeMessages]
type pageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeSuccess bool

const (
	PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeSuccessTrue PageruleSettingAvailablePageRulesSettingsListAvailablePageRulesSettingsResponseEnvelopeSuccess = true
)
