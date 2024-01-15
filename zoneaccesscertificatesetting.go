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

// List all mTLS hostname settings for this zone.
func (r *ZoneAccessCertificateSettingService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneAccessCertificateSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAccessCertificateSettingListResponse struct {
	Errors     []ZoneAccessCertificateSettingListResponseError    `json:"errors"`
	Messages   []ZoneAccessCertificateSettingListResponseMessage  `json:"messages"`
	Result     []ZoneAccessCertificateSettingListResponseResult   `json:"result"`
	ResultInfo ZoneAccessCertificateSettingListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneAccessCertificateSettingListResponseSuccess `json:"success"`
	JSON    zoneAccessCertificateSettingListResponseJSON    `json:"-"`
}

// zoneAccessCertificateSettingListResponseJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateSettingListResponse]
type zoneAccessCertificateSettingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateSettingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateSettingListResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneAccessCertificateSettingListResponseErrorJSON `json:"-"`
}

// zoneAccessCertificateSettingListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneAccessCertificateSettingListResponseError]
type zoneAccessCertificateSettingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateSettingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateSettingListResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneAccessCertificateSettingListResponseMessageJSON `json:"-"`
}

// zoneAccessCertificateSettingListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneAccessCertificateSettingListResponseMessage]
type zoneAccessCertificateSettingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateSettingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateSettingListResponseResult struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                                             `json:"hostname,required"`
	JSON     zoneAccessCertificateSettingListResponseResultJSON `json:"-"`
}

// zoneAccessCertificateSettingListResponseResultJSON contains the JSON metadata
// for the struct [ZoneAccessCertificateSettingListResponseResult]
type zoneAccessCertificateSettingListResponseResultJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ZoneAccessCertificateSettingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateSettingListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       zoneAccessCertificateSettingListResponseResultInfoJSON `json:"-"`
}

// zoneAccessCertificateSettingListResponseResultInfoJSON contains the JSON
// metadata for the struct [ZoneAccessCertificateSettingListResponseResultInfo]
type zoneAccessCertificateSettingListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateSettingListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessCertificateSettingListResponseSuccess bool

const (
	ZoneAccessCertificateSettingListResponseSuccessTrue ZoneAccessCertificateSettingListResponseSuccess = true
)
