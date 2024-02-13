// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SSLRecommendationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSSLRecommendationService] method
// instead.
type SSLRecommendationService struct {
	Options []option.RequestOption
}

// NewSSLRecommendationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSSLRecommendationService(opts ...option.RequestOption) (r *SSLRecommendationService) {
	r = &SSLRecommendationService{}
	r.Options = opts
	return
}

// Retrieve the SSL/TLS Recommender's recommendation for a zone.
func (r *SSLRecommendationService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SSLRecommendationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLRecommendationListResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/recommendation", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLRecommendationListResponse struct {
	// Identifier of a recommedation result.
	ID         string                             `json:"id"`
	ModifiedOn time.Time                          `json:"modified_on" format:"date-time"`
	Value      SSLRecommendationListResponseValue `json:"value"`
	JSON       sslRecommendationListResponseJSON  `json:"-"`
}

// sslRecommendationListResponseJSON contains the JSON metadata for the struct
// [SSLRecommendationListResponse]
type sslRecommendationListResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommendationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLRecommendationListResponseValue string

const (
	SSLRecommendationListResponseValueFlexible SSLRecommendationListResponseValue = "flexible"
	SSLRecommendationListResponseValueFull     SSLRecommendationListResponseValue = "full"
	SSLRecommendationListResponseValueStrict   SSLRecommendationListResponseValue = "strict"
)

type SSLRecommendationListResponseEnvelope struct {
	Errors   []SSLRecommendationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLRecommendationListResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLRecommendationListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SSLRecommendationListResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslRecommendationListResponseEnvelopeJSON    `json:"-"`
}

// sslRecommendationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLRecommendationListResponseEnvelope]
type sslRecommendationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommendationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLRecommendationListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    sslRecommendationListResponseEnvelopeErrorsJSON `json:"-"`
}

// sslRecommendationListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SSLRecommendationListResponseEnvelopeErrors]
type sslRecommendationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommendationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLRecommendationListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    sslRecommendationListResponseEnvelopeMessagesJSON `json:"-"`
}

// sslRecommendationListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SSLRecommendationListResponseEnvelopeMessages]
type sslRecommendationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommendationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLRecommendationListResponseEnvelopeSuccess bool

const (
	SSLRecommendationListResponseEnvelopeSuccessTrue SSLRecommendationListResponseEnvelopeSuccess = true
)
