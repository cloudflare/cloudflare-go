// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// RecommendationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecommendationService] method instead.
type RecommendationService struct {
	Options []option.RequestOption
}

// NewRecommendationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecommendationService(opts ...option.RequestOption) (r *RecommendationService) {
	r = &RecommendationService{}
	r.Options = opts
	return
}

// Retrieve the SSL/TLS Recommender's recommendation for a zone.
func (r *RecommendationService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *RecommendationGetResponse, err error) {
	var env RecommendationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/recommendation", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RecommendationGetResponse struct {
	// Identifier of a recommedation result.
	ID         string                         `json:"id"`
	ModifiedOn time.Time                      `json:"modified_on" format:"date-time"`
	Value      RecommendationGetResponseValue `json:"value"`
	JSON       recommendationGetResponseJSON  `json:"-"`
}

// recommendationGetResponseJSON contains the JSON metadata for the struct
// [RecommendationGetResponse]
type recommendationGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecommendationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recommendationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RecommendationGetResponseValue string

const (
	RecommendationGetResponseValueFlexible RecommendationGetResponseValue = "flexible"
	RecommendationGetResponseValueFull     RecommendationGetResponseValue = "full"
	RecommendationGetResponseValueStrict   RecommendationGetResponseValue = "strict"
)

func (r RecommendationGetResponseValue) IsKnown() bool {
	switch r {
	case RecommendationGetResponseValueFlexible, RecommendationGetResponseValueFull, RecommendationGetResponseValueStrict:
		return true
	}
	return false
}

type RecommendationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   RecommendationGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RecommendationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    recommendationGetResponseEnvelopeJSON    `json:"-"`
}

// recommendationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecommendationGetResponseEnvelope]
type recommendationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecommendationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recommendationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RecommendationGetResponseEnvelopeSuccess bool

const (
	RecommendationGetResponseEnvelopeSuccessTrue RecommendationGetResponseEnvelopeSuccess = true
)

func (r RecommendationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RecommendationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
