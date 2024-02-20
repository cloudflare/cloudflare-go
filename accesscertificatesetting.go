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

// List all mTLS hostname settings for this account or zone.
func (r *AccessCertificateSettingService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessCertificateSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateSettingListResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an mTLS certificate's hostname settings.
func (r *AccessCertificateSettingService) Replace(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessCertificateSettingReplaceParams, opts ...option.RequestOption) (res *[]AccessCertificateSettingReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateSettingReplaceResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessCertificateSettingListResponse struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                                   `json:"hostname,required"`
	JSON     accessCertificateSettingListResponseJSON `json:"-"`
}

// accessCertificateSettingListResponseJSON contains the JSON metadata for the
// struct [AccessCertificateSettingListResponse]
type accessCertificateSettingListResponseJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingReplaceResponse struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                                      `json:"hostname,required"`
	JSON     accessCertificateSettingReplaceResponseJSON `json:"-"`
}

// accessCertificateSettingReplaceResponseJSON contains the JSON metadata for the
// struct [AccessCertificateSettingReplaceResponse]
type accessCertificateSettingReplaceResponseJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccessCertificateSettingReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingListResponseEnvelope struct {
	Errors   []AccessCertificateSettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateSettingListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessCertificateSettingListResponse                 `json:"result,required,nullable"`
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

type AccessCertificateSettingReplaceParams struct {
	Settings param.Field[[]AccessCertificateSettingReplaceParamsSetting] `json:"settings,required"`
}

func (r AccessCertificateSettingReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingReplaceParamsSetting struct {
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

func (r AccessCertificateSettingReplaceParamsSetting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingReplaceResponseEnvelope struct {
	Errors   []AccessCertificateSettingReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateSettingReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessCertificateSettingReplaceResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessCertificateSettingReplaceResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessCertificateSettingReplaceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessCertificateSettingReplaceResponseEnvelopeJSON       `json:"-"`
}

// accessCertificateSettingReplaceResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessCertificateSettingReplaceResponseEnvelope]
type accessCertificateSettingReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingReplaceResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accessCertificateSettingReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateSettingReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessCertificateSettingReplaceResponseEnvelopeErrors]
type accessCertificateSettingReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingReplaceResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accessCertificateSettingReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateSettingReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AccessCertificateSettingReplaceResponseEnvelopeMessages]
type accessCertificateSettingReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateSettingReplaceResponseEnvelopeSuccess bool

const (
	AccessCertificateSettingReplaceResponseEnvelopeSuccessTrue AccessCertificateSettingReplaceResponseEnvelopeSuccess = true
)

type AccessCertificateSettingReplaceResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       accessCertificateSettingReplaceResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessCertificateSettingReplaceResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AccessCertificateSettingReplaceResponseEnvelopeResultInfo]
type accessCertificateSettingReplaceResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingReplaceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
