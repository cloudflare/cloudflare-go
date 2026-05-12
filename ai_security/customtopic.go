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

// CustomTopicService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomTopicService] method instead.
type CustomTopicService struct {
	Options []option.RequestOption
}

// NewCustomTopicService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomTopicService(opts ...option.RequestOption) (r *CustomTopicService) {
	r = &CustomTopicService{}
	r.Options = opts
	return
}

// Set the AI Security for Apps custom topic categories for a zone.
//
// A maximum of 20 custom topics can be configured per zone. Each topic label must
// be 2–20 characters using only lowercase letters (a–z), digits (0–9), and
// hyphens. Each topic description must be 2–50 printable ASCII characters.
//
// Changes can take up to a minute to propagate to the zone.
func (r *CustomTopicService) Update(ctx context.Context, params CustomTopicUpdateParams, opts ...option.RequestOption) (res *CustomTopicUpdateResponse, err error) {
	var env CustomTopicUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/ai-security/custom-topics", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get the AI Security for Apps custom topic categories for a zone.
func (r *CustomTopicService) Get(ctx context.Context, query CustomTopicGetParams, opts ...option.RequestOption) (res *CustomTopicGetResponse, err error) {
	var env CustomTopicGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/ai-security/custom-topics", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CustomTopicUpdateResponse struct {
	// Custom topic categories for AI Security for Apps content detection.
	Topics []CustomTopicUpdateResponseTopic `json:"topics"`
	JSON   customTopicUpdateResponseJSON    `json:"-"`
}

// customTopicUpdateResponseJSON contains the JSON metadata for the struct
// [CustomTopicUpdateResponse]
type customTopicUpdateResponseJSON struct {
	Topics      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTopicUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTopicUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CustomTopicUpdateResponseTopic struct {
	// Unique label identifier. Must contain only lowercase letters (a–z), digits
	// (0–9), and hyphens.
	Label string `json:"label" api:"required"`
	// Description of the topic category. Must contain only printable ASCII characters.
	Topic string                             `json:"topic" api:"required"`
	JSON  customTopicUpdateResponseTopicJSON `json:"-"`
}

// customTopicUpdateResponseTopicJSON contains the JSON metadata for the struct
// [CustomTopicUpdateResponseTopic]
type customTopicUpdateResponseTopicJSON struct {
	Label       apijson.Field
	Topic       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTopicUpdateResponseTopic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTopicUpdateResponseTopicJSON) RawJSON() string {
	return r.raw
}

type CustomTopicGetResponse struct {
	// Custom topic categories for AI Security for Apps content detection.
	Topics []CustomTopicGetResponseTopic `json:"topics"`
	JSON   customTopicGetResponseJSON    `json:"-"`
}

// customTopicGetResponseJSON contains the JSON metadata for the struct
// [CustomTopicGetResponse]
type customTopicGetResponseJSON struct {
	Topics      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTopicGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTopicGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomTopicGetResponseTopic struct {
	// Unique label identifier. Must contain only lowercase letters (a–z), digits
	// (0–9), and hyphens.
	Label string `json:"label" api:"required"`
	// Description of the topic category. Must contain only printable ASCII characters.
	Topic string                          `json:"topic" api:"required"`
	JSON  customTopicGetResponseTopicJSON `json:"-"`
}

// customTopicGetResponseTopicJSON contains the JSON metadata for the struct
// [CustomTopicGetResponseTopic]
type customTopicGetResponseTopicJSON struct {
	Label       apijson.Field
	Topic       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTopicGetResponseTopic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTopicGetResponseTopicJSON) RawJSON() string {
	return r.raw
}

type CustomTopicUpdateParams struct {
	// Defines the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Custom topic categories for AI Security for Apps content detection.
	Topics param.Field[[]CustomTopicUpdateParamsTopic] `json:"topics"`
}

func (r CustomTopicUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomTopicUpdateParamsTopic struct {
	// Unique label identifier. Must contain only lowercase letters (a–z), digits
	// (0–9), and hyphens.
	Label param.Field[string] `json:"label" api:"required"`
	// Description of the topic category. Must contain only printable ASCII characters.
	Topic param.Field[string] `json:"topic" api:"required"`
}

func (r CustomTopicUpdateParamsTopic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomTopicUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors" api:"required"`
	Messages []shared.ResponseInfo     `json:"messages" api:"required"`
	Result   CustomTopicUpdateResponse `json:"result" api:"required"`
	// Defines whether the API call was successful.
	Success CustomTopicUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    customTopicUpdateResponseEnvelopeJSON    `json:"-"`
}

// customTopicUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomTopicUpdateResponseEnvelope]
type customTopicUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTopicUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTopicUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type CustomTopicUpdateResponseEnvelopeSuccess bool

const (
	CustomTopicUpdateResponseEnvelopeSuccessTrue CustomTopicUpdateResponseEnvelopeSuccess = true
)

func (r CustomTopicUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomTopicUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomTopicGetParams struct {
	// Defines the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type CustomTopicGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors" api:"required"`
	Messages []shared.ResponseInfo  `json:"messages" api:"required"`
	Result   CustomTopicGetResponse `json:"result" api:"required"`
	// Defines whether the API call was successful.
	Success CustomTopicGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    customTopicGetResponseEnvelopeJSON    `json:"-"`
}

// customTopicGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomTopicGetResponseEnvelope]
type customTopicGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTopicGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTopicGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type CustomTopicGetResponseEnvelopeSuccess bool

const (
	CustomTopicGetResponseEnvelopeSuccessTrue CustomTopicGetResponseEnvelopeSuccess = true
)

func (r CustomTopicGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomTopicGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
