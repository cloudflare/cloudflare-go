// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// AccessCertificateSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccessCertificateSettingService] method instead.
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
func (r *AccessCertificateSettingService) Update(ctx context.Context, params AccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *[]AccessSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateSettingUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
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
func (r *AccessCertificateSettingService) List(ctx context.Context, query AccessCertificateSettingListParams, opts ...option.RequestOption) (res *[]AccessSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateSettingListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessSettings struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string             `json:"hostname,required"`
	JSON     accessSettingsJSON `json:"-"`
}

// accessSettingsJSON contains the JSON metadata for the struct [AccessSettings]
type accessSettingsJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccessSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSettingsJSON) RawJSON() string {
	return r.raw
}

type AccessSettingsParam struct {
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

func (r AccessSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]AccessSettingsParam] `json:"settings,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingUpdateResponseEnvelope struct {
	Errors   []AccessCertificateSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessSettings                                         `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessCertificateSettingUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessCertificateSettingUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessCertificateSettingUpdateResponseEnvelopeJSON       `json:"-"`
}

// accessCertificateSettingUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessCertificateSettingUpdateResponseEnvelope]
type accessCertificateSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

type AccessCertificateSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessCertificateSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateSettingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessCertificateSettingUpdateResponseEnvelopeErrors]
type accessCertificateSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accessCertificateSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateSettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessCertificateSettingUpdateResponseEnvelopeMessages]
type accessCertificateSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateSettingUpdateResponseEnvelopeSuccess bool

const (
	AccessCertificateSettingUpdateResponseEnvelopeSuccessTrue AccessCertificateSettingUpdateResponseEnvelopeSuccess = true
)

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

type AccessCertificateSettingListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessCertificateSettingListResponseEnvelope struct {
	Errors   []AccessCertificateSettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateSettingListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessSettings                                       `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessCertificateSettingListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessCertificateSettingListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessCertificateSettingListResponseEnvelopeJSON       `json:"-"`
}

// accessCertificateSettingListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessCertificateSettingListResponseEnvelope]
type accessCertificateSettingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateSettingListResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accessCertificateSettingListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateSettingListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessCertificateSettingListResponseEnvelopeErrors]
type accessCertificateSettingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateSettingListResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessCertificateSettingListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateSettingListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessCertificateSettingListResponseEnvelopeMessages]
type accessCertificateSettingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateSettingListResponseEnvelopeSuccess bool

const (
	AccessCertificateSettingListResponseEnvelopeSuccessTrue AccessCertificateSettingListResponseEnvelopeSuccess = true
)

type AccessCertificateSettingListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       accessCertificateSettingListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessCertificateSettingListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AccessCertificateSettingListResponseEnvelopeResultInfo]
type accessCertificateSettingListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateSettingListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
