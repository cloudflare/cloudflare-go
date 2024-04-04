// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_client_auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	return
}

// Enable or disable zone-level authenticated origin pulls. 'enabled' should be set
// true either before/after the certificate is uploaded to see the certificate in
// use.
func (r *SettingService) Update(ctx context.Context, params SettingUpdateParams, opts ...option.RequestOption) (res *SettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get whether zone-level authenticated origin pulls is enabled or not. It is false
// by default.
func (r *SettingService) Get(ctx context.Context, query SettingGetParams, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingUpdateResponse struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                      `json:"enabled"`
	JSON    settingUpdateResponseJSON `json:"-"`
}

// settingUpdateResponseJSON contains the JSON metadata for the struct
// [SettingUpdateResponse]
type settingUpdateResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponse struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                   `json:"enabled"`
	JSON    settingGetResponseJSON `json:"-"`
}

// settingGetResponseJSON contains the JSON metadata for the struct
// [SettingGetResponse]
type settingGetResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r SettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   SettingUpdateResponse        `json:"result,required"`
	// Whether the API call was successful
	Success SettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingUpdateResponseEnvelopeJSON    `json:"-"`
}

// settingUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingUpdateResponseEnvelope]
type settingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingUpdateResponseEnvelopeSuccess bool

const (
	SettingUpdateResponseEnvelopeSuccessTrue SettingUpdateResponseEnvelopeSuccess = true
)

func (r SettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   SettingGetResponse           `json:"result,required"`
	// Whether the API call was successful
	Success SettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingGetResponseEnvelopeJSON    `json:"-"`
}

// settingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelope]
type settingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingGetResponseEnvelopeSuccess bool

const (
	SettingGetResponseEnvelopeSuccessTrue SettingGetResponseEnvelopeSuccess = true
)

func (r SettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
