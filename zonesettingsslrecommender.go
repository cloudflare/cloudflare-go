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

// ZoneSettingSSLRecommenderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingSSLRecommenderService] method instead.
type ZoneSettingSSLRecommenderService struct {
	Options []option.RequestOption
}

// NewZoneSettingSSLRecommenderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSSLRecommenderService(opts ...option.RequestOption) (r *ZoneSettingSSLRecommenderService) {
	r = &ZoneSettingSSLRecommenderService{}
	r.Options = opts
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *ZoneSettingSSLRecommenderService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingSSLRecommenderUpdateParams, opts ...option.RequestOption) (res *ZoneSettingSSLRecommenderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *ZoneSettingSSLRecommenderService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingSSLRecommenderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingSSLRecommenderUpdateResponse struct {
	Errors   []ZoneSettingSSLRecommenderUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingSSLRecommenderUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result ZoneSettingSSLRecommenderUpdateResponseResult `json:"result"`
	JSON   zoneSettingSSLRecommenderUpdateResponseJSON   `json:"-"`
}

// zoneSettingSSLRecommenderUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSSLRecommenderUpdateResponse]
type zoneSettingSSLRecommenderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLRecommenderUpdateResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingSSLRecommenderUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingSSLRecommenderUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingSSLRecommenderUpdateResponseError]
type zoneSettingSSLRecommenderUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLRecommenderUpdateResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingSSLRecommenderUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingSSLRecommenderUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingSSLRecommenderUpdateResponseMessage]
type zoneSettingSSLRecommenderUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSSLRecommenderUpdateResponseResult struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingSSLRecommenderUpdateResponseResultID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                              `json:"enabled"`
	JSON    zoneSettingSSLRecommenderUpdateResponseResultJSON `json:"-"`
}

// zoneSettingSSLRecommenderUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingSSLRecommenderUpdateResponseResult]
type zoneSettingSSLRecommenderUpdateResponseResultJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSSLRecommenderUpdateResponseResultID string

const (
	ZoneSettingSSLRecommenderUpdateResponseResultIDSSLRecommender ZoneSettingSSLRecommenderUpdateResponseResultID = "ssl_recommender"
)

type ZoneSettingSSLRecommenderListResponse struct {
	Errors   []ZoneSettingSSLRecommenderListResponseError   `json:"errors,required"`
	Messages []ZoneSettingSSLRecommenderListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result ZoneSettingSSLRecommenderListResponseResult `json:"result"`
	JSON   zoneSettingSSLRecommenderListResponseJSON   `json:"-"`
}

// zoneSettingSSLRecommenderListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSSLRecommenderListResponse]
type zoneSettingSSLRecommenderListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLRecommenderListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingSSLRecommenderListResponseErrorJSON `json:"-"`
}

// zoneSettingSSLRecommenderListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingSSLRecommenderListResponseError]
type zoneSettingSSLRecommenderListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLRecommenderListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingSSLRecommenderListResponseMessageJSON `json:"-"`
}

// zoneSettingSSLRecommenderListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingSSLRecommenderListResponseMessage]
type zoneSettingSSLRecommenderListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSSLRecommenderListResponseResult struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingSSLRecommenderListResponseResultID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                            `json:"enabled"`
	JSON    zoneSettingSSLRecommenderListResponseResultJSON `json:"-"`
}

// zoneSettingSSLRecommenderListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingSSLRecommenderListResponseResult]
type zoneSettingSSLRecommenderListResponseResultJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSSLRecommenderListResponseResultID string

const (
	ZoneSettingSSLRecommenderListResponseResultIDSSLRecommender ZoneSettingSSLRecommenderListResponseResultID = "ssl_recommender"
)

type ZoneSettingSSLRecommenderUpdateParams struct {
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Value param.Field[ZoneSettingSSLRecommenderUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSSLRecommenderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSSLRecommenderUpdateParamsValue struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[ZoneSettingSSLRecommenderUpdateParamsValueID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingSSLRecommenderUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSSLRecommenderUpdateParamsValueID string

const (
	ZoneSettingSSLRecommenderUpdateParamsValueIDSSLRecommender ZoneSettingSSLRecommenderUpdateParamsValueID = "ssl_recommender"
)
