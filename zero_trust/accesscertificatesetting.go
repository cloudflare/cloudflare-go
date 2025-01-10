// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AccessCertificateSettingService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessCertificateSettingService] method instead.
type AccessCertificateSettingService struct {
	Options []option.RequestOption
}

// NewAccessCertificateSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessCertificateSettingService(opts ...option.RequestOption) (r *AccessCertificateSettingService) {
	r = &AccessCertificateSettingService{}
	r.Options = opts
	return
}

// Updates an mTLS certificate's hostname settings.
func (r *AccessCertificateSettingService) Update(ctx context.Context, params AccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *[]CertificateSettings, err error) {
	var env AccessCertificateSettingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all mTLS hostname settings for this account or zone.
func (r *AccessCertificateSettingService) Get(ctx context.Context, query AccessCertificateSettingGetParams, opts ...option.RequestOption) (res *[]CertificateSettings, err error) {
	var env AccessCertificateSettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CertificateSettings struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                  `json:"hostname,required"`
	JSON     certificateSettingsJSON `json:"-"`
}

// certificateSettingsJSON contains the JSON metadata for the struct
// [CertificateSettings]
type certificateSettingsJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CertificateSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateSettingsJSON) RawJSON() string {
	return r.raw
}

type CertificateSettingsParam struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork param.Field[bool] `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding param.Field[bool] `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname param.Field[string] `json:"hostname,required"`
}

func (r CertificateSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]CertificateSettingsParam] `json:"settings,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AccessCertificateSettingUpdateResponseEnvelopeSuccess    `json:"success,required"`
	Result     []CertificateSettings                                    `json:"result"`
	ResultInfo AccessCertificateSettingUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessCertificateSettingUpdateResponseEnvelopeJSON       `json:"-"`
}

// accessCertificateSettingUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessCertificateSettingUpdateResponseEnvelope]
type accessCertificateSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateSettingUpdateResponseEnvelopeSuccess bool

const (
	AccessCertificateSettingUpdateResponseEnvelopeSuccessTrue AccessCertificateSettingUpdateResponseEnvelopeSuccess = true
)

func (r AccessCertificateSettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateSettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCertificateSettingUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       accessCertificateSettingUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessCertificateSettingUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AccessCertificateSettingUpdateResponseEnvelopeResultInfo]
type accessCertificateSettingUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateSettingGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessCertificateSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AccessCertificateSettingGetResponseEnvelopeSuccess    `json:"success,required"`
	Result     []CertificateSettings                                 `json:"result"`
	ResultInfo AccessCertificateSettingGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessCertificateSettingGetResponseEnvelopeJSON       `json:"-"`
}

// accessCertificateSettingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessCertificateSettingGetResponseEnvelope]
type accessCertificateSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateSettingGetResponseEnvelopeSuccess bool

const (
	AccessCertificateSettingGetResponseEnvelopeSuccessTrue AccessCertificateSettingGetResponseEnvelopeSuccess = true
)

func (r AccessCertificateSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCertificateSettingGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       accessCertificateSettingGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessCertificateSettingGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AccessCertificateSettingGetResponseEnvelopeResultInfo]
type accessCertificateSettingGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
