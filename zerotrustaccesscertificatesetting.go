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

// ZeroTrustAccessCertificateSettingService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustAccessCertificateSettingService] method instead.
type ZeroTrustAccessCertificateSettingService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessCertificateSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessCertificateSettingService(opts ...option.RequestOption) (r *ZeroTrustAccessCertificateSettingService) {
	r = &ZeroTrustAccessCertificateSettingService{}
	r.Options = opts
	return
}

// Updates an mTLS certificate's hostname settings.
func (r *ZeroTrustAccessCertificateSettingService) Update(ctx context.Context, params ZeroTrustAccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *[]AccessSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCertificateSettingUpdateResponseEnvelope
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
func (r *ZeroTrustAccessCertificateSettingService) List(ctx context.Context, query ZeroTrustAccessCertificateSettingListParams, opts ...option.RequestOption) (res *[]AccessSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCertificateSettingListResponseEnvelope
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

type ZeroTrustAccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]AccessSettingsParam] `json:"settings,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r ZeroTrustAccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessCertificateSettingUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessSettings                                                  `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessCertificateSettingUpdateResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessCertificateSettingUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessCertificateSettingUpdateResponseEnvelope]
type zeroTrustAccessCertificateSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustAccessCertificateSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCertificateSettingUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeErrors]
type zeroTrustAccessCertificateSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zeroTrustAccessCertificateSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCertificateSettingUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeMessages]
type zeroTrustAccessCertificateSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       zeroTrustAccessCertificateSettingUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessCertificateSettingUpdateResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeResultInfo]
type zeroTrustAccessCertificateSettingUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateSettingUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateSettingListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessCertificateSettingListResponseEnvelope struct {
	Errors   []ZeroTrustAccessCertificateSettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCertificateSettingListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessSettings                                                `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessCertificateSettingListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessCertificateSettingListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessCertificateSettingListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessCertificateSettingListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateSettingListResponseEnvelope]
type zeroTrustAccessCertificateSettingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateSettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateSettingListResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustAccessCertificateSettingListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCertificateSettingListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessCertificateSettingListResponseEnvelopeErrors]
type zeroTrustAccessCertificateSettingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateSettingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateSettingListResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustAccessCertificateSettingListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCertificateSettingListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessCertificateSettingListResponseEnvelopeMessages]
type zeroTrustAccessCertificateSettingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateSettingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessCertificateSettingListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCertificateSettingListResponseEnvelopeSuccessTrue ZeroTrustAccessCertificateSettingListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCertificateSettingListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                             `json:"total_count"`
	JSON       zeroTrustAccessCertificateSettingListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessCertificateSettingListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessCertificateSettingListResponseEnvelopeResultInfo]
type zeroTrustAccessCertificateSettingListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateSettingListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
