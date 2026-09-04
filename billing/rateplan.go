// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// RatePlanService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRatePlanService] method instead.
type RatePlanService struct {
	Options []option.RequestOption
}

// NewRatePlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRatePlanService(opts ...option.RequestOption) (r *RatePlanService) {
	r = &RatePlanService{}
	r.Options = opts
	return
}

// Gets a rate plan's details by its public key (e.g., 'teams_free',
// 'cf_pro_20_20'). This is a public catalog endpoint, so authentication is not
// enforced and credentials are accepted but not required.
func (r *RatePlanService) Get(ctx context.Context, publicKey string, opts ...option.RequestOption) (res *RatePlanGetResponse, err error) {
	var env RatePlanGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if publicKey == "" {
		err = errors.New("missing required public_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/rate_plans/%s", publicKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RatePlanGetResponse struct {
	// The uppercase rate plan public key.
	ID string `json:"id"`
	// Pricing components that make up this rate plan.
	Components []map[string]interface{} `json:"components"`
	// Currency of the rate plan pricing.
	Currency string `json:"currency"`
	// Human-readable description of the rate plan.
	PublicName string                  `json:"public_name"`
	JSON       ratePlanGetResponseJSON `json:"-"`
}

// ratePlanGetResponseJSON contains the JSON metadata for the struct
// [RatePlanGetResponse]
type ratePlanGetResponseJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	PublicName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseJSON) RawJSON() string {
	return r.raw
}

type RatePlanGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   RatePlanGetResponse   `json:"result" api:"required"`
	// Whether the API call was successful
	Success RatePlanGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    ratePlanGetResponseEnvelopeJSON    `json:"-"`
}

// ratePlanGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RatePlanGetResponseEnvelope]
type ratePlanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RatePlanGetResponseEnvelopeSuccess bool

const (
	RatePlanGetResponseEnvelopeSuccessTrue RatePlanGetResponseEnvelopeSuccess = true
)

func (r RatePlanGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RatePlanGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
