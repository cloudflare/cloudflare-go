// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Put Rules
func (r *RuleService) Update(ctx context.Context, params RuleUpdateParams, opts ...option.RequestOption) (res *[]RuleUpdateResponse, err error) {
	var env RuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cloud_connector/rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rules
func (r *RuleService) List(ctx context.Context, query RuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cloud_connector/rules", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Rules
func (r *RuleService) ListAutoPaging(ctx context.Context, query RuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RuleListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type RuleUpdateResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Parameters of Cloud Connector Rule
	Parameters RuleUpdateResponseParameters `json:"parameters"`
	// Cloud Provider type
	Provider RuleUpdateResponseProvider `json:"provider"`
	JSON     ruleUpdateResponseJSON     `json:"-"`
}

// ruleUpdateResponseJSON contains the JSON metadata for the struct
// [RuleUpdateResponse]
type ruleUpdateResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Parameters  apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Parameters of Cloud Connector Rule
type RuleUpdateResponseParameters struct {
	// Host to perform Cloud Connection to
	Host string                           `json:"host"`
	JSON ruleUpdateResponseParametersJSON `json:"-"`
}

// ruleUpdateResponseParametersJSON contains the JSON metadata for the struct
// [RuleUpdateResponseParameters]
type ruleUpdateResponseParametersJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseParametersJSON) RawJSON() string {
	return r.raw
}

// Cloud Provider type
type RuleUpdateResponseProvider string

const (
	RuleUpdateResponseProviderAwsS3        RuleUpdateResponseProvider = "aws_s3"
	RuleUpdateResponseProviderR2           RuleUpdateResponseProvider = "r2"
	RuleUpdateResponseProviderGcpStorage   RuleUpdateResponseProvider = "gcp_storage"
	RuleUpdateResponseProviderAzureStorage RuleUpdateResponseProvider = "azure_storage"
)

func (r RuleUpdateResponseProvider) IsKnown() bool {
	switch r {
	case RuleUpdateResponseProviderAwsS3, RuleUpdateResponseProviderR2, RuleUpdateResponseProviderGcpStorage, RuleUpdateResponseProviderAzureStorage:
		return true
	}
	return false
}

type RuleListResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Parameters of Cloud Connector Rule
	Parameters RuleListResponseParameters `json:"parameters"`
	// Cloud Provider type
	Provider RuleListResponseProvider `json:"provider"`
	JSON     ruleListResponseJSON     `json:"-"`
}

// ruleListResponseJSON contains the JSON metadata for the struct
// [RuleListResponse]
type ruleListResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Parameters  apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseJSON) RawJSON() string {
	return r.raw
}

// Parameters of Cloud Connector Rule
type RuleListResponseParameters struct {
	// Host to perform Cloud Connection to
	Host string                         `json:"host"`
	JSON ruleListResponseParametersJSON `json:"-"`
}

// ruleListResponseParametersJSON contains the JSON metadata for the struct
// [RuleListResponseParameters]
type ruleListResponseParametersJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseParametersJSON) RawJSON() string {
	return r.raw
}

// Cloud Provider type
type RuleListResponseProvider string

const (
	RuleListResponseProviderAwsS3        RuleListResponseProvider = "aws_s3"
	RuleListResponseProviderR2           RuleListResponseProvider = "r2"
	RuleListResponseProviderGcpStorage   RuleListResponseProvider = "gcp_storage"
	RuleListResponseProviderAzureStorage RuleListResponseProvider = "azure_storage"
)

func (r RuleListResponseProvider) IsKnown() bool {
	switch r {
	case RuleListResponseProviderAwsS3, RuleListResponseProviderR2, RuleListResponseProviderGcpStorage, RuleListResponseProviderAzureStorage:
		return true
	}
	return false
}

type RuleUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// List of Cloud Connector rules
	Body []RuleUpdateParamsBody `json:"body,required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleUpdateParamsBody struct {
	ID          param.Field[string] `json:"id"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Expression  param.Field[string] `json:"expression"`
	// Parameters of Cloud Connector Rule
	Parameters param.Field[RuleUpdateParamsBodyParameters] `json:"parameters"`
	// Cloud Provider type
	Provider param.Field[RuleUpdateParamsBodyProvider] `json:"provider"`
}

func (r RuleUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Parameters of Cloud Connector Rule
type RuleUpdateParamsBodyParameters struct {
	// Host to perform Cloud Connection to
	Host param.Field[string] `json:"host"`
}

func (r RuleUpdateParamsBodyParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloud Provider type
type RuleUpdateParamsBodyProvider string

const (
	RuleUpdateParamsBodyProviderAwsS3        RuleUpdateParamsBodyProvider = "aws_s3"
	RuleUpdateParamsBodyProviderR2           RuleUpdateParamsBodyProvider = "r2"
	RuleUpdateParamsBodyProviderGcpStorage   RuleUpdateParamsBodyProvider = "gcp_storage"
	RuleUpdateParamsBodyProviderAzureStorage RuleUpdateParamsBodyProvider = "azure_storage"
)

func (r RuleUpdateParamsBodyProvider) IsKnown() bool {
	switch r {
	case RuleUpdateParamsBodyProviderAwsS3, RuleUpdateParamsBodyProviderR2, RuleUpdateParamsBodyProviderGcpStorage, RuleUpdateParamsBodyProviderAzureStorage:
		return true
	}
	return false
}

type RuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	// List of Cloud Connector rules
	Result []RuleUpdateResponse           `json:"result"`
	JSON   ruleUpdateResponseEnvelopeJSON `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

func (r RuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
