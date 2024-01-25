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

// ZoneSslRecommendationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSslRecommendationService]
// method instead.
type ZoneSslRecommendationService struct {
	Options []option.RequestOption
}

// NewZoneSslRecommendationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSslRecommendationService(opts ...option.RequestOption) (r *ZoneSslRecommendationService) {
	r = &ZoneSslRecommendationService{}
	r.Options = opts
	return
}

// Retrieve the SSL/TLS Recommender's recommendation for a zone.
func (r *ZoneSslRecommendationService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSslRecommendationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/recommendation", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSslRecommendationListResponse struct {
	Errors   []ZoneSslRecommendationListResponseError   `json:"errors"`
	Messages []ZoneSslRecommendationListResponseMessage `json:"messages"`
	Result   ZoneSslRecommendationListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSslRecommendationListResponseSuccess `json:"success"`
	JSON    zoneSslRecommendationListResponseJSON    `json:"-"`
}

// zoneSslRecommendationListResponseJSON contains the JSON metadata for the struct
// [ZoneSslRecommendationListResponse]
type zoneSslRecommendationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslRecommendationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslRecommendationListResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSslRecommendationListResponseErrorJSON `json:"-"`
}

// zoneSslRecommendationListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSslRecommendationListResponseError]
type zoneSslRecommendationListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslRecommendationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslRecommendationListResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSslRecommendationListResponseMessageJSON `json:"-"`
}

// zoneSslRecommendationListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSslRecommendationListResponseMessage]
type zoneSslRecommendationListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslRecommendationListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslRecommendationListResponseResult struct {
	// Identifier of a recommedation result.
	ID         string                                       `json:"id"`
	ModifiedOn time.Time                                    `json:"modified_on" format:"date-time"`
	Value      ZoneSslRecommendationListResponseResultValue `json:"value"`
	JSON       zoneSslRecommendationListResponseResultJSON  `json:"-"`
}

// zoneSslRecommendationListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSslRecommendationListResponseResult]
type zoneSslRecommendationListResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslRecommendationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslRecommendationListResponseResultValue string

const (
	ZoneSslRecommendationListResponseResultValueFlexible ZoneSslRecommendationListResponseResultValue = "flexible"
	ZoneSslRecommendationListResponseResultValueFull     ZoneSslRecommendationListResponseResultValue = "full"
	ZoneSslRecommendationListResponseResultValueStrict   ZoneSslRecommendationListResponseResultValue = "strict"
)

// Whether the API call was successful
type ZoneSslRecommendationListResponseSuccess bool

const (
	ZoneSslRecommendationListResponseSuccessTrue ZoneSslRecommendationListResponseSuccess = true
)
