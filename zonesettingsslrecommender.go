// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ZoneSettingSSLRecommenderService) Edit(ctx context.Context, params ZoneSettingSSLRecommenderEditParams, opts ...option.RequestOption) (res *ZoneSettingSSLRecommenderEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSSLRecommenderEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *ZoneSettingSSLRecommenderService) Get(ctx context.Context, query ZoneSettingSSLRecommenderGetParams, opts ...option.RequestOption) (res *ZoneSettingSSLRecommenderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSSLRecommenderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSSLRecommenderEditResponse struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingSSLRecommenderEditResponseID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                      `json:"enabled"`
	JSON    zoneSettingSSLRecommenderEditResponseJSON `json:"-"`
}

// zoneSettingSSLRecommenderEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSSLRecommenderEditResponse]
type zoneSettingSSLRecommenderEditResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLRecommenderEditResponseJSON) RawJSON() string {
	return r.raw
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSSLRecommenderEditResponseID string

const (
	ZoneSettingSSLRecommenderEditResponseIDSSLRecommender ZoneSettingSSLRecommenderEditResponseID = "ssl_recommender"
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSSLRecommenderGetResponse struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingSSLRecommenderGetResponseID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                     `json:"enabled"`
	JSON    zoneSettingSSLRecommenderGetResponseJSON `json:"-"`
}

// zoneSettingSSLRecommenderGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSSLRecommenderGetResponse]
type zoneSettingSSLRecommenderGetResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLRecommenderGetResponseJSON) RawJSON() string {
	return r.raw
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSSLRecommenderGetResponseID string

const (
	ZoneSettingSSLRecommenderGetResponseIDSSLRecommender ZoneSettingSSLRecommenderGetResponseID = "ssl_recommender"
)

type ZoneSettingSSLRecommenderEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Value param.Field[ZoneSettingSSLRecommenderEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingSSLRecommenderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSSLRecommenderEditParamsValue struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[ZoneSettingSSLRecommenderEditParamsValueID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingSSLRecommenderEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSSLRecommenderEditParamsValueID string

const (
	ZoneSettingSSLRecommenderEditParamsValueIDSSLRecommender ZoneSettingSSLRecommenderEditParamsValueID = "ssl_recommender"
)

type ZoneSettingSSLRecommenderEditResponseEnvelope struct {
	Errors   []ZoneSettingSSLRecommenderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSSLRecommenderEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result ZoneSettingSSLRecommenderEditResponse             `json:"result"`
	JSON   zoneSettingSSLRecommenderEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSSLRecommenderEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingSSLRecommenderEditResponseEnvelope]
type zoneSettingSSLRecommenderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLRecommenderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSSLRecommenderEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingSSLRecommenderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSSLRecommenderEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingSSLRecommenderEditResponseEnvelopeErrors]
type zoneSettingSSLRecommenderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLRecommenderEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSSLRecommenderEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingSSLRecommenderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSSLRecommenderEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingSSLRecommenderEditResponseEnvelopeMessages]
type zoneSettingSSLRecommenderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLRecommenderEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSSLRecommenderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingSSLRecommenderGetResponseEnvelope struct {
	Errors   []ZoneSettingSSLRecommenderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSSLRecommenderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result ZoneSettingSSLRecommenderGetResponse             `json:"result"`
	JSON   zoneSettingSSLRecommenderGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSSLRecommenderGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingSSLRecommenderGetResponseEnvelope]
type zoneSettingSSLRecommenderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLRecommenderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSSLRecommenderGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingSSLRecommenderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSSLRecommenderGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingSSLRecommenderGetResponseEnvelopeErrors]
type zoneSettingSSLRecommenderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLRecommenderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSSLRecommenderGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingSSLRecommenderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSSLRecommenderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingSSLRecommenderGetResponseEnvelopeMessages]
type zoneSettingSSLRecommenderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLRecommenderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLRecommenderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
