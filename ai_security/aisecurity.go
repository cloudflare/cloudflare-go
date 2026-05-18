// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// AISecurityService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAISecurityService] method instead.
type AISecurityService struct {
	Options      []option.RequestOption
	CustomTopics *CustomTopicService
}

// NewAISecurityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAISecurityService(opts ...option.RequestOption) (r *AISecurityService) {
	r = &AISecurityService{}
	r.Options = opts
	r.CustomTopics = NewCustomTopicService(opts...)
	return
}

// Enable or disable AI Security for Apps for a zone.
//
// Changes can take up to a minute to propagate to the zone.
func (r *AISecurityService) Update(ctx context.Context, params AISecurityUpdateParams, opts ...option.RequestOption) (res *AISecurityUpdateResponse, err error) {
	var env AISecurityUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/ai-security/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get whether AI Security for Apps is enabled or disabled for a zone.
func (r *AISecurityService) Get(ctx context.Context, query AISecurityGetParams, opts ...option.RequestOption) (res *AISecurityGetResponse, err error) {
	var env AISecurityGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/ai-security/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// AI Security for Apps enablement status for a zone.
type AISecurityUpdateResponse struct {
	// Whether AI Security for Apps is enabled on the zone.
	Enabled bool                         `json:"enabled"`
	JSON    aiSecurityUpdateResponseJSON `json:"-"`
}

// aiSecurityUpdateResponseJSON contains the JSON metadata for the struct
// [AISecurityUpdateResponse]
type aiSecurityUpdateResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AISecurityUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiSecurityUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// AI Security for Apps enablement status for a zone.
type AISecurityGetResponse struct {
	// Whether AI Security for Apps is enabled on the zone.
	Enabled bool                      `json:"enabled"`
	JSON    aiSecurityGetResponseJSON `json:"-"`
}

// aiSecurityGetResponseJSON contains the JSON metadata for the struct
// [AISecurityGetResponse]
type aiSecurityGetResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AISecurityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiSecurityGetResponseJSON) RawJSON() string {
	return r.raw
}

type AISecurityUpdateParams struct {
	// Defines the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Whether AI Security for Apps is enabled on the zone.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AISecurityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AISecurityUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// AI Security for Apps enablement status for a zone.
	Result AISecurityUpdateResponse `json:"result" api:"required"`
	// Defines whether the API call was successful.
	Success AISecurityUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    aiSecurityUpdateResponseEnvelopeJSON    `json:"-"`
}

// aiSecurityUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AISecurityUpdateResponseEnvelope]
type aiSecurityUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AISecurityUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiSecurityUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AISecurityUpdateResponseEnvelopeSuccess bool

const (
	AISecurityUpdateResponseEnvelopeSuccessTrue AISecurityUpdateResponseEnvelopeSuccess = true
)

func (r AISecurityUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AISecurityUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AISecurityGetParams struct {
	// Defines the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type AISecurityGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// AI Security for Apps enablement status for a zone.
	Result AISecurityGetResponse `json:"result" api:"required"`
	// Defines whether the API call was successful.
	Success AISecurityGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    aiSecurityGetResponseEnvelopeJSON    `json:"-"`
}

// aiSecurityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AISecurityGetResponseEnvelope]
type aiSecurityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AISecurityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiSecurityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AISecurityGetResponseEnvelopeSuccess bool

const (
	AISecurityGetResponseEnvelopeSuccessTrue AISecurityGetResponseEnvelopeSuccess = true
)

func (r AISecurityGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AISecurityGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
