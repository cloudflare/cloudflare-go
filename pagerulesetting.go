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
func (r *PageruleSettingService) List(ctx context.Context, query PageruleSettingListParams, opts ...option.RequestOption) (res *[]PageruleSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleSettingListResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageruleSettingListResponse = interface{}

type PageruleSettingListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleSettingListResponseEnvelope struct {
	Errors   []PageruleSettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageruleSettingListResponseEnvelopeMessages `json:"messages,required"`
	// Settings available for the zone.
	Result []PageruleSettingListResponse `json:"result,required"`
	// Whether the API call was successful
	Success PageruleSettingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageruleSettingListResponseEnvelopeJSON    `json:"-"`
}

// pageruleSettingListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageruleSettingListResponseEnvelope]
type pageruleSettingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleSettingListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageruleSettingListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    pageruleSettingListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageruleSettingListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageruleSettingListResponseEnvelopeErrors]
type pageruleSettingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleSettingListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageruleSettingListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    pageruleSettingListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageruleSettingListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageruleSettingListResponseEnvelopeMessages]
type pageruleSettingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleSettingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleSettingListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleSettingListResponseEnvelopeSuccess bool

const (
	PageruleSettingListResponseEnvelopeSuccessTrue PageruleSettingListResponseEnvelopeSuccess = true
)
