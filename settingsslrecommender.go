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
func (r *SettingSSLRecommenderService) Update(ctx context.Context, zoneID string, body SettingSSLRecommenderUpdateParams, opts ...option.RequestOption) (res *SettingSSLRecommenderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLRecommenderUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *SettingSSLRecommenderService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingSSLRecommenderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLRecommenderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SettingSSLRecommenderUpdateResponse struct {
	// Enrollment value for SSL/TLS Recommender.
	ID SettingSSLRecommenderUpdateResponseID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                    `json:"enabled"`
	JSON    settingSSLRecommenderUpdateResponseJSON `json:"-"`
}

// settingSSLRecommenderUpdateResponseJSON contains the JSON metadata for the
// struct [SettingSSLRecommenderUpdateResponse]
type settingSSLRecommenderUpdateResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment value for SSL/TLS Recommender.
type SettingSSLRecommenderUpdateResponseID string

const (
	SettingSSLRecommenderUpdateResponseIDSSLRecommender SettingSSLRecommenderUpdateResponseID = "ssl_recommender"
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

type SettingSSLRecommenderUpdateParams struct {
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Value param.Field[SettingSSLRecommenderUpdateParamsValue] `json:"value,required"`
}

func (r SettingSSLRecommenderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SettingSSLRecommenderUpdateParamsValue struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[SettingSSLRecommenderUpdateParamsValueID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingSSLRecommenderUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment value for SSL/TLS Recommender.
type SettingSSLRecommenderUpdateParamsValueID string

const (
	SettingSSLRecommenderUpdateParamsValueIDSSLRecommender SettingSSLRecommenderUpdateParamsValueID = "ssl_recommender"
)

type SettingSSLRecommenderUpdateResponseEnvelope struct {
	Errors   []SettingSSLRecommenderUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSSLRecommenderUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result SettingSSLRecommenderUpdateResponse             `json:"result"`
	JSON   settingSSLRecommenderUpdateResponseEnvelopeJSON `json:"-"`
}

// settingSSLRecommenderUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingSSLRecommenderUpdateResponseEnvelope]
type settingSSLRecommenderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLRecommenderUpdateResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingSSLRecommenderUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSSLRecommenderUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSSLRecommenderUpdateResponseEnvelopeErrors]
type settingSSLRecommenderUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLRecommenderUpdateResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingSSLRecommenderUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSSLRecommenderUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingSSLRecommenderUpdateResponseEnvelopeMessages]
type settingSSLRecommenderUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLRecommenderUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
