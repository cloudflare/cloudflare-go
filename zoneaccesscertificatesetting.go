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

// ZoneAccessCertificateSettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneAccessCertificateSettingService] method instead.
type ZoneAccessCertificateSettingService struct {
	Options []option.RequestOption
}

// NewZoneAccessCertificateSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAccessCertificateSettingService(opts ...option.RequestOption) (r *ZoneAccessCertificateSettingService) {
	r = &ZoneAccessCertificateSettingService{}
	r.Options = opts
	return
}

// Updates an mTLS certificates hostname settings.
func (r *ZoneAccessCertificateSettingService) Update(ctx context.Context, identifier string, uuid string, body ZoneAccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *SingleResponseHostname, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/%s/settings", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all mTLS hostname settings for this zone.
func (r *ZoneAccessCertificateSettingService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ResponseCollectionHostname, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseCollectionHostname struct {
	Errors     []ResponseCollectionHostnameError    `json:"errors"`
	Messages   []ResponseCollectionHostnameMessage  `json:"messages"`
	Result     []ResponseCollectionHostnameResult   `json:"result"`
	ResultInfo ResponseCollectionHostnameResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionHostnameSuccess `json:"success"`
	JSON    responseCollectionHostnameJSON    `json:"-"`
}

// responseCollectionHostnameJSON contains the JSON metadata for the struct
// [ResponseCollectionHostname]
type responseCollectionHostnameJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionHostnameError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionHostnameErrorJSON `json:"-"`
}

// responseCollectionHostnameErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionHostnameError]
type responseCollectionHostnameErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionHostnameError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionHostnameMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionHostnameMessageJSON `json:"-"`
}

// responseCollectionHostnameMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionHostnameMessage]
type responseCollectionHostnameMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionHostnameMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionHostnameResult struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                               `json:"hostname,required"`
	JSON     responseCollectionHostnameResultJSON `json:"-"`
}

// responseCollectionHostnameResultJSON contains the JSON metadata for the struct
// [ResponseCollectionHostnameResult]
type responseCollectionHostnameResultJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ResponseCollectionHostnameResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionHostnameResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionHostnameResultInfoJSON `json:"-"`
}

// responseCollectionHostnameResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionHostnameResultInfo]
type responseCollectionHostnameResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionHostnameResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionHostnameSuccess bool

const (
	ResponseCollectionHostnameSuccessTrue ResponseCollectionHostnameSuccess = true
)

type SingleResponseHostname struct {
	Errors   []SingleResponseHostnameError   `json:"errors"`
	Messages []SingleResponseHostnameMessage `json:"messages"`
	Result   SingleResponseHostnameResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseHostnameSuccess `json:"success"`
	JSON    singleResponseHostnameJSON    `json:"-"`
}

// singleResponseHostnameJSON contains the JSON metadata for the struct
// [SingleResponseHostname]
type singleResponseHostnameJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseHostnameError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseHostnameErrorJSON `json:"-"`
}

// singleResponseHostnameErrorJSON contains the JSON metadata for the struct
// [SingleResponseHostnameError]
type singleResponseHostnameErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseHostnameError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseHostnameMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseHostnameMessageJSON `json:"-"`
}

// singleResponseHostnameMessageJSON contains the JSON metadata for the struct
// [SingleResponseHostnameMessage]
type singleResponseHostnameMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseHostnameMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseHostnameResult struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                           `json:"hostname,required"`
	JSON     singleResponseHostnameResultJSON `json:"-"`
}

// singleResponseHostnameResultJSON contains the JSON metadata for the struct
// [SingleResponseHostnameResult]
type singleResponseHostnameResultJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *SingleResponseHostnameResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleResponseHostnameSuccess bool

const (
	SingleResponseHostnameSuccessTrue SingleResponseHostnameSuccess = true
)

type ZoneAccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]ZoneAccessCertificateSettingUpdateParamsSetting] `json:"settings,required"`
}

func (r ZoneAccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessCertificateSettingUpdateParamsSetting struct {
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

func (r ZoneAccessCertificateSettingUpdateParamsSetting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
