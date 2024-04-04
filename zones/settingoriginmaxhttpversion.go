// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingOriginMaxHTTPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingOriginMaxHTTPVersionService] method instead.
type SettingOriginMaxHTTPVersionService struct {
	Options []option.RequestOption
}

// NewSettingOriginMaxHTTPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOriginMaxHTTPVersionService(opts ...option.RequestOption) (r *SettingOriginMaxHTTPVersionService) {
	r = &SettingOriginMaxHTTPVersionService{}
	r.Options = opts
	return
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except ENT
// where it is "1"
func (r *SettingOriginMaxHTTPVersionService) Edit(ctx context.Context, params SettingOriginMaxHTTPVersionEditParams, opts ...option.RequestOption) (res *SettingOriginMaxHTTPVersionEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOriginMaxHTTPVersionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except ENT
// where it is "1"
func (r *SettingOriginMaxHTTPVersionService) Get(ctx context.Context, query SettingOriginMaxHTTPVersionGetParams, opts ...option.RequestOption) (res *SettingOriginMaxHTTPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOriginMaxHTTPVersionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except ENT
// where it is "1"
type SettingOriginMaxHTTPVersionEditResponse struct {
	// Value of the zone setting.
	ID shared.UnnamedSchemaRef142 `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Origin Max HTTP Version Setting.
	Value SettingOriginMaxHTTPVersionEditResponseValue `json:"value,required"`
	JSON  settingOriginMaxHTTPVersionEditResponseJSON  `json:"-"`
}

// settingOriginMaxHTTPVersionEditResponseJSON contains the JSON metadata for the
// struct [SettingOriginMaxHTTPVersionEditResponse]
type settingOriginMaxHTTPVersionEditResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginMaxHTTPVersionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginMaxHTTPVersionEditResponseJSON) RawJSON() string {
	return r.raw
}

// Value of the Origin Max HTTP Version Setting.
type SettingOriginMaxHTTPVersionEditResponseValue string

const (
	SettingOriginMaxHTTPVersionEditResponseValue2 SettingOriginMaxHTTPVersionEditResponseValue = "2"
	SettingOriginMaxHTTPVersionEditResponseValue1 SettingOriginMaxHTTPVersionEditResponseValue = "1"
)

func (r SettingOriginMaxHTTPVersionEditResponseValue) IsKnown() bool {
	switch r {
	case SettingOriginMaxHTTPVersionEditResponseValue2, SettingOriginMaxHTTPVersionEditResponseValue1:
		return true
	}
	return false
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except ENT
// where it is "1"
type SettingOriginMaxHTTPVersionGetResponse struct {
	// Value of the zone setting.
	ID shared.UnnamedSchemaRef142 `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Origin Max HTTP Version Setting.
	Value SettingOriginMaxHTTPVersionGetResponseValue `json:"value,required"`
	JSON  settingOriginMaxHTTPVersionGetResponseJSON  `json:"-"`
}

// settingOriginMaxHTTPVersionGetResponseJSON contains the JSON metadata for the
// struct [SettingOriginMaxHTTPVersionGetResponse]
type settingOriginMaxHTTPVersionGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginMaxHTTPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginMaxHTTPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Value of the Origin Max HTTP Version Setting.
type SettingOriginMaxHTTPVersionGetResponseValue string

const (
	SettingOriginMaxHTTPVersionGetResponseValue2 SettingOriginMaxHTTPVersionGetResponseValue = "2"
	SettingOriginMaxHTTPVersionGetResponseValue1 SettingOriginMaxHTTPVersionGetResponseValue = "1"
)

func (r SettingOriginMaxHTTPVersionGetResponseValue) IsKnown() bool {
	switch r {
	case SettingOriginMaxHTTPVersionGetResponseValue2, SettingOriginMaxHTTPVersionGetResponseValue1:
		return true
	}
	return false
}

type SettingOriginMaxHTTPVersionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Origin Max HTTP Version Setting.
	Value param.Field[SettingOriginMaxHTTPVersionEditParamsValue] `json:"value,required"`
}

func (r SettingOriginMaxHTTPVersionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Origin Max HTTP Version Setting.
type SettingOriginMaxHTTPVersionEditParamsValue string

const (
	SettingOriginMaxHTTPVersionEditParamsValue2 SettingOriginMaxHTTPVersionEditParamsValue = "2"
	SettingOriginMaxHTTPVersionEditParamsValue1 SettingOriginMaxHTTPVersionEditParamsValue = "1"
)

func (r SettingOriginMaxHTTPVersionEditParamsValue) IsKnown() bool {
	switch r {
	case SettingOriginMaxHTTPVersionEditParamsValue2, SettingOriginMaxHTTPVersionEditParamsValue1:
		return true
	}
	return false
}

type SettingOriginMaxHTTPVersionEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
	// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
	// requests to your origin. (Refer to
	// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
	// for more information.). The default value is "2" for all plan types except ENT
	// where it is "1"
	Result SettingOriginMaxHTTPVersionEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success SettingOriginMaxHTTPVersionEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingOriginMaxHTTPVersionEditResponseEnvelopeJSON    `json:"-"`
}

// settingOriginMaxHTTPVersionEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingOriginMaxHTTPVersionEditResponseEnvelope]
type settingOriginMaxHTTPVersionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginMaxHTTPVersionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginMaxHTTPVersionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingOriginMaxHTTPVersionEditResponseEnvelopeSuccess bool

const (
	SettingOriginMaxHTTPVersionEditResponseEnvelopeSuccessTrue SettingOriginMaxHTTPVersionEditResponseEnvelopeSuccess = true
)

func (r SettingOriginMaxHTTPVersionEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingOriginMaxHTTPVersionEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingOriginMaxHTTPVersionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOriginMaxHTTPVersionGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
	// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
	// requests to your origin. (Refer to
	// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
	// for more information.). The default value is "2" for all plan types except ENT
	// where it is "1"
	Result SettingOriginMaxHTTPVersionGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success SettingOriginMaxHTTPVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingOriginMaxHTTPVersionGetResponseEnvelopeJSON    `json:"-"`
}

// settingOriginMaxHTTPVersionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingOriginMaxHTTPVersionGetResponseEnvelope]
type settingOriginMaxHTTPVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginMaxHTTPVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginMaxHTTPVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingOriginMaxHTTPVersionGetResponseEnvelopeSuccess bool

const (
	SettingOriginMaxHTTPVersionGetResponseEnvelopeSuccessTrue SettingOriginMaxHTTPVersionGetResponseEnvelopeSuccess = true
)

func (r SettingOriginMaxHTTPVersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingOriginMaxHTTPVersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
