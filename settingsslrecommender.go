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

// SettingSSLRecommenderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingSSLRecommenderService]
// method instead.
type SettingSSLRecommenderService struct {
	Options []option.RequestOption
}

// NewSettingSSLRecommenderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingSSLRecommenderService(opts ...option.RequestOption) (r *SettingSSLRecommenderService) {
	r = &SettingSSLRecommenderService{}
	r.Options = opts
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *SettingSSLRecommenderService) Edit(ctx context.Context, params SettingSSLRecommenderEditParams, opts ...option.RequestOption) (res *SettingSSLRecommenderEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLRecommenderEditResponseEnvelope
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
func (r *SettingSSLRecommenderService) Get(ctx context.Context, query SettingSSLRecommenderGetParams, opts ...option.RequestOption) (res *SettingSSLRecommenderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLRecommenderGetResponseEnvelope
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
type SettingSSLRecommenderEditResponse struct {
	// Enrollment value for SSL/TLS Recommender.
	ID SettingSSLRecommenderEditResponseID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                  `json:"enabled"`
	JSON    settingSSLRecommenderEditResponseJSON `json:"-"`
}

// settingSSLRecommenderEditResponseJSON contains the JSON metadata for the struct
// [SettingSSLRecommenderEditResponse]
type settingSSLRecommenderEditResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment value for SSL/TLS Recommender.
type SettingSSLRecommenderEditResponseID string

const (
	SettingSSLRecommenderEditResponseIDSSLRecommender SettingSSLRecommenderEditResponseID = "ssl_recommender"
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SettingSSLRecommenderGetResponse struct {
	// Enrollment value for SSL/TLS Recommender.
	ID SettingSSLRecommenderGetResponseID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                 `json:"enabled"`
	JSON    settingSSLRecommenderGetResponseJSON `json:"-"`
}

// settingSSLRecommenderGetResponseJSON contains the JSON metadata for the struct
// [SettingSSLRecommenderGetResponse]
type settingSSLRecommenderGetResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment value for SSL/TLS Recommender.
type SettingSSLRecommenderGetResponseID string

const (
	SettingSSLRecommenderGetResponseIDSSLRecommender SettingSSLRecommenderGetResponseID = "ssl_recommender"
)

type SettingSSLRecommenderEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Value param.Field[SettingSSLRecommenderEditParamsValue] `json:"value,required"`
}

func (r SettingSSLRecommenderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SettingSSLRecommenderEditParamsValue struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[SettingSSLRecommenderEditParamsValueID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingSSLRecommenderEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment value for SSL/TLS Recommender.
type SettingSSLRecommenderEditParamsValueID string

const (
	SettingSSLRecommenderEditParamsValueIDSSLRecommender SettingSSLRecommenderEditParamsValueID = "ssl_recommender"
)

type SettingSSLRecommenderEditResponseEnvelope struct {
	Errors   []SettingSSLRecommenderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSSLRecommenderEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result SettingSSLRecommenderEditResponse             `json:"result"`
	JSON   settingSSLRecommenderEditResponseEnvelopeJSON `json:"-"`
}

// settingSSLRecommenderEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSSLRecommenderEditResponseEnvelope]
type settingSSLRecommenderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLRecommenderEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingSSLRecommenderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSSLRecommenderEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSSLRecommenderEditResponseEnvelopeErrors]
type settingSSLRecommenderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLRecommenderEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingSSLRecommenderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSSLRecommenderEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingSSLRecommenderEditResponseEnvelopeMessages]
type settingSSLRecommenderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLRecommenderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingSSLRecommenderGetResponseEnvelope struct {
	Errors   []SettingSSLRecommenderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSSLRecommenderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result SettingSSLRecommenderGetResponse             `json:"result"`
	JSON   settingSSLRecommenderGetResponseEnvelopeJSON `json:"-"`
}

// settingSSLRecommenderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSSLRecommenderGetResponseEnvelope]
type settingSSLRecommenderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLRecommenderGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingSSLRecommenderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSSLRecommenderGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSSLRecommenderGetResponseEnvelopeErrors]
type settingSSLRecommenderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLRecommenderGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingSSLRecommenderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSSLRecommenderGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingSSLRecommenderGetResponseEnvelopeMessages]
type settingSSLRecommenderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
