// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingSSLRecommenderService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingSSLRecommenderService] method instead.
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
func (r *SettingSSLRecommenderService) Edit(ctx context.Context, params SettingSSLRecommenderEditParams, opts ...option.RequestOption) (res *SSLRecommender, err error) {
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
func (r *SettingSSLRecommenderService) Get(ctx context.Context, query SettingSSLRecommenderGetParams, opts ...option.RequestOption) (res *SSLRecommender, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLRecommenderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", query.ZoneID)
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
type SSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID SSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool               `json:"enabled"`
	JSON    sslRecommenderJSON `json:"-"`
}

// sslRecommenderJSON contains the JSON metadata for the struct [SSLRecommender]
type sslRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslRecommenderJSON) RawJSON() string {
	return r.raw
}

// Enrollment value for SSL/TLS Recommender.
type SSLRecommenderID string

const (
	SSLRecommenderIDSSLRecommender SSLRecommenderID = "ssl_recommender"
)

func (r SSLRecommenderID) IsKnown() bool {
	switch r {
	case SSLRecommenderIDSSLRecommender:
		return true
	}
	return false
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SSLRecommenderParam struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[SSLRecommenderID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SSLRecommenderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingSSLRecommenderEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Value param.Field[SSLRecommenderParam] `json:"value,required"`
}

func (r SettingSSLRecommenderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingSSLRecommenderEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result SSLRecommender                                `json:"result"`
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

func (r settingSSLRecommenderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSSLRecommenderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingSSLRecommenderGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result SSLRecommender                               `json:"result"`
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

func (r settingSSLRecommenderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
